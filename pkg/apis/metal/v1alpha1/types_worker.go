// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerConfig contains configuration settings for the worker nodes.
type WorkerConfig struct {
	metav1.TypeMeta
	// ExtraIgnition contains additional Ignition for Worker nodes.
	// +optional
	ExtraIgnition *IgnitionConfig `json:"extraIgnition,omitempty"`
	// ExtraServerLabels is a map of additional labels that are applied to the ServerClaim for Server selection.
	// +optional
	ExtraServerLabels map[string]string `json:"extraServerLabels,omitempty"`
	// AddressesFromNetworks is a list of references to Network resources that should be used to assign IP addresses to the worker nodes.
	// +optional
	AddressesFromNetworks []*AddressesFromNetworks `json:"addressesFromNetworks,omitempty"`
	// MedaData is a key-value map of additional data which should be passed to the Machine.
	// +optional
	MetaData map[string]string `json:"metaData,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerStatus contains information about created worker resources.
type WorkerStatus struct {
	metav1.TypeMeta `json:",inline"`

	// MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
	// gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
	// a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing `Worker`
	// resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
	// reconciliation is possible.
	// +optional
	MachineImages []MachineImage `json:"machineImages,omitempty"`
}

// MachineImage is a mapping from logical names and versions to metal-specific identifiers.
type MachineImage struct {
	// Name is the logical name of the machine image.
	Name string `json:"name"`
	// Version is the logical version of the machine image.
	Version string `json:"version"`
	// Image is the path to the image.
	Image string `json:"image"`
	// Architecture is the CPU architecture of the machine image.
	// +optional
	Architecture *string `json:"architecture,omitempty"`
}

// IgnitionConfig contains ignition settings.
type IgnitionConfig struct {
	// Raw contains an inline ignition config, which is merged with the config from the os extension.
	// +optional
	Raw string `json:"raw,omitempty"`

	// SecretRef is a reference to a secret containing the ignition config.
	// +optional
	SecretRef *corev1.LocalObjectReference `json:"secretRef,omitempty"`

	// Override configures, if ignition keys set by the os-extension are overridden
	// by extra ignition.
	// +optional
	Override bool `json:"override,omitempty"`
}

// SubnetRef is a reference to the IP subnet.
type SubnetRef struct {
	// Name is the name of the network.
	Name string `json:"name"`
	// APIGroup is the group of the IP pool
	APIGroup string `json:"apiGroup"`
	// Kind is the kind of the IP pool
	Kind string `json:"kind"`
}

// AddressesFromNetworks is a reference to a network resource.
type AddressesFromNetworks struct {
	// Key is the name of metadata key for the network.
	Key string `json:"key"`
	// SubnetRef is a reference to the IP subnet.
	SubnetRef *SubnetRef `json:"subnetRef"`
}
