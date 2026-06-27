// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nomadsecretrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NomadSecretRoleConfig struct {
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
	// The mount path for the Nomad backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#backend NomadSecretRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#role NomadSecretRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Specifies if the token should be global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#global NomadSecretRole#global}
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#id NomadSecretRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#namespace NomadSecretRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Comma separated list of Nomad policies the token is going to be created against.
	//
	// These need to be created beforehand in Nomad.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#policies NomadSecretRole#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Specifies the type of token to create when using this role. Valid values are "client" or "management".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/nomad_secret_role#type NomadSecretRole#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

