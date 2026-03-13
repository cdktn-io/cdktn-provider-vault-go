// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendexternalcaorderchallengefulfilled

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendExternalCaOrderChallengeFulfilledConfig struct {
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
	// The type of ACME challenge that was fulfilled. Valid values are `http-01`, `dns-01`, `tls-alpn-01`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#challenge_type PkiSecretBackendExternalCaOrderChallengeFulfilled#challenge_type}
	ChallengeType *string `field:"required" json:"challengeType" yaml:"challengeType"`
	// The identifier (domain name) for which the challenge was fulfilled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#identifier PkiSecretBackendExternalCaOrderChallengeFulfilled#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#mount PkiSecretBackendExternalCaOrderChallengeFulfilled#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The unique identifier for the ACME order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#order_id PkiSecretBackendExternalCaOrderChallengeFulfilled#order_id}
	OrderId *string `field:"required" json:"orderId" yaml:"orderId"`
	// Name of the role associated with the order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#role_name PkiSecretBackendExternalCaOrderChallengeFulfilled#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order_challenge_fulfilled#namespace PkiSecretBackendExternalCaOrderChallengeFulfilled#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

