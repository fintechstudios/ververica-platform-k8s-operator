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

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	vpAPIHelpers "github.com/fintechstudios/ververica-platform-k8s-controller/controllers/vp_api_helpers"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	// "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"
	// "github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VpSavepointReconciler reconciles a VpSavepoint object
type VpSavepointReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient *vpAPI.APIClient
}

// getLogger creates a logger for the controller with the request name
func (r *VpSavepointReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpsavepoint", req.NamespacedName)
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpSavepointReconciler) updateResource(resource *ververicaplatformv1beta1.VpSavepoint, deployment *vpAPI.Savepoint) error {
	// ctx := context.Background()
	//
	// metadata, err := converters.DeploymentMetadataToNative(*deployment.Metadata)
	//
	// if err != nil {
	// 	return err
	// }
	// resource.Spec.Metadata = metadata
	//
	// spec, err := converters.DeploymentSpecToNative(*deployment.Spec)
	// if err != nil {
	// 	return err
	// }
	// resource.Spec.Spec = spec
	//
	// state, err := converters.DeploymentStateToNative(deployment.Status.State)
	// if err != nil {
	// 	return err
	// }
	// resource.Status.State = state
	//
	// if err := r.Update(ctx, resource); err != nil {
	// 	return err
	// }
	//
	// if err := r.Status().Update(ctx, resource); err != nil {
	// 	return err
	// }

	return nil
}

func (r *VpSavepointReconciler) handleCreate(req ctrl.Request, vpSavepoint ververicaplatformv1beta1.VpSavepoint) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	namespace := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)

	savepoint, res, err := r.VPAPIClient.SavepointsApi.CreateSavepoint(ctx, namespace, vpAPI.Savepoint{
		Kind:       "Savepoint",
		ApiVersion: "v1",
		Metadata: &vpAPI.SavepointMetadata{
			Id:              "",
			CreatedAt:       time.Time{},
			ModifiedAt:      time.Time{},
			DeploymentId:    "",
			JobId:           "",
			Origin:          "",
			Annotations:     nil,
			ResourceVersion: 0,
			Namespace:       "",
		},
		Status: &vpAPI.SavepointStatus{
			State:   "",
			Failure: nil,
		},
		Spec:       &vpAPI.SavepointSpec{
			SavepointLocation: "",
			FlinkSavepointId:  "",
		},
	})

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		return ctrl.Result{Requeue: false}, err
	}

	if err != nil {
		log.Error(err, "Error creating VP Savepoint")
		return ctrl.Result{}, err
	}

	// Now update the k8s resource and status as well
	if err := r.updateResource(&vpSavepoint, &savepoint); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpsavepoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpsavepoints/status,verbs=get;update;patch

func (r *VpSavepointReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	var vpSavepoint ververicaplatformv1beta1.VpSavepoint
	// If it's gone, it's gone!
	if err := r.Get(ctx, req.NamespacedName, &vpSavepoint); err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if !vpSavepoint.ObjectMeta.DeletionTimestamp.IsZero() {
		// Being deleted
		log.Info("Delete event")
		log.Info("Warn: Deletion is not supported through the Ververica Platform. All savepoints must be manually cleaned.")
		return ctrl.Result{}, nil
	}

	namespace := utils.GetNamespaceOrDefault(vpSavepoint.Spec.Metadata.Namespace)
	depId := vpSavepoint.Spec.Metadata.DeploymentID
	if len(depId) == 0 {
		// no deployment id has been explicitly set
		// try to find one
		depName := vpSavepoint.Spec.DeploymentName
		deployment, err := vpAPIHelpers.GetDeploymentByName(r.VPAPIClient, ctx, namespace, depName)

		if utils.IsNotFoundError(err) {
			log.Error(err, "No deployment by name %s", depName)
			// TODO: update status with message
			return ctrl.Result{Requeue: false}, err
		}

		if err != nil {
			return ctrl.Result{}, err
		}

		depId = deployment.Metadata.Id
	}

	savepointId := vpSavepoint.Spec.Metadata.ID
	savepoint, _, err := r.VPAPIClient.SavepointsApi.GetSavepoint(ctx, namespace, savepointId)
	if err != nil {
		if utils.IsNotFoundError(err) {
			log.Info("Savepoint by id not found - creating", "id", )
			return r.handleCreate(req, vpSavepoint)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log = log.WithValues("savepoint", savepoint.Metadata.Id)
	log.Info("Update event")
	log.Info("Warn: Updates are not supported through the Ververica Platform.")
	return ctrl.Result{}, nil
}

func (r *VpSavepointReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VpSavepoint{}).
		Complete(r)
}
