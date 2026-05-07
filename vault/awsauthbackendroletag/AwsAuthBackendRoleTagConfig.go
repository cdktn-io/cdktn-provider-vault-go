// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendroletag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendRoleTagConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#role AwsAuthBackendRoleTag#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Allows migration of the underlying instance where the client resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#allow_instance_migration AwsAuthBackendRoleTag#allow_instance_migration}
	AllowInstanceMigration interface{} `field:"optional" json:"allowInstanceMigration" yaml:"allowInstanceMigration"`
	// AWS auth backend to read tags from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#backend AwsAuthBackendRoleTag#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Only allow a single token to be granted per instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#disallow_reauthentication AwsAuthBackendRoleTag#disallow_reauthentication}
	DisallowReauthentication interface{} `field:"optional" json:"disallowReauthentication" yaml:"disallowReauthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#id AwsAuthBackendRoleTag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Instance ID for which this tag is intended.
	//
	// The created tag can only be used by the instance with the given ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#instance_id AwsAuthBackendRoleTag#instance_id}
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The maximum allowed lifetime of tokens issued using this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#max_ttl AwsAuthBackendRoleTag#max_ttl}
	MaxTtl *string `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#namespace AwsAuthBackendRoleTag#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Policies to be associated with the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role_tag#policies AwsAuthBackendRoleTag#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

