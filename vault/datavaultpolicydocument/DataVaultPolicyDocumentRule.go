// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpolicydocument


type DataVaultPolicyDocumentRule struct {
	// A list of capabilities to apply to the specified path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#capabilities DataVaultPolicyDocument#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// A path in Vault that this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#path DataVaultPolicyDocument#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// allowed_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#allowed_parameter DataVaultPolicyDocument#allowed_parameter}
	AllowedParameter interface{} `field:"optional" json:"allowedParameter" yaml:"allowedParameter"`
	// denied_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#denied_parameter DataVaultPolicyDocument#denied_parameter}
	DeniedParameter interface{} `field:"optional" json:"deniedParameter" yaml:"deniedParameter"`
	// Description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#description DataVaultPolicyDocument#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum allowed TTL that clients can specify for a wrapped response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#max_wrapping_ttl DataVaultPolicyDocument#max_wrapping_ttl}
	MaxWrappingTtl *string `field:"optional" json:"maxWrappingTtl" yaml:"maxWrappingTtl"`
	// The minimum allowed TTL that clients can specify for a wrapped response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#min_wrapping_ttl DataVaultPolicyDocument#min_wrapping_ttl}
	MinWrappingTtl *string `field:"optional" json:"minWrappingTtl" yaml:"minWrappingTtl"`
	// A list of parameters that must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#required_parameters DataVaultPolicyDocument#required_parameters}
	RequiredParameters *[]*string `field:"optional" json:"requiredParameters" yaml:"requiredParameters"`
	// A list of event types to subscribe to when using `subscribe` capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/policy_document#subscribe_event_types DataVaultPolicyDocument#subscribe_event_types}
	SubscribeEventTypes *[]*string `field:"optional" json:"subscribeEventTypes" yaml:"subscribeEventTypes"`
}

