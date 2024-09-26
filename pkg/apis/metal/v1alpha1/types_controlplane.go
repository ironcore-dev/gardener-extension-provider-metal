// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlPlaneConfig contains configuration settings for the control plane.
type ControlPlaneConfig struct {
	metav1.TypeMeta `json:",inline"`

	// CloudControllerManager contains configuration settings for the cloud-controller-manager.
	// +optional
	CloudControllerManager *CloudControllerManagerConfig `json:"cloudControllerManager,omitempty"`

	// LoadBalancerConfig contains configuration settings for the shoot loadbalancing.
	// +optional
	LoadBalancerConfig *LoadBalancerConfig `json:"loadBalancerConfig,omitempty"`
}

// CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.
type CloudControllerManagerConfig struct {
	// FeatureGates contains information about enabled feature gates.
	// +optional
	FeatureGates map[string]bool `json:"featureGates,omitempty"`
}

// LoadBalancerConfig contains configuration settings for the shoot loadbalancing.
type LoadBalancerConfig struct {
	// MetallbConfig contains configuration settings for metallb.
	// +optional
	MetallbConfig *MetallbConfig `json:"metallbConfig,omitempty"`
}

// MetallbConfig contains configuration settings for metallb.
type MetallbConfig struct {
	// IPAddressPool contains IP address pools for metallb.
	// +optional
	IPAddressPool []string `json:"ipAddressPool,omitempty"`

	// EnableSpeaker enables the metallb speaker.
	// +optional
	EnableSpeaker bool `json:"enableSpeaker,omitempty"`

	// EnableL2Advertisement enables L2 advertisement.
	// +optional
	EnableL2Advertisement bool `json:"enableL2Advertisement,omitempty"`
}
