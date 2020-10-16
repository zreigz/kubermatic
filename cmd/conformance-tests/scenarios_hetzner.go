/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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
	"fmt"

	"k8c.io/kubermatic/v2/pkg/semver"
	apimodels "k8c.io/kubermatic/v2/pkg/test/e2e/api/utils/apiclient/models"
)

// Returns a matrix of (version x operating system)
func getHetznerScenarios(versions []*semver.Semver) []testScenario {
	var scenarios []testScenario
	for _, v := range versions {
		// Ubuntu
		scenarios = append(scenarios, &hetznerScenario{
			version: v,
			nodeOsSpec: apimodels.OperatingSystemSpec{
				Ubuntu: &apimodels.UbuntuSpec{},
			},
		})
		// CentOS
		scenarios = append(scenarios, &hetznerScenario{
			version: v,
			nodeOsSpec: apimodels.OperatingSystemSpec{
				Centos: &apimodels.CentOSSpec{},
			},
		})
	}

	return scenarios
}

type hetznerScenario struct {
	version    *semver.Semver
	nodeOsSpec apimodels.OperatingSystemSpec
}

func (s *hetznerScenario) Name() string {
	return fmt.Sprintf("hetzner-%s-%s", getOSNameFromSpec(s.nodeOsSpec), s.version.String())
}

func (s *hetznerScenario) Cluster(secrets secrets) *apimodels.CreateClusterSpec {
	return &apimodels.CreateClusterSpec{
		Cluster: &apimodels.Cluster{
			Type: "kubernetes",
			Spec: &apimodels.ClusterSpec{
				Cloud: &apimodels.CloudSpec{
					DatacenterName: "hetzner-nbg1",
					Hetzner: &apimodels.HetznerCloudSpec{
						Token: secrets.Hetzner.Token,
					},
				},
				Version: s.version.String(),
			},
		},
	}
}

func (s *hetznerScenario) MachineDeployments(num int, _ secrets) ([]apimodels.MachineDeployment, error) {
	replicas := int32(num)
	nodeType := "cx31"

	return []apimodels.MachineDeployment{
		{
			Spec: &apimodels.MachineDeploymentSpec{
				Replicas: &replicas,
				Template: &apimodels.MachineSpec{
					Cloud: &apimodels.MachineCloudSpec{
						Hetzner: &apimodels.HetznerMachineSpec{
							Type: &nodeType,
						},
					},
					Versions: &apimodels.MachineVersionInfo{
						Kubelet: s.version.String(),
					},
					OperatingSystem: &s.nodeOsSpec,
				},
			},
		},
	}, nil
}

func (s *hetznerScenario) OS() apimodels.OperatingSystemSpec {
	return s.nodeOsSpec
}
