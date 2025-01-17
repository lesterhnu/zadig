/*
Copyright 2022 The KodeRover Authors.

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

package step

type StepDeploySpec struct {
	Env                string     `bson:"env"                              json:"env"                                 yaml:"env"`
	ServiceName        string     `bson:"service_name"                     json:"service_name"                        yaml:"service_name"`
	ServiceType        string     `bson:"service_type"                     json:"service_type"                        yaml:"service_type"`
	ServiceModule      string     `bson:"service_module"                   json:"service_module"                      yaml:"service_module"`
	SkipCheckRunStatus bool       `bson:"skip_check_run_status"            json:"skip_check_run_status"               yaml:"skip_check_run_status"`
	Image              string     `bson:"image"                            json:"image"                               yaml:"image"`
	ClusterID          string     `bson:"cluster_id"                       json:"cluster_id"                          yaml:"cluster_id"`
	Timeout            int        `bson:"timeout"                          json:"timeout"                             yaml:"timeout"`
	ReplaceResources   []Resource `bson:"replace_resources"                json:"replace_resources"                   yaml:"replace_resources"`
}

type Resource struct {
	Name      string `bson:"name"                              json:"name"                                 yaml:"name"`
	Kind      string `bson:"kind"                              json:"kind"                                 yaml:"kind"`
	Container string `bson:"container"                         json:"container"                            yaml:"container"`
	Origin    string `bson:"origin"                            json:"origin"                               yaml:"origin"`
}
