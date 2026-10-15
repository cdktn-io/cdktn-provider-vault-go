// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tpmauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TpmAuthBackendConfigConfig struct {
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
	// Path of the enabled TPM auth backend mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_config#mount TpmAuthBackendConfig#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// How long each CA is valid once it becomes active, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_config#ca_lifetime TpmAuthBackendConfig#ca_lifetime}
	CaLifetime *float64 `field:"optional" json:"caLifetime" yaml:"caLifetime"`
	// How long before hard expiry the active CA stops signing new certificates, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_config#ca_soft_expiry TpmAuthBackendConfig#ca_soft_expiry}
	CaSoftExpiry *float64 `field:"optional" json:"caSoftExpiry" yaml:"caSoftExpiry"`
	// Default lifetime for issued client certificates when the role does not specify `cert_ttl`, in seconds.
	//
	// When unset or `0`, the effective TTL falls back to CA soft expiry. Vault always caps the actual certificate lifetime by the remaining validity of the signing CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_config#default_cert_ttl TpmAuthBackendConfig#default_cert_ttl}
	DefaultCertTtl *float64 `field:"optional" json:"defaultCertTtl" yaml:"defaultCertTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_config#namespace TpmAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

