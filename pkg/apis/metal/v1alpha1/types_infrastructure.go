// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// IgnitionConfig contains ignition settings.
type IgnitionConfig struct {
	// Raw contains an inline ignition config, which is merged with the config from the os extension.
	// +optional
	Raw string `json:"raw,omitempty"`
}

// WorkerConfig contains settings per pool, which are specific to the metal-operator.
type WorkerConfig struct {
	// ExtraIgnition contains additional ignition configuration.
	// +optional
	ExtraIgnition *IgnitionConfig `json:"extraIgnition,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InfrastructureConfig infrastructure configuration resource
type InfrastructureConfig struct {
	metav1.TypeMeta `json:",inline"`

	// Worker contains settings per worker pool specific to the metal-operator
	// +optional
	Worker map[string]WorkerConfig `json:"worker,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InfrastructureStatus contains information about created infrastructure resources.
type InfrastructureStatus struct {
	metav1.TypeMeta `json:",inline"`
}
