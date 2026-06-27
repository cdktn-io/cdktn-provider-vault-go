// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultawsaccesscredentials

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultAwsAccessCredentialsConfig struct {
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
	// AWS Secret Backend to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#backend DataVaultAwsAccessCredentials#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// AWS Secret Role to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#role DataVaultAwsAccessCredentials#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#id DataVaultAwsAccessCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#namespace DataVaultAwsAccessCredentials#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Region the read credentials belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#region DataVaultAwsAccessCredentials#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// ARN to use if multiple are available in the role. Required if the role has multiple ARNs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#role_arn DataVaultAwsAccessCredentials#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// User specified Time-To-Live for the STS token. Uses the Role defined default_sts_ttl when not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#ttl DataVaultAwsAccessCredentials#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Type of credentials to read. Must be either 'creds' for Access Key and Secret Key, or 'sts' for STS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/aws_access_credentials#type DataVaultAwsAccessCredentials#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

