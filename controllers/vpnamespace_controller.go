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

	"github.com/antihax/optional"
	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"

	appManager "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// VpNamespaceReconciler reconciles a VpNamespace object
type VpNamespaceReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient *appManager.APIClient
}

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VpNamespaceReconciler) updateResource(resource *v1beta1.VpNamespace, namespace *appManager.Namespace) error {
	ctx := context.Background()

	resource.Name = namespace.Metadata.Name
	resource.Spec.Metadata = v1beta1.VpNamespaceMetadata{
		Name:            namespace.Metadata.Name,
		ID:              namespace.Metadata.Id,
		CreatedAt:       &metav1.Time{Time: namespace.Metadata.CreatedAt},
		ModifiedAt:      &metav1.Time{Time: namespace.Metadata.ModifiedAt},
		ResourceVersion: namespace.Metadata.ResourceVersion,
	}

	var err error
	if resource.Status.State, err = converters.NamespaceStateToNative(namespace.Status.State); err != nil {
		return err
	}

	if err := r.Update(ctx, resource); err != nil {
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
	ctx := context.Background()
	log := r.getLogger(req)

	// create it
	namespace, _, err := r.VPAPIClient.NamespacesApi.PostNamespace(ctx, &appManager.PostNamespaceOpts{
		Body: optional.NewInterface(appManager.Namespace{
			ApiVersion: "v1",
			Metadata: &appManager.NamespaceMetadata{
				Name: req.Name,
			},
		}),
	})

	if err != nil {
		log.Error(err, "Error creating VP namespace")
		return ctrl.Result{}, err
	}
	log.Info("Created namespace", "namespace", namespace)

	// Now update the k8s resource and status as well
	if err := r.updateResource(&vpNamespace, &namespace); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
func (r *VpNamespaceReconciler) handleUpdate(req ctrl.Request, vpNamespace v1beta1.VpNamespace, namespace appManager.Namespace) (ctrl.Result, error) {
	// Now update the k8s resource and status as well
	err := r.updateResource(&vpNamespace, &namespace)
	return ctrl.Result{}, err
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpNamespaceReconciler) handleDelete(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)
	// Let's make sure it's deleted from the ververica platform
	// Should be idempotent, so retrying shouldn't matter
	namespace, _, err := r.VPAPIClient.NamespacesApi.DeleteNamespace(ctx, req.Name)

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting namespace", "name", namespace.Metadata.Id)

	if namespace.Status.State == "MARKED_FOR_DELETION" {
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

	namespace, _, err := r.VPAPIClient.NamespacesApi.GetNamespace(ctx, req.Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Not found, let's create it
			return r.handleCreate(req, vpNamespace)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log.Info("Update event", "vp namespace", namespace.Metadata.Name)
	return r.handleUpdate(req, vpNamespace, namespace)
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VpNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.VpNamespace{}).
		Complete(r)
}
