// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendcrlconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendCrlConfigConfig struct {
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
	// The path of the PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#backend PkiSecretBackendCrlConfig#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Enables or disables periodic rebuilding of the CRL upon expiry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#auto_rebuild PkiSecretBackendCrlConfig#auto_rebuild}
	AutoRebuild interface{} `field:"optional" json:"autoRebuild" yaml:"autoRebuild"`
	// Grace period before CRL expiry to attempt rebuild of CRL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#auto_rebuild_grace_period PkiSecretBackendCrlConfig#auto_rebuild_grace_period}
	AutoRebuildGracePeriod *string `field:"optional" json:"autoRebuildGracePeriod" yaml:"autoRebuildGracePeriod"`
	// Enable cross-cluster revocation request queues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#cross_cluster_revocation PkiSecretBackendCrlConfig#cross_cluster_revocation}
	CrossClusterRevocation interface{} `field:"optional" json:"crossClusterRevocation" yaml:"crossClusterRevocation"`
	// Interval to check for new revocations on, to regenerate the delta CRL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#delta_rebuild_interval PkiSecretBackendCrlConfig#delta_rebuild_interval}
	DeltaRebuildInterval *string `field:"optional" json:"deltaRebuildInterval" yaml:"deltaRebuildInterval"`
	// Disables or enables CRL building.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#disable PkiSecretBackendCrlConfig#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// Enables or disables building of delta CRLs with up-to-date revocation information, augmenting the last complete CRL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#enable_delta PkiSecretBackendCrlConfig#enable_delta}
	EnableDelta interface{} `field:"optional" json:"enableDelta" yaml:"enableDelta"`
	// Specifies the time until expiration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#expiry PkiSecretBackendCrlConfig#expiry}
	Expiry *string `field:"optional" json:"expiry" yaml:"expiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#id PkiSecretBackendCrlConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum number of entries a CRL can contain.
	//
	// This option exists to prevent accidental runaway issuance/revocation from overloading Vault. If set to -1, the limit is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#max_crl_entries PkiSecretBackendCrlConfig#max_crl_entries}
	MaxCrlEntries *float64 `field:"optional" json:"maxCrlEntries" yaml:"maxCrlEntries"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#namespace PkiSecretBackendCrlConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Disables or enables the OCSP responder in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#ocsp_disable PkiSecretBackendCrlConfig#ocsp_disable}
	OcspDisable interface{} `field:"optional" json:"ocspDisable" yaml:"ocspDisable"`
	// The amount of time an OCSP response can be cached for, useful for OCSP stapling refresh durations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#ocsp_expiry PkiSecretBackendCrlConfig#ocsp_expiry}
	OcspExpiry *string `field:"optional" json:"ocspExpiry" yaml:"ocspExpiry"`
	// Enables unified CRL and OCSP building.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#unified_crl PkiSecretBackendCrlConfig#unified_crl}
	UnifiedCrl interface{} `field:"optional" json:"unifiedCrl" yaml:"unifiedCrl"`
	// Enables serving the unified CRL and OCSP on the existing, previously cluster-local paths.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_crl_config#unified_crl_on_existing_paths PkiSecretBackendCrlConfig#unified_crl_on_existing_paths}
	UnifiedCrlOnExistingPaths interface{} `field:"optional" json:"unifiedCrlOnExistingPaths" yaml:"unifiedCrlOnExistingPaths"`
}

