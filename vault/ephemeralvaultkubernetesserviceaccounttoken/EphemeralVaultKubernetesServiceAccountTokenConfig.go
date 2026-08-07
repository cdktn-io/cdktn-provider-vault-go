// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultkubernetesserviceaccounttoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultKubernetesServiceAccountTokenConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// The Kubernetes secret backend to generate service account tokens from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#backend EphemeralVaultKubernetesServiceAccountToken#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The name of the Kubernetes namespace in which to generate the credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#kubernetes_namespace EphemeralVaultKubernetesServiceAccountToken#kubernetes_namespace}
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#role EphemeralVaultKubernetesServiceAccountToken#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#cluster_role_binding EphemeralVaultKubernetesServiceAccountToken#cluster_role_binding}
	ClusterRoleBinding interface{} `field:"optional" json:"clusterRoleBinding" yaml:"clusterRoleBinding"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#mount_id EphemeralVaultKubernetesServiceAccountToken#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#namespace EphemeralVaultKubernetesServiceAccountToken#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The TTL of the generated Kubernetes service account token, specified in seconds or as a Go duration format string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/kubernetes_service_account_token#ttl EphemeralVaultKubernetesServiceAccountToken#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

