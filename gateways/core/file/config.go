/*
Copyright 2018 BlackRock, Inc.

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

package file

import (
	"github.com/ghodss/yaml"
	"github.com/rs/zerolog"

	"github.com/argoproj/argo-events/gateways/common"
)

// FileEventSourceExecutor implements Eventing
type FileEventSourceExecutor struct {
	Log zerolog.Logger
}

// fileWatcher contains configuration information for this gateway
// +k8s:openapi-gen=true
type fileWatcher struct {
	common.WatchPathConfig `json:",inline"`

	// Type of file operations to watch
	// Refer https://github.com/fsnotify/fsnotify/blob/master/fsnotify.go for more information
	Type string `json:"type"`
}

func parseEventSource(eventSource string) (interface{}, error) {
	var f *fileWatcher
	err := yaml.Unmarshal([]byte(eventSource), &f)
	if err != nil {
		return nil, err
	}
	return f, err
}