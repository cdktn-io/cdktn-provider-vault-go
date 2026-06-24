// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alicloudsecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AlicloudSecretBackendRoleConfig struct {
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
	// Path of the AliCloud Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#mount AlicloudSecretBackendRole#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#name AlicloudSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// inline_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#inline_policies AlicloudSecretBackendRole#inline_policies}
	InlinePolicies interface{} `field:"optional" json:"inlinePolicies" yaml:"inlinePolicies"`
	// The maximum allowed lifetime of credentials issued using this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#max_ttl AlicloudSecretBackendRole#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#namespace AlicloudSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// remote_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#remote_policies AlicloudSecretBackendRole#remote_policies}
	RemotePolicies interface{} `field:"optional" json:"remotePolicies" yaml:"remotePolicies"`
	// ARN of the RAM role to assume.
	//
	// If provided, inline_policies and remote_policies should be blank. The trusted principal of the role must be configured to allow assumption by the access key and secret configured in the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#role_arn AlicloudSecretBackendRole#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Duration in seconds after which the issued credentials should expire.
	//
	// Defaults to 0, in which case the value will fallback to the system/mount defaults.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/alicloud_secret_backend_role#ttl AlicloudSecretBackendRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

