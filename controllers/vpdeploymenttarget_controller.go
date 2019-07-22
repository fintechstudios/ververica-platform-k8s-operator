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
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	"time"

	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
)

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VPDeploymentTargetReconciler) updateResource(req ctrl.Request, resource *ververicaplatformv1beta1.VPDeploymentTarget, namespace *vpAPI.DeploymentTarget) error {
	ctx := context.Background()

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	if err := r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// getLogger creates a logger for the controller with the request name
func (r *VPDeploymentTargetReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpdeploymenttarget", req.NamespacedName)
}

// handleCreate creates VP resources
func (r *VPDeploymentTargetReconciler) handleCreate(req ctrl.Request) (time.Duration, error) {
	//ctx := context.Background()
	//log := r.getLogger(req)
	//var vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget
	//if err := r.Get(ctx, req.NamespacedName, &vpDepTarget); err != nil {
	//	return 0, err
	//}
	//
	//var namespace string
	//if vpDepTarget.Spec.Metadata.Namespace != "" {
	//	namespace = vpDepTarget.Spec.Metadata.Namespace
	//} else {
	//	namespace = "default"
	//}

	// create it
	//depTarget, _, err := r.VPAPIClient.DeploymentTargetsApi.CreateDeploymentTarget(ctx, namespace, &vpAPI.DeploymentTarget{
	//	ApiVersion: "v1",
	//	Metadata: &vpAPI.DeploymentTargetMetadata{
	//		Name: req.Name,
	//		Labels: vpDepTarget.Spec.Metadata.Labels,
	//		Annotations: vpDepTarget.Spec.Metadata.Annotations,
	//	},
	//})
	//
	//if err != nil {
	//	log.Error(err, "Error creating VP namespace")
	//	return 0, err
	//}
	//log.Info("Created namespace", "namespace", namespace)
	//
	//// Now update the k8s resource and status as well
	//if err := r.updateResource(req, &vpDepTarget, &depTarget); err != nil {
	//	return 0, err
	//}

	return 0, nil
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VPDeploymentTargetReconciler) handleDelete(req ctrl.Request) (time.Duration, error) {
	// Cannot be deleted if there are still associated deployments
	return 0, nil
}

// VPDeploymentTargetReconciler reconciles a VPDeploymentTarget object
type VPDeploymentTargetReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient vpAPI.APIClient
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeploymenttargets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeploymenttargets/status,verbs=get;update;patch

// Reconcile tries to make the current state more like the desired state
func (r *VPDeploymentTargetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	var vpDepTarget ververicaplatformv1beta1.VPDeploymentTarget
	if err := r.Get(ctx, req.NamespacedName, &vpDepTarget); err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if vpDepTarget.ObjectMeta.DeletionTimestamp.IsZero() {
		// Not being deleted, add the finalizer
		if utils.AddFinalizerToObjectMeta(&vpDepTarget.ObjectMeta) {
			log.Info("Adding Finalizer")
			if err := r.Update(ctx, &vpDepTarget); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// Being deleted
		log.Info("Deletion event", "name", req.Name)
		dur, err := r.handleDelete(req)
		res, err := utils.EventHandlerResponse(dur, err)
		// TODO: not super happy with this flow, but will keep thinking on it
		if err != nil || res.RequeueAfter > 0 {
			// if fail to delete the external dependency here, return with error
			// so that it can be retried
			return res, err
		}
		// otherwise, we're all good, just remove the finalizer
		if utils.RemoveFinalizerFromObjectMeta(&vpDepTarget.ObjectMeta) {
			if err := r.Update(ctx, &vpDepTarget); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager is a helper function to initial on manager boot
func (r *VPDeploymentTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VPDeploymentTarget{}).
		Complete(r)
}
