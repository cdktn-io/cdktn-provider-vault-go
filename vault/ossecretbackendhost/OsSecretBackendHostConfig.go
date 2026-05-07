// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ossecretbackendhost

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OsSecretBackendHostConfig struct {
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
	// Address of the host (hostname or IP).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#address OsSecretBackendHost#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Path where the OS secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#mount OsSecretBackendHost#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#name OsSecretBackendHost#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Custom metadata for the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#custom_metadata OsSecretBackendHost#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Disable automated password rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#disable_automated_rotation OsSecretBackendHost#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#namespace OsSecretBackendHost#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the password policy to use for password generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#password_policy OsSecretBackendHost#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Port to connect to on the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#port OsSecretBackendHost#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// How often to rotate passwords, in seconds. Mutually exclusive with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#rotation_period OsSecretBackendHost#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// Cron schedule for password rotation. Mutually exclusive with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#rotation_schedule OsSecretBackendHost#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// Window of time for password rotation, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#rotation_window OsSecretBackendHost#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// SSH host key for the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/os_secret_backend_host#ssh_host_key OsSecretBackendHost#ssh_host_key}
	SshHostKey *string `field:"optional" json:"sshHostKey" yaml:"sshHostKey"`
}

