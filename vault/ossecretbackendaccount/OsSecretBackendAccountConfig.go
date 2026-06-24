// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ossecretbackendaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OsSecretBackendAccountConfig struct {
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
	// Name of the host this account belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#host OsSecretBackendAccount#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Path where the OS secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#mount OsSecretBackendAccount#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#name OsSecretBackendAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Password for the account.
	//
	// This is write-only, will not be read back from Vault,
	// 	and can only be set during resource creation. To update the password after creation, use the Vault CLI
	// 	or API to call the reset endpoint directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#password_wo OsSecretBackendAccount#password_wo}
	PasswordWo *string `field:"required" json:"passwordWo" yaml:"passwordWo"`
	// Username for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#username OsSecretBackendAccount#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Custom metadata for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#custom_metadata OsSecretBackendAccount#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Disable automated password rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#disable_automated_rotation OsSecretBackendAccount#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#namespace OsSecretBackendAccount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Reference to a parent account for rotation management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#parent_account_ref OsSecretBackendAccount#parent_account_ref}
	ParentAccountRef *string `field:"optional" json:"parentAccountRef" yaml:"parentAccountRef"`
	// Name of the password policy to use for password generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#password_policy OsSecretBackendAccount#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// How often to rotate passwords, in seconds. Mutually exclusive with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#rotation_period OsSecretBackendAccount#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// Cron schedule for password rotation. Mutually exclusive with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#rotation_schedule OsSecretBackendAccount#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// Window of time for password rotation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#rotation_window OsSecretBackendAccount#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Verify the connection to the host with the provided credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account#verify_connection OsSecretBackendAccount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

