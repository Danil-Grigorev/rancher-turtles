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

package framework

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/cluster-api/test/framework"
	capiexec "sigs.k8s.io/cluster-api/test/framework/exec"
)

// Apply wraps `kubectl apply ...` and prints the output so we can see what gets applied to the cluster.
func Apply(ctx context.Context, p framework.ClusterProxy, resources []byte, args ...string) error {
	Expect(ctx).NotTo(BeNil(), "ctx is required for Apply")
	Expect(resources).NotTo(BeNil(), "resources is required for Apply")

	if err := KubectlApply(ctx, p.GetKubeconfigPath(), resources, args...); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return errors.New(fmt.Sprintf("%s: stderr: %s", err.Error(), exitErr.Stderr))
		}
		return err
	}

	return nil
}

// KubectlApply shells out to kubectl apply.
//
// TODO: Remove this usage of kubectl and replace with a function from apply.go using the controller-runtime client.
func KubectlApply(ctx context.Context, kubeconfigPath string, resources []byte, args ...string) error {
	aargs := append([]string{"apply", "--kubeconfig", kubeconfigPath, "-f", "-"}, args...)
	rbytes := bytes.NewReader(resources)
	applyCmd := capiexec.NewCommand(
		capiexec.WithCommand(kubectlPath()),
		capiexec.WithArgs(aargs...),
		capiexec.WithStdin(rbytes),
	)
	fmt.Printf("Running kubectl %s\n", strings.Join(aargs, " "))
	stdout, stderr, err := applyCmd.Run(ctx)
	if len(stderr) > 0 {
		fmt.Printf("stderr:\n%s\n", string(stderr))
	}
	if len(stdout) > 0 {
		fmt.Printf("stdout:\n%s\n", string(stdout))
	}
	return err
}

func kubectlPath() string {
	if kubectlPath, ok := os.LookupEnv("CAPI_KUBECTL_PATH"); ok {
		return kubectlPath
	}
	return "kubectl"
}