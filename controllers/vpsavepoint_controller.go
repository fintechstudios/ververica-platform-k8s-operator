/*
Copyright 2019 FinTech Studios, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1/converters"
	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-operator/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/annotations"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/app-manager"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/polling"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/utils"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VpSavepointReconciler reconciles a VpSavepoint object
type VpSavepointReconciler struct {
	client.Client
	Log                 logr.Logger
	AppManagerAPIClient *appManagerApi.APIClient
	AppManagerAuthStore *appmanager.AuthStore
	pollerMap           map[string]*polling.Poller
}

func (r *VpSavepointReconciler) addStatusPollerForResource(req ctrl.Request, vpSavepoint *v1beta1.VpSavepoint) {
	log := r.getLogger(req)

	if r.pollerMap == nil {
		r.pollerMap = make(map[string]*polling.Poller)
	}

	if r.pollerMap[req.String()] != nil {
		log.Info("A status poller already exists, removing...")
		r.removeStatusPollerForResource(req)
	}

	nsName := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	vpID := annotations.Get(vpSavepoint.Annotations, annotations.ID)

	// On each polling callback, push the update through the k8s client
	poller := polling.NewPoller(func() interface{} {
		ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
		savepoint, _, err := r.AppManagerAPIClient.SavepointsApi.GetSavepoint(ctx, nsName, vpID)
		if err != nil {
			log.Error(err, "Error while polling savepoint")
			return nil
		}

		var vpSavepointUpdated v1beta1.VpSavepoint
		if err = r.Get(ctx, req.NamespacedName, &vpSavepointUpdated); err != nil {
			if utils.IsNotFoundError(err) {
				// TODO: should we force stop polling?
				log.Error(err, "VpSavepoint not found while polling")
			} else {
				log.Error(err, "Error getting VpSavepoint while polling")
			}
			return savepoint
		}

		if err = r.updateResource(&vpSavepointUpdated, &savepoint); err != nil {
			log.Error(err, "Error while updating VpSavepoint from poller")
		}

		return nil
	}, time.Second*5)

	r.pollerMap[req.String()] = poller
	poller.Start()
}

func (r *VpSavepointReconciler) removeStatusPollerForResource(req ctrl.Request) bool {
	if r.pollerMap == nil {
		return false
	}
	log := r.getLogger(req)
	poller := r.pollerMap[req.String()]
	if poller == nil {
		return false
	}

	log.Info("Stopping poller")
	poller.Stop()
	delete(r.pollerMap, req.String())
	return true
}

// getLogger creates a logger for the controller with the request name
func (r *VpSavepointReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpsavepoint", req.NamespacedName)
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpSavepointReconciler) updateResource(resource *v1beta1.VpSavepoint, savepoint *appManagerApi.Savepoint) error {
	ctx := context.Background()

	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	annotations.Set(resource.Annotations,
		annotations.Pair(annotations.ID, savepoint.Metadata.Id),
		annotations.Pair(annotations.DeploymentID, savepoint.Metadata.DeploymentId),
		annotations.Pair(annotations.JobID, savepoint.Metadata.JobId))

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	state, err := converters.SavepointStateToNative(savepoint.Status.State)
	if err != nil {
		return err
	}
	resource.Status.State = state

	if err := r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

func (r *VpSavepointReconciler) handleCreate(req ctrl.Request, vpSavepoint v1beta1.VpSavepoint) (ctrl.Result, error) {
	log := r.getLogger(req)

	nsName := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}
	depID := vpSavepoint.Spec.Metadata.DeploymentID
	if depID == "" {
		// no deployment id has been explicitly set
		// try to find one
		depName := vpSavepoint.Spec.DeploymentName
		deployment, err := appmanager.GetDeploymentByName(ctx, r.AppManagerAPIClient, nsName, depName)

		if utils.IsNotFoundError(err) {
			log.Info("No deployment by name %s", depName)
			return ctrl.Result{}, err
		}

		if err != nil {
			return ctrl.Result{}, err
		}

		depID = deployment.Metadata.Id
	}

	createdSavepoint, res, err := r.AppManagerAPIClient.SavepointsApi.CreateSavepoint(ctx, nsName, appManagerApi.Savepoint{
		Kind:       "Savepoint",
		ApiVersion: "v1",
		Metadata: &appManagerApi.SavepointMetadata{
			DeploymentId: depID,
			Namespace:    nsName,
		},
	})

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		log.Error(err, "Bad request when creating savepoint")
		return ctrl.Result{Requeue: false}, nil
	}

	if err != nil {
		log.Info("Error creating VP Savepoint")
		return ctrl.Result{}, err
	}

	// Now update the k8s resource and status as well
	if err := r.updateResource(&vpSavepoint, &createdSavepoint); err != nil {
		return ctrl.Result{}, err
	}

	// Create a poller to keep the savepoint up to date
	r.addStatusPollerForResource(req, &vpSavepoint)

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpsavepoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpsavepoints/status,verbs=get;update;patch

// Reconcile is the main entrypoint for the reconciliation loop
func (r *VpSavepointReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	var vpSavepoint v1beta1.VpSavepoint
	// If it's gone, it's gone!
	if err := r.Get(ctx, req.NamespacedName, &vpSavepoint); err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if !vpSavepoint.ObjectMeta.DeletionTimestamp.IsZero() {
		// Being deleted
		log.Info("Delete event")
		r.removeStatusPollerForResource(req)
		log.Info("Warning: Deletion is not supported through the Ververica Platform. All savepoints must be manually cleaned up.")
		return ctrl.Result{}, nil
	}

	// Id has not been set - must be creating
	if annotations.Has(vpSavepoint.Annotations, annotations.ID) {
		log.Info("Creating savepoint")
		return r.handleCreate(req, vpSavepoint)
	}

	savepointID := annotations.Get(vpSavepoint.Annotations, annotations.ID)
	nsName := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	appManagerCtx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	savepoint, _, err := r.AppManagerAPIClient.SavepointsApi.GetSavepoint(appManagerCtx, nsName, savepointID)
	if err != nil {
		if utils.IsNotFoundError(err) {
			log.Info("Savepoint by id not found - creating", "id", savepointID)
			return r.handleCreate(req, vpSavepoint)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log = log.WithValues("savepoint", savepoint.Metadata.Id)
	log.Info("Update event")
	log.Info("Warning: Updates are not supported through the Ververica Platform.")
	return ctrl.Result{}, nil
}

// SetupWithManager registers the controller
func (r *VpSavepointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.VpSavepoint{}).
		Complete(r)
}
