// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapsecretbackenddynamicrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LdapSecretBackendDynamicRoleConfig struct {
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
	// A templatized LDIF string used to create a user account. May contain multiple entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#creation_ldif LdapSecretBackendDynamicRole#creation_ldif}
	CreationLdif *string `field:"required" json:"creationLdif" yaml:"creationLdif"`
	// A templatized LDIF string used to delete the user account once its TTL has expired.
	//
	// This may contain multiple LDIF entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#deletion_ldif LdapSecretBackendDynamicRole#deletion_ldif}
	DeletionLdif *string `field:"required" json:"deletionLdif" yaml:"deletionLdif"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#role_name LdapSecretBackendDynamicRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Specifies the TTL for the leases associated with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#default_ttl LdapSecretBackendDynamicRole#default_ttl}
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#id LdapSecretBackendDynamicRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum TTL for the leases associated with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#max_ttl LdapSecretBackendDynamicRole#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The path where the LDAP secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#mount LdapSecretBackendDynamicRole#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#namespace LdapSecretBackendDynamicRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A templatized LDIF string used to attempt to rollback any changes in the event that execution of the creation_ldif results in an error.
	//
	// This may contain multiple LDIF entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#rollback_ldif LdapSecretBackendDynamicRole#rollback_ldif}
	RollbackLdif *string `field:"optional" json:"rollbackLdif" yaml:"rollbackLdif"`
	// A template used to generate a dynamic username.
	//
	// This will be used to fill in the .Username field within the creation_ldif string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_dynamic_role#username_template LdapSecretBackendDynamicRole#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

