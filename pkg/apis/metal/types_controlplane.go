// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package metal

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControlPlaneConfig contains configuration settings for the control plane.
type ControlPlaneConfig struct {
	metav1.TypeMeta

	// CloudControllerManager contains configuration settings for the cloud-controller-manager.
	CloudControllerManager *CloudControllerManagerConfig

	// LoadBalancerConfig contains configuration settings for the shoot loadbalancing.
	LoadBalancerConfig *LoadBalancerConfig
}

// CloudControllerNetworking contains configuration settings for CCM networking.
type CloudControllerNetworking struct {
	// ConfigureNodeAddresses enables the configuration of node addresses.
	ConfigureNodeAddresses bool
}

// CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.
type CloudControllerManagerConfig struct {
	// FeatureGates contains information about enabled feature gates.
	FeatureGates map[string]bool

	// CloudControllerNetworking contains configuration settings for CCM networking.
	CloudControllerNetworking *CloudControllerNetworking
}

// LoadBalancerConfig contains configuration settings for the shoot loadbalancing.
type LoadBalancerConfig struct {
	// MetallbConfig contains configuration settings for metallb.
	MetallbConfig *MetallbConfig

	// CalicoBgpConfig contains configuration settings for calico.
	CalicoBgpConfig *CalicoBgpConfig
}

// MetallbConfig contains configuration settings for metallb.
type MetallbConfig struct {
	// IPAddressPool contains IP address pools for metallb.
	IPAddressPool []string

	// EnableSpeaker enables the metallb speaker.
	EnableSpeaker bool

	// EnableL2Advertisement enables L2 advertisement.
	EnableL2Advertisement bool
}

// CalicoBgpConfig contains BGP configuration settings for calico.
type CalicoBgpConfig struct {
	// ASNumber is the default AS number used by a node.
	ASNumber int

	// ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs.
	ServiceLoadBalancerIPs []string

	// ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs.
	ServiceExternalIPs []string

	// ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated.
	ServiceClusterIPs []string

	// BGPPeer contains configuration for BGPPeer resource.
	BgpPeer []BgpPeer
}

// BgpPeer contains configuration for BGPPeer resource.
type BgpPeer struct {
	// PeerIP contains IP address of BGP peer followed by an optional port number to peer with.
	PeerIP string

	// ASNumber contains the AS number of the BGP peer.
	ASNumber int

	// NodeSelector is a key-value pair to select nodes that should have this peering.
	NodeSelector string
}
