// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultkubernetesserviceaccounttoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultKubernetesServiceAccountTokenConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Kubernetes secret backend to generate service account tokens from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#backend DataVaultKubernetesServiceAccountToken#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The name of the Kubernetes namespace in which to generate the credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#kubernetes_namespace DataVaultKubernetesServiceAccountToken#kubernetes_namespace}
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#role DataVaultKubernetesServiceAccountToken#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// If true, generate a ClusterRoleBinding to grant permissions across the whole cluster instead of within a namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#cluster_role_binding DataVaultKubernetesServiceAccountToken#cluster_role_binding}
	ClusterRoleBinding interface{} `field:"optional" json:"clusterRoleBinding" yaml:"clusterRoleBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#id DataVaultKubernetesServiceAccountToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#namespace DataVaultKubernetesServiceAccountToken#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The TTL of the generated Kubernetes service account token, specified in seconds or as a Go duration format string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/kubernetes_service_account_token#ttl DataVaultKubernetesServiceAccountToken#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

