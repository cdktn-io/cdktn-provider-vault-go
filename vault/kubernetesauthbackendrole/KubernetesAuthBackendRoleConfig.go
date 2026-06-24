// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesAuthBackendRoleConfig struct {
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
	// List of service account names able to access this role.
	//
	// If set to `["*"]` all names are allowed, both this and bound_service_account_namespaces can not be "*".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#bound_service_account_names KubernetesAuthBackendRole#bound_service_account_names}
	BoundServiceAccountNames *[]*string `field:"required" json:"boundServiceAccountNames" yaml:"boundServiceAccountNames"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#role_name KubernetesAuthBackendRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#alias_metadata KubernetesAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Configures how identity aliases are generated. Valid choices are: serviceaccount_uid, serviceaccount_name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#alias_name_source KubernetesAuthBackendRole#alias_name_source}
	AliasNameSource *string `field:"optional" json:"aliasNameSource" yaml:"aliasNameSource"`
	// Optional Audience claim to verify in the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#audience KubernetesAuthBackendRole#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Unique name of the kubernetes backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#backend KubernetesAuthBackendRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// List of namespaces allowed to access this role.
	//
	// If set to `["*"]` all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#bound_service_account_namespaces KubernetesAuthBackendRole#bound_service_account_namespaces}
	BoundServiceAccountNamespaces *[]*string `field:"optional" json:"boundServiceAccountNamespaces" yaml:"boundServiceAccountNamespaces"`
	// A label selector for Kubernetes namespaces allowed to access this role.
	//
	// Accepts either a JSON or YAML object. The value should be of type LabelSelector. Currently, label selectors with matchExpressions are not supported. To use label selectors, Vault must have permission to read namespaces on the Kubernetes cluster. If set with bound_service_account_namespaces, the conditions are ORed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#bound_service_account_namespace_selector KubernetesAuthBackendRole#bound_service_account_namespace_selector}
	BoundServiceAccountNamespaceSelector *string `field:"optional" json:"boundServiceAccountNamespaceSelector" yaml:"boundServiceAccountNamespaceSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#id KubernetesAuthBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#namespace KubernetesAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_bound_cidrs KubernetesAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_explicit_max_ttl KubernetesAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_max_ttl KubernetesAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_no_default_policy KubernetesAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_num_uses KubernetesAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_period KubernetesAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_policies KubernetesAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_ttl KubernetesAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_auth_backend_role#token_type KubernetesAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

