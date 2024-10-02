// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package controlplane

import (
	"os"
	"path/filepath"

	"github.com/gardener/gardener/extensions/pkg/controller"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	secretsmanager "github.com/gardener/gardener/pkg/utils/secrets/manager"
	fakesecretsmanager "github.com/gardener/gardener/pkg/utils/secrets/manager/fake"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"

	apismetal "github.com/ironcore-dev/gardener-extension-provider-metal/pkg/apis/metal"
	"github.com/ironcore-dev/gardener-extension-provider-metal/pkg/internal"
	"github.com/ironcore-dev/gardener-extension-provider-metal/pkg/metal"
)

var _ = Describe("Valueprovider Reconcile", func() {
	ns, vp, cluster := SetupTest()

	var (
		fakeClient         client.Client
		fakeSecretsManager secretsmanager.Interface
	)

	BeforeEach(func(ctx SpecContext) {
		curDir, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		Expect(os.Chdir(filepath.Join("..", "..", ".."))).To(Succeed())
		DeferCleanup(os.Chdir, curDir)

		fakeClient = fakeclient.NewClientBuilder().Build()
		fakeSecretsManager = fakesecretsmanager.New(fakeClient, ns.Name)
		Expect(fakeClient.Create(ctx, &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "cloud-controller-manager-server", Namespace: ns.Name}})).To(Succeed())
	})

	Describe("#GetConfigChartValues", func() {
		It("should return correct config chart values", func(ctx SpecContext) {
			cp := &extensionsv1alpha1.ControlPlane{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "control-plane",
					Namespace: ns.Name,
				},
				Spec: extensionsv1alpha1.ControlPlaneSpec{
					Region: "foo",
					SecretRef: corev1.SecretReference{
						Name:      "my-infra-creds",
						Namespace: ns.Name,
					},
					DefaultSpec: extensionsv1alpha1.DefaultSpec{
						Type: metal.Type,
						ProviderConfig: &runtime.RawExtension{
							Raw: encode(&apismetal.ControlPlaneConfig{
								CloudControllerManager: &apismetal.CloudControllerManagerConfig{
									FeatureGates: map[string]bool{
										"CustomResourceValidation": true,
									},
								},
							}),
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cp)).To(Succeed())

			By("ensuring that the provider ConfigMap has been created")
			config := &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: ns.Name,
					Name:      internal.CloudProviderConfigMapName,
				},
			}
			Eventually(Get(config)).Should(Succeed())
			Expect(config.Data).To(HaveKey("cloudprovider.conf"))
			cloudProviderConfig := map[string]any{}
			Expect(yaml.Unmarshal([]byte(config.Data["cloudprovider.conf"]), &cloudProviderConfig)).NotTo(HaveOccurred())
			Expect(cloudProviderConfig["clusterName"]).To(Equal(cluster.Name))
		})
	})

	Describe("#GetControlPlaneShootCRDsChartValues", func() {
		It("should return correct config chart values", func(ctx SpecContext) {
			values, err := vp.GetControlPlaneShootCRDsChartValues(ctx, nil, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(values).To(Equal(map[string]any{}))
		})
	})

	Describe("#GetControlPlaneChartValues", func() {
		It("should return correct config chart values", func(ctx SpecContext) {
			cp := &extensionsv1alpha1.ControlPlane{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "control-plane",
					Namespace: ns.Name,
				},
				Spec: extensionsv1alpha1.ControlPlaneSpec{
					Region: "foo",
					SecretRef: corev1.SecretReference{
						Name:      "my-infra-creds",
						Namespace: ns.Name,
					},
					DefaultSpec: extensionsv1alpha1.DefaultSpec{
						Type: metal.Type,
						ProviderConfig: &runtime.RawExtension{
							Raw: encode(&apismetal.ControlPlaneConfig{
								CloudControllerManager: &apismetal.CloudControllerManagerConfig{
									FeatureGates: map[string]bool{
										"CustomResourceValidation": true,
									},
								},
							}),
						},
					},
				},
			}
			providerCloudProfile := &apismetal.CloudProfileConfig{}
			providerCloudProfileJson, err := json.Marshal(providerCloudProfile)
			Expect(err).NotTo(HaveOccurred())
			networkProviderConfig := &unstructured.Unstructured{Object: map[string]any{
				"kind":       "FooNetworkConfig",
				"apiVersion": "v1alpha1",
				"overlay": map[string]any{
					"enabled": false,
				},
			}}
			networkProviderConfigData, err := runtime.Encode(unstructured.UnstructuredJSONScheme, networkProviderConfig)
			Expect(err).NotTo(HaveOccurred())
			cluster := &controller.Cluster{
				CloudProfile: &gardencorev1beta1.CloudProfile{
					Spec: gardencorev1beta1.CloudProfileSpec{
						ProviderConfig: &runtime.RawExtension{
							Raw: providerCloudProfileJson,
						},
					},
				},
				Shoot: &gardencorev1beta1.Shoot{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: ns.Name,
						Name:      "my-shoot",
					},
					Spec: gardencorev1beta1.ShootSpec{
						Networking: &gardencorev1beta1.Networking{
							ProviderConfig: &runtime.RawExtension{Raw: networkProviderConfigData},
							Pods:           ptr.To[string]("10.0.0.0/16"),
						},
						Kubernetes: gardencorev1beta1.Kubernetes{
							Version: "1.26.0",
							VerticalPodAutoscaler: &gardencorev1beta1.VerticalPodAutoscaler{
								Enabled: true,
							},
						},
					},
				},
			}

			checksums := map[string]string{
				metal.CloudProviderConfigName: "8bafb35ff1ac60275d62e1cbd495aceb511fb354f74a20f7d06ecb48b3a68432",
			}
			values, err := vp.GetControlPlaneChartValues(ctx, cp, cluster, fakeSecretsManager, checksums, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(values).To(Equal(map[string]any{
				"global": map[string]any{
					"genericTokenKubeconfigSecretName": "generic-token-kubeconfig",
				},
				"cloud-controller-manager": map[string]any{
					"enabled":     true,
					"replicas":    1,
					"clusterName": ns.Name,
					"podAnnotations": map[string]any{
						"checksum/secret-cloud-provider-config": "8bafb35ff1ac60275d62e1cbd495aceb511fb354f74a20f7d06ecb48b3a68432",
					},
					"podLabels": map[string]any{
						"maintenance.gardener.cloud/restart": "true",
					},
					"tlsCipherSuites": []string{
						"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
						"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
						"TLS_AES_128_GCM_SHA256",
						"TLS_AES_256_GCM_SHA384",
						"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
						"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
						"TLS_CHACHA20_POLY1305_SHA256",
						"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
						"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
					},
					"secrets": map[string]any{
						"server": "cloud-controller-manager-server",
					},
					"featureGates": map[string]bool{
						"CustomResourceValidation": true,
					},
					"podNetwork":           "10.0.0.0/16",
					"configureCloudRoutes": true,
				},
			}))
		})
	})

	Describe("#GetControlPlaneShootChartValues", func() {
		It("should return correct shoot system chart values without metallb", func(ctx SpecContext) {
			cp := &extensionsv1alpha1.ControlPlane{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "control-plane",
					Namespace: ns.Name,
				},
				Spec: extensionsv1alpha1.ControlPlaneSpec{
					Region: "foo",
					SecretRef: corev1.SecretReference{
						Name:      "my-infra-creds",
						Namespace: ns.Name,
					},
					DefaultSpec: extensionsv1alpha1.DefaultSpec{
						Type: metal.Type,
						ProviderConfig: &runtime.RawExtension{
							Raw: encode(&apismetal.ControlPlaneConfig{
								CloudControllerManager: &apismetal.CloudControllerManagerConfig{
									FeatureGates: map[string]bool{
										"CustomResourceValidation": true,
									},
								},
							}),
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cp)).To(Succeed())

			providerCloudProfile := &apismetal.CloudProfileConfig{}
			providerCloudProfileJson, err := json.Marshal(providerCloudProfile)
			Expect(err).NotTo(HaveOccurred())
			networkProviderConfig := &unstructured.Unstructured{Object: map[string]any{
				"kind":       "FooNetworkConfig",
				"apiVersion": "v1alpha1",
				"overlay": map[string]any{
					"enabled": false,
				},
			}}
			networkProviderConfigData, err := runtime.Encode(unstructured.UnstructuredJSONScheme, networkProviderConfig)
			Expect(err).NotTo(HaveOccurred())
			cluster := &controller.Cluster{
				CloudProfile: &gardencorev1beta1.CloudProfile{
					Spec: gardencorev1beta1.CloudProfileSpec{
						ProviderConfig: &runtime.RawExtension{
							Raw: providerCloudProfileJson,
						},
					},
				},
				Shoot: &gardencorev1beta1.Shoot{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: ns.Name,
						Name:      "my-shoot",
					},
					Spec: gardencorev1beta1.ShootSpec{
						Networking: &gardencorev1beta1.Networking{
							ProviderConfig: &runtime.RawExtension{Raw: networkProviderConfigData},
							Pods:           ptr.To[string]("10.0.0.0/16"),
						},
						Kubernetes: gardencorev1beta1.Kubernetes{
							Version: "1.26.0",
							VerticalPodAutoscaler: &gardencorev1beta1.VerticalPodAutoscaler{
								Enabled: true,
							},
						},
					},
				},
			}

			values, err := vp.GetControlPlaneShootChartValues(ctx, cp, cluster, fakeSecretsManager, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(values).To(Equal(map[string]any{
				"cloud-controller-manager": map[string]any{"enabled": true},
				"metallb": map[string]any{
					"enabled": false,
				},
			}))
		})
	})

	Describe("#GetControlPlaneShootChartValues", func() {
		It("should return correct shoot system chart values with metallb", func(ctx SpecContext) {
			cp := &extensionsv1alpha1.ControlPlane{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "control-plane",
					Namespace: ns.Name,
				},
				Spec: extensionsv1alpha1.ControlPlaneSpec{
					Region: "foo",
					SecretRef: corev1.SecretReference{
						Name:      "my-infra-creds",
						Namespace: ns.Name,
					},
					DefaultSpec: extensionsv1alpha1.DefaultSpec{
						Type: metal.Type,
						ProviderConfig: &runtime.RawExtension{
							Raw: encode(&apismetal.ControlPlaneConfig{
								CloudControllerManager: &apismetal.CloudControllerManagerConfig{
									FeatureGates: map[string]bool{
										"CustomResourceValidation": true,
									},
								},
								LoadBalancerConfig: &apismetal.LoadBalancerConfig{
									MetallbConfig: &apismetal.MetallbConfig{
										IPAddressPool: []string{"10.10.10.0/24", "10.20.20.10-10.20.20.30"},
									},
								},
							}),
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, cp)).To(Succeed())

			providerCloudProfile := &apismetal.CloudProfileConfig{}
			providerCloudProfileJson, err := json.Marshal(providerCloudProfile)
			Expect(err).NotTo(HaveOccurred())
			networkProviderConfig := &unstructured.Unstructured{Object: map[string]any{
				"kind":       "FooNetworkConfig",
				"apiVersion": "v1alpha1",
				"overlay": map[string]any{
					"enabled": false,
				},
			}}
			networkProviderConfigData, err := runtime.Encode(unstructured.UnstructuredJSONScheme, networkProviderConfig)
			Expect(err).NotTo(HaveOccurred())
			cluster := &controller.Cluster{
				CloudProfile: &gardencorev1beta1.CloudProfile{
					Spec: gardencorev1beta1.CloudProfileSpec{
						ProviderConfig: &runtime.RawExtension{
							Raw: providerCloudProfileJson,
						},
					},
				},
				Shoot: &gardencorev1beta1.Shoot{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: ns.Name,
						Name:      "my-shoot",
					},
					Spec: gardencorev1beta1.ShootSpec{
						Networking: &gardencorev1beta1.Networking{
							ProviderConfig: &runtime.RawExtension{Raw: networkProviderConfigData},
							Pods:           ptr.To[string]("10.0.0.0/16"),
						},
						Kubernetes: gardencorev1beta1.Kubernetes{
							Version: "1.26.0",
							VerticalPodAutoscaler: &gardencorev1beta1.VerticalPodAutoscaler{
								Enabled: true,
							},
						},
					},
				},
			}

			values, err := vp.GetControlPlaneShootChartValues(ctx, cp, cluster, fakeSecretsManager, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(values).To(Equal(map[string]any{
				"cloud-controller-manager": map[string]any{"enabled": true},
				"metallb": map[string]any{
					"enabled": true,
					"speaker": map[string]any{
						"enabled": false,
					},
					"l2Advertisement": map[string]any{
						"enabled": false,
					},
					"ipAddressPool": []string{"10.10.10.0/24", "10.20.20.10-10.20.20.30"},
				},
			}))
		})
	})
})

func encode(obj runtime.Object) []byte {
	data, _ := json.Marshal(obj)
	return data
}
