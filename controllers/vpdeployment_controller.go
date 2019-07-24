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
	"io/ioutil"

	"encoding/json"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	vpAPI "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
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

func (r *VpDeploymentReconciler) getDeploymentTargetId(resource *ververicaplatformv1beta1.VpDeployment) (string, error) {
	if len(resource.Spec.Spec.DeploymentTargetId) > 0 {
		// an id has been set, just return it
		return resource.Spec.Spec.DeploymentTargetId, nil
	}
	name := resource.Spec.Spec.DeploymentTargetName
	if len(name) == 0 {
		return "", errors.New("must set spec.spec.deploymentTargetName if spec.spec.deploymentTargetId is not specified")
	}

	ctx := context.Background()
	namespace := utils.GetNamespaceOrDefault(resource.Spec.Metadata.Namespace)
	depTarget, _, err := r.VPAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, namespace, resource.Spec.Spec.DeploymentTargetName)

	if err != nil {
		return "", err
	}

	return depTarget.Metadata.Id, nil
}

// updateResource takes a k8s resource and a VP resource and merges them
func (r *VpDeploymentReconciler) updateResource(req ctrl.Request, resource *ververicaplatformv1beta1.VpDeployment, deployment *vpAPI.Deployment) error {
	ctx := context.Background()

	resource.Name = deployment.Metadata.Name
	resource.Spec.Metadata = &ververicaplatformv1beta1.VpDeploymentMetadata{
		Name:            deployment.Metadata.Name,
		Namespace:       deployment.Metadata.Namespace,
		Id:              deployment.Metadata.Id,
		CreatedAt:       &metav1.Time{Time: deployment.Metadata.CreatedAt},
		ModifiedAt:      &metav1.Time{Time: deployment.Metadata.ModifiedAt},
		ResourceVersion: deployment.Metadata.ResourceVersion,
		Labels:          deployment.Metadata.Labels,
		Annotations:     deployment.Metadata.Annotations,
	}
	//
	//resource.Spec.Spec = ververicaplatformv1beta1.VpDeploymentTargetSpec{
	//	Kubernetes: ververicaplatformv1beta1.VpKubernetesTarget{
	//		Namespace: deployment.Spec.Kubernetes.Namespace,
	//	},
	//}

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	return nil
}

// handleCreate creates VP resources
func (r *VpDeploymentReconciler) handleCreate(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// TODO: must make this idempotent

	// TODO: redo this to take advantage of marshalling
	//vpJson, err := json.Marshal(vpDeployment)
	//log.Info(string(vpJson))
	//vpSpecJson, err := json.Marshal(vpDeployment.Spec)
	//log.Info(string(vpSpecJson))

	vpTemplate := vpDeployment.Spec.Spec.Template
	vpArtifact := vpTemplate.Spec.Artifact
	vpPods := vpTemplate.Spec.Kubernetes.Pods
	templateResources := make(map[string]vpAPI.ResourceSpec)
	for k, v := range vpTemplate.Spec.Resources {
		res := vpAPI.ResourceSpec{}
		if v.Memory != nil {
			res.Memory = *v.Memory
		}
		res.Cpu = float64(v.Cpu.MilliValue()) / 1000 // convert back to a plain float
		templateResources[k] = res
	}

	templatePods := vpAPI.Pods{
		Annotations:  vpPods.Annotations,
		NodeSelector: vpPods.NodeSelector,
	}

	templateSpec := vpAPI.DeploymentTemplateSpec{
		Artifact: &vpAPI.Artifact{
			Kind:               vpArtifact.Kind,
			JarUri:             vpArtifact.JarUri,
			MainArgs:           vpArtifact.MainArgs,
			EntryClass:         vpArtifact.EntryClass,
			FlinkVersion:       vpArtifact.FlinkVersion,
			FlinkImageRegistry: vpArtifact.FlinkImageRegistry,
			FlinkImageTag:      vpArtifact.FlinkImageTag,
		},
		Resources:          templateResources,
		FlinkConfiguration: vpTemplate.Spec.FlinkConfiguration,
		Logging: &vpAPI.Logging{
			Log4jLoggers: vpTemplate.Spec.Logging.Log4jLoggers,
		},
		Kubernetes: &vpAPI.KubernetesOptions{
			Pods: &templatePods,
		},
	}

	if vpTemplate.Spec.Parallelism != nil {
		templateSpec.Parallelism = *vpTemplate.Spec.Parallelism
	}

	if vpTemplate.Spec.NumberOfTaskManagers != nil {
		templateSpec.NumberOfTaskManagers = *vpTemplate.Spec.NumberOfTaskManagers
	}

	templateAnnotations := vpAPI.DeploymentTemplateMetadata{
		Annotations: vpDeployment.Spec.Spec.Template.Metadata.Annotations,
	}

	template := vpAPI.DeploymentTemplate{
		Metadata: &templateAnnotations,
		Spec:     &templateSpec,
	}

	depTargetId, err := r.getDeploymentTargetId(&vpDeployment)

	if err != nil {
		return ctrl.Result{}, err
	}

	depSpec := &vpAPI.DeploymentSpec{
		DeploymentTargetId: depTargetId,
		State:              vpDeployment.Spec.Spec.State,
		UpgradeStrategy: &vpAPI.DeploymentUpgradeStrategy{
			Kind: vpDeployment.Spec.Spec.UpgradeStrategy.Kind,
		},
		RestoreStrategy: &vpAPI.DeploymentRestoreStrategy{
			Kind:                  vpDeployment.Spec.Spec.RestoreStrategy.Kind,
			AllowNonRestoredState: vpDeployment.Spec.Spec.RestoreStrategy.AllowNonRestoredState,
		},
		Template: &template,
	}

	if vpDeployment.Spec.Spec.MaxSavepointCreationAttempts != nil {
		depSpec.MaxSavepointCreationAttempts = *vpDeployment.Spec.Spec.MaxSavepointCreationAttempts
	}

	if vpDeployment.Spec.Spec.MaxJobCreationAttempts != nil {
		depSpec.MaxJobCreationAttempts = *vpDeployment.Spec.Spec.MaxJobCreationAttempts
	}

	if vpDeployment.Spec.Spec.StartFromSavepoint != nil {
		depSpec.StartFromSavepoint = &vpAPI.DeploymentStartFromSavepoint{
			Kind: vpDeployment.Spec.Spec.StartFromSavepoint.Kind,
		}
	} else {
		depSpec.StartFromSavepoint = &vpAPI.DeploymentStartFromSavepoint{
			Kind: "NONE",
		}
	}

	dep := vpAPI.Deployment{
		ApiVersion: "v1",
		Metadata: &vpAPI.DeploymentMetadata{
			Name:        req.Name,
			Namespace:   vpDeployment.Spec.Metadata.Namespace,
			Labels:      vpDeployment.Spec.Metadata.Labels,
			Annotations: vpDeployment.Spec.Metadata.Annotations,
		},
		Spec: depSpec,
	}

	// create it
	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	res, err := r.VPAPIClient.
		DeploymentsApi.
		CreateDeployment(ctx, namespace, dep)

	if res != nil && res.StatusCode == 400 {
		// Bad Request, should not requeue
		return ctrl.Result{Requeue: false}, err
	}

	if err != nil {
		log.Error(err, "Error creating VP Deployment Target")
		return ctrl.Result{}, err
	}

	// TODO: the dep data is already in the res, but for some reason need to un-marshal it
	// 		 most likely a problem with the Swagger
	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	var createdDep vpAPI.Deployment
	if err := json.Unmarshal(body, &createdDep); err != nil {
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
// updates are not supported on Deployment Targets in the VP API, so just need to mirror the latest state
func (r *VpDeploymentReconciler) handleUpdate(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment, deployment vpAPI.Deployment) (ctrl.Result, error) {
	// Now update the k8s resource
	err := r.updateResource(req, &vpDeployment, &deployment)
	return ctrl.Result{}, err
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpDeploymentReconciler) handleDelete(req ctrl.Request, vpDeployment ververicaplatformv1beta1.VpDeployment) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.getLogger(req)

	// First must make sure the deployment is canceled, then must delete it

	// Let's make sure it's deleted from the ververica platform
	deployment, _, err := r.VPAPIClient.DeploymentTargetsApi.DeleteDeploymentTarget(ctx, vpDeployment.Spec.Metadata.Namespace, req.Name)

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting Deployment", "name", deployment.Metadata.Name)
	// Should happen instantaneously
	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments/status,verbs=get;update;patch

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
		if utils.AddFinalizerToObjectMeta(&vpDeployment.ObjectMeta) {
			log.Info("Adding Finalizer")
			if err := r.Update(ctx, &vpDeployment); err != nil {
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
		if utils.RemoveFinalizerFromObjectMeta(&vpDeployment.ObjectMeta) {
			if err := r.Update(ctx, &vpDeployment); err != nil {
				return ctrl.Result{}, err
			}
		}

		return res, nil
	}

	namespace := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	id := vpDeployment.Spec.Metadata.Id
	if len(id) == 0 {
		log.Info("Create event")
		// Creation, as no id has yet been set
		// TODO: should we check if one already exists by name in the VP and update if so?
		return r.handleCreate(req, vpDeployment)
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

func (r *VpDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ververicaplatformv1beta1.VpDeployment{}).
		Complete(r)
}
