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
	"errors"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"
	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
	appManager "github.com/fintechstudios/ververica-platform-k8s-controller/controllers/app-manager"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var InvalidDeploymentTargetNoTargetName = errors.New("must set spec.deploymentTargetName if spec.spec.deploymentTargetId is not specified")

// VpDeploymentReconciler reconciles a VpDeployment object
type VpDeploymentReconciler struct {
	client.Client
	Log                 logr.Logger
	AppManagerApiClient *appManagerApi.APIClient
	AppManagerAuthStore *appManager.AuthStore
}

// getLogger creates a logger for the controller with the request name
func (r *VpDeploymentReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpdeployment", req.NamespacedName)
}

// getDeploymentTargetID gets the id of a deployment
func (r *VpDeploymentReconciler) getDeploymentTargetID(vpDeployment v1beta1.VpDeployment, ctx context.Context) (string, error) {
	if len(vpDeployment.Spec.Spec.DeploymentTargetID) > 0 {
		// an id has been set, just return it
		return vpDeployment.Spec.Spec.DeploymentTargetID, nil
	}
	name := vpDeployment.Spec.DeploymentTargetName
	if len(name) == 0 {
		return "", InvalidDeploymentTargetNoTargetName
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	depTarget, _, err := r.AppManagerApiClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, nsName, vpDeployment.Spec.DeploymentTargetName)

	if err != nil {
		return "", err
	}

	return depTarget.Metadata.Id, nil
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpDeploymentReconciler) updateResource(resource *v1beta1.VpDeployment, deployment *appManagerApi.Deployment) error {
	ctx := context.Background()

	metadata, err := converters.DeploymentMetadataToNative(*deployment.Metadata)

	if err != nil {
		return err
	}
	resource.Spec.Metadata = metadata

	spec, err := converters.DeploymentSpecToNative(*deployment.Spec)
	if err != nil {
		return err
	}
	resource.Spec.Spec = spec

	state, err := converters.DeploymentStateToNative(deployment.Status.State)
	if err != nil {
		return err
	}
	resource.Status.State = state

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	if err := r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// handleCreate creates VP resources
func (r *VpDeploymentReconciler) handleCreate(req ctrl.Request, vpDeployment v1beta1.VpDeployment) (ctrl.Result, error) {
	log := r.getLogger(req)

	// See if there already exists a deployment by that name
	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	deployment, err := converters.DeploymentFromNative(vpDeployment)

	if err != nil {
		return ctrl.Result{}, err
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	deployment.Spec.DeploymentTargetId, err = r.getDeploymentTargetID(vpDeployment, ctx)
	if err != nil {
		return ctrl.Result{}, err
	}

	deployment.Metadata.Name = req.Name

	// create it
	createdDep, res, err := r.AppManagerApiClient.
		DeploymentsApi.
		CreateDeployment(ctx, namespace, deployment)

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		return ctrl.Result{Requeue: false}, err
	}

	if err != nil {
		log.Info("Error creating Deployment")
		return ctrl.Result{}, err
	}

	log.Info("Created deployment", "deployment", createdDep.Metadata.Id)

	// Now update the k8s resource and status as well
	if err := r.updateResource(&vpDeployment, &createdDep); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
// it also patches the deployment in the Ververica Platform, which could trigger a state transition
// which we should wait for, if possible
func (r *VpDeploymentReconciler) handleUpdate(req ctrl.Request, vpDeployment v1beta1.VpDeployment, currentDeployment appManagerApi.Deployment) (ctrl.Result, error) {
	log := r.getLogger(req)
	log.Info("Patching VP Deployment")

	desiredDeployment, err := converters.DeploymentFromNative(vpDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	// Patches with no changes to the spec should not trigger
	// sequential patches with the same spec will not trigger a new transition
	// but will bump the resource version, making a direct equality check impossible
	updatedDep, res, err := r.AppManagerApiClient.DeploymentsApi.UpdateDeployment(ctx, nsName, currentDeployment.Metadata.Id, desiredDeployment)

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		log.Error(err, "Error patching Deployment")
		return ctrl.Result{Requeue: false}, nil
	}

	if err != nil {
		return ctrl.Result{}, err
	}

	vpDeployment.Status.State, err = converters.DeploymentStateToNative(updatedDep.Status.State)

	if err != nil {
		return ctrl.Result{}, err
	}

	// Don't trigger a full update - should figure out how to truly make this idempotent
	// and still store all the fun stuff in K8s
	if err = r.Status().Update(context.Background(), &vpDeployment); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpDeploymentReconciler) handleDelete(req ctrl.Request, vpDeployment v1beta1.VpDeployment) (ctrl.Result, error) {
	log := r.getLogger(req)

	// First must make sure the deployment is canceled, then must delete it.
	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	var deployment appManagerApi.Deployment
	if len(vpDeployment.Spec.Metadata.ID) > 0 {
		deployment, _, err = r.AppManagerApiClient.DeploymentsApi.GetDeployment(ctx, nsName, vpDeployment.Spec.Metadata.ID)
	} else {
		deployment, err = appManager.GetDeploymentByName(r.AppManagerApiClient, ctx, nsName, vpDeployment.Name)
	}

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	// If the desired state is cancelled, we're good - just have to wait
	if deployment.Status.State != string(v1beta1.CancelledState) {
		// If the desired state is not cancelled, we're not good - must cancel and then wait
		if deployment.Spec.State != string(v1beta1.CancelledState) {
			// must cancel it
			log.Info("Cancelling Deployment")
			deployment.Spec.State = string(v1beta1.CancelledState)
			deployment, _, err = r.AppManagerApiClient.DeploymentsApi.UpdateDeployment(ctx, vpDeployment.Spec.Metadata.Namespace, vpDeployment.Spec.Metadata.ID, deployment)

			if err != nil {
				return ctrl.Result{}, utils.IgnoreNotFoundError(err)
			}
		}
		// Just have to wait now
		err = r.updateResource(&vpDeployment, &deployment)
		// Can take a while to tear down
		return ctrl.Result{RequeueAfter: time.Second * 30}, err
	}

	deployment, _, err = r.AppManagerApiClient.DeploymentsApi.DeleteDeployment(ctx, nsName, deployment.Metadata.Id)
	if err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting Deployment", "name", deployment.Metadata.Name)
	// Should happen instantaneously
	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments/status,verbs=get;update;patch

// Reconcile is the main reconciliation loop handler
func (r *VpDeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	var vpDeployment v1beta1.VpDeployment
	// If it's gone, it's gone!
	if err := r.Get(ctx, req.NamespacedName, &vpDeployment); err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	if vpDeployment.ObjectMeta.DeletionTimestamp.IsZero() {
		// Not being deleted, add the finalizer
		if utils.AddFinalizer(&vpDeployment.ObjectMeta) {
			log.Info("Adding Finalizer")
			if err := r.Update(ctx, &vpDeployment); err != nil {
				return ctrl.Result{}, err
			}

			if err := r.Get(ctx, req.NamespacedName, &vpDeployment); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		log.Info("Delete event", "name", req.Name)
		res, err := r.handleDelete(req, vpDeployment)
		if utils.IsRequeueResponse(res, err) {
			return res, err
		}
		// otherwise, we're all good, just remove the finalizer
		if utils.RemoveFinalizer(&vpDeployment.ObjectMeta) {
			if err := r.Update(ctx, &vpDeployment); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	appManagerCtx, err := r.AppManagerAuthStore.ContextForNamespace(nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue:false}, nil
	}
	id := vpDeployment.Spec.Metadata.ID
	if len(id) == 0 {
		// no id has been set
		deployment, err := appManager.GetDeploymentByName(r.AppManagerApiClient, appManagerCtx, nsName, req.Name)

		if utils.IsNotFoundError(err) {
			log.Info("Create event")
			return r.handleCreate(req, vpDeployment)
		}

		if err != nil {
			return ctrl.Result{}, err
		}

		log.Info("No id set for deployment", "deployment", deployment.Metadata.Name)
		// Update in k8s but don't patch - should be handled by the update loop
		err = r.updateResource(&vpDeployment, &deployment)
		return ctrl.Result{}, err
	}

	deployment, _, err := r.AppManagerApiClient.DeploymentsApi.GetDeployment(appManagerCtx, nsName, id)
	if err != nil {
		if utils.IsNotFoundError(err) {
			log.Info("Deployment by id not found - creating", "id", id)
			// Not found, means incorrect id is set but let's create it anyways
			return r.handleCreate(req, vpDeployment)
		}
		// Other error, not good!
		return ctrl.Result{}, err
	}

	log.Info("Update event")
	return r.handleUpdate(req, vpDeployment, deployment)
}

// SetupWithManager hooks the reconciler into the main manager
func (r *VpDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.VpDeployment{}).
		Complete(r)
}
