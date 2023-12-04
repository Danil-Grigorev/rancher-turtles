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

package sync

import (
	"context"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rancher-sandbox/rancher-turtles/internal/test"

	managementv3 "github.com/rancher-sandbox/rancher-turtles/internal/rancher/management/v3"
	provisioningv1 "github.com/rancher-sandbox/rancher-turtles/internal/rancher/provisioning/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	testEnv *envtest.Environment
	cfg     *rest.Config
	cl      client.Client
	ctx     = context.Background()
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rancher Turtles sync logic")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("bootstrapping test environment")
	var err error
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "hack", "crd", "bases"),
		},
		ErrorIfCRDPathMissing: true,
		Scheme:                runtime.NewScheme(),
	}

	utilruntime.Must(clusterv1.AddToScheme(testEnv.Scheme))
	utilruntime.Must(provisioningv1.AddToScheme(testEnv.Scheme))
	utilruntime.Must(managementv3.AddToScheme(testEnv.Scheme))

	cfg, cl, err = test.StartEnvTest(testEnv)
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())
	Expect(cl).NotTo(BeNil())

})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	Expect(test.StopEnvTest(testEnv)).To(Succeed())
})
