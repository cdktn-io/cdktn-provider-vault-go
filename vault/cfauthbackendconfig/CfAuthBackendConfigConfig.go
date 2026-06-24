// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cfauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CfAuthBackendConfigConfig struct {
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
	// CF's full API address, used for verifying that a given `CF_INSTANCE_CERT` shows an application ID, space ID, and organization ID that presently exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_api_addr CfAuthBackendConfig#cf_api_addr}
	CfApiAddr *string `field:"required" json:"cfApiAddr" yaml:"cfApiAddr"`
	// The password for authenticating to the CF API. This attribute is write-only and is never stored in Terraform state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_password_wo CfAuthBackendConfig#cf_password_wo}
	CfPasswordWo *string `field:"required" json:"cfPasswordWo" yaml:"cfPasswordWo"`
	// Version counter for 'cf_password_wo'. Increment this value to trigger an update when only the write-only password changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_password_wo_version CfAuthBackendConfig#cf_password_wo_version}
	CfPasswordWoVersion *float64 `field:"required" json:"cfPasswordWoVersion" yaml:"cfPasswordWoVersion"`
	// The username for authenticating to the CF API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_username CfAuthBackendConfig#cf_username}
	CfUsername *string `field:"required" json:"cfUsername" yaml:"cfUsername"`
	// The root CA certificate(s) to be used for verifying that the `CF_INSTANCE_CERT` presented for logging in was issued by the proper authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#identity_ca_certificates CfAuthBackendConfig#identity_ca_certificates}
	IdentityCaCertificates *[]*string `field:"required" json:"identityCaCertificates" yaml:"identityCaCertificates"`
	// Mount path for the CF auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#mount CfAuthBackendConfig#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The certificate(s) presented by the CF API. Configures Vault to trust these certificates when making API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_api_trusted_certificates CfAuthBackendConfig#cf_api_trusted_certificates}
	CfApiTrustedCertificates *[]*string `field:"optional" json:"cfApiTrustedCertificates" yaml:"cfApiTrustedCertificates"`
	// The timeout for the CF API in seconds.
	//
	// Defaults to `0` (no timeout). Removing this field from config resets the value to `0` in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#cf_timeout CfAuthBackendConfig#cf_timeout}
	CfTimeout *float64 `field:"optional" json:"cfTimeout" yaml:"cfTimeout"`
	// The maximum number of seconds in the future when a signature could have been created.
	//
	// Defaults to `60`. This field is `Computed`: if removed from config, Vault retains the previously set value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#login_max_seconds_not_after CfAuthBackendConfig#login_max_seconds_not_after}
	LoginMaxSecondsNotAfter *float64 `field:"optional" json:"loginMaxSecondsNotAfter" yaml:"loginMaxSecondsNotAfter"`
	// The maximum number of seconds in the past when a signature could have been created.
	//
	// Defaults to `300`. This field is `Computed`: if removed from config, Vault retains the previously set value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#login_max_seconds_not_before CfAuthBackendConfig#login_max_seconds_not_before}
	LoginMaxSecondsNotBefore *float64 `field:"optional" json:"loginMaxSecondsNotBefore" yaml:"loginMaxSecondsNotBefore"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cf_auth_backend_config#namespace CfAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

