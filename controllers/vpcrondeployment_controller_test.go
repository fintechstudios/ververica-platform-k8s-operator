package controllers

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/annotations"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/scheduling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

var _ = Describe("VpCronDeployment Controller", func() {
	var clockTime *time.Time
	var clock scheduling.Clock
	var reconciler VpCronDeploymentReconciler

	BeforeEach(func() {
		clockTime = timeMustParse(time.RFC3339, "2019-12-09T14:27:58.328Z")
		clock = scheduling.NewFixedClock(clockTime)

		reconciler = VpCronDeploymentReconciler{
			Client: k8sClient,
			Log:    logger,
			Clock:  clock,
			Scheme: scheme.Scheme,
		}
	})

	Describe("isVpDeploymentFinished", func() {
		type testCase struct {
			state    v1beta2.VpDeploymentState
			expected bool
		}

		DescribeTable("should understand terminal states",
			func(given testCase) {
				isFinished, state := isVpDeploymentFinished(&v1beta2.VpDeployment{
					Status: &v1beta2.VpDeploymentStatus{
						State: given.state,
					},
				})
				Expect(isFinished).To(Equal(given.expected))
				if isFinished {
					Expect(state).To(Equal(given.state))
				} else {
					Expect(state).To(BeEmpty())
				}
			},
			Entry("empty", testCase{
				expected: false,
				state:    "",
			}),
			Entry("cancelled", testCase{
				expected: true,
				state:    v1beta2.CancelledState,
			}),
			Entry("suspended", testCase{
				expected: true,
				state:    v1beta2.SuspendedState,
			}),
			Entry("finished", testCase{
				expected: true,
				state:    v1beta2.FinishedState,
			}),
			Entry("failed", testCase{
				expected: true,
				state:    v1beta2.FailedState,
			}),
			Entry("running", testCase{
				expected: false,
				state:    v1beta2.RunningState,
			}),
			Entry("transitioning", testCase{
				expected: false,
				state:    v1beta2.TransitioningState,
			}))
	})

	Describe("splitDeploymentList", func() {
		var makeVpDep = func(name string, state v1beta2.VpDeploymentState) v1beta2.VpDeployment {
			return v1beta2.VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: "default",
				},
				Status: &v1beta2.VpDeploymentStatus{
					State: state,
				},
			}
		}

		It("should split an empty slice", func() {
			deps := v1beta2.VpDeploymentList{
				Items: make([]v1beta2.VpDeployment, 0),
			}
			active, successful, failed := splitDeploymentList(&deps)
			Expect(active).To(HaveLen(0))
			Expect(successful).To(HaveLen(0))
			Expect(failed).To(HaveLen(0))
		})

		It("should split all states", func() {
			deps := v1beta2.VpDeploymentList{
				Items: []v1beta2.VpDeployment{
					makeVpDep("failed-1", v1beta2.FailedState),
					makeVpDep("failed-2", v1beta2.FailedState),
					makeVpDep("active-1", v1beta2.TransitioningState),
					makeVpDep("successful-1", v1beta2.FinishedState),
					makeVpDep("active-2", v1beta2.RunningState),
					makeVpDep("failed-3", v1beta2.CancelledState),
					makeVpDep("successful-2", v1beta2.FinishedState),
					makeVpDep("failed-4", v1beta2.SuspendedState),
				},
			}
			active, successful, failed := splitDeploymentList(&deps)
			Expect(active).To(HaveLen(2))
			Expect(successful).To(HaveLen(2))
			Expect(failed).To(HaveLen(4))
			for _, d := range active {
				Expect(d.Name).To(ContainSubstring("active"))
			}
			for _, d := range successful {
				Expect(d.Name).To(ContainSubstring("successful"))
			}
			for _, d := range failed {
				Expect(d.Name).To(ContainSubstring("failed"))
			}
		})
	})

	Describe("getNextCronSchedule", func() {
		var makeVpCronDep = func(
			name string,
			cronSchedule string,
			startingDeadlineSeconds *int64,
			creationTime metav1.Time,
			lastScheduledTime *metav1.Time,
		) *v1beta2.VpCronDeployment {
			return &v1beta2.VpCronDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:              name,
					Namespace:         "default",
					CreationTimestamp: creationTime,
				},
				Spec: v1beta2.VpCronDeploymentSpec{
					Schedule:                cronSchedule,
					StartingDeadlineSeconds: startingDeadlineSeconds,
				},
				Status: v1beta2.VpCronDeploymentStatus{
					LastScheduleTime: lastScheduledTime,
				},
			}
		}

		It("should report errors for invalid schedules", func() {
			cronDep := makeVpCronDep(
				"test",
				"not-a-schedule",
				pointer.Int64Ptr(10),
				metav1.Time{},
				nil,
			)
			_, _, err := getNextCronSchedule(cronDep, clock.Now())
			Expect(err).To(HaveOccurred())
		})

		It("should schedule the next time from now if already passed without scheduling, with no misses", func() {
			now := *timeMustParse(time.RFC3339, "2019-12-09T14:27:58.328Z")
			cronDep := makeVpCronDep(
				"test",
				"@daily",
				nil,
				metav1.Time{Time: now},
				&metav1.Time{Time: now.Add(time.Minute)},
			)
			missedT, nextT, err := getNextCronSchedule(cronDep, now)
			Expect(err).NotTo(HaveOccurred())
			Expect(missedT).To(BeZero())
			// should be the next day at midnight after "now"
			expected := *timeMustParse(time.RFC3339, "2019-12-10T00:00:00.000Z")
			Expect(nextT.Equal(expected)).To(BeTrue())
		})

		It("should schedule the next time from earliest schedule, with the last miss", func() {
			now := *timeMustParse(time.RFC3339, "2019-12-09T14:10:30.000Z")
			cronDep := makeVpCronDep(
				"test",
				"* * * * *",             // every minute always!
				pointer.Int64Ptr(10*60), // 10 min deadline
				metav1.Time{Time: now.Add(-11 * time.Minute)}, // created 11 minutes ago
				nil,
			)
			missedT, nextT, err := getNextCronSchedule(cronDep, now)
			Expect(err).NotTo(HaveOccurred())
			expectedMiss := *timeMustParse(time.RFC3339, "2019-12-09T14:10:00.000Z")
			Expect(missedT.Equal(expectedMiss)).To(BeTrue())
			expectedNext := *timeMustParse(time.RFC3339, "2019-12-09T14:11:00.000Z")
			Expect(nextT.Equal(expectedNext)).To(BeTrue())
		})

		It("should return an error if there are too many misses", func() {
			now := *timeMustParse(time.RFC3339, "2019-12-09T14:10:30.000Z")
			cronDep := makeVpCronDep(
				"test",
				"* * * * *", // every minute always!
				nil,         // no deadline
				metav1.Time{Time: now.Add(-(maxMissedStartTimes + 1) * time.Minute)}, // created maxMissedStartTimes + 1 minutes ago
				nil,
			)
			_, _, err := getNextCronSchedule(cronDep, now)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("too many missed start times"))
		})
	})

	Context("schedule time parsing", func() {
		var makeVpDep = func(name string, timeVal *string) *v1beta2.VpDeployment {
			dep := &v1beta2.VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: "default",
				},
			}

			if timeVal != nil {
				dep.Annotations = annotations.Create(
					annotations.Pair(scheduledTimeAnnotation, *timeVal))
			}

			return dep
		}

		Describe("getScheduledTimeForDeployment", func() {
			It("should return nil if the annotation is not present", func() {
				t, err := getScheduledTimeForDeployment(makeVpDep("not-present", nil))
				Expect(t).To(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return the parse error if the annotation is in the wrong format", func() {
				t, err := getScheduledTimeForDeployment(makeVpDep("invalid", pointer.StringPtr("not-a-time-string")))
				Expect(t).To(BeNil())
				Expect(err).To(HaveOccurred())
			})

			It("should parse the time", func() {
				timeVal := "2019-12-09T14:27:58.328Z"
				t, err := getScheduledTimeForDeployment(makeVpDep("invalid", &timeVal))
				Expect(t).ToNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
				tExpected := timeMustParse(time.RFC3339, timeVal)
				Expect(t.Equal(*tExpected)).To(BeTrue())
			})
		})

		Describe("getMostRecentScheduledTime", func() {
			It("should ignore errors", func() {
				depList := v1beta2.VpDeploymentList{
					Items: []v1beta2.VpDeployment{
						*makeVpDep("invalid", pointer.StringPtr("not-a-time")),
						*makeVpDep("not-set", nil),
					},
				}
				t := getMostRecentScheduledTime(nil, &depList) // nolint:staticcheck
				Expect(t).To(BeNil())
			})

			It("should return nil for an empty list", func() {
				depList := v1beta2.VpDeploymentList{
					Items: []v1beta2.VpDeployment{},
				}
				t := getMostRecentScheduledTime(nil, &depList) // nolint:staticcheck
				Expect(t).To(BeNil())
			})

			It("should get the earliest time", func() {
				earliest := "2019-12-09T14:27:58.328Z"
				depList := v1beta2.VpDeploymentList{
					Items: []v1beta2.VpDeployment{
						*makeVpDep("latest", pointer.StringPtr("2019-12-08T14:27:58.328Z")),
						*makeVpDep("earliest", &earliest),
						*makeVpDep("not-set", nil),
					},
				}
				t := getMostRecentScheduledTime(nil, &depList) // nolint:staticcheck
				Expect(t).ToNot(BeNil())
				tExpected := timeMustParse(time.RFC3339, earliest)
				Expect(t.Equal(*tExpected)).To(BeTrue())
			})
		})
	})

	Describe("getDeploymentsToDeleteByAge", func() {
		var makeVpDep = func(name string, time *metav1.Time) *v1beta2.VpDeployment {
			return &v1beta2.VpDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: "default",
				},
				Status: &v1beta2.VpDeploymentStatus{
					StartTime: time,
				},
			}
		}

		It("should return an empty slice given an empty slice", func() {
			toDelete := getDeploymentsToDeleteByAge(nil, 2)
			Expect(toDelete).To(HaveLen(0))
		})

		It("should truncate in order if none have started", func() {
			deps := []*v1beta2.VpDeployment{
				makeVpDep("not-started-1", nil),
				makeVpDep("not-started-2", nil),
				makeVpDep("not-started-3", nil),
			}
			toDelete := getDeploymentsToDeleteByAge(deps, 2)
			Expect(toDelete).To(HaveLen(1))
			dep := toDelete[0]
			Expect(dep.Name).To(Equal("not-started-1"))
		})

		It("should truncate remove the oldest first", func() {
			oldest := makeVpDep("dep-2", metaTimeMustParse(time.RFC3339, "2019-12-09T14:27:58.328Z"))
			youngest := makeVpDep("dep-1", metaTimeMustParse(time.RFC3339, "2019-12-12T14:27:58.328Z"))
			middle := makeVpDep("dep-3", metaTimeMustParse(time.RFC3339, "2019-12-11T14:27:58.328Z"))
			deps := []*v1beta2.VpDeployment{
				oldest,
				youngest,
				middle,
			}
			toDelete := getDeploymentsToDeleteByAge(deps, 1)
			Expect(toDelete).To(HaveLen(2))
			Expect(toDelete).To(ContainElements(oldest, middle))
		})
	})

	It("should reconcile non-existent objs", func() {
		_, err := reconciler.Reconcile(ctrl.Request{
			NamespacedName: types.NamespacedName{
				Name:      "not-found",
				Namespace: "default",
			},
		})
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("buildVpDeploymentForCronDep", func() {
		var makeVpCronDep = func(
			name string,
			template v1beta2.VpDeploymentSpecTemplate,
		) *v1beta2.VpCronDeployment {
			return &v1beta2.VpCronDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: "default",
				},
				Spec: v1beta2.VpCronDeploymentSpec{
					VpDeploymentTemplate: template,
				},
			}
		}

		It("should build a VpDeployment from the template with the scheduled time", func() {
			template := v1beta2.VpDeploymentSpecTemplate{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "a-namespace",
					Labels: map[string]string{
						"test-label": "is-present!",
					},
					Annotations: annotations.Create(
						annotations.Pair("test-annotation", "is-also-present!")),
				},
				Spec: v1beta2.VpDeploymentObjectSpec{
					Spec: v1beta2.VpDeploymentSpec{
						State:                        v1beta2.RunningState,
						MaxSavepointCreationAttempts: pointer.Int32Ptr(5),
						MaxJobCreationAttempts:       pointer.Int32Ptr(2),
					},
					DeploymentTargetName: "a-dep-target",
				},
			}
			cronDep := makeVpCronDep("cool-cron", template)
			scheduledTime := timeMustParse(time.RFC3339, "2019-12-11T14:27:58.328Z")
			dep, err := reconciler.buildVpDeploymentForCronDep(cronDep, *scheduledTime)
			Expect(err).NotTo(HaveOccurred())
			Expect(dep.Name).To(MatchRegexp("cool-cron-[0-9]+"))
			Expect(dep.Namespace).To(Equal(cronDep.Namespace))
			Expect(dep.Labels["test-label"]).To(Equal("is-present!"))
			Expect(annotations.Has(dep.Annotations, scheduledTimeAnnotation)).To(BeTrue())
			Expect(annotations.Get(dep.Annotations, scheduledTimeAnnotation)).To(Equal(scheduledTime.Format(scheduledTimeFormat)))
			Expect(annotations.Get(dep.Annotations, "test-annotation")).To(Equal("is-also-present!"))
			Expect(dep.Spec.Spec.State).To(Equal(template.Spec.Spec.State))
		})
	})
})
