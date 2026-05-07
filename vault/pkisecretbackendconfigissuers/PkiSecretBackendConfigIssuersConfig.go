// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigissuers

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigIssuersConfig struct {
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
	// Full path where PKI backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_issuers#backend PkiSecretBackendConfigIssuers#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Specifies the default issuer by ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_issuers#default PkiSecretBackendConfigIssuers#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// Specifies whether a root creation or an issuer import operation updates the default issuer to the newly added issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_issuers#default_follows_latest_issuer PkiSecretBackendConfigIssuers#default_follows_latest_issuer}
	DefaultFollowsLatestIssuer interface{} `field:"optional" json:"defaultFollowsLatestIssuer" yaml:"defaultFollowsLatestIssuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_issuers#id PkiSecretBackendConfigIssuers#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_issuers#namespace PkiSecretBackendConfigIssuers#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

