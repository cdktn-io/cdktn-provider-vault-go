// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitytpmgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityTpmGroupConfig struct {
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
	// Set of TPM IDs that are members of this TPM group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm_group#member_tpm_ids IdentityTpmGroup#member_tpm_ids}
	MemberTpmIds *[]*string `field:"optional" json:"memberTpmIds" yaml:"memberTpmIds"`
	// Metadata to associate with the TPM group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm_group#metadata IdentityTpmGroup#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Name of the TPM group. Vault generates a name if one is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm_group#name IdentityTpmGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm_group#namespace IdentityTpmGroup#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

