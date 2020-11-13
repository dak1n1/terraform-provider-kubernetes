package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"strings"
	"testing"

	gversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var testAccProvider *schema.Provider
var testAccExternalProviders map[string]resource.ExternalProvider
var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"kubernetes": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func init() {
	testAccProvider = Provider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"kubernetes": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
	testAccExternalProviders = map[string]resource.ExternalProvider{
		"kubernetes-local": {
			VersionConstraint: "9.9.9",
			Source:            "localhost/test/kubernetes",
		},
		"kubernetes-released": {
			VersionConstraint: "~> 1.13.2",
			Source:            "hashicorp/kubernetes",
		},
		"aws": {
			Source: "hashicorp/aws",
		},
		"google": {
			Source: "hashicorp/google",
		},
		"azurerm": {
			Source: "hashicorp/azurerm",
		},
	}
}

func TestProvider(t *testing.T) {
	provider := Provider()
	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ schema.Provider = *Provider()
}

func testAccPreCheck(t *testing.T) {
	ctx := context.TODO()
	hasFileCfg := (os.Getenv("KUBE_CTX_AUTH_INFO") != "" && os.Getenv("KUBE_CTX_CLUSTER") != "") ||
		os.Getenv("KUBE_CTX") != "" ||
		os.Getenv("KUBECONFIG") != "" ||
		os.Getenv("KUBE_CONFIG") != ""
	hasUserCredentials := os.Getenv("KUBE_USER") != "" && os.Getenv("KUBE_PASSWORD") != ""
	hasClientCert := os.Getenv("KUBE_CLIENT_CERT_DATA") != "" && os.Getenv("KUBE_CLIENT_KEY_DATA") != ""
	hasStaticCfg := (os.Getenv("KUBE_HOST") != "" &&
		os.Getenv("KUBE_CLUSTER_CA_CERT_DATA") != "") &&
		(hasUserCredentials || hasClientCert || os.Getenv("KUBE_TOKEN") != "")

	if !hasFileCfg && !hasStaticCfg && !hasUserCredentials {
		t.Fatalf("File config (KUBE_CTX_AUTH_INFO and KUBE_CTX_CLUSTER) or static configuration"+
			"(%s) or (%s) must be set for acceptance tests",
			strings.Join([]string{
				"KUBE_HOST",
				"KUBE_USER",
				"KUBE_PASSWORD",
				"KUBE_CLUSTER_CA_CERT_DATA",
			}, ", "),
			strings.Join([]string{
				"KUBE_HOST",
				"KUBE_CLIENT_CERT_DATA",
				"KUBE_CLIENT_KEY_DATA",
				"KUBE_CLUSTER_CA_CERT_DATA",
			}, ", "),
		)
	}

	diags := testAccProvider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if diags.HasError() {
		t.Fatal(diags[0].Summary)
	}
	return
}

func getClusterVersion() (*gversion.Version, error) {
	meta := testAccProvider.Meta()

	if meta == nil {
		return nil, fmt.Errorf("Provider not initialized, unable to check cluster version")
	}

	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return nil, err
	}
	serverVersion, err := conn.ServerVersion()

	if err != nil {
		return nil, err
	}

	return gversion.NewVersion(serverVersion.String())
}

func skipIfClusterVersionLessThan(t *testing.T, vs string) {
	if clusterVersionLessThan(vs) {
		t.Skip(fmt.Sprintf("This test will only run on cluster versions %v and above", vs))
	}
}

func skipIfNoLoadBalancersAvailable(t *testing.T) {
	isInGke, err := isRunningInGke()
	if err != nil {
		t.Fatal(err)
	}
	isInEks, err := isRunningInEks()
	if err != nil {
		t.Fatal(err)
	}
	if !isInGke && !isInEks {
		t.Skip("The Kubernetes endpoint must come from an environment which supports " +
			"load balancer provisioning for this test to run - skipping")
	}
}

func skipIfNotRunningInGke(t *testing.T) {
	isInGke, err := isRunningInGke()
	if err != nil {
		t.Fatal(err)
	}
	if !isInGke {
		t.Skip("The Kubernetes endpoint must come from GKE for this test to run - skipping")
	}
	if os.Getenv("GOOGLE_PROJECT") == "" || os.Getenv("GOOGLE_REGION") == "" || os.Getenv("GOOGLE_ZONE") == "" {
		t.Fatal("GOOGLE_PROJECT, GOOGLE_REGION, and GOOGLE_ZONE must be set for GoogleCloud tests")
	}
}

func skipIfNotRunningInAks(t *testing.T) {
	isInAks, err := isRunningInAks()
	if err != nil {
		t.Fatal(err)
	}
	if !isInAks {
		t.Skip("The Kubernetes endpoint must come from AKS for this test to run - skipping")
	}
	location := os.Getenv("TF_VAR_location")
	subscription := os.Getenv("ARM_SUBSCRIPTION_ID")
	if location == "" || subscription == "" {
		t.Fatal("TF_VAR_location and ARM_SUBSCRIPTION_ID must be set for Azure tests")
	}
}

func skipIfNotRunningInEks(t *testing.T) {
	isInEks, err := isRunningInEks()
	if err != nil {
		t.Fatal(err)
	}
	if !isInEks {
		t.Skip("The Kubernetes endpoint must come from EKS for this test to run - skipping")
	}
	if os.Getenv("AWS_DEFAULT_REGION") == "" || os.Getenv("AWS_ZONE") == "" || os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		t.Fatal("AWS_DEFAULT_REGION, AWS_ZONE, AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set for AWS tests")
	}
}

func skipIfNotRunningInMinikube(t *testing.T) {
	isInMinikube, err := isRunningInMinikube()
	if err != nil {
		t.Fatal(err)
	}
	if !isInMinikube {
		t.Skip("The Kubernetes endpoint must come from Minikube for this test to run - skipping")
	}
}

func skipIfRunningInMinikube(t *testing.T) {
	isInMinikube, err := isRunningInMinikube()
	if err != nil {
		t.Fatal(err)
	}
	if isInMinikube {
		t.Skip("This test requires multiple Kubernetes nodes - skipping")
	}
}

func skipIfUnsupportedSecurityContextRunAsGroup(t *testing.T) {
	skipIfClusterVersionLessThan(t, "1.14.0")
}

func isRunningInMinikube() (bool, error) {
	node, err := getFirstNode()
	if err != nil {
		return false, err
	}

	labels := node.GetLabels()
	if v, ok := labels["kubernetes.io/hostname"]; ok && v == "minikube" {
		return true, nil
	}
	return false, nil
}

func isRunningInGke() (bool, error) {
	node, err := getFirstNode()
	if err != nil {
		return false, err
	}

	labels := node.GetLabels()
	if _, ok := labels["cloud.google.com/gke-nodepool"]; ok {
		return true, nil
	}
	return false, nil
}

func isRunningInEks() (bool, error) {
	// EKS nodes don't have any unique labels, so check for the AWS
	// specific config map created by our test-infra.
	meta := testAccProvider.Meta()
	if meta == nil {
		return false, errors.New("Provider not initialized, unable to fetch provider metadata")
	}
	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return false, err
	}
	ctx := context.TODO()
	_, err = conn.CoreV1().ConfigMaps("kube-system").Get(ctx, "aws-auth", metav1.GetOptions{})
	if err != nil {
		return false, nil
	}
	return true, nil
}

func isRunningInAks() (bool, error) {
	node, err := getFirstNode()
	if err != nil {
		return false, err
	}

	labels := node.GetLabels()
	if _, ok := labels["kubernetes.azure.com/cluster"]; ok {
		return true, nil
	}
	return false, nil
}

func getFirstNode() (api.Node, error) {
	meta := testAccProvider.Meta()
	if meta == nil {
		return api.Node{}, errors.New("Provider not initialized, unable to get cluster node")
	}
	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return api.Node{}, err
	}
	ctx := context.TODO()

	resp, err := conn.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return api.Node{}, err
	}

	if len(resp.Items) < 1 {
		return api.Node{}, errors.New("Expected at least 1 node, none found")
	}

	return resp.Items[0], nil
}

func clusterVersionLessThan(vs string) bool {
	cv, err := getClusterVersion()

	if err != nil {
		return false
	}

	v, err := gversion.NewVersion(vs)

	if err != nil {
		return false
	}

	return cv.LessThan(v)
}

func requiredProviders() string {
	return fmt.Sprintf(`terraform {
  required_providers {
    kubernetes-local = {
      source  = "localhost/test/kubernetes"
      version = "9.9.9"
    }
    kubernetes-released = {
      source  = "hashicorp/kubernetes"
      version = "~> 1.13.2"
    }
  }
}
`)
}
