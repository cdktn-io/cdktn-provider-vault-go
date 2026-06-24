// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityEntityConfig struct {
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
	// Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#disabled IdentityEntity#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Manage policies externally through `vault_identity_entity_policies`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#external_policies IdentityEntity#external_policies}
	ExternalPolicies interface{} `field:"optional" json:"externalPolicies" yaml:"externalPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#id IdentityEntity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata to be associated with the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#metadata IdentityEntity#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Name of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#name IdentityEntity#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#namespace IdentityEntity#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Policies to be tied to the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity#policies IdentityEntity#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

