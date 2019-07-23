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
	"github.com/antihax/optional"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// VPNamespaceReconciler reconciles a VPNamespace object
type VPNamespaceReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient vpAPI.APIClient
}

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VPNamespaceReconciler) updateResource(req ctrl.Request, resource *ververicaplatformv1beta1.VPNamespace, namespace *vpAPI.Namespace) error {
	ctx := context.Background()

	resource.Name = namespace.Metadata.Name
	//time.Parse(time.RFC3339, namespace.Metadata.CreatedAt)
	resource.Spec.Metadata = ververicaplatformv1beta1.VPNamespaceMetadata{
		Name:            namespace.Metadata.Name,
		Id:              namespace.Metadata.Id,
		CreatedAt:       &metav1.Time{Time: namespace.Metadata.CreatedAt},
		ModifiedAt:      &metav1.Time{Time: namespace.Metadata.ModifiedAt},
		ResourceVersion: namespace.Metadata.ResourceVersion,
	}
	resource.Status.State = namespace.Status.State

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	if err := r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// getLogger creates a logger for the controller with the request name
func (r *VPNamespaceReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpnamespace", req.NamespacedName)
}

// handleCreate creates VP resources
func (r *VPNamespaceReconciler) handleCreate(req ctrl.Request) (time.Duration, error) {
	ctx := context.Background()
	log := r.getLogger(req)
	var vpNamespace ververicaplatformv1beta1.VPNamespace
	if err := r.Get(ctx, req.NamespacedName, &vpNamespace); err != nil {
		return 0, err
	}

	// create it
	namespace, _, err := r.VPAPIClient.NamespacesApi.PostNamespace(ctx, &vpAPI.PostNamespaceOpts{
		Body: optional.NewInterface(vpAPI.Namespace{
			ApiVersion: "v1",
			Metadata: &vpAPI.NamespaceMetadata{
				Name: req.Name,
			},
		}),
	})

	if err != nil {
		log.Error(err, "Error creating VP namespace")
		return 0, err
	}
	log.Info("Created namespace", "namespace", namespace)

	// Now update the k8s resource and status as well
	if err := r.updateResource(req, &vpNamespace, &namespace); err != nil {
		return 0, err
	}

	return 0, nil
}

// handle update updates the k8s resource when it already exists in the VP
func (r *VPNamespaceReconciler) handleUpdate(req ctrl.Request, namespace vpAPI.Namespace) (time.Duration, error) {
	ctx := context.Background()
	var vpNamespace ververicaplatformv1beta1.VPNamespace
	if err := r.Get(ctx, req.NamespacedName, &vpNamespace); err != nil {
		return 0, err
	}

	// Now update the k8s resource and status as well
	if err := r.updateResource(req, &vpNamespace, &namespace); err != nil {
		return 0, err
	}

	return 0, nil
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VPNamespaceReconciler) handleDelete(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)
	// Let's make sure it's deleted from the ververica platform
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
func (r *VPNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// otherwise, let's check if it exists in the cluster
	// If so, it's been deleted
	var vpNamespace ververicaplatformv1beta1.VPNamespace
	if err := r.Get(ctx, req.NamespacedName, &vpNamespace); err != nil {
		log.Info("Not Found event", "name", req.Name)
		// If it is not stored, must make sure it is deleted from VP as well
		res, err := r.handleDelete(req)
		return res, err
	}

	if vpNamespace.ObjectMeta.DeletionTimestamp.IsZero() {
		// Not being deleted, add the finalizer
		if utils.AddFinalizerToObjectMeta(&vpNamespace.ObjectMeta) {
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
		// TODO: not super happy with this flow, but will keep thinking on it
		if err != nil || res.RequeueAfter > 0 || res.Requeue {
			// if fail to delete the external dependency here,
			// requeue so that it can be retried
			return res, err
		}
		// otherwise, we're all good, just remove the finalizer
		if utils.RemoveFinalizerFromObjectMeta(&vpNamespace.ObjectMeta) {
			if err := r.Update(ctx, &vpNamespace); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	namespace, _, err := r.VPAPIClient.NamespacesApi.GetNamespace(nil, req.Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Not found, let's create it
			dur, err := r.handleCreate(req)
			return utils.EventHandlerResponse(dur, err)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log.Info("Update event", "vp namespace", namespace.Metadata.Name)
	dur, err := r.handleUpdate(req, namespace)
	return utils.EventHandlerResponse(dur, err)
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VPNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VPNamespace{}).
		Complete(r)
}
