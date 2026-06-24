// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awssecretbackendstaticrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsSecretBackendStaticRoleConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#name AwsSecretBackendStaticRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// How often Vault should rotate the password of the user entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#rotation_period AwsSecretBackendStaticRole#rotation_period}
	RotationPeriod *float64 `field:"required" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The username of the existing AWS IAM user to manage password rotation for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#username AwsSecretBackendStaticRole#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The ARN of the role to assume when managing the static role.
	//
	// This is required for cross-account role management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#assume_role_arn AwsSecretBackendStaticRole#assume_role_arn}
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Session name to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#assume_role_session_name AwsSecretBackendStaticRole#assume_role_session_name}
	AssumeRoleSessionName *string `field:"optional" json:"assumeRoleSessionName" yaml:"assumeRoleSessionName"`
	// The path where the AWS secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#backend AwsSecretBackendStaticRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// External ID to use when assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#external_id AwsSecretBackendStaticRole#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#id AwsSecretBackendStaticRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/aws_secret_backend_static_role#namespace AwsSecretBackendStaticRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

