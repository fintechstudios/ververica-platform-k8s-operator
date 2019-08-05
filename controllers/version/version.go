/*
Copyright 2017 The Kubernetes Authors.
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

// Modified version of the github.com/kubernetes-sigs/kubebuilder version package

package version

import (
	"fmt"
	"runtime"
)

var (
	controllerVersion = "unknown"
	goos              = runtime.GOOS
	goarch            = runtime.GOARCH
	gitCommit         = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	buildDate         = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Version is a simple representation of the current application and runtime version
type Version struct {
	ControllerVersion string `json:"controllerVersion"`
	GitCommit         string `json:"gitCommit"`
	BuildDate         string `json:"buildDate"`
	GoVersion         string `json:"goVersion"`
	GoOs              string `json:"goOs"`
	GoArch            string `json:"goArch"`
}

// GetVersion constructs the current version
func GetVersion() Version {
	return Version{
		ControllerVersion: controllerVersion,
		GitCommit:         gitCommit,
		BuildDate:         buildDate,
		GoVersion:         runtime.Version(),
		GoOs:              goos,
		GoArch:            goarch,
	}
}

// ToString gets a simple string representation of the version
func (v Version) ToString() string {
	return fmt.Sprintf("%#v", v)
}
