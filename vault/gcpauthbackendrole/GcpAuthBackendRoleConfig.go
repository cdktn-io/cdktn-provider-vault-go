// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpAuthBackendRoleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#role GcpAuthBackendRole#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#type GcpAuthBackendRole#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#add_group_aliases GcpAuthBackendRole#add_group_aliases}.
	AddGroupAliases interface{} `field:"optional" json:"addGroupAliases" yaml:"addGroupAliases"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#alias_metadata GcpAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#allow_gce_inference GcpAuthBackendRole#allow_gce_inference}.
	AllowGceInference interface{} `field:"optional" json:"allowGceInference" yaml:"allowGceInference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#backend GcpAuthBackendRole#backend}.
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_instance_groups GcpAuthBackendRole#bound_instance_groups}.
	BoundInstanceGroups *[]*string `field:"optional" json:"boundInstanceGroups" yaml:"boundInstanceGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_labels GcpAuthBackendRole#bound_labels}.
	BoundLabels *[]*string `field:"optional" json:"boundLabels" yaml:"boundLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_projects GcpAuthBackendRole#bound_projects}.
	BoundProjects *[]*string `field:"optional" json:"boundProjects" yaml:"boundProjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_regions GcpAuthBackendRole#bound_regions}.
	BoundRegions *[]*string `field:"optional" json:"boundRegions" yaml:"boundRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_service_accounts GcpAuthBackendRole#bound_service_accounts}.
	BoundServiceAccounts *[]*string `field:"optional" json:"boundServiceAccounts" yaml:"boundServiceAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#bound_zones GcpAuthBackendRole#bound_zones}.
	BoundZones *[]*string `field:"optional" json:"boundZones" yaml:"boundZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#id GcpAuthBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#max_jwt_exp GcpAuthBackendRole#max_jwt_exp}.
	MaxJwtExp *string `field:"optional" json:"maxJwtExp" yaml:"maxJwtExp"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#namespace GcpAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The role_id is the stable, unique identifier for the role generated by vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#role_id GcpAuthBackendRole#role_id}
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_bound_cidrs GcpAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_explicit_max_ttl GcpAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_max_ttl GcpAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_no_default_policy GcpAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_num_uses GcpAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_period GcpAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_policies GcpAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_ttl GcpAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_auth_backend_role#token_type GcpAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

