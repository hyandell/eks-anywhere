// +build e2e

package e2e

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runOIDCFlow(test *framework.E2ETest) {
	test.GenerateClusterConfig()
	test.CreateCluster()
	test.ValidateOIDC()
	test.StopIfFailed()
	test.DeleteCluster()
}

func TestDockerKubernetes120OIDC(t *testing.T) {
	test := framework.NewE2ETest(t,
		framework.NewDocker(t),
		framework.WithOIDC(),
		framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube120)),
	)
	runOIDCFlow(test)
}

func TestDockerKubernetes121OIDC(t *testing.T) {
	test := framework.NewE2ETest(t,
		framework.NewDocker(t),
		framework.WithOIDC(),
		framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube120)),
	)
	runOIDCFlow(test)
}

func TestVSphereKubernetes120OIDC(t *testing.T) {
	test := framework.NewE2ETest(
		t,
		framework.NewVSphere(t, framework.WithUbuntu120()),
		framework.WithOIDC(),
		framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube120)),
	)
	runOIDCFlow(test)
}

func TestVSphereKubernetes121OIDC(t *testing.T) {
	test := framework.NewE2ETest(
		t,
		framework.NewVSphere(t, framework.WithUbuntu121()),
		framework.WithOIDC(),
		framework.WithClusterFiller(api.WithKubernetesVersion(v1alpha1.Kube121)),
	)
	runOIDCFlow(test)
}