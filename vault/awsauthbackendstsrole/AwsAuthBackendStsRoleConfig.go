// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendstsrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendStsRoleConfig struct {
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
	// AWS account ID to be associated with STS role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#account_id AwsAuthBackendStsRole#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// AWS ARN for STS role to be assumed when interacting with the account specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#sts_role AwsAuthBackendStsRole#sts_role}
	StsRole *string `field:"required" json:"stsRole" yaml:"stsRole"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#backend AwsAuthBackendStsRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// External ID expected by the STS role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#external_id AwsAuthBackendStsRole#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#id AwsAuthBackendStsRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_auth_backend_sts_role#namespace AwsAuthBackendStsRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

