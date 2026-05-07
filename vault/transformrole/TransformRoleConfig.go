// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransformRoleConfig struct {
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
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/transform_role#name TransformRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/transform_role#path TransformRole#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/transform_role#id TransformRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/transform_role#namespace TransformRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A comma separated string or slice of transformations to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/transform_role#transformations TransformRole#transformations}
	Transformations *[]*string `field:"optional" json:"transformations" yaml:"transformations"`
}

