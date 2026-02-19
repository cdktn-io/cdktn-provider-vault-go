// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package genericendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GenericEndpointConfig struct {
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
	// JSON-encoded data to write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#data_json GenericEndpoint#data_json}
	DataJson *string `field:"required" json:"dataJson" yaml:"dataJson"`
	// Full path where to the endpoint that will be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#path GenericEndpoint#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Don't attempt to delete the path from Vault if true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#disable_delete GenericEndpoint#disable_delete}
	DisableDelete interface{} `field:"optional" json:"disableDelete" yaml:"disableDelete"`
	// Don't attempt to read the path from Vault if true; drift won't be detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#disable_read GenericEndpoint#disable_read}
	DisableRead interface{} `field:"optional" json:"disableRead" yaml:"disableRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#id GenericEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When reading, disregard fields not present in data_json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#ignore_absent_fields GenericEndpoint#ignore_absent_fields}
	IgnoreAbsentFields interface{} `field:"optional" json:"ignoreAbsentFields" yaml:"ignoreAbsentFields"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#namespace GenericEndpoint#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Top-level fields returned by write to persist in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/generic_endpoint#write_fields GenericEndpoint#write_fields}
	WriteFields *[]*string `field:"optional" json:"writeFields" yaml:"writeFields"`
}

