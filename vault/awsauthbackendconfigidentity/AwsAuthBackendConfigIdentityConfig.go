// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendconfigidentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendConfigIdentityConfig struct {
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
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#backend AwsAuthBackendConfigIdentity#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Configures how to generate the identity alias when using the ec2 auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#ec2_alias AwsAuthBackendConfigIdentity#ec2_alias}
	Ec2Alias *string `field:"optional" json:"ec2Alias" yaml:"ec2Alias"`
	// The metadata to include on the token returned by the login endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#ec2_metadata AwsAuthBackendConfigIdentity#ec2_metadata}
	Ec2Metadata *[]*string `field:"optional" json:"ec2Metadata" yaml:"ec2Metadata"`
	// How to generate the identity alias when using the iam auth method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#iam_alias AwsAuthBackendConfigIdentity#iam_alias}
	IamAlias *string `field:"optional" json:"iamAlias" yaml:"iamAlias"`
	// The metadata to include on the token returned by the login endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#iam_metadata AwsAuthBackendConfigIdentity#iam_metadata}
	IamMetadata *[]*string `field:"optional" json:"iamMetadata" yaml:"iamMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#id AwsAuthBackendConfigIdentity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_config_identity#namespace AwsAuthBackendConfigIdentity#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

