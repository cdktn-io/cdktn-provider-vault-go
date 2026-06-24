// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitcmac

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultTransitCmacConfig struct {
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
	// Name of the CMAC key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#name DataVaultTransitCmac#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Transit secret backend the key belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#path DataVaultTransitCmac#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Specifies a list of items for processing.
	//
	// When this parameter is set, any supplied 'input' or 'context' parameters will be ignored. Responses are returned in the 'batch_results' array component of the 'data' element of the response. Any batch output will preserve the order of the batch input. If the input data value of an item is invalid, the corresponding item in the 'batch_results' will have the key 'error' with a value describing the error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#batch_input DataVaultTransitCmac#batch_input}
	BatchInput interface{} `field:"optional" json:"batchInput" yaml:"batchInput"`
	// The results returned from Vault if using batch_input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#batch_results DataVaultTransitCmac#batch_results}
	BatchResults interface{} `field:"optional" json:"batchResults" yaml:"batchResults"`
	// The CMAC returned from Vault if using input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#cmac DataVaultTransitCmac#cmac}
	Cmac *string `field:"optional" json:"cmac" yaml:"cmac"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#id DataVaultTransitCmac#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the base64 encoded input data. One of input or batch_input must be supplied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#input DataVaultTransitCmac#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The version of the key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#key_version DataVaultTransitCmac#key_version}
	KeyVersion *float64 `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// Specifies the MAC length to use (POST body parameter). The mac_length cannot be larger than the cipher's block size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#mac_length DataVaultTransitCmac#mac_length}
	MacLength *float64 `field:"optional" json:"macLength" yaml:"macLength"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#namespace DataVaultTransitCmac#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the MAC length to use (URL parameter).
	//
	// If provided, this value overrides mac_length. The url_mac_length cannot be larger than the cipher's block size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac#url_mac_length DataVaultTransitCmac#url_mac_length}
	UrlMacLength *float64 `field:"optional" json:"urlMacLength" yaml:"urlMacLength"`
}

