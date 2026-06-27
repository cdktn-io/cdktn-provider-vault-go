// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpsecretimpersonatedaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpSecretImpersonatedAccountConfig struct {
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
	// Path where the GCP secrets engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#backend GcpSecretImpersonatedAccount#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the Impersonated Account to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#impersonated_account GcpSecretImpersonatedAccount#impersonated_account}
	ImpersonatedAccount *string `field:"required" json:"impersonatedAccount" yaml:"impersonatedAccount"`
	// Email of the GCP service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#service_account_email GcpSecretImpersonatedAccount#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#id GcpSecretImpersonatedAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#namespace GcpSecretImpersonatedAccount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// List of OAuth scopes to assign to `access_token` secrets generated under this impersonated account (`access_token` impersonated accounts only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#token_scopes GcpSecretImpersonatedAccount#token_scopes}
	TokenScopes *[]*string `field:"optional" json:"tokenScopes" yaml:"tokenScopes"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_impersonated_account#ttl GcpSecretImpersonatedAccount#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

