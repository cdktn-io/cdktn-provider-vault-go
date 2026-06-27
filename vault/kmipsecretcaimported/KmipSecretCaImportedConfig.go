// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmipsecretcaimported

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KmipSecretCaImportedConfig struct {
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
	// CA certificate in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#ca_pem KmipSecretCaImported#ca_pem}
	CaPem *string `field:"required" json:"caPem" yaml:"caPem"`
	// Name to identify the CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#name KmipSecretCaImported#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path where KMIP backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#path KmipSecretCaImported#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#namespace KmipSecretCaImported#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The field in the certificate to use for the role (CN, O, OU, or UID).
	//
	// Must specify exactly one of role_name or role_field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#role_field KmipSecretCaImported#role_field}
	RoleField *string `field:"optional" json:"roleField" yaml:"roleField"`
	// The role name to associate with this CA. Must specify exactly one of role_name or role_field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#role_name KmipSecretCaImported#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// The field in the certificate to use for the scope (CN, O, OU, or UID).
	//
	// Must specify exactly one of scope_name or scope_field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#scope_field KmipSecretCaImported#scope_field}
	ScopeField *string `field:"optional" json:"scopeField" yaml:"scopeField"`
	// The scope name to associate with this CA. Must specify exactly one of scope_name or scope_field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kmip_secret_ca_imported#scope_name KmipSecretCaImported#scope_name}
	ScopeName *string `field:"optional" json:"scopeName" yaml:"scopeName"`
}

