// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpsecretroleset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpSecretRolesetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#backend GcpSecretRoleset#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#binding GcpSecretRoleset#binding}
	Binding interface{} `field:"required" json:"binding" yaml:"binding"`
	// Name of the GCP project that this roleset's service account will belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#project GcpSecretRoleset#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Name of the RoleSet to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#roleset GcpSecretRoleset#roleset}
	Roleset *string `field:"required" json:"roleset" yaml:"roleset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#id GcpSecretRoleset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#namespace GcpSecretRoleset#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Type of secret generated for this role set. Defaults to `access_token`. Accepted values: `access_token`, `service_account_key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#secret_type GcpSecretRoleset#secret_type}
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
	// List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_roleset#token_scopes GcpSecretRoleset#token_scopes}
	TokenScopes *[]*string `field:"optional" json:"tokenScopes" yaml:"tokenScopes"`
}

