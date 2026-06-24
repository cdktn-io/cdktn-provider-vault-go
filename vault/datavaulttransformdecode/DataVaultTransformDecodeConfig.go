// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransformdecode

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultTransformDecodeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#path DataVaultTransformDecode#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#role_name DataVaultTransformDecode#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Specifies a list of items to be decoded in a single batch.
	//
	// If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#batch_input DataVaultTransformDecode#batch_input}
	BatchInput interface{} `field:"optional" json:"batchInput" yaml:"batchInput"`
	// The result of decoding batch_input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#batch_results DataVaultTransformDecode#batch_results}
	BatchResults interface{} `field:"optional" json:"batchResults" yaml:"batchResults"`
	// The result of decoding a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#decoded_value DataVaultTransformDecode#decoded_value}
	DecodedValue *string `field:"optional" json:"decodedValue" yaml:"decodedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#id DataVaultTransformDecode#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#namespace DataVaultTransformDecode#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The transformation to perform.
	//
	// If no value is provided and the role contains a single transformation, this value will be inferred from the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#transformation DataVaultTransformDecode#transformation}
	Transformation *string `field:"optional" json:"transformation" yaml:"transformation"`
	// The tweak value to use. Only applicable for FPE transformations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#tweak DataVaultTransformDecode#tweak}
	Tweak *string `field:"optional" json:"tweak" yaml:"tweak"`
	// The value in which to decode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transform_decode#value DataVaultTransformDecode#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

