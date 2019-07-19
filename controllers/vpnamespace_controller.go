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

	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	ververicaplatformapi "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// VPNamespaceReconciler reconciles a VPNamespace object
type VPNamespaceReconciler struct {
	client.Client
	Log                logr.Logger
	VervericaAPIClient ververicaplatformapi.APIClient
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpnamespaces/status,verbs=get;update;patch

// Reconcile tries to make the current state more like the desired state
func (r *VPNamespaceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("vpnamespace", req.NamespacedName)

	// otherwise, let's check if it exists in the cluster
	// If so, it's been deleted
	var vpResNamespace ververicaplatformv1beta1.VPNamespace
	if err := r.Get(ctx, req.NamespacedName, &vpResNamespace); err != nil {
		log.Info("Deletion event", "name", req.Name)
		// Let's make sure it's deleted from the ververica platform
		_, _, err := r.VervericaAPIClient.NamespacesApi.DeleteNamespace(ctx, req.Name)
		if err != nil {
			// If it's already gone, great!
			return ctrl.Result{}, utils.IgnoreNotFoundError(err)
		}
		log.Info("Deleted namespace", "name", req.Name)
	} else {
		// Let's make sure it's created in the ververica platform

		// if it's already created
		namespace, _, err := r.VervericaAPIClient.NamespacesApi.GetNamespace(nil, req.Name)

		if err == nil {
			log.Info("Namespace exists. Maybe update event?")
			// Now update the k8s resource and status as well
			vpResNamespace.Status.State = namespace.Status.State

			if err := r.Status().Update(ctx, &vpResNamespace); err != nil {
				log.Error(err, "Unable to update VPNamespace status")
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}

		log.Info("Creation event", "vp namespace", req.Name)
		if utils.IsNotFoundError(err) {
			// create it
			namespace, _, err := r.VervericaAPIClient.NamespacesApi.PostNamespace(ctx, &ververicaplatformapi.PostNamespaceOpts{
				Body: optional.NewInterface(ververicaplatformapi.Namespace{
					ApiVersion: "v1",
					Metadata: &ververicaplatformapi.NamespaceMetadata{
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
			vpResNamespace.Status.State = namespace.Status.State

			if err := r.Status().Update(ctx, &vpResNamespace); err != nil {
				log.Error(err, "Unable to update VPNamespace status")
				return ctrl.Result{}, err
			}
		} else {
			// Requeue
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VPNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VPNamespace{}).
		Complete(r)
}
