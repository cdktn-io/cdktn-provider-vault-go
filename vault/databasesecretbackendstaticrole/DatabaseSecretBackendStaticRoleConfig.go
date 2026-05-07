// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendstaticrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretBackendStaticRoleConfig struct {
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
	// The path of the Database Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#backend DatabaseSecretBackendStaticRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Database connection to use for this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#db_name DatabaseSecretBackendStaticRole#db_name}
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// Unique name for the static role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#name DatabaseSecretBackendStaticRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The database username that this role corresponds to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#username DatabaseSecretBackendStaticRole#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The configuration for the credential type.Full documentation for the allowed values can be found under "https://developer.hashicorp.com/vault/api-docs/secret/databases#credential_config".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#credential_config DatabaseSecretBackendStaticRole#credential_config}
	CredentialConfig *map[string]*string `field:"optional" json:"credentialConfig" yaml:"credentialConfig"`
	// The credential type for the user, can be one of "password", "rsa_private_key" or "client_certificate".The configuration can be done in `credential_config`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#credential_type DatabaseSecretBackendStaticRole#credential_type}
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#id DatabaseSecretBackendStaticRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#namespace DatabaseSecretBackendStaticRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The password corresponding to the username in the database.
	//
	// This is a write-only field. Requires Vault 1.19+. Deprecates 'self_managed_password' which was introduced in Vault 1.18.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#password_wo DatabaseSecretBackendStaticRole#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// The version of the password_wo field. Used for tracking changes to the write-only password field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#password_wo_version DatabaseSecretBackendStaticRole#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// The amount of time Vault should wait before rotating the password, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#rotation_period DatabaseSecretBackendStaticRole#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// A cron-style string that will define the schedule on which rotations should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#rotation_schedule DatabaseSecretBackendStaticRole#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// Database statements to execute to rotate the password for the configured database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#rotation_statements DatabaseSecretBackendStaticRole#rotation_statements}
	RotationStatements *[]*string `field:"optional" json:"rotationStatements" yaml:"rotationStatements"`
	// The amount of time in seconds in which the rotations are allowed to occur starting from a given rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#rotation_window DatabaseSecretBackendStaticRole#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// The password corresponding to the username in the database.
	//
	// Required when using the Rootless Password Rotation workflow for static roles. Deprecated in favor of password_wo field introduced in Vault 1.19.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#self_managed_password DatabaseSecretBackendStaticRole#self_managed_password}
	SelfManagedPassword *string `field:"optional" json:"selfManagedPassword" yaml:"selfManagedPassword"`
	// Skip rotation of the password on import. When not set, inherits from connection's skip_static_role_import_rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_static_role#skip_import_rotation DatabaseSecretBackendStaticRole#skip_import_rotation}
	SkipImportRotation interface{} `field:"optional" json:"skipImportRotation" yaml:"skipImportRotation"`
}

