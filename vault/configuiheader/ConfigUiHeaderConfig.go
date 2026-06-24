// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configuiheader

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigUiHeaderConfig struct {
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
	// The name of the custom header. Cannot start with `X-Vault-`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/config_ui_header#name ConfigUiHeader#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set of values for the header. At least one value is required. Duplicates are automatically ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/config_ui_header#values ConfigUiHeader#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

