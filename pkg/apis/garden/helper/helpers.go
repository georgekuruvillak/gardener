// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helper

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver"

	"github.com/gardener/gardener/pkg/apis/garden"
)

// DetermineCloudProviderInProfile takes a CloudProfile specification and returns the cloud provider this profile is used for.
// If it is not able to determine it, an error will be returned.
func DetermineCloudProviderInProfile(spec garden.CloudProfileSpec) (garden.CloudProvider, error) {
	var (
		cloud     garden.CloudProvider
		numClouds = 0
	)

	if spec.AWS != nil {
		numClouds++
		cloud = garden.CloudProviderAWS
	}
	if spec.Azure != nil {
		numClouds++
		cloud = garden.CloudProviderAzure
	}
	if spec.GCP != nil {
		numClouds++
		cloud = garden.CloudProviderGCP
	}
	if spec.OpenStack != nil {
		numClouds++
		cloud = garden.CloudProviderOpenStack
	}
	if spec.Alicloud != nil {
		numClouds++
		cloud = garden.CloudProviderAlicloud
	}
	if spec.Packet != nil {
		numClouds++
		cloud = garden.CloudProviderPacket
	}

	if numClouds != 1 {
		return "", errors.New("cloud profile must only contain exactly one field of alicloud/aws/azure/gcp/openstack/packet")
	}
	return cloud, nil
}

// DetermineCloudProviderInShoot takes a Shoot cloud object and returns the cloud provider this profile is used for.
// If it is not able to determine it, an error will be returned.
func DetermineCloudProviderInShoot(cloudObj garden.Cloud) (garden.CloudProvider, error) {
	var (
		cloud     garden.CloudProvider
		numClouds = 0
	)

	if cloudObj.AWS != nil {
		numClouds++
		cloud = garden.CloudProviderAWS
	}
	if cloudObj.Azure != nil {
		numClouds++
		cloud = garden.CloudProviderAzure
	}
	if cloudObj.GCP != nil {
		numClouds++
		cloud = garden.CloudProviderGCP
	}
	if cloudObj.OpenStack != nil {
		numClouds++
		cloud = garden.CloudProviderOpenStack
	}
	if cloudObj.Alicloud != nil {
		numClouds++
		cloud = garden.CloudProviderAlicloud
	}
	if cloudObj.Packet != nil {
		numClouds++
		cloud = garden.CloudProviderPacket
	}

	if numClouds != 1 {
		return "", errors.New("cloud object must only contain exactly one field of aws/azure/gcp/openstack/packet")
	}
	return cloud, nil
}

// DetermineLatestMachineImageVersions determines the latest versions (semVer) of the given machine images from a slice of machine images
func DetermineLatestMachineImageVersions(images []garden.MachineImage) (map[string]garden.MachineImageVersion, error) {
	resultMapVersions := make(map[string]garden.MachineImageVersion)

	for _, image := range images {
		latestMachineImageVersion, err := DetermineLatestMachineImageVersion(image)
		if err != nil {
			return nil, err
		}
		resultMapVersions[image.Name] = latestMachineImageVersion
	}
	return resultMapVersions, nil
}

// DetermineLatestMachineImageVersion determines the latest MachineImageVersion from a MachineImage
func DetermineLatestMachineImageVersion(image garden.MachineImage) (garden.MachineImageVersion, error) {
	var (
		latestSemVerVersion       *semver.Version
		latestMachineImageVersion garden.MachineImageVersion
	)

	for _, imageVersion := range image.Versions {
		v, err := semver.NewVersion(imageVersion.Version)
		if err != nil {
			return garden.MachineImageVersion{}, fmt.Errorf("error while parsing machine image version '%s' of machine image '%s': version not valid: %s", imageVersion.Version, image.Name, err.Error())
		}
		if latestSemVerVersion == nil || v.GreaterThan(latestSemVerVersion) {
			latestSemVerVersion = v
			latestMachineImageVersion = imageVersion
		}
	}
	return latestMachineImageVersion, nil
}

// ShootWantsBasicAuthentication returns true if basic authentication is not configured or
// if it is set explicitly to 'true'.
func ShootWantsBasicAuthentication(kubeAPIServerConfig *garden.KubeAPIServerConfig) bool {
	if kubeAPIServerConfig == nil {
		return true
	}
	if kubeAPIServerConfig.EnableBasicAuthentication == nil {
		return true
	}
	return *kubeAPIServerConfig.EnableBasicAuthentication
}
