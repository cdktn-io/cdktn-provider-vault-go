// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapsecretbackendstaticrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LdapSecretBackendStaticRoleConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#role_name LdapSecretBackendStaticRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The username of the existing LDAP entry to manage password rotation for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#username LdapSecretBackendStaticRole#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#disable_automated_rotation LdapSecretBackendStaticRole#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Distinguished name (DN) of the existing LDAP entry to manage password rotation for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#dn LdapSecretBackendStaticRole#dn}
	Dn *string `field:"optional" json:"dn" yaml:"dn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#id LdapSecretBackendStaticRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The path where the LDAP secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#mount LdapSecretBackendStaticRole#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#namespace LdapSecretBackendStaticRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Password for the static role. This is required for Vault to manage an existing account and enable rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#password_wo LdapSecretBackendStaticRole#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Version counter for write-only password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#password_wo_version LdapSecretBackendStaticRole#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#rotation_period LdapSecretBackendStaticRole#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The rotation policy to use for this credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#rotation_policy LdapSecretBackendStaticRole#rotation_policy}
	RotationPolicy *string `field:"optional" json:"rotationPolicy" yaml:"rotationPolicy"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#rotation_schedule LdapSecretBackendStaticRole#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#rotation_window LdapSecretBackendStaticRole#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Skip rotation of the password on import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend_static_role#skip_import_rotation LdapSecretBackendStaticRole#skip_import_rotation}
	SkipImportRotation interface{} `field:"optional" json:"skipImportRotation" yaml:"skipImportRotation"`
}

