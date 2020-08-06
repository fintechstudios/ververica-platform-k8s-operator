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
	"fmt"
	"sort"
	"time"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/scheduling"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/reference"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	"github.com/robfig/cron/v3"
)

// based on: https://kubebuilder.io/cronjob-tutorial/controller-implementation.html

// VpCronDeploymentReconciler reconciles a VpCronDeployment object
type VpCronDeploymentReconciler struct {
	client.Client
	scheduling.Clock // embed a clock that can be swapped out for testing
	Log              logr.Logger
	Scheme           *runtime.Scheme
}

var (
	scheduledTimeAnnotation = annotations.NewNamespacedAnnotationName("vpcrondeployment", "scheduled-at")
	apiGroupVersionStr      = v1beta2.GroupVersion.String()
)

const (
	maxMissedStartTimes = 100
	scheduledTimeFormat = time.RFC3339
	deploymentOwnerKey  = ".metadata.controller"
)

// isVpDeploymentFinished determines whether the vpdeployment has reached a terminal state
func isVpDeploymentFinished(vpdeployment *v1beta2.VpDeployment) (bool, v1beta2.VpDeploymentState) {
	switch vpdeployment.Status.State {
	// done if something/ someone else has stopped it
	case v1beta2.SuspendedState:
		fallthrough
	case v1beta2.CancelledState:
		fallthrough
	// done if it has failed
	case v1beta2.FailedState:
		fallthrough
	// finished successfully!
	case v1beta2.FinishedState:
		return true, vpdeployment.Status.State
	default:
		return false, ""
	}
}

// getScheduledTimeForDeployment parses the annotation storing when the next scheduled
// time for this deployment should be, or nil if it has not yet been set
func getScheduledTimeForDeployment(vpdeployment *v1beta2.VpDeployment) (*time.Time, error) {
	if !annotations.Has(vpdeployment.Annotations, scheduledTimeAnnotation) {
		return nil, nil
	}

	raw := annotations.Get(vpdeployment.Annotations, scheduledTimeAnnotation)
	parsed, err := time.Parse(scheduledTimeFormat, raw)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// fifoSortDeploymentSlice sorts a slice of VpDeployments based on their start time
// into a FIFO list, newest to oldest
func fifoSortDeploymentSlice(vpdeployments []*v1beta2.VpDeployment) {
	sort.Slice(vpdeployments, func(i, j int) bool {
		if vpdeployments[i].Status.StartTime != nil {
			return vpdeployments[j].Status.StartTime == nil
		}
		return vpdeployments[i].Status.StartTime.Before(vpdeployments[j].Status.StartTime)
	})
}

// getNextSchedule calculates when the next scheduled deployment run should be
// and if there is one that has yet to be processed
func getNextSchedule(cronDep *v1beta2.VpCronDeployment, now time.Time) (lastMissed time.Time, next time.Time, err error) {
	sched, err := cron.ParseStandard(cronDep.Spec.Schedule)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	// for optimization purposes, cheat a bit and start from our last observed run time
	var earliestTime time.Time
	if cronDep.Status.LastScheduleTime != nil {
		earliestTime = cronDep.Status.LastScheduleTime.Time
	} else {
		earliestTime = cronDep.ObjectMeta.CreationTimestamp.Time
	}
	// see if there's a deadline, and if the earliest time needs to be moved up
	if cronDep.Spec.StartingDeadlineSeconds != nil {
		deadline := now.Add(-time.Second * time.Duration(*cronDep.Spec.StartingDeadlineSeconds))
		if deadline.After(earliestTime) {
			earliestTime = deadline
		}
	}
	// if we're passed due, schedule it for now
	if earliestTime.After(now) {
		return time.Time{}, sched.Next(now), nil
	}

	numStarts := 0
	for t := sched.Next(earliestTime); !t.After(now); t = sched.Next(t) {
		lastMissed = t

		// An object might miss several starts. For example, if the
		// controller gets wedged on Friday at 5:01pm when everyone has
		// gone home, and someone comes in on Tuesday AM and discovers
		// the problem and restarts the controller, then all the hourly
		// jobs, more than 80 of them for one hourly scheduledJob, should
		// all start running with no further intervention (if the scheduledJob
		// allows concurrency and late starts).
		//
		// However, if there is a bug somewhere, or incorrect clock
		// on controller's server or apiservers (for setting creationTimestamp)
		// then there could be so many missed start times (it could be off
		// by decades or more), that it would eat up all the CPU and memory
		// of this controller. In that case, we want to not try to list
		// all the missed start times.
		numStarts++
		if numStarts > maxMissedStartTimes {
			return time.Time{}, time.Time{}, fmt.Errorf("too many missed start times (> %d). Set or decrease .spec.startingDeadlineSeconds or check clock skew", maxMissedStartTimes)
		}
	}
	return lastMissed, sched.Next(now), nil
}

// splitDeploymentList takes a list of deployments and categorizes them into sublists of
// active, successful(ly finished), and failed. It also calculates the most recently scheduled time.
func splitDeploymentList(ctx context.Context, childDeployments *v1beta2.VpDeploymentList) (activeDeps, successfulDeps, failedDeps []*v1beta2.VpDeployment, mostRecentTime *time.Time) {
	logger := log.FromContext(ctx)
	for _, dep := range childDeployments.Items {
		finished, finishType := isVpDeploymentFinished(&dep)
		if !finished {
			activeDeps = append(activeDeps, &dep)
		} else {
			switch finishType {
			case v1beta2.FailedState:
				failedDeps = append(failedDeps, &dep)
			case v1beta2.FinishedState:
				successfulDeps = append(successfulDeps, &dep)
			}
		}

		scheduledTime, err := getScheduledTimeForDeployment(&dep)
		if err != nil {
			logger.Error(err, "unable to parse scheduled time for deployment", "deployment", &dep)
			continue
		}
		if scheduledTime != nil {
			if mostRecentTime == nil {
				mostRecentTime = scheduledTime
			} else if mostRecentTime.Before(*scheduledTime) {
				mostRecentTime = scheduledTime
			}
		}
	}

	return
}

// getDeploymentsToDeleteByAge calculates the sublist of deployments to delete, keeping the youngest, given a max limit to keep.
func getDeploymentsToDeleteByAge(deployments []*v1beta2.VpDeployment, limit int32) []*v1beta2.VpDeployment {
	var toDelete []*v1beta2.VpDeployment
	// fifo queue, delete oldest first
	fifoSortDeploymentSlice(deployments)

	maxIndexToKeep := int32(len(deployments)) - limit
	for i, dep := range deployments {
		if int32(i) >= maxIndexToKeep {
			break
		}
		toDelete = append(toDelete, dep)
	}

	return toDelete
}

// buildVpDeploymentForCronDep creates the fully-formed deployment from the template, including the labels and
// the controller reference
func (r *VpCronDeploymentReconciler) buildVpDeploymentForCronDep(cronDep *v1beta2.VpCronDeployment, scheduledTime time.Time) (*v1beta2.VpDeployment, error) {
	// We want job names for a given nominal start time to have a deterministic name
	// to avoid the same job being created twice
	name := fmt.Sprintf("%s-%d", cronDep.Name, scheduledTime.Unix())

	// copy all annotations and labels
	annotationSet := annotations.Set(
		annotations.EnsureExist(annotations.Copy(cronDep.Spec.VpDeploymentTemplate.Metadata.Annotations)),
		annotations.Pair(scheduledTimeAnnotation, scheduledTime.Format(scheduledTimeFormat)))

	labelSet := make(map[string]string)
	if cronDep.Spec.VpDeploymentTemplate.Metadata.Labels != nil {
		for k, v := range cronDep.Spec.VpDeploymentTemplate.Metadata.Labels {
			labelSet[k] = v
		}
	}

	dep := &v1beta2.VpDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      labelSet,
			Annotations: annotationSet,
			Name:        name,
			Namespace:   cronDep.Namespace,
		},
		Spec: *cronDep.Spec.VpDeploymentTemplate.DeepCopy(),
	}

	if err := ctrl.SetControllerReference(cronDep, dep, r.Scheme); err != nil {
		return nil, err
	}

	return dep, nil
}

// getLogger creates a logger for the controller with the request name
func (r *VpCronDeploymentReconciler) getLogger(req ctrl.Request) logr.Logger {
	return r.Log.WithValues("vpcrondeployment", req.NamespacedName)
}

// runScheduling
func (r *VpCronDeploymentReconciler) runScheduling(ctx context.Context, vpCronDep *v1beta2.VpCronDeployment, childDeployments *v1beta2.VpDeploymentList) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	activeDeps, successfulDeps, failedDeps, mostRecentTime := splitDeploymentList(ctx, childDeployments)

	if mostRecentTime != nil {
		vpCronDep.Status.LastScheduleTime = &metav1.Time{Time: *mostRecentTime}
	} else {
		vpCronDep.Status.LastScheduleTime = nil
	}

	// reset which deployments are active
	vpCronDep.Status.Active = nil
	for _, activeDep := range activeDeps {
		depRef, err := reference.GetReference(r.Scheme, activeDep)
		if err != nil {
			logger.Error(err, "unable to make reference to active deployment", "deployment", activeDep)
			continue
		}
		vpCronDep.Status.Active = append(vpCronDep.Status.Active, *depRef)
	}
	logger.V(1).Info(
		"deployment count",
		"active", len(activeDeps),
		"failed", len(failedDeps),
		"successful", len(successfulDeps))

	// update the crondep
	if err := r.Status().Update(ctx, vpCronDep); err != nil {
		logger.Error(err, "unable to update VpCronDeployment status")
		return ctrl.Result{}, err
	}

	// clean up old jobs

	// delete in the background
	deletePropPolicy := client.PropagationPolicy(metav1.DeletePropagationBackground)

	if vpCronDep.Spec.FailedDeploymentsHistoryLimit != nil {
		depsToDelete := getDeploymentsToDeleteByAge(failedDeps, *vpCronDep.Spec.FailedDeploymentsHistoryLimit)
		for _, dep := range depsToDelete {
			if err := r.Delete(ctx, dep, deletePropPolicy); err != nil {
				logger.Error(err, "unable to delete old failed deployment", "deployment", dep)
			} else {
				logger.V(0).Info("deleted old failed deployment", "deployment", dep)
			}
		}
	}

	if vpCronDep.Spec.SuccessfulDeploymentsHistoryLimit != nil {
		depsToDelete := getDeploymentsToDeleteByAge(successfulDeps, *vpCronDep.Spec.SuccessfulDeploymentsHistoryLimit)
		for _, dep := range depsToDelete {
			if err := r.Delete(ctx, dep, deletePropPolicy); err != nil {
				logger.Error(err, "unable to delete old failed deployment", "deployment", dep)
			} else {
				logger.V(0).Info("deleted old failed deployment", "deployment", dep)
			}
		}
	}

	// Scheduling phase

	// check if this cron deployment is currently suspended
	if vpCronDep.Spec.Suspend != nil && *vpCronDep.Spec.Suspend {
		logger.V(1).Info("skipping suspended", "crondeployment", vpCronDep)
		return ctrl.Result{}, nil
	}

	// get the next scheduled run
	missedRunTime, nextRunTime, err := getNextSchedule(vpCronDep, r.Now())
	if err != nil {
		logger.Error(err, "unable to determine schedule", "crondeployment", vpCronDep)
		// non-recoverable error, needs an external update before it can be processed
		return ctrl.Result{}, nil
	}
	// let the k8s schedule requeue this until the next run
	scheduledResult := ctrl.Result{RequeueAfter: nextRunTime.Sub(r.Now())}
	logger = logger.WithValues("now", r.Now(), "next run", nextRunTime)

	// if there's a deployment scheduled and within the deadline, deploy it!
	if missedRunTime.IsZero() {
		logger.V(1).Info("no upcoming scheduled times, sleeping until next", "next", nextRunTime)
		return scheduledResult, nil
	}
	// make sure we're not too late to start the run
	logger = logger.WithValues("current run", missedRunTime)
	isTooLate := false
	if vpCronDep.Spec.StartingDeadlineSeconds != nil {
		deadlineDur := time.Second * time.Duration(*vpCronDep.Spec.StartingDeadlineSeconds)
		isTooLate = missedRunTime.Add(deadlineDur).Before(r.Now())
	}
	if isTooLate {
		logger.V(1).Info("missed starting deadline for last run, sleeping until next")
		return scheduledResult, nil
	}
	// handle the concurrency policy
	if vpCronDep.Spec.ConcurrencyPolicy == v1beta2.ForbidConcurrent && len(activeDeps) > 0 {
		logger.V(1).Info("concurrency policy blocks concurrent runs, skipping", "num active", len(activeDeps))
		return scheduledResult, nil
	}

	// delete all current jobs and then replace
	if vpCronDep.Spec.ConcurrencyPolicy == v1beta2.ReplaceConcurrent {
		for _, activeDep := range activeDeps {
			if err := r.Delete(ctx, activeDep, deletePropPolicy); client.IgnoreNotFound(err) != nil {
				logger.Error(err, "unable to delete active deployment", "deployment", activeDep)
				return ctrl.Result{}, err
			}
		}
	}

	// create the new deployment
	dep, err := r.buildVpDeploymentForCronDep(vpCronDep, missedRunTime)
	if err != nil {
		return ctrl.Result{}, err
	}
	if err := r.Create(ctx, dep); err != nil {
		logger.Error(err, "unable to create VpDeployment for VpCronDeployment", "deployment", dep)
	}
	logger.V(1).Info("created VpDeployment for VpCronDeployment", "deployment", dep)

	return scheduledResult, nil
}

// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpcrondeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpcrondeployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ververicaplatform.fintechstudios.com,resources=vpdeployments,verbs=get

func (r *VpCronDeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	logger := r.getLogger(req)
	ctx := log.IntoContext(context.Background(), logger)

	var vpCronDep v1beta2.VpCronDeployment
	// If it's gone, it's gone!
	if err := r.Get(ctx, req.NamespacedName, &vpCronDep); err != nil {
		logger.Error(err, "unable to fetch VpCronDeployment")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// list all active deployments
	var childDeployments v1beta2.VpDeploymentList
	err := r.List(
		ctx,
		&childDeployments,
		client.InNamespace(req.Namespace),
		client.MatchingFields{deploymentOwnerKey: req.Name})

	if err != nil {
		logger.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}

	return r.runScheduling(ctx, &vpCronDep, &childDeployments)
}

func (r *VpCronDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// if not given a Clock implementation, use the real one
	if r.Clock == nil {
		r.Clock = scheduling.NewClock()
	}

	// setup indexing for the owner field
	if err := mgr.GetFieldIndexer().
		IndexField(
			context.Background(),
			&v1beta2.VpDeployment{},
			deploymentOwnerKey,
			func(rawObj runtime.Object) []string {
				dep := rawObj.(*v1beta2.VpDeployment)
				owner := metav1.GetControllerOf(dep)
				if owner == nil {
					return nil
				}
				// make sure this is the controller
				if owner.APIVersion != apiGroupVersionStr || owner.Kind != "VpCronDeployment" {
					return nil
				}
				// single owner, which is the VpCronDeployment
				return []string{owner.Name}
			}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1beta2.VpCronDeployment{}).
		Owns(&v1beta2.VpDeployment{}).
		Complete(r)
}
