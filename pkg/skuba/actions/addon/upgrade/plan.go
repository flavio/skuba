/*
 * Copyright (c) 2019 SUSE LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package addons

import (
	"fmt"

	"github.com/pkg/errors"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/SUSE/skuba/internal/pkg/skuba/addons"
	"github.com/SUSE/skuba/internal/pkg/skuba/kubeadm"
	"github.com/SUSE/skuba/internal/pkg/skuba/kubernetes"
	"github.com/SUSE/skuba/internal/pkg/skuba/upgrade/addon"
)

func Plan(client clientset.Interface) error {
	currentClusterVersion, err := kubeadm.GetCurrentClusterVersion(client)
	if err != nil {
		return err
	}
	clusterConfiguration, err := kubeadm.GetClusterConfiguration(client)
	if err != nil {
		return errors.Wrap(err, "Could not fetch cluster configuration")
	}

	addonConfiguration := addons.AddonConfiguration{
		ClusterVersion: currentClusterVersion,
		ControlPlane:   clusterConfiguration.ControlPlaneEndpoint,
		ClusterName:    clusterConfiguration.ClusterName,
	}

	// check local addons cluster folder configuration is up-to-date
	match, err := addons.CheckLocalAddonsBaseManifests(addonConfiguration)
	if err != nil {
		return err
	}
	if !match {
		fmt.Println("Current local addons cluster folder configuration is out-of-date.")
		fmt.Println("Please run \"skuba addon refresh localconfig\" before you perform addon upgrade.")
		return nil
	}

	currentVersion := currentClusterVersion.String()
	latestVersion := kubernetes.LatestVersion().String()
	allNodesVersioningInfo, err := kubernetes.AllNodesVersioningInfo(client)
	if err != nil {
		return err
	}
	allNodesMatchClusterVersion := kubernetes.AllNodesMatchClusterVersionWithVersioningInfo(allNodesVersioningInfo, currentClusterVersion)
	fmt.Printf("Current Kubernetes cluster version: %s\n", currentVersion)
	fmt.Printf("Latest Kubernetes version: %s\n", latestVersion)
	fmt.Println()

	if !allNodesMatchClusterVersion {
		return errors.Errorf("Not all nodes match clusterVersion %s", currentVersion)
	}

	updatedAddons, err := addon.UpdatedAddons(client, currentClusterVersion)
	if err != nil {
		return err
	}
	if addon.HasAddonUpdate(updatedAddons) {
		fmt.Printf("Addon upgrades for %s:\n", currentVersion)
		addon.PrintAddonUpdates(updatedAddons)

		dryRun := true
		if err := addons.DeployAddons(client, addonConfiguration, dryRun); err != nil {
			return errors.Wrap(err, "Failed to plan addons")
		}
	} else {
		fmt.Printf("Congratulations! Addons for %s are already at the latest version available\n", currentVersion)
	}

	return nil
}
