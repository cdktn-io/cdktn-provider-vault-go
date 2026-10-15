// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaulttoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultTokenConfig struct {
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
	// The display name of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#display_name EphemeralVaultToken#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Name of the entity alias to associate with during token creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#entity_alias EphemeralVaultToken#entity_alias}
	EntityAlias *string `field:"optional" json:"entityAlias" yaml:"entityAlias"`
	// The explicit max TTL of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#explicit_max_ttl EphemeralVaultToken#explicit_max_ttl}
	ExplicitMaxTtl *string `field:"optional" json:"explicitMaxTtl" yaml:"explicitMaxTtl"`
	// The ID of the client token.
	//
	// This is an input field, not a resource identifier. Can only be specified by a root token. The ID provided may not contain a '.' character and should not start with the 's.' prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#id EphemeralVaultToken#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata to be associated with the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#metadata EphemeralVaultToken#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#mount_id EphemeralVaultToken#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#namespace EphemeralVaultToken#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Flag to disable the default policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#no_default_policy EphemeralVaultToken#no_default_policy}
	NoDefaultPolicy interface{} `field:"optional" json:"noDefaultPolicy" yaml:"noDefaultPolicy"`
	// Flag to create a token without parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#no_parent EphemeralVaultToken#no_parent}
	NoParent interface{} `field:"optional" json:"noParent" yaml:"noParent"`
	// The number of allowed uses of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#num_uses EphemeralVaultToken#num_uses}
	NumUses *float64 `field:"optional" json:"numUses" yaml:"numUses"`
	// The period of the token for periodic tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#period EphemeralVaultToken#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// List of policies to attach to the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#policies EphemeralVaultToken#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Flag to allow the token to be renewed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#renewable EphemeralVaultToken#renewable}
	Renewable interface{} `field:"optional" json:"renewable" yaml:"renewable"`
	// The token role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#role_name EphemeralVaultToken#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// The TTL period of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#ttl EphemeralVaultToken#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// The token type. Can be 'batch' or 'service'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#type EphemeralVaultToken#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The TTL period of the wrapped token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token#wrapping_ttl EphemeralVaultToken#wrapping_ttl}
	WrappingTtl *string `field:"optional" json:"wrappingTtl" yaml:"wrappingTtl"`
}

