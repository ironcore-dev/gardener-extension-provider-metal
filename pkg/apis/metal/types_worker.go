// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package metal

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerConfig contains configuration settings for the worker nodes.
type WorkerConfig struct {
	metav1.TypeMeta
	// ExtraIgnition contains additional Ignition for Worker nodes.
	ExtraIgnition *IgnitionConfig
	// ExtraServerLabels is a map of extra labels that are applied to the ServerClaim for Server selection.
	ExtraServerLabels map[string]string
	// AddressesFromNetworks is a list of references to Network resources that should be used to assign IP addresses to the worker nodes.
	AddressesFromNetworks []*AddressesFromNetworks
	// MetaData is a key-value map of additional data which should be passed to the Machine.
	MetaData map[string]string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerStatus contains information about created worker resources.
type WorkerStatus struct {
	metav1.TypeMeta

	// MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
	// gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
	// a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing `Worker`
	// resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
	// reconciliation is possible.
	MachineImages []MachineImage
}

// MachineImage is a mapping from logical names and versions to metal-specific identifiers.
type MachineImage struct {
	// Name is the logical name of the machine image.
	Name string
	// Version is the logical version of the machine image.
	Version string
	// Image is the path to the image.
	Image string
	// Architecture is the CPU architecture of the machine image.
	Architecture *string
}

// IgnitionConfig contains ignition settings.
type IgnitionConfig struct {
	Raw       string
	SecretRef *corev1.LocalObjectReference
	Override  bool
}

// SubnetRef is a reference to the IP subnet.
type SubnetRef struct {
	// Name is the name of the network.
	Name string
	// APIGroup is the group of the IP pool
	APIGroup string
	// Kind is the kind of the IP pool
	Kind string
}

// AddressesFromNetworks is a reference to a network resource.
type AddressesFromNetworks struct {
	// Key is the name of metadata key for the network.
	Key string
	// SubnetRef is a reference to the IP subnet.
	SubnetRef *SubnetRef
}
