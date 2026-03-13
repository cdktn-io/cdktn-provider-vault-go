// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesSecretBackendRoleConfig struct {
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
	// The mount path for the Kubernetes secrets engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#backend KubernetesSecretBackendRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#name KubernetesSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Kubernetes namespaces this role can generate credentials for.
	//
	// If set to '*' all namespaces are allowed. If set with`allowed_kubernetes_namespace_selector`, the conditions are `OR`ed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#allowed_kubernetes_namespaces KubernetesSecretBackendRole#allowed_kubernetes_namespaces}
	AllowedKubernetesNamespaces *[]*string `field:"optional" json:"allowedKubernetesNamespaces" yaml:"allowedKubernetesNamespaces"`
	// A label selector for Kubernetes namespaces in which credentials can begenerated.
	//
	// Accepts either a JSON or YAML object. The value should be of typeLabelSelector. If set with `allowed_kubernetes_namespace`, the conditions are `OR`ed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#allowed_kubernetes_namespace_selector KubernetesSecretBackendRole#allowed_kubernetes_namespace_selector}
	AllowedKubernetesNamespaceSelector *string `field:"optional" json:"allowedKubernetesNamespaceSelector" yaml:"allowedKubernetesNamespaceSelector"`
	// Additional annotations to apply to all generated Kubernetes objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#extra_annotations KubernetesSecretBackendRole#extra_annotations}
	ExtraAnnotations *map[string]*string `field:"optional" json:"extraAnnotations" yaml:"extraAnnotations"`
	// Additional labels to apply to all generated Kubernetes objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#extra_labels KubernetesSecretBackendRole#extra_labels}
	ExtraLabels *map[string]*string `field:"optional" json:"extraLabels" yaml:"extraLabels"`
	// The Role or ClusterRole rules to use when generating a role.
	//
	// Accepts either JSON or YAML formatted rules. Mutually exclusive with 'service_account_name' and 'kubernetes_role_name'. If set, the entire chain of Kubernetes objects will be generated when credentials are requested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#generated_role_rules KubernetesSecretBackendRole#generated_role_rules}
	GeneratedRoleRules *string `field:"optional" json:"generatedRoleRules" yaml:"generatedRoleRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#id KubernetesSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The pre-existing Role or ClusterRole to bind a generated service account to.
	//
	// Mutually exclusive with 'service_account_name' and 'generated_role_rules'. If set, Kubernetes token, service account, and role binding objects will be created when credentials are requested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#kubernetes_role_name KubernetesSecretBackendRole#kubernetes_role_name}
	KubernetesRoleName *string `field:"optional" json:"kubernetesRoleName" yaml:"kubernetesRoleName"`
	// Specifies whether the Kubernetes role is a Role or ClusterRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#kubernetes_role_type KubernetesSecretBackendRole#kubernetes_role_type}
	KubernetesRoleType *string `field:"optional" json:"kubernetesRoleType" yaml:"kubernetesRoleType"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#namespace KubernetesSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name template to use when generating service accounts, roles and role bindings.
	//
	// If unset, a default template is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#name_template KubernetesSecretBackendRole#name_template}
	NameTemplate *string `field:"optional" json:"nameTemplate" yaml:"nameTemplate"`
	// The pre-existing service account to generate tokens for.
	//
	// Mutually exclusive with 'kubernetes_role_name' and 'generated_role_rules'. If set, only a Kubernetes token will be created when credentials are requested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#service_account_name KubernetesSecretBackendRole#service_account_name}
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// The default audiences for generated Kubernetes tokens.
	//
	// If not set, defaults to the Kubernetes cluster's default audiences. Requires Vault 1.15+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#token_default_audiences KubernetesSecretBackendRole#token_default_audiences}
	TokenDefaultAudiences *[]*string `field:"optional" json:"tokenDefaultAudiences" yaml:"tokenDefaultAudiences"`
	// The default TTL for generated Kubernetes tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#token_default_ttl KubernetesSecretBackendRole#token_default_ttl}
	TokenDefaultTtl *float64 `field:"optional" json:"tokenDefaultTtl" yaml:"tokenDefaultTtl"`
	// The maximum TTL for generated Kubernetes tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kubernetes_secret_backend_role#token_max_ttl KubernetesSecretBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
}

