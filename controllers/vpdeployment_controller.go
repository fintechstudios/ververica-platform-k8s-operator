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
	"strconv"
	"time"

	"github.com/antihax/optional"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1/converters"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/appmanager-api-client"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/annotations"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/appmanager"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/polling"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers/utils"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ErrorInvalidDeploymentTargetNoTargetName = errors.New("must set spec.deploymentTargetName if spec.spec.deploymentTargetId is not specified")

const statusPollingInterval = 30 * time.Second
const eventsLastPolledFormat = time.RFC3339Nano
const eventPollingInterval = 10 * time.Second
var lastEventTimestampAnnotation = annotations.NewAnnotationName("last-event-timestamp")

func eventAnnotations(event appmanagerapi.Event) map[string]string {
	return annotations.Create(
		annotations.Pair(annotations.ID, event.Metadata.Id),
		annotations.Pair(annotations.ResourceVersion, strconv.Itoa(int(event.Metadata.ResourceVersion))),
		annotations.Pair(annotations.Namespace, event.Metadata.Namespace),
		annotations.Pair(annotations.DeploymentID, event.Metadata.DeploymentId),
		annotations.Pair(annotations.JobID, event.Metadata.JobId))
}

// VpDeploymentReconciler reconciles a VpDeployment object
type VpDeploymentReconciler struct {
	client.Client
	Log                 logr.Logger
	AppManagerAPIClient *appmanagerapi.APIClient
	AppManagerAuthStore *appmanager.AuthStore
	pollerMap           map[string]*polling.Poller
	manager             *ctrl.Manager
}

// getLogger creates a logger for the controller with the request name
func (r *VpDeploymentReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpdeployment", req.NamespacedName)
}

// getDeploymentTargetID gets the id of a deployment
func (r *VpDeploymentReconciler) getDeploymentTargetID(ctx context.Context, vpDeployment v1beta1.VpDeployment) (string, error) {
	if annotations.Has(vpDeployment.Annotations, annotations.DeploymentTargetID) {
		return annotations.Get(vpDeployment.Annotations, annotations.DeploymentTargetID), nil
	}

	if len(vpDeployment.Spec.Spec.DeploymentTargetID) > 0 {
		// an id has been set, just return it
		return vpDeployment.Spec.Spec.DeploymentTargetID, nil
	}
	name := vpDeployment.Spec.DeploymentTargetName
	if len(name) == 0 {
		return "", ErrorInvalidDeploymentTargetNoTargetName
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	depTarget, _, err := r.AppManagerAPIClient.DeploymentTargetsApi.GetDeploymentTarget(ctx, nsName, vpDeployment.Spec.DeploymentTargetName)

	if err != nil {
		return "", err
	}

	return depTarget.Metadata.Id, nil
}

func (r *VpDeploymentReconciler) ensurePollersAreRunning(req ctrl.Request, vpDeployment *v1beta1.VpDeployment) {
	if !r.pollerIsRunning(req, "event") {
		r.addEventPollerForResource(req, vpDeployment)
	}
	if !r.pollerIsRunning(req, "status") {
		r.addStatusPollerForResource(req, vpDeployment)
	}
}

func (r *VpDeploymentReconciler) setPoller(req ctrl.Request, pollerType string, poller *polling.Poller) {
	if r.pollerMap == nil {
		r.pollerMap = make(map[string]*polling.Poller)
	}
	r.pollerMap[req.String()+"-"+pollerType] = poller
}

func (r *VpDeploymentReconciler) removePoller(req ctrl.Request, pollerType string) bool {
	if r.pollerMap == nil {
		return false
	}

	log := r.getLogger(req).WithValues("poller", pollerType)
	poller := r.pollerMap[req.String()+"-"+pollerType]
	if poller == nil {
		return false
	}
	log.Info("Stopping poller")
	poller.StopAndBlock()
	delete(r.pollerMap, req.String()+"-"+pollerType)
	return true
}

func (r *VpDeploymentReconciler) pollerIsRunning(req ctrl.Request, pollerType string) bool {
	if r.pollerMap[req.String()+"-"+pollerType] == nil {
		return false
	}

	return !r.pollerMap[req.String()+"-"+pollerType].IsStopped()
}

func (r *VpDeploymentReconciler) removePollers(req ctrl.Request) {
	r.removePoller(req, "status")
	r.removePoller(req, "event")
}

func (r *VpDeploymentReconciler) getEventPollerFunc(req ctrl.Request, namespace, id string) polling.PollerFunc {
	log := r.getLogger(req).WithValues("poller", "event")

	return func() interface{} {
		log.Info("Polling")
		ctx, _ := r.AppManagerAuthStore.ContextForNamespace(context.Background(), namespace)

		events, _, err := r.AppManagerAPIClient.EventsApi.GetEvents(ctx, namespace, &appmanagerapi.GetEventsOpts{
			DeploymentId: optional.NewInterface(id),
			JobId:        optional.EmptyInterface(),
		})

		if err != nil {
			log.Error(err, "Error while polling events")
			return nil
		}

		var vpDeployment v1beta1.VpDeployment
		if err := r.Get(context.Background(), req.NamespacedName, &vpDeployment); err != nil {
			log.Error(err, "Error while getting latest k8s object")
			return nil
		}

		// Since the VVP API doesn't support polling from a specific point-in-time,
		// record the last event time on the k8s obj

		var lastPolledTime *time.Time
		if annotations.Has(vpDeployment.Annotations, lastEventTimestampAnnotation) {
			timeStr := annotations.Get(vpDeployment.Annotations, lastEventTimestampAnnotation)
			var t time.Time
			if t, err = time.Parse(eventsLastPolledFormat, timeStr); err != nil {
				log.WithValues("timestamp", timeStr).Error(err, "Error parsing annotation time")
				annotations.Remove(vpDeployment.Annotations, lastEventTimestampAnnotation)
				// update the k8s object
				if err = r.Update(context.Background(), &vpDeployment); err != nil {
					log.Error(err, "Unable to update deployment")
				}
				return nil
			}
			lastPolledTime = &t
		}

		var maxTime *time.Time
		for _, event := range events.Items {
			eventTime := event.Metadata.CreatedAt
			// filter out all created events before the last time polled, or where the event time is unset
			if eventTime.IsZero() ||
				(lastPolledTime != nil &&
					(lastPolledTime.Equal(eventTime) || lastPolledTime.After(eventTime))) {
				continue
			}

			// record the latest event to have occurred
			if maxTime == nil || maxTime.Before(eventTime) {
				maxTime = &eventTime
			}

			recorder := (*r.manager).GetEventRecorderFor("ververica-platform-k8s-operator")
			recorder.AnnotatedEventf(&vpDeployment,
				eventAnnotations(event),
				"Normal",
				event.Metadata.Name,
				event.Spec.Message)
		}

		// update if there is a new max time and
		if maxTime != nil && (lastPolledTime == nil || !maxTime.Equal(*lastPolledTime)) {
			timeStr := maxTime.Format(eventsLastPolledFormat)
			log.WithValues(
				"last", annotations.Get(vpDeployment.Annotations, lastEventTimestampAnnotation),
				"latest", timeStr).
				Info("Updating last event time polled")
			annotations.Set(vpDeployment.Annotations,
				annotations.Pair(lastEventTimestampAnnotation, timeStr))

			// update the k8s object
			if err = r.Update(context.Background(), &vpDeployment); err != nil {
				log.Error(err, "Unable to update deployment")
				return nil
			}
		}

		return nil
	}
}

func (r *VpDeploymentReconciler) addEventPollerForResource(req ctrl.Request, vpDeployment *v1beta1.VpDeployment) {
	log := r.getLogger(req).WithValues("poller", "event")
	if r.pollerIsRunning(req, "event") {
		log.Info("A status poller already exists, removing...")
		r.removePoller(req, "event")
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	vpID := annotations.Get(vpDeployment.Annotations, annotations.ID)

	// On each polling callback, push the update through the k8s client
	poller := polling.NewPoller(r.getEventPollerFunc(req, nsName, vpID), eventPollingInterval)

	r.setPoller(req, "event", poller)
	poller.Start()
}

func (r *VpDeploymentReconciler) getStatusPollerFunc(req ctrl.Request, namespace, id string) polling.PollerFunc {
	log := r.getLogger(req).WithValues("poller", "status")
	return func() interface{} {
		log.Info("Polling")
		ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), namespace)

		if err != nil {
			log.Error(err, "Error getting authorized context")
			return nil
		}

		deployment, _, err := r.AppManagerAPIClient.DeploymentsApi.GetDeployment(ctx, namespace, id)
		if err != nil {
			log.Error(err, "Error while polling deployment")
			return nil
		}

		var vpDeploymentUpdated v1beta1.VpDeployment
		if err = r.Get(ctx, req.NamespacedName, &vpDeploymentUpdated); err != nil {
			if utils.IsNotFoundError(err) {
				log.Error(err, "VpDeployment not found while polling")
			} else {
				log.Error(err, "Error getting VpDeployment while polling")
			}
			return deployment
		}

		if err = r.updateResource(&vpDeploymentUpdated, &deployment); err != nil {
			log.Error(err, "Error while updating VpSavepoint from poller")
		}

		return nil
	}
}

func (r *VpDeploymentReconciler) addStatusPollerForResource(req ctrl.Request, vpDeployment *v1beta1.VpDeployment) {
	log := r.getLogger(req)
	if r.pollerIsRunning(req, "status") {
		log.Info("A status poller already exists, removing...")
		r.removePoller(req, "status")
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	vpID := annotations.Get(vpDeployment.Annotations, annotations.ID)
	poller := polling.NewPoller(r.getStatusPollerFunc(req, nsName, vpID), statusPollingInterval)

	r.setPoller(req, "status", poller)
	poller.Start()
}

// updateResource takes a k8s resource and a VP resource and syncs them in k8s - does a full update
func (r *VpDeploymentReconciler) updateResource(resource *v1beta1.VpDeployment, deployment *appmanagerapi.Deployment) error {
	ctx := context.Background()

	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}
	// save dynamic information as annotations
	annotations.Set(resource.Annotations,
		annotations.Pair(annotations.ID, deployment.Metadata.Id),
		annotations.Pair(annotations.ResourceVersion, strconv.Itoa(int(deployment.Metadata.ResourceVersion))),
		annotations.Pair(annotations.DeploymentTargetID, deployment.Spec.DeploymentTargetId))

	if err := r.Update(ctx, resource); err != nil {
		return err
	}

	state, err := converters.DeploymentStateToNative(deployment.Status.State)
	if err != nil {
		return err
	}
	resource.Status.State = state

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
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	deployment.Spec.DeploymentTargetId, err = r.getDeploymentTargetID(ctx, vpDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	deployment.Metadata.Name = req.Name

	// create it
	createdDep, res, err := r.AppManagerAPIClient.
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

	// Create a poller to keep the savepoint up to date
	r.ensurePollersAreRunning(req, &vpDeployment)

	return ctrl.Result{}, nil
}

// handleUpdate updates the k8s resource when it already exists in the VP
// it also patches the deployment in the Ververica Platform, which could trigger a state transition
// which we should wait for, if possible
func (r *VpDeploymentReconciler) handleUpdate(req ctrl.Request, vpDeployment v1beta1.VpDeployment, currentDeployment appmanagerapi.Deployment) (ctrl.Result, error) {
	log := r.getLogger(req)
	log.Info("Patching VP Deployment")

	desiredDeployment, err := converters.DeploymentFromNative(vpDeployment)
	if err != nil {
		return ctrl.Result{}, err
	}

	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	// Patches with no changes to the spec should not trigger
	// sequential patches with the same spec will not trigger a new transition
	// but will bump the resource version, making a direct equality check impossible
	updatedDep, res, err := r.AppManagerAPIClient.DeploymentsApi.UpdateDeployment(ctx, nsName, currentDeployment.Metadata.Id, desiredDeployment)

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

	r.ensurePollersAreRunning(req, &vpDeployment)

	return ctrl.Result{}, nil
}

// handleDelete will ensure that the Ververica Platform namespace is also cleaned up
func (r *VpDeploymentReconciler) handleDelete(req ctrl.Request, vpDeployment v1beta1.VpDeployment) (ctrl.Result, error) {
	log := r.getLogger(req)

	// First must make sure the deployment is canceled, then must delete it.
	nsName := utils.GetNamespaceOrDefault(vpDeployment.Spec.Metadata.Namespace)
	ctx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	var deployment appmanagerapi.Deployment
	if annotations.Has(vpDeployment.ObjectMeta.Annotations, annotations.ID) {
		id := annotations.Get(vpDeployment.ObjectMeta.Annotations, annotations.ID)
		deployment, _, err = r.AppManagerAPIClient.DeploymentsApi.GetDeployment(ctx, nsName, id)
	} else {
		deployment, err = appmanager.GetDeploymentByName(ctx, r.AppManagerAPIClient, nsName, vpDeployment.Name)
	}

	if err != nil {
		// If it's already gone, great!
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	// If the desired state is cancelled, we're good - just have to wait
	if deployment.Status.State != string(v1beta1.CancelledState) {
		// If the desired state is not cancelled, we're ~not~ good - must cancel and then wait
		if deployment.Spec.State != string(v1beta1.CancelledState) {
			// must cancel it
			log.Info("Cancelling Deployment")
			deployment.Spec.State = string(v1beta1.CancelledState)
			deployment, _, err = r.AppManagerAPIClient.DeploymentsApi.UpdateDeployment(ctx, vpDeployment.Spec.Metadata.Namespace, deployment.Metadata.Id, deployment)

			if err != nil {
				return ctrl.Result{}, utils.IgnoreNotFoundError(err)
			}
		}
		// Just have to wait now
		err = r.updateResource(&vpDeployment, &deployment)
		if err != nil {
			return ctrl.Result{}, err
		}
		// Can take a while to tear down
		log.Info("Requeue-ing after 30 seconds")
		return ctrl.Result{RequeueAfter: time.Second * 30}, nil
	}

	deployment, _, err = r.AppManagerAPIClient.DeploymentsApi.DeleteDeployment(ctx, nsName, deployment.Metadata.Id)
	if err != nil {
		return ctrl.Result{}, utils.IgnoreNotFoundError(err)
	}

	log.Info("Deleting Deployment", "name", deployment.Metadata.Name)
	// Should happen instantaneously
	r.removePollers(req)
	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events/status,verbs=get

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
		r.removePoller(req, "status")
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
	appManagerCtx, err := r.AppManagerAuthStore.ContextForNamespace(context.Background(), nsName)
	if err != nil {
		log.Error(err, "cannot create context")
		return ctrl.Result{Requeue: false}, nil
	}

	if !annotations.Has(vpDeployment.ObjectMeta.Annotations, annotations.ID) {
		// no id has been set
		deployment, err := appmanager.GetDeploymentByName(appManagerCtx, r.AppManagerAPIClient, nsName, req.Name)

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

	id := annotations.Get(vpDeployment.ObjectMeta.Annotations, annotations.ID)

	deployment, _, err := r.AppManagerAPIClient.DeploymentsApi.GetDeployment(appManagerCtx, nsName, id)
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
	r.manager = &mgr
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta1.VpDeployment{}).
		Complete(r)
}
