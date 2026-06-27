// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keymgmtawskms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KeymgmtAwsKmsConfig struct {
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
	// Refers to the name of an AWS region. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#key_collection KeymgmtAwsKms#key_collection}
	KeyCollection *string `field:"required" json:"keyCollection" yaml:"keyCollection"`
	// Path of the Key Management secrets engine mount.
	//
	// Must match the `path` of a `vault_mount` resource with `type = "keymgmt"`. Use `vault_mount.keymgmt.path` here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#mount KeymgmtAwsKms#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Specifies the name of the AWS KMS provider. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#name KeymgmtAwsKms#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The credentials to use for authentication with AWS KMS.
	//
	// Supplying values for this parameter is optional, as credentials may also be specified as environment variables. Credentials provided to this parameter will take precedence over credentials provided via environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#credentials_wo KeymgmtAwsKms#credentials_wo}
	CredentialsWo *map[string]*string `field:"optional" json:"credentialsWo" yaml:"credentialsWo"`
	// Version counter for the write-only `credentials_wo` field.
	//
	// Since write-only values are not stored in state, Terraform cannot detect when credentials change. Increment this value whenever you update `credentials_wo` to ensure the new credentials are sent to Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#credentials_wo_version KeymgmtAwsKms#credentials_wo_version}
	CredentialsWoVersion *float64 `field:"optional" json:"credentialsWoVersion" yaml:"credentialsWoVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_aws_kms#namespace KeymgmtAwsKms#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

