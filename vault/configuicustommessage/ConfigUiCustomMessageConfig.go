// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configuicustommessage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigUiCustomMessageConfig struct {
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
	// The base64-encoded content of the custom message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#message_base64 ConfigUiCustomMessage#message_base64}
	MessageBase64 *string `field:"required" json:"messageBase64" yaml:"messageBase64"`
	// The starting time of the active period of the custom message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#start_time ConfigUiCustomMessage#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// The title of the custom message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#title ConfigUiCustomMessage#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A flag indicating whether the custom message is displayed pre-login (false) or post-login (true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#authenticated ConfigUiCustomMessage#authenticated}
	Authenticated interface{} `field:"optional" json:"authenticated" yaml:"authenticated"`
	// The ending time of the active period of the custom message. Can be omitted for non-expiring message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#end_time ConfigUiCustomMessage#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#link ConfigUiCustomMessage#link}
	Link *ConfigUiCustomMessageLink `field:"optional" json:"link" yaml:"link"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#namespace ConfigUiCustomMessage#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A map containing additional options for the custom message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#options ConfigUiCustomMessage#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The display type of custom message. Allowed values are banner and modal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/config_ui_custom_message#type ConfigUiCustomMessage#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

