//go:build e2e
// +build e2e

/*
Copyright © 2023 - 2024 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package import_gitops

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher/turtles/test/e2e"
	"github.com/rancher/turtles/test/testenv"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/klog/v2"
	capiframework "sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/cluster-api/test/framework/clusterctl"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Test suite flags.
var (
	flagVals *e2e.FlagValues
)

// Test suite global vars.
var (
	// e2eConfig to be used for this test, read from configPath.
	e2eConfig *clusterctl.E2EConfig

	// hostName is the host name for the Rancher Manager server.
	hostName string

	ctx = context.Background()

	setupClusterResult    *testenv.SetupTestClusterResult
	bootstrapClusterProxy capiframework.ClusterProxy
	gitAddress            string
)

func init() {
	flagVals = &e2e.FlagValues{}
	e2e.InitFlags(flagVals)
}

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)

	ctrl.SetLogger(klog.Background())

	RunSpecs(t, "rancher-turtles-e2e-import-gitops")
}

var _ = SynchronizedBeforeSuite(
	func() []byte {
		e2eConfig = e2e.LoadE2EConfig(flagVals.ConfigPath)

		setupClusterResult = testenv.SetupTestCluster(ctx, testenv.SetupTestClusterInput{
			E2EConfig: e2eConfig,
			Scheme:    e2e.InitScheme(),
		})

		testenv.RancherDeployIngress(ctx, testenv.RancherDeployIngressInput{
			BootstrapClusterProxy:    setupClusterResult.BootstrapClusterProxy,
			CustomIngress:            e2e.NginxIngress,
			DefaultIngressClassPatch: e2e.IngressClassPatch,
		})

		rancherHookResult := testenv.DeployRancher(ctx, testenv.DeployRancherInput{
			BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
			RancherPatches:        [][]byte{e2e.RancherSettingPatch},
		})

		if shortTestOnly() {
			testenv.DeployChartMuseum(ctx, testenv.DeployChartMuseumInput{
				BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
			})

			testenv.DeployRancherTurtles(ctx, testenv.DeployRancherTurtlesInput{
				BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
				CAPIProvidersYAML:     e2e.CapiProviders,
				AdditionalValues: map[string]string{
					"rancherTurtles.features.addon-provider-fleet.enabled": "true",
					"rancherTurtles.features.managementv3-cluster.enabled": "false", // disable the default management.cattle.io/v3 controller
				},
				WaitForDeployments: append(
					testenv.DefaultDeployments,
					testenv.NamespaceName{
						Name:      "caapf-controller-manager",
						Namespace: e2e.RancherTurtlesNamespace,
					},
				),
				ConfigurationPatches: [][]byte{e2e.AddonProviderFleetHostNetworkPatch},
			})
		} else {
			testenv.DeployRancherTurtles(ctx, testenv.DeployRancherTurtlesInput{
				BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
				CAPIProvidersYAML:     e2e.CapiProviders,
				AdditionalValues: map[string]string{
					"rancherTurtles.features.managementv3-cluster.enabled": "false", // disable the default management.cattle.io/v3 controller
				},
				WaitForDeployments: testenv.DefaultDeployments,
			})
		}

		if !shortTestOnly() && !localTestOnly() {
			By("Running full tests, deploying additional infrastructure providers")

			testenv.CAPIOperatorDeployProvider(ctx, testenv.CAPIOperatorDeployProviderInput{
				BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
				CAPIProvidersSecretsYAML: [][]byte{
					e2e.AWSProviderSecret,
					e2e.AzureIdentitySecret,
					e2e.GCPProviderSecret,
				},
				CAPIProvidersYAML: e2e.FullProviders,
				WaitForDeployments: []testenv.NamespaceName{
					{
						Name:      "capa-controller-manager",
						Namespace: "capa-system",
					},
					{
						Name:      "capz-controller-manager",
						Namespace: "capz-system",
					},
					{
						Name:      "capg-controller-manager",
						Namespace: "capg-system",
					},
				},
			})
		} else if Label(e2e.LocalTestLabel).MatchesLabelFilter(GinkgoLabelFilter()) {
			By("Running local vSphere tests, deploying vSphere infrastructure provider")

			testenv.CAPIOperatorDeployProvider(ctx, testenv.CAPIOperatorDeployProviderInput{
				BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
				CAPIProvidersSecretsYAML: [][]byte{
					e2e.VSphereProviderSecret,
				},
				CAPIProvidersYAML: e2e.CapvProvider,
				WaitForDeployments: []testenv.NamespaceName{
					{
						Name:      "capv-controller-manager",
						Namespace: "capv-system",
					},
				},
			})
		}

		giteaResult := testenv.DeployGitea(ctx, testenv.DeployGiteaInput{
			BootstrapClusterProxy: setupClusterResult.BootstrapClusterProxy,
			ValuesFile:            e2e.GiteaValues,
			CustomIngressConfig:   e2e.GiteaIngress,
		})

		data, err := json.Marshal(e2e.Setup{
			ClusterName:     setupClusterResult.ClusterName,
			KubeconfigPath:  setupClusterResult.KubeconfigPath,
			GitAddress:      giteaResult.GitAddress,
			E2EConfig:       e2eConfig,
			RancherHostname: rancherHookResult.Hostname,
		})
		Expect(err).ToNot(HaveOccurred())
		return data
	},
	func(sharedData []byte) {
		setup := e2e.Setup{}
		Expect(json.Unmarshal(sharedData, &setup)).To(Succeed())

		gitAddress = setup.GitAddress
		e2eConfig = setup.E2EConfig
		hostName = setup.RancherHostname

		bootstrapClusterProxy = capiframework.NewClusterProxy(setup.ClusterName, setup.KubeconfigPath, e2e.InitScheme(), capiframework.WithMachineLogCollector(capiframework.DockerLogCollector{}))
		Expect(bootstrapClusterProxy).ToNot(BeNil(), "cluster proxy should not be nil")
	},
)

var _ = SynchronizedAfterSuite(
	func() {
	},
	func() {
		testenv.UninstallGitea(ctx, testenv.UninstallGiteaInput{
			BootstrapClusterProxy: bootstrapClusterProxy,
		})

		testenv.UninstallRancherTurtles(ctx, testenv.UninstallRancherTurtlesInput{
			BootstrapClusterProxy: bootstrapClusterProxy,
		})

		testenv.CleanupTestCluster(ctx, testenv.CleanupTestClusterInput{
			SetupTestClusterResult: *setupClusterResult,
		})
	},
)

func shortTestOnly() bool {
	return GinkgoLabelFilter() == e2e.ShortTestLabel
}

func localTestOnly() bool {
	return GinkgoLabelFilter() == e2e.LocalTestLabel
}
