//go:build e2e
// +build e2e

/*
Copyright 2023 SUSE.

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

package e2e

import (
	"fmt"
	"path/filepath"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher-sandbox/rancher-turtles/internal/rancher"
	"sigs.k8s.io/cluster-api/controllers/remote"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	azureClusterName      = "test-azure"
	azureClusterNamespace = "default"
)

var _ = Describe("Import Azure CAPI cluster functionality should work", func() {
	var (
		rancherClusterHandler *rancher.ClusterHandler
		rancherCluster        *rancher.Cluster
		capiCluster           *clusterv1.Cluster
		rancherClusterKey     client.ObjectKey
		testCount             int
	)

	BeforeEach(func() {
		testCount += 1
		rancherClusterHandler = rancher.NewClusterHandler(ctx, bootstrapClusterProxy.GetClient())
		capiCluster = &clusterv1.Cluster{ObjectMeta: metav1.ObjectMeta{
			Namespace: azureClusterNamespace,
			Name:      azureClusterName,
		}}
		rancherCluster = &rancher.Cluster{ObjectMeta: metav1.ObjectMeta{
			Namespace: azureClusterNamespace,
			Name:      fmt.Sprintf("%s-capi", azureClusterName),
		}}
		rancherClusterKey = client.ObjectKey{
			Namespace: rancherCluster.Namespace,
			Name:      rancherCluster.Name,
		}

		By("Creating an Azure CAPI cluster with calico CNI")
		Eventually(func() error {
			return bootstrapClusterProxy.Apply(ctx, azureCAPITestdata)
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-controllers")...).Should(Succeed())
	})

	AfterEach(func() {
		By("Dumping child cluster logs")
		bootstrapClusterProxy.CollectWorkloadClusterLogs(ctx, capiCluster.Namespace, capiCluster.Name, filepath.Join(artifactFolder, "clusters", capiCluster.Name, strconv.Itoa(testCount)))

		By("Removing CAPI cluster record")
		Eventually(func() bool {
			return apierrors.IsNotFound(bootstrapClusterProxy.GetClient().Delete(ctx, capiCluster))
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-azure")...).Should(BeTrue())

		By("Waiting for the rancher cluster record to be removed")
		Eventually(func() bool {
			_, err := rancherClusterHandler.Get(rancherClusterKey)
			return apierrors.IsNotFound(err)
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-azure")...).Should(BeTrue())
	})

	It("should successfully create a rancher cluster from a CAPI cluster", func() {
		By("Waiting for the CAPI cluster to be connectable")
		Eventually(func() error {
			remoteClient, err := remote.NewClusterClient(ctx, capiCluster.Name, bootstrapClusterProxy.GetClient(), client.ObjectKeyFromObject(capiCluster))
			if err != nil {
				return err
			}
			namespaces := &corev1.NamespaceList{}
			return remoteClient.List(ctx, namespaces)
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-rancher")...).Should(Succeed())

		By("Waiting for the rancher cluster record to appear")
		Eventually(func() error {
			_, err := rancherClusterHandler.Get(rancherClusterKey)
			return err
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-rancher")...).Should(Succeed())

		By("Waiting for the rancher cluster to have a deployed agent")
		Eventually(func() bool {
			cluster, err := rancherClusterHandler.Get(rancherClusterKey)
			return err == nil && cluster.Status.AgentDeployed == true
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-rancher")...).Should(BeTrue())

		By("Waiting for the rancher cluster to be ready")
		Eventually(func() bool {
			cluster, err := rancherClusterHandler.Get(rancherClusterKey)
			return err == nil && cluster.Status.Ready == true
		}, e2eConfig.GetIntervals(bootstrapClusterProxy.GetName(), "wait-rancher")...).Should(BeTrue())
	})

})
