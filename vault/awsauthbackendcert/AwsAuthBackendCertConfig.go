// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendcert

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendCertConfig struct {
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
	// Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#aws_public_cert AwsAuthBackendCert#aws_public_cert}
	AwsPublicCert *string `field:"required" json:"awsPublicCert" yaml:"awsPublicCert"`
	// Name of the certificate to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#cert_name AwsAuthBackendCert#cert_name}
	CertName *string `field:"required" json:"certName" yaml:"certName"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#backend AwsAuthBackendCert#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#id AwsAuthBackendCert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#namespace AwsAuthBackendCert#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The type of document that can be verified using the certificate. Must be either "pkcs7" or "identity".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_cert#type AwsAuthBackendCert#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

