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

	appManager "github.com/fintechstudios/ververica-platform-k8s-controller/controllers/app-manager"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/converters"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	platformApiClient "github.com/fintechstudios/ververica-platform-k8s-controller/platform-api-client"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// VpNamespaceReconciler reconciles a VpNamespace object
type VpNamespaceReconciler struct {
	client.Client
	Log                 logr.Logger
	AppManagerAuthStore *appManager.AuthStore
	PlatformApiClient   *platformApiClient.APIClient
}

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VpNamespaceReconciler) updateResource(resource *v1beta1.VpNamespace, namespace *platformApiClient.Namespace) error {
	ctx := context.Background()

	var err error
	if resource.Status.LifecyclePhase, err = converters.NamespaceLifecyclePhaseToNative(*namespace.LifecyclePhase); err != nil {
		return err
	}

	if err := r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// getLogger creates a logger for the controller with the request name
func (r *VpNamespaceReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpnamespace", req.NamespacedName)
}

// handleCreate creates VP resources
func (r *VpNamespaceReconciler) handleCreate(req ctrl.Request, vpNamespace v1beta1.VpNamespace) (ctrl.Result, error) {
	log := r.getLogger(req)
	ctx := context.Background()
	// create it
	createRes, _, err := r.PlatformApiClient.NamespacesApi.CreateNamespace(ctx, platformApiClient.Namespace{
		Name:         "namespaces/" + req.Name,
		RoleBindings: converters.NamespaceRoleBindingsFromNative(vpNamespace.Spec.RoleBindings),
	})

	if err != nil {
		log.Info("Error creating VP namespace")
		return ctrl.Result{}, err
	}
	log.Info("Created namespace", "namespace", createRes.Namespace)

	// Now update the k8s resource and status as well
	if err := r.updateResource(&vpNamespace, createRes.Namespace); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
func (r *VpNamespaceReconciler) handleUpdate(req ctrl.Request, vpNamespace v1beta1.VpNamespace, currentNamespace platformApiClient.Namespace) (ctrl.Result, error) {
	ctx := context.Background()

	// lifecyclePhase and createTime must be left nil
	updatedNamespace := platformApiClient.Namespace{
		Name:         "namespaces/" + req.Name,
		RoleBindings: converters.NamespaceRoleBindingsFromNative(vpNamespace.Spec.RoleBindings),
	}
	updateRes, _, err := r.PlatformApiClient.NamespacesApi.UpdateNamespace(ctx, updatedNamespace, req.Name)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.updateResource(&vpNamespace, updateRes.Namespace)
	return ctrl.Result{}, err
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpNamespaceReconciler) handleDelete(req ctrl.Request) (ctrl.Result, error) {
	log := r.getLogger(req)
	ctx := context.Background()
	// Let's make sure it's deleted from the ververica platform
	// Should be idempotent, so retrying shouldn't matter
	namespaceRes, _, err := r.PlatformApiClient.NamespacesApi.DeleteNamespace(ctx, "namespaces/"+req.Name)

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting namespace")

	lifecylePhase, err := converters.NamespaceLifecyclePhaseToNative(*namespaceRes.Namespace.LifecyclePhase)
	if err != nil {
		return ctrl.Result{}, err
	}

	if lifecylePhase == v1beta1.TerminatingNamespaceLifecyclePhase {
		// Requeue for 5 seconds to wait for the namespace to be deleted
		log.Info("Requeueing deletion request for 5 seconds")
		return ctrl.Result{RequeueAfter: time.Second * 5}, nil
	}

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpnamespaces/status,verbs=get;update;patch

// Reconcile tries to make the current state more like the desired state
func (r *VpNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// otherwise, let's check if it exists in the cluster
	// If so, it's been deleted
	var vpNamespace v1beta1.VpNamespace
	if err := r.Get(ctx, req.NamespacedName, &vpNamespace); err != nil {
		log.Info("Not Found event", "name", req.Name)
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if vpNamespace.ObjectMeta.DeletionTimestamp.IsZero() {
		// Not being deleted, add the finalizer
		if utils.AddFinalizer(&vpNamespace.ObjectMeta) {
			log.Info("Adding Finalizer")
			if err := r.Update(ctx, &vpNamespace); err != nil {
				return ctrl.Result{}, err
			}
		}
		// Continue on processing
	} else {
		// Being deleted
		log.Info("Deletion event", "name", req.Name)
		res, err := r.handleDelete(req)
		if utils.IsRequeueResponse(res, err) {
			// if the external dependency is still deleting or there was an error,
			// requeue so that it can be retried
			return res, err
		}
		// otherwise, we're all good, just remove the finalizer
		if utils.RemoveFinalizer(&vpNamespace.ObjectMeta) {
			if err := r.Update(ctx, &vpNamespace); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	namespaceRes, _, err := r.PlatformApiClient.NamespacesApi.GetNamespace(context.Background(), req.Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Not found, let's create it
			return r.handleCreate(req, vpNamespace)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log.Info("Update event", "vp namespace", namespaceRes.Namespace.Name)
	return r.handleUpdate(req, vpNamespace, *namespaceRes.Namespace)
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VpNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.VpNamespace{}).
		Complete(r)
}
