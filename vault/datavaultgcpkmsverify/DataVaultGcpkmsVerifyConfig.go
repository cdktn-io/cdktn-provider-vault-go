// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultgcpkmsverify

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultGcpkmsVerifyConfig struct {
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
	// Base64-encoded digest to verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#digest DataVaultGcpkmsVerify#digest}
	Digest *string `field:"required" json:"digest" yaml:"digest"`
	// Name of the Vault key to use for verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#key_name DataVaultGcpkmsVerify#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Version of the key to use for verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#key_version DataVaultGcpkmsVerify#key_version}
	KeyVersion *float64 `field:"required" json:"keyVersion" yaml:"keyVersion"`
	// Path where the GCP KMS secrets engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#mount DataVaultGcpkmsVerify#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Base64-encoded signature to verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#signature DataVaultGcpkmsVerify#signature}
	Signature *string `field:"required" json:"signature" yaml:"signature"`
	// Target namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/data-sources/gcpkms_verify#namespace DataVaultGcpkmsVerify#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

