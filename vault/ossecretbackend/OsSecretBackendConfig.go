// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ossecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OsSecretBackendConfig struct {
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
	// Path where the OS secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/os_secret_backend#mount OsSecretBackend#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Maximum number of versions to keep for secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/os_secret_backend#max_versions OsSecretBackend#max_versions}
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/os_secret_backend#namespace OsSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Trust SSH host keys on first use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/os_secret_backend#ssh_host_key_trust_on_first_use OsSecretBackend#ssh_host_key_trust_on_first_use}
	SshHostKeyTrustOnFirstUse interface{} `field:"optional" json:"sshHostKeyTrustOnFirstUse" yaml:"sshHostKeyTrustOnFirstUse"`
}

