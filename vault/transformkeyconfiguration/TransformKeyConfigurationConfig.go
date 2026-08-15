// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformkeyconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransformKeyConfigurationConfig struct {
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
	// The name of the transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#name TransformKeyConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The mount path for the transform backend, for example the path given in "$ vault secrets enable -path=transform transform".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#path TransformKeyConfiguration#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Amount of time the key should live before being automatically rotated.
	//
	// A value of 0 disables automatic rotation for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#auto_rotate_period TransformKeyConfiguration#auto_rotate_period}
	AutoRotatePeriod *float64 `field:"optional" json:"autoRotatePeriod" yaml:"autoRotatePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#id TransformKeyConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Minimum key version that vault uses to decode values for the transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#min_decryption_version TransformKeyConfiguration#min_decryption_version}
	MinDecryptionVersion *float64 `field:"optional" json:"minDecryptionVersion" yaml:"minDecryptionVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/resources/transform_key_configuration#namespace TransformKeyConfiguration#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

