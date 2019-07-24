/*

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

package main

import (
	"flag"
	"os"

	ververicaplatformv1beta1 "github.com/fintechstudios/ververica-platform-k8s-controller/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-controller/controllers"
	ververicaplatformapi "github.com/fintechstudios/ververica-platform-k8s-controller/ververica-platform-api"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {

	_ = ververicaplatformv1beta1.AddToScheme(scheme)
	ververicaplatformv1beta1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var (
		metricsAddr          = flag.String("metrics-addr", ":8080", "The address the metric endpoint binds to.")
		enableLeaderElection = flag.Bool("enable-leader-election", false,
			"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
		enableDebugMode = flag.Bool("debug", true, "Enable debug mode for logging.")

		watchNamespace = flag.String("watch-namespace", apiv1.NamespaceAll,
			`Namespace to watch for resources. Default is to watch all namespaces`)

		ververicaPlatformURL = flag.String("ververica-platform-url", "http://localhost:8081/api",
			"The URL to the Ververica Platform API, without a trailing slash.")
	)
	flag.Parse()
	setupLog.Info("Watching namespace", "namespace", watchNamespace)

	ctrl.SetLogger(zap.Logger(*enableDebugMode))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: *metricsAddr,
		LeaderElection:     *enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// Build the Ververica Platform API Client
	ververicaAPIClient := ververicaplatformapi.NewAPIClient(&ververicaplatformapi.Configuration{
		BasePath:      *ververicaPlatformURL,
		DefaultHeader: make(map[string]string), // TODO: allow users to pass these in dynamically
		UserAgent:     "VervericaPlatformK8sController/1.0.0/go",
	})

	setupLog.Info("Created VP API client", "client", ververicaAPIClient)

	err = (&controllers.VPNamespaceReconciler{
		Client:      mgr.GetClient(),
		Log:         ctrl.Log.WithName("controllers").WithName("VPNamespace"),
		VPAPIClient: *ververicaAPIClient,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VPNamespace")
		os.Exit(1)
	}
	err = (&controllers.VPDeploymentTargetReconciler{
		Client:      mgr.GetClient(),
		Log:         ctrl.Log.WithName("controllers").WithName("VPDeploymentTarget"),
		VPAPIClient: *ververicaAPIClient,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VPDeploymentTarget")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
