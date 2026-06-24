// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alicloudsecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AlicloudSecretBackendConfig struct {
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
	// The AliCloud Access Key ID to use when generating new credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend#access_key AlicloudSecretBackend#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// Path of the AliCloud secrets engine mount.
	//
	// Must match the `path` of a `vault_mount` resource with `type = "alicloud"`. Use `vault_mount.alicloud.path` here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend#mount AlicloudSecretBackend#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Write-only AliCloud Secret Access Key. This value will never be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend#secret_key_wo AlicloudSecretBackend#secret_key_wo}
	SecretKeyWo *string `field:"required" json:"secretKeyWo" yaml:"secretKeyWo"`
	// A version counter for the write-only `secret_key_wo` field.
	//
	// Incrementing this value will trigger an update to the secret key in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend#secret_key_wo_version AlicloudSecretBackend#secret_key_wo_version}
	SecretKeyWoVersion *float64 `field:"required" json:"secretKeyWoVersion" yaml:"secretKeyWoVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend#namespace AlicloudSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

