// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpkisecretbackendexternalcaorderchallenge

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultPkiSecretBackendExternalCaOrderChallengeConfig struct {
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
	// The type of ACME challenge to retrieve. Valid values are `http-01`, `dns-01`, `tls-alpn-01`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#challenge_type DataVaultPkiSecretBackendExternalCaOrderChallenge#challenge_type}
	ChallengeType *string `field:"required" json:"challengeType" yaml:"challengeType"`
	// The identifier (domain name) for which to retrieve the challenge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#identifier DataVaultPkiSecretBackendExternalCaOrderChallenge#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#mount DataVaultPkiSecretBackendExternalCaOrderChallenge#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The unique identifier for the ACME order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#order_id DataVaultPkiSecretBackendExternalCaOrderChallenge#order_id}
	OrderId *string `field:"required" json:"orderId" yaml:"orderId"`
	// Name of the role associated with the order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#role_name DataVaultPkiSecretBackendExternalCaOrderChallenge#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_external_ca_order_challenge#namespace DataVaultPkiSecretBackendExternalCaOrderChallenge#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

