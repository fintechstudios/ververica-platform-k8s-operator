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

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"
	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/app-manager"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VpSavepointReconciler reconciles a VpSavepoint object
type VpSavepointReconciler struct {
	client.Client
	Log                 logr.Logger
	AppManagerApiClient *appManagerApi.APIClient
	AppManagerAuthStore *appManager.AuthStore
}

// getLogger creates a logger for the controller with the request name
func (r *VpSavepointReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpsavepoint", req.NamespacedName)
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpSavepointReconciler) updateResource(vpSavepoint *v1beta1.VpSavepoint, savepoint *appManagerApi.Savepoint) error {
	ctx := context.Background()

	metadata, err := converters.SavepointMetadataToNative(*savepoint.Metadata)

	if err != nil {
		return err
	}
	vpSavepoint.Spec.Metadata = metadata

	spec, err := converters.SavepointSpecToNative(*savepoint.Spec)
	if err != nil {
		return err
	}
	vpSavepoint.Spec.Spec = spec

	state, err := converters.SavepointStateToNative(savepoint.Status.State)
	if err != nil {
		return err
	}
	vpSavepoint.Status.State = state

	if err := r.Update(ctx, vpSavepoint); err != nil {
		return err
	}

	if err := r.Status().Update(ctx, vpSavepoint); err != nil {
		return err
	}

	return nil
}

func (r *VpSavepointReconciler) handleCreate(req ctrl.Request, vpSavepoint v1beta1.VpSavepoint) (ctrl.Result, error) {
	log := r.getLogger(req)

	nsName := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue:false}, nil
	}
	depId := vpSavepoint.Spec.Metadata.DeploymentID
	if depId == "" {
		// no deployment id has been explicitly set
		// try to find one
		depName := vpSavepoint.Spec.DeploymentName
		deployment, err := appManager.GetDeploymentByName(r.AppManagerApiClient, ctx, nsName, depName)

		if utils.IsNotFoundError(err) {
			log.Info("No deployment by name %s", depName)
			return ctrl.Result{}, err
		}

		if err != nil {
			return ctrl.Result{}, err
		}

		depId = deployment.Metadata.Id
	}

	createdSavepoint, res, err := r.AppManagerApiClient.SavepointsApi.CreateSavepoint(ctx, nsName, appManagerApi.Savepoint{
		Kind:       "Savepoint",
		ApiVersion: "v1",
		Metadata: &appManagerApi.SavepointMetadata{
			DeploymentId: depId,
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
		log.Info("Warning: Deletion is not supported through the Ververica Platform. All savepoints must be manually cleaned.")
		return ctrl.Result{}, nil
	}

	// Id has not been set - must be creating
	savepointId := vpSavepoint.Spec.Metadata.ID
	if vpSavepoint.Spec.Metadata.ID == "" {
		log.Info("Creating savepoint")
		return r.handleCreate(req, vpSavepoint)
	}

	nsName := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	appManagerCtx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue:false}, nil
	}

	savepoint, _, err := r.AppManagerApiClient.SavepointsApi.GetSavepoint(appManagerCtx, nsName, savepointId)
	if err != nil {
		if utils.IsNotFoundError(err) {
			log.Info("Savepoint by id not found - creating", "id", savepointId)
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
