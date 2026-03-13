// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendlogin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendLoginConfig struct {
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
	// AWS Auth Backend to read the token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#backend AwsAuthBackendLogin#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// The HTTP method used in the signed request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#iam_http_request_method AwsAuthBackendLogin#iam_http_request_method}
	IamHttpRequestMethod *string `field:"optional" json:"iamHttpRequestMethod" yaml:"iamHttpRequestMethod"`
	// The Base64-encoded body of the signed request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#iam_request_body AwsAuthBackendLogin#iam_request_body}
	IamRequestBody *string `field:"optional" json:"iamRequestBody" yaml:"iamRequestBody"`
	// The Base64-encoded, JSON serialized representation of the sts:GetCallerIdentity HTTP request headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#iam_request_headers AwsAuthBackendLogin#iam_request_headers}
	IamRequestHeaders *string `field:"optional" json:"iamRequestHeaders" yaml:"iamRequestHeaders"`
	// The Base64-encoded HTTP URL used in the signed request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#iam_request_url AwsAuthBackendLogin#iam_request_url}
	IamRequestUrl *string `field:"optional" json:"iamRequestUrl" yaml:"iamRequestUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#id AwsAuthBackendLogin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Base64-encoded EC2 instance identity document to authenticate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#identity AwsAuthBackendLogin#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#namespace AwsAuthBackendLogin#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The nonce to be used for subsequent login requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#nonce AwsAuthBackendLogin#nonce}
	Nonce *string `field:"optional" json:"nonce" yaml:"nonce"`
	// PKCS7 signature of the identity document to authenticate with, with all newline characters removed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#pkcs7 AwsAuthBackendLogin#pkcs7}
	Pkcs7 *string `field:"optional" json:"pkcs7" yaml:"pkcs7"`
	// AWS Auth Role to read the token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#role AwsAuthBackendLogin#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Base64-encoded SHA256 RSA signature of the instance identtiy document to authenticate with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/aws_auth_backend_login#signature AwsAuthBackendLogin#signature}
	Signature *string `field:"optional" json:"signature" yaml:"signature"`
}

