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
	"github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1/converters"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VpDeploymentReconciler reconciles a VpDeployment object
type VpDeploymentReconciler struct {
	client.Client
	Log         logr.Logger
	VPAPIClient vpAPI.APIClient
}

// getLogger creates a logger for the controller with the request name
func (r *VpDeploymentReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpdeployment", req.NamespacedName)
}

func (r *VpDeploymentReconciler) getDeploymentTargetID(resource ververicaplatformv1beta1.VpDeployment) (string, error) {
	if len(resource.Spec.Spec.DeploymentTargetID) > 0 {
		// an id has been set, just return it
		return resource.Spec.Spec.DeploymentTargetID, nil
	}
	name := resource.Spec.DeploymentTargetName
	if len(name) == 0 {
		return "", errors.New("must set spec.deploymentTargetName if spec.spec.deploymentTargetId is not specified")
	}

	ctx := context.Background()
	namespace := utils.GetNamespaceOrDefault(resource.Spec.Metadata.Namespace)
	depTarget, _, err := r.VPAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, namespace, resource.Spec.DeploymentTargetName)

	if err != nil {
		return "", err
	}

	return depTarget.Metadata.Id, nil
}

func (r *VpDeploymentReconciler) getDeploymentByName(ctx context.Context, namespace string, name string) (*vpAPI.Deployment, error) {
	if len(name) == 0 {
		return nil, errors.New("name must not be empty")
	}

	deploymentsList, _, err := r.VPAPIClient.DeploymentsApi.GetDeployments(ctx, namespace, nil)

	if err != nil {
		return nil, err
	}

	for _, deployment := range deploymentsList.Items {
		if deployment.Metadata.Name == name {
			return &deployment, nil
		}
	}

	// no errors but not found
	// TODO: consider making this into a not-found error, so we could pass back the struct and not a pointer
	return nil, nil
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpDeploymentReconciler) updateResource(req ctrl.Request, resource *ververicaplatformv1beta1.VpDeployment, deployment *vpAPI.Deployment) error {
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

	if err = r.Update(ctx, resource); err != nil {
		return err
	}

	if err = r.Status().Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// handleCreate creates VP resources
func (r *VpDeploymentReconciler) handleCreate(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// See if there already exists a deployment by that name
	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	dep, err := r.getDeploymentByName(ctx, namespace, vpDeployment.Name)

	if err != nil {
		log.Error(err, "while fetching deployments list")
		return ctrl.Result{}, err
	}

	if dep != nil && dep.Metadata.Name == vpDeployment.Name {
		return ctrl.Result{}, errors.New("deployment names must be unique per namespace")
	}

	deployment, err := converters.DeploymentFromNative(vpDeployment)

	if err != nil {
		return ctrl.Result{}, err
	}

	deployment.Spec.DeploymentTargetId, err = r.getDeploymentTargetID(vpDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	deployment.Metadata.Name = req.Name

	// create it
	createdDep, res, err := r.VPAPIClient.
		DeploymentsApi.
		CreateDeployment(ctx, namespace, deployment)

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		return ctrl.Result{Requeue: false}, err
	}

	if err != nil {
		log.Error(err, "Error creating VP Deployment")
		return ctrl.Result{}, err
	}

	log.Info("Created deployment", "deployment", createdDep)

	// Now update the k8s resource and status as well
	if err := r.updateResource(req, &vpDeployment, &createdDep); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
// it also patches the deployment in the Ververica Platform, which could trigger a state transition
// which we should wait for, if possible
func (r *VpDeploymentReconciler) handleUpdate(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment, currentDeployment vpAPI.Deployment) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)
	log.Info("Patching VP Deployment")

	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	desiredDeployment, err := converters.DeploymentFromNative(vpDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Patches with no changes to the spec should not trigger
	// another transition but will bump the resource version
	updatedDep, res, err := r.VPAPIClient.DeploymentsApi.UpdateDeployment(ctx, namespace, currentDeployment.Metadata.Id, desiredDeployment)

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		return ctrl.Result{Requeue: false}, err
	}

	if err != nil {
		log.Error(err, "Error patching VP Deployment")
		return ctrl.Result{}, err
	}
	vpDeployment.Status.State, err = converters.DeploymentStateToNative(updatedDep.Status.State)

	if err != nil {
		return ctrl.Result{}, err
	}

	// Don't trigger a full update - should figure out how to truly make this idempotent
	// and still store all the fun stuff in K8s
	if err = r.Status().Update(ctx, &vpDeployment); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpDeploymentReconciler) handleDelete(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// First must make sure the deployment is canceled, then must delete it

	// Let's make sure it's deleted from the ververica platform
	// must make sure the namespace and id are set, or will return a list of deployments...
	var namespace = utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)

	var (
		deployment vpAPI.Deployment
		err        error
	)
	if len(vpDeployment.Spec.Metadata.Id) > 0 {
		deployment, _, err = r.VPAPIClient.DeploymentsApi.GetDeployment(ctx, namespace, vpDeployment.Spec.Metadata.Id)
	} else {
		// TODO: this can definitely be cleaned up
		depPtr, err := r.getDeploymentByName(ctx, namespace, vpDeployment.Name)
		if err != nil {
			return ctrl.Result{}, err
		}
		if depPtr == nil {
			return ctrl.Result{}, nil
		}
		deployment = *depPtr
	}

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	// If the desired state is cancelled, we're good - just have to wait
	if deployment.Status.State != string(ververicaplatformv1beta1.CancelledState) {
		// If the desired state is not cancelled, we're not good - must cancel and then wait
		if deployment.Spec.State != string(ververicaplatformv1beta1.CancelledState) {
			// must cancel it
			log.Info("Cancelling Deployment")
			deployment.Spec.State = string(ververicaplatformv1beta1.CancelledState)
			deployment, _, err = r.VPAPIClient.DeploymentsApi.UpdateDeployment(ctx, vpDeployment.Spec.Metadata.Namespace, vpDeployment.Spec.Metadata.Id, deployment)

			if err != nil {
				return ctrl.Result{}, utils.IgnoreNotFoundError(err)
			}
		}
		// Just have to wait now
		err = r.updateResource(req, &vpDeployment, &deployment)
		// Can take a while to tear down
		return ctrl.Result{RequeueAfter: time.Second * 30}, err
	}

	deployment, _, err = r.VPAPIClient.DeploymentsApi.DeleteDeployment(ctx, namespace, deployment.Metadata.Id)
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

	var vpDeployment ververicaplatformv1beta1.VpDeployment
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
		// Being deleted
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

	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	id := vpDeployment.Spec.Metadata.Id
	if len(id) == 0 {
		deployment, err := r.getDeploymentByName(ctx, namespace, req.Name)

		if err != nil {
			return ctrl.Result{}, err
		}

		if deployment == nil {
			log.Info("Create event")
			return r.handleCreate(req, vpDeployment)
		}

		// no id, but hasn't
		log.Info("No id set for deployment", "deployment", *deployment)
		// Update in k8s but don't patch - should be handled by the update loop
		err = r.updateResource(req, &vpDeployment, deployment)
		return ctrl.Result{}, err
	}

	deployment, _, err := r.VPAPIClient.DeploymentsApi.GetDeployment(ctx, namespace, id)
	if err != nil {
		if utils.IsNotFoundError(err) {
			// Not found, let's create it
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
		For(&ververicaplatformv1beta1.VpDeployment{}).
		Complete(r)
}
