// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configuidefaultauth

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigUiDefaultAuthConfig struct {
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
	// The default authentication method.
	//
	// Uses `OneOf` validator to ensure only valid auth methods are accepted: github, jwt, ldap, oidc, okta, radius, saml, token, userpass.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#default_auth_type ConfigUiDefaultAuth#default_auth_type}
	DefaultAuthType *string `field:"required" json:"defaultAuthType" yaml:"defaultAuthType"`
	// Unique identifier for the configuration.
	//
	// Can contain letters, numbers, underscores, and dashes. Uses `RequiresReplace()` plan modifier - changing this forces resource recreation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#name ConfigUiDefaultAuth#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of backup authentication methods.
	//
	// Uses `ListAttribute` with `ElementType: StringType` to preserve order of backup methods. Each must be a valid auth type. Vault presents these in the "Sign in with other methods" tab. **Note:** Removing this field from configuration will clear it in Vault by sending an empty array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#backup_auth_types ConfigUiDefaultAuth#backup_auth_types}
	BackupAuthTypes *[]*string `field:"optional" json:"backupAuthTypes" yaml:"backupAuthTypes"`
	// If true, child namespaces will not inherit default_auth_type and backup_auth_types from this configuration.
	//
	// **Note:** Removing this field from configuration will reset it to `false` in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#disable_inheritance ConfigUiDefaultAuth#disable_inheritance}
	DisableInheritance interface{} `field:"optional" json:"disableInheritance" yaml:"disableInheritance"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#namespace ConfigUiDefaultAuth#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Target namespace for the configuration. Empty string or omitted applies to root namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_ui_default_auth#namespace_path ConfigUiDefaultAuth#namespace_path}
	NamespacePath *string `field:"optional" json:"namespacePath" yaml:"namespacePath"`
}

