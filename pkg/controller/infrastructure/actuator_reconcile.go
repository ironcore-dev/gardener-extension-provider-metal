// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/gardener/gardener/extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/go-logr/logr"

	metalv1alpha1 "github.com/ironcore-dev/gardener-extension-provider-ironcore-metal/pkg/apis/metal/v1alpha1"
)

var (
	infrastructureConfig metalv1alpha1.InfrastructureConfig
)

// Reconcile implements infrastructure actuator reconciliation
func (a *actuator) Reconcile(ctx context.Context, log logr.Logger, infra *extensionsv1alpha1.Infrastructure, cluster *controller.Cluster) error {
	err := json.Unmarshal(cluster.Shoot.Spec.Provider.InfrastructureConfig.Raw, &infrastructureConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal infrastructure config: %w", err)
	}

	var newNodes []string
	if infrastructureConfig.Networks != nil {
		for _, network := range infrastructureConfig.Networks {
			if network.Name == "" {
				return fmt.Errorf("network name is required")
			}
			newNodes = append(newNodes, network.CIDR)
		}
		if !slices.Equal(infra.Status.Networking.Nodes, newNodes) {
			infra.Status.Networking.Nodes = newNodes
		}
	}

	return a.reconcile(ctx, log, infra, cluster)
}

func (a *actuator) reconcile(ctx context.Context, log logr.Logger, infra *extensionsv1alpha1.Infrastructure, cluster *controller.Cluster) error {
	return nil
}
