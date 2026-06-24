// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configcontrolgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigControlGroupConfig struct {
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
	// The maximum ttl for a control group wrapping token.
	//
	// This can be provided in seconds or duration (for example, 2h).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/config_control_group#max_ttl ConfigControlGroup#max_ttl}
	MaxTtl *string `field:"required" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/config_control_group#namespace ConfigControlGroup#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

