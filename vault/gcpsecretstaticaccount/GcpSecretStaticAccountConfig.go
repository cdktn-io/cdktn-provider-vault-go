// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpsecretstaticaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpSecretStaticAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#backend GcpSecretStaticAccount#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Email of the GCP service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#service_account_email GcpSecretStaticAccount#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Name of the Static Account to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#static_account GcpSecretStaticAccount#static_account}
	StaticAccount *string `field:"required" json:"staticAccount" yaml:"staticAccount"`
	// binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#binding GcpSecretStaticAccount#binding}
	Binding interface{} `field:"optional" json:"binding" yaml:"binding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#id GcpSecretStaticAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#namespace GcpSecretStaticAccount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Type of secret generated for this static account. Defaults to `access_token`. Accepted values: `access_token`, `service_account_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#secret_type GcpSecretStaticAccount#secret_type}
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
	// List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_secret_static_account#token_scopes GcpSecretStaticAccount#token_scopes}
	TokenScopes *[]*string `field:"optional" json:"tokenScopes" yaml:"tokenScopes"`
}

