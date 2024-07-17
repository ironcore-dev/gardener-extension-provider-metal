//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&InfrastructureConfig{}, func(obj interface{}) { SetObjectDefaults_InfrastructureConfig(obj.(*InfrastructureConfig)) })
	return nil
}

func SetObjectDefaults_InfrastructureConfig(in *InfrastructureConfig) {
	if in.NetworkRef != nil {
		if in.NetworkRef.Name == "" {
			in.NetworkRef.Name = ""
		}
	}
}
