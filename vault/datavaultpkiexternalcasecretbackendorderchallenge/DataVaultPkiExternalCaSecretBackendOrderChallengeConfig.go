// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpkiexternalcasecretbackendorderchallenge

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultPkiExternalCaSecretBackendOrderChallengeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#challenge_type DataVaultPkiExternalCaSecretBackendOrderChallenge#challenge_type}
	ChallengeType *string `field:"required" json:"challengeType" yaml:"challengeType"`
	// The identifier (domain name) for which to retrieve the challenge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#identifier DataVaultPkiExternalCaSecretBackendOrderChallenge#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#mount DataVaultPkiExternalCaSecretBackendOrderChallenge#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The unique identifier for the ACME order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#order_id DataVaultPkiExternalCaSecretBackendOrderChallenge#order_id}
	OrderId *string `field:"required" json:"orderId" yaml:"orderId"`
	// Name of the role associated with the order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#role_name DataVaultPkiExternalCaSecretBackendOrderChallenge#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/pki_external_ca_secret_backend_order_challenge#namespace DataVaultPkiExternalCaSecretBackendOrderChallenge#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

