// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultawsaccesscredentials

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultAwsAccessCredentialsConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Mount path for the AWS secret engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#mount EphemeralVaultAwsAccessCredentials#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// AWS Secret Role to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#role EphemeralVaultAwsAccessCredentials#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#mount_id EphemeralVaultAwsAccessCredentials#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#namespace EphemeralVaultAwsAccessCredentials#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Region the read credentials belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#region EphemeralVaultAwsAccessCredentials#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ARN to use if multiple are available in the role. Required if the role has multiple ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#role_arn EphemeralVaultAwsAccessCredentials#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// User specified Time-To-Live for the STS token. Uses the Role defined default_sts_ttl when not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#ttl EphemeralVaultAwsAccessCredentials#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Type of credentials to read. Must be either 'creds' for Access Key and Secret Key, or 'sts' for STS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/aws_access_credentials#type EphemeralVaultAwsAccessCredentials#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

