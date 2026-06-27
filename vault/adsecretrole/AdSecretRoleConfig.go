// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package adsecretrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AdSecretRoleConfig struct {
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
	// The mount path for the AD backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#backend AdSecretRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#role AdSecretRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The username/logon name for the service account with which this role will be associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#service_account_name AdSecretRole#service_account_name}
	ServiceAccountName *string `field:"required" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#id AdSecretRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#namespace AdSecretRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// In seconds, the default password time-to-live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/ad_secret_role#ttl AdSecretRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

