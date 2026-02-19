// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginAws struct {
	// The Vault role to use when logging into Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The AWS access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_access_key_id VaultProvider#aws_access_key_id}
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// The IAM endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_iam_endpoint VaultProvider#aws_iam_endpoint}
	AwsIamEndpoint *string `field:"optional" json:"awsIamEndpoint" yaml:"awsIamEndpoint"`
	// The name of the AWS profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_profile VaultProvider#aws_profile}
	AwsProfile *string `field:"optional" json:"awsProfile" yaml:"awsProfile"`
	// The AWS region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_region VaultProvider#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The ARN of the AWS Role to assume.Used during STS AssumeRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_role_arn VaultProvider#aws_role_arn}
	AwsRoleArn *string `field:"optional" json:"awsRoleArn" yaml:"awsRoleArn"`
	// Specifies the name to attach to the AWS role session. Used during STS AssumeRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_role_session_name VaultProvider#aws_role_session_name}
	AwsRoleSessionName *string `field:"optional" json:"awsRoleSessionName" yaml:"awsRoleSessionName"`
	// The AWS secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_secret_access_key VaultProvider#aws_secret_access_key}
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// The AWS session token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_session_token VaultProvider#aws_session_token}
	AwsSessionToken *string `field:"optional" json:"awsSessionToken" yaml:"awsSessionToken"`
	// Path to the AWS shared credentials file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_shared_credentials_file VaultProvider#aws_shared_credentials_file}
	AwsSharedCredentialsFile *string `field:"optional" json:"awsSharedCredentialsFile" yaml:"awsSharedCredentialsFile"`
	// The STS endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_sts_endpoint VaultProvider#aws_sts_endpoint}
	AwsStsEndpoint *string `field:"optional" json:"awsStsEndpoint" yaml:"awsStsEndpoint"`
	// Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#aws_web_identity_token_file VaultProvider#aws_web_identity_token_file}
	AwsWebIdentityTokenFile *string `field:"optional" json:"awsWebIdentityTokenFile" yaml:"awsWebIdentityTokenFile"`
	// The Vault header value to include in the STS signing request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#header_value VaultProvider#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

