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

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	ververicaplatformapi "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
)

// DeploymentTargetReconciler reconciles a DeploymentTarget object
type DeploymentTargetReconciler struct {
	client.Client
	Log                logr.Logger
	VervericaAPIClient ververicaplatformapi.APIClient
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=deploymenttargets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=deploymenttargets/status,verbs=get;update;patch

// Reconcile tries to make the current state more like the desired state
func (r *DeploymentTargetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("deploymenttarget", req.NamespacedName)

	log.Info("Built api client", nil)
	// your logic here
	target, _, err := r.VervericaAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, "default", req.Name)

	if err != nil {
		// TODO: think about ignoring not-found errors, as they won't
		// 		 be immediately solved by re-queueing
		return ctrl.Result{}, err
	}
	log.Info("Found target", target)

	return ctrl.Result{}, nil
}

// SetupWithManager is a helper function to initial on manager boot
func (r *DeploymentTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.DeploymentTarget{}).
		Complete(r)
}
