// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package infrastructure

import (
	"context"
	"encoding/json"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"

	metalv1alpha1 "github.com/ironcore-dev/gardener-extension-provider-metal/pkg/apis/metal/v1alpha1"
)

var (
	shootVersionMajorMinor = "1.2"
	shootVersion           = shootVersionMajorMinor + ".3"
)

var _ = Describe("Actuator Reconcile", func() {
	var (
		ctx     context.Context
		log     logr.Logger
		infra   *extensionsv1alpha1.Infrastructure
		cluster *extensionscontroller.Cluster
		act     *actuator
	)

	BeforeEach(func() {
		ctx = context.TODO()
		log = logr.Discard()

		infra = &extensionsv1alpha1.Infrastructure{
			Status: extensionsv1alpha1.InfrastructureStatus{
				Networking: &extensionsv1alpha1.InfrastructureStatusNetworking{
					Nodes: []string{},
				},
			},
		}

		infrastructureConfig := metalv1alpha1.InfrastructureConfig{
			Networks: []metalv1alpha1.Networks{
				{Name: "worker-network-1", CIDR: "10.10.10.0/24", ID: "1"},
				{Name: "worker-network-2", CIDR: "10.10.20.0/24", ID: "2"},
			},
		}
		infrastructureConfigRaw, _ := json.Marshal(infrastructureConfig)

		cluster = &extensionscontroller.Cluster{
			Shoot: &gardencorev1beta1.Shoot{
				Spec: gardencorev1beta1.ShootSpec{
					Kubernetes: gardencorev1beta1.Kubernetes{
						Version: shootVersion,
					},
					Provider: gardencorev1beta1.Provider{
						InfrastructureConfig: &runtime.RawExtension{
							Raw: infrastructureConfigRaw,
						},
					},
				},
			},
		}

		act = &actuator{}
	})

	Describe("#Reconcile", func() {
		It("should update infra.Status.Networking.Nodes", func() {
			err := act.Reconcile(ctx, log, infra, cluster)
			Expect(err).NotTo(HaveOccurred())

			// Verify that infra.Status.Networking.Nodes is updated
			expectedNodes := []string{"10.10.10.0/24", "10.10.20.0/24"}
			Expect(infra.Status.Networking.Nodes).To(Equal(expectedNodes))
		})
	})
})
