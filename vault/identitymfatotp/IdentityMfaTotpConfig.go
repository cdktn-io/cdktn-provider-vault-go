// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfatotp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityMfaTotpConfig struct {
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
	// The name of the key's issuing organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#issuer IdentityMfaTotp#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#algorithm IdentityMfaTotp#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The number of digits in the generated TOTP token. This value can either be 6 or 8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#digits IdentityMfaTotp#digits}
	Digits *float64 `field:"optional" json:"digits" yaml:"digits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#id IdentityMfaTotp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the size in bytes of the generated key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#key_size IdentityMfaTotp#key_size}
	KeySize *float64 `field:"optional" json:"keySize" yaml:"keySize"`
	// The maximum number of consecutive failed validation attempts allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#max_validation_attempts IdentityMfaTotp#max_validation_attempts}
	MaxValidationAttempts *float64 `field:"optional" json:"maxValidationAttempts" yaml:"maxValidationAttempts"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#namespace IdentityMfaTotp#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The length of time in seconds used to generate a counter for the TOTP token calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#period IdentityMfaTotp#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The pixel size of the generated square QR code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#qr_size IdentityMfaTotp#qr_size}
	QrSize *float64 `field:"optional" json:"qrSize" yaml:"qrSize"`
	// The number of delay periods that are allowed when validating a TOTP token.
	//
	// This value can either be 0 or 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_mfa_totp#skew IdentityMfaTotp#skew}
	Skew *float64 `field:"optional" json:"skew" yaml:"skew"`
}

