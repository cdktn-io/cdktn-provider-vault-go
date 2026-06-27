// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendidentitywhitelist

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsAuthBackendIdentityWhitelistConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_identity_whitelist#backend AwsAuthBackendIdentityWhitelist#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// If true, disables the periodic tidying of the identiy whitelist entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_identity_whitelist#disable_periodic_tidy AwsAuthBackendIdentityWhitelist#disable_periodic_tidy}
	DisablePeriodicTidy interface{} `field:"optional" json:"disablePeriodicTidy" yaml:"disablePeriodicTidy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_identity_whitelist#id AwsAuthBackendIdentityWhitelist#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_identity_whitelist#namespace AwsAuthBackendIdentityWhitelist#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The amount of extra time that must have passed beyond the roletag expiration, before it's removed from backend storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/aws_auth_backend_identity_whitelist#safety_buffer AwsAuthBackendIdentityWhitelist#safety_buffer}
	SafetyBuffer *float64 `field:"optional" json:"safetyBuffer" yaml:"safetyBuffer"`
}

