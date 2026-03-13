// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cfauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CfAuthBackendRoleConfig struct {
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
	// Mount path for the CF auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#mount CfAuthBackendRole#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the CF auth role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#name CfAuthBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A map of string to string that will be set as metadata on the identity alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#alias_metadata CfAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// An optional set of application IDs an instance must be a member of to qualify for this role.
	//
	// To clear this constraint, omit the field entirely rather than setting it to an empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#bound_application_ids CfAuthBackendRole#bound_application_ids}
	BoundApplicationIds *[]*string `field:"optional" json:"boundApplicationIds" yaml:"boundApplicationIds"`
	// An optional set of instance IDs an instance must be a member of to qualify for this role.
	//
	// To clear this constraint, omit the field entirely rather than setting it to an empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#bound_instance_ids CfAuthBackendRole#bound_instance_ids}
	BoundInstanceIds *[]*string `field:"optional" json:"boundInstanceIds" yaml:"boundInstanceIds"`
	// An optional set of organization IDs an instance must be a member of to qualify for this role.
	//
	// To clear this constraint, omit the field entirely rather than setting it to an empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#bound_organization_ids CfAuthBackendRole#bound_organization_ids}
	BoundOrganizationIds *[]*string `field:"optional" json:"boundOrganizationIds" yaml:"boundOrganizationIds"`
	// An optional set of space IDs an instance must be a member of to qualify for this role.
	//
	// To clear this constraint, omit the field entirely rather than setting it to an empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#bound_space_ids CfAuthBackendRole#bound_space_ids}
	BoundSpaceIds *[]*string `field:"optional" json:"boundSpaceIds" yaml:"boundSpaceIds"`
	// If set to `true`, disables the default behavior that logging in must be performed from an acceptable IP address described by the presented certificate.
	//
	// Defaults to `false`. To reset to the default, omit this field from config rather than setting it to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#disable_ip_matching CfAuthBackendRole#disable_ip_matching}
	DisableIpMatching interface{} `field:"optional" json:"disableIpMatching" yaml:"disableIpMatching"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#namespace CfAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_bound_cidrs CfAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_explicit_max_ttl CfAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_max_ttl CfAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_no_default_policy CfAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_num_uses CfAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_period CfAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_policies CfAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_ttl CfAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/cf_auth_backend_role#token_type CfAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

