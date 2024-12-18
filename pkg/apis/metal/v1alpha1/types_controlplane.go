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

// CloudControllerNetworking contains configuration settings for CCM networking.
type CloudControllerNetworking struct {
	// ConfigureNodeAddresses enables the configuration of node addresses.
	// +optional
	ConfigureNodeAddresses bool `json:"configureNodeAddresses,omitempty"`
}

// CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.
type CloudControllerManagerConfig struct {
	// FeatureGates contains information about enabled feature gates.
	// +optional
	FeatureGates map[string]bool `json:"featureGates,omitempty"`

	// CloudControllerNetworking contains configuration settings for CCM networking.
	// +optional
	CloudControllerNetworking *CloudControllerNetworking `json:"cloudControllerNetworking,omitempty"`
}

// LoadBalancerConfig contains configuration settings for the shoot loadbalancing.
type LoadBalancerConfig struct {
	// MetallbConfig contains configuration settings for metallb.
	// +optional
	MetallbConfig *MetallbConfig `json:"metallbConfig,omitempty"`

	// CalicoBgpConfig contains configuration settings for calico.
	// +optional
	CalicoBgpConfig *CalicoBgpConfig `json:"calicoBgpConfig,omitempty"`
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

// CalicoBgpConfig contains BGP configuration settings for calico.
type CalicoBgpConfig struct {
	// ASNumber is the default AS number used by a node.
	// +required
	ASNumber int `json:"asNumber"`

	// ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs.
	// +optional
	ServiceLoadBalancerIPs []string `json:"serviceLoadBalancerIPs,omitempty"`

	// ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs.
	// +optional
	ServiceExternalIPs []string `json:"serviceExternalIPs,omitempty"`

	// ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated.
	// +optional
	ServiceClusterIPs []string `json:"serviceClusterIPs,omitempty"`

	// BGPPeer contains configuration for BGPPeer resource.
	// +optional
	BgpPeer []BgpPeer `json:"bgpPeer,omitempty"`
}

// BgpPeer contains configuration for BGPPeer resource.
type BgpPeer struct {
	// PeerIP contains IP address of BGP peer followed by an optional port number to peer with.
	// +required
	PeerIP string `json:"peerIP"`

	// ASNumber contains the AS number of the BGP peer.
	// +required
	ASNumber int `json:"asNumber"`

	// NodeSelector is a key-value pair to select nodes that should have this peering.
	// +optional
	NodeSelector string `json:"nodeSelector,omitempty"`
}
