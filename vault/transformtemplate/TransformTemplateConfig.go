// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransformTemplateConfig struct {
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
	// The name of the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#name TransformTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#path TransformTemplate#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The alphabet to use for this template. This is only used during FPE transformations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#alphabet TransformTemplate#alphabet}
	Alphabet *string `field:"optional" json:"alphabet" yaml:"alphabet"`
	// The map of regular expression templates used to customize decoded outputs. Only applicable to FPE transformations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#decode_formats TransformTemplate#decode_formats}
	DecodeFormats *map[string]*string `field:"optional" json:"decodeFormats" yaml:"decodeFormats"`
	// The regular expression template used for encoding values. Only applicable to FPE transformations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#encode_format TransformTemplate#encode_format}
	EncodeFormat *string `field:"optional" json:"encodeFormat" yaml:"encodeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#id TransformTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#namespace TransformTemplate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#pattern TransformTemplate#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/transform_template#type TransformTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

