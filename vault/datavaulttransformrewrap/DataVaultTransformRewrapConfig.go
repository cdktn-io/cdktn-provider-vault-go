// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransformrewrap

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultTransformRewrapConfig struct {
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
	// Path to backend from which to retrieve data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#path DataVaultTransformRewrap#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#role_name DataVaultTransformRewrap#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Specifies a list of items to be re-encoded in a single batch.
	//
	// If this parameter is set, the top-level parameters 'value', 'transformation', 'decode_transformation', 'tweak', and 'decode_tweak' will be ignored. Each batch item within the list can specify these parameters instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#batch_input DataVaultTransformRewrap#batch_input}
	BatchInput interface{} `field:"optional" json:"batchInput" yaml:"batchInput"`
	// The FPE transformation to use to decode the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#decode_transformation DataVaultTransformRewrap#decode_transformation}
	DecodeTransformation *string `field:"optional" json:"decodeTransformation" yaml:"decodeTransformation"`
	// The tweak value to use for decoding. Only applicable for FPE transformations with a supplied tweak source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#decode_tweak DataVaultTransformRewrap#decode_tweak}
	DecodeTweak *string `field:"optional" json:"decodeTweak" yaml:"decodeTweak"`
	// The result of rewrapping a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#encoded_value DataVaultTransformRewrap#encoded_value}
	EncodedValue *string `field:"optional" json:"encodedValue" yaml:"encodedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#id DataVaultTransformRewrap#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#namespace DataVaultTransformRewrap#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The FPE transformation to use for re-encrypting the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#transformation DataVaultTransformRewrap#transformation}
	Transformation *string `field:"optional" json:"transformation" yaml:"transformation"`
	// The tweak value to use for re-encrypting. Only applicable for FPE transformations with a supplied tweak source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#tweak DataVaultTransformRewrap#tweak}
	Tweak *string `field:"optional" json:"tweak" yaml:"tweak"`
	// The value to rewrap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap#value DataVaultTransformRewrap#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

