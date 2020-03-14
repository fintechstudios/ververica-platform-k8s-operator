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
	"context"
	"flag"
	"fmt"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform"
	"os"
	"runtime"

	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta1"
	"github.com/fintechstudios/ververica-platform-k8s-operator/controllers"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager"
	appmanagerapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	platformapi "github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/platform-api"
	dotenv "github.com/joho/godotenv"
	apiv1 "k8s.io/api/core/v1"
	k8s "k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	operatorVersion = "unknown"
	goos            = runtime.GOOS
	goarch          = runtime.GOARCH
	gitCommit       = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	buildDate       = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Version is a simple representation of the current application and runtime operatorVersion
type Version struct {
	OperatorVersion string `json:"operatorVersion"`
	GitCommit       string `json:"gitCommit"`
	BuildDate       string `json:"buildDate"`
	GoVersion       string `json:"goVersion"`
	GoOs            string `json:"goOs"`
	GoArch          string `json:"goArch"`
}

// GetVersion constructs the current operatorVersion
func GetVersion() Version {
	return Version{
		OperatorVersion: operatorVersion,
		GitCommit:       gitCommit,
		BuildDate:       buildDate,
		GoVersion:       runtime.Version(),
		GoOs:            goos,
		GoArch:          goarch,
	}
}

// String gets a simple string representation of the operatorVersion
func (v Version) String() string {
	return fmt.Sprintf("%#v", v)
}

var (
	scheme   = k8s.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = v1beta1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var (
		metricsAddr          = flag.String("metrics-addr", ":8080", "The address the metric endpoint binds to.")
		enableLeaderElection = flag.Bool("enable-leader-election", false,
			"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
		enableDebugMode = flag.Bool("debug", false, "Enable debug mode for logging.")
		watchNamespace  = flag.String("watch-namespace", apiv1.NamespaceAll,
			`Namespace to watch for resources. Default is to watch all namespaces`)
		vvpURL = flag.String("vvp-url", "http://localhost:8081",
			"The URL to the Ververica Platform API, without a trailing slash. Should include the protocol and host.")
		envFile = flag.String("env-file", "", "The path to an environment (`.env`) file to be loaded")
	)
	flag.Parse()

	if *envFile == "" {
		// ignore error if just trying to autoload
		_ = dotenv.Load()
	} else {
		err := dotenv.Load(*envFile)
		if err != nil {
			setupLog.Error(err, "unable to load env file")
			os.Exit(1)
		}
	}

	setupLog.Info("Watching namespace", "namespace", watchNamespace)

	ctrl.SetLogger(zap.New(zap.UseDevMode(*enableDebugMode)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: *metricsAddr,
		LeaderElection:     *enableLeaderElection,
		Namespace:          *watchNamespace,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	version := GetVersion()
	setupLog.Info("Starting Ververica Platform K8s controller",
		"operatorVersion", version.String())

	// Create clients
	userAgent := fmt.Sprintf("VervericaPlatformK8sOperator/%s/go-%s", version.OperatorVersion, version.GoVersion)
	platformAPIConfig := &platformapi.Configuration{
		BasePath:      *vvpURL,
		DefaultHeader: make(map[string]string),
		UserAgent:     userAgent,
	}
	platformClient := platform.NewClient(platformAPIConfig)

	appManagerConfig := &appmanagerapi.Configuration{
		BasePath:      *vvpURL,
		DefaultHeader: make(map[string]string),
		UserAgent:     userAgent,
	}
	appManagerAuthStore := appmanager.NewAuthStore(&platform.TokenManager{PlatformClient: platformClient})
	appManagerClient := appmanager.NewClient(appManagerConfig, appManagerAuthStore)

	cleanup := func(ctx context.Context) {
		tokens, err := appManagerAuthStore.RemoveAllCreatedTokens(ctx)
		if err != nil {
			setupLog.Error(err, "error cleaning up")
		}
		setupLog.Info("Removed %i auth tokens", len(tokens))
	}

	err = (&controllers.VpNamespaceReconciler{
		Client:         mgr.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("VpNamespace"),
		PlatformClient: platformClient,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VpNamespace")
		os.Exit(1)
	}
	err = (&controllers.VpDeploymentTargetReconciler{
		Client:           mgr.GetClient(),
		Log:              ctrl.Log.WithName("controllers").WithName("VpDeploymentTarget"),
		AppManagerClient: appManagerClient,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VpDeploymentTarget")
		os.Exit(1)
	}
	if err = (&controllers.VpDeploymentReconciler{
		Client:           mgr.GetClient(),
		Log:              ctrl.Log.WithName("controllers").WithName("VpDeployment"),
		AppManagerClient: appManagerClient,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VpDeployment")
		os.Exit(1)
	}
	if err = (&controllers.VpSavepointReconciler{
		Client:           mgr.GetClient(),
		Log:              ctrl.Log.WithName("controllers").WithName("VpSavepoint"),
		AppManagerClient: appManagerClient,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VpSavepoint")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	// after the manager has quit, make sure to clean up created resources
	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		cleanup(context.Background())
		os.Exit(1)
	}
	cleanup(context.Background())
}
