// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package adsecretlibrary

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AdSecretLibraryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#backend AdSecretLibrary#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The name of the set of service accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#name AdSecretLibrary#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The names of all the service accounts that can be checked out from this set.
	//
	// These service accounts must already exist in Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#service_account_names AdSecretLibrary#service_account_names}
	ServiceAccountNames *[]*string `field:"required" json:"serviceAccountNames" yaml:"serviceAccountNames"`
	// Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#disable_check_in_enforcement AdSecretLibrary#disable_check_in_enforcement}
	DisableCheckInEnforcement interface{} `field:"optional" json:"disableCheckInEnforcement" yaml:"disableCheckInEnforcement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#id AdSecretLibrary#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#max_ttl AdSecretLibrary#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#namespace AdSecretLibrary#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ad_secret_library#ttl AdSecretLibrary#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

