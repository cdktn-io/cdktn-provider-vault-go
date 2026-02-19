// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigautotidy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigAutoTidyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#backend PkiSecretBackendConfigAutoTidy#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Specifies whether automatic tidy is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#enabled PkiSecretBackendConfigAutoTidy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The amount of time that must pass after creation that an account with no orders is marked revoked, and the amount of time after being marked revoked or deactivated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#acme_account_safety_buffer PkiSecretBackendConfigAutoTidy#acme_account_safety_buffer}
	AcmeAccountSafetyBuffer *string `field:"optional" json:"acmeAccountSafetyBuffer" yaml:"acmeAccountSafetyBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#id PkiSecretBackendConfigAutoTidy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Interval at which to run an auto-tidy operation.
	//
	// This is the time between tidy invocations (after one finishes to the start of the next).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#interval_duration PkiSecretBackendConfigAutoTidy#interval_duration}
	IntervalDuration *string `field:"optional" json:"intervalDuration" yaml:"intervalDuration"`
	// The amount of extra time that must have passed beyond issuer's expiration before it is removed from the backend storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#issuer_safety_buffer PkiSecretBackendConfigAutoTidy#issuer_safety_buffer}
	IssuerSafetyBuffer *string `field:"optional" json:"issuerSafetyBuffer" yaml:"issuerSafetyBuffer"`
	// This configures whether stored certificate are counted upon initialization of the backend, and whether during normal operation, a running count of certificates stored is maintained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#maintain_stored_certificate_counts PkiSecretBackendConfigAutoTidy#maintain_stored_certificate_counts}
	MaintainStoredCertificateCounts interface{} `field:"optional" json:"maintainStoredCertificateCounts" yaml:"maintainStoredCertificateCounts"`
	// The maximum amount of time auto-tidy will be delayed after startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#max_startup_backoff_duration PkiSecretBackendConfigAutoTidy#max_startup_backoff_duration}
	MaxStartupBackoffDuration *string `field:"optional" json:"maxStartupBackoffDuration" yaml:"maxStartupBackoffDuration"`
	// The minimum amount of time auto-tidy will be delayed after startup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#min_startup_backoff_duration PkiSecretBackendConfigAutoTidy#min_startup_backoff_duration}
	MinStartupBackoffDuration *string `field:"optional" json:"minStartupBackoffDuration" yaml:"minStartupBackoffDuration"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#namespace PkiSecretBackendConfigAutoTidy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The amount of time to wait between processing certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#pause_duration PkiSecretBackendConfigAutoTidy#pause_duration}
	PauseDuration *string `field:"optional" json:"pauseDuration" yaml:"pauseDuration"`
	// This configures whether the stored certificate count is published to the metrics consumer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#publish_stored_certificate_count_metrics PkiSecretBackendConfigAutoTidy#publish_stored_certificate_count_metrics}
	PublishStoredCertificateCountMetrics interface{} `field:"optional" json:"publishStoredCertificateCountMetrics" yaml:"publishStoredCertificateCountMetrics"`
	// The amount of time that must pass from the cross-cluster revocation request being initiated to when it will be slated for removal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#revocation_queue_safety_buffer PkiSecretBackendConfigAutoTidy#revocation_queue_safety_buffer}
	RevocationQueueSafetyBuffer *string `field:"optional" json:"revocationQueueSafetyBuffer" yaml:"revocationQueueSafetyBuffer"`
	// The amount of extra time that must have passed beyond certificate expiration before it is removed from the backend storage and/or revocation list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#safety_buffer PkiSecretBackendConfigAutoTidy#safety_buffer}
	SafetyBuffer *string `field:"optional" json:"safetyBuffer" yaml:"safetyBuffer"`
	// Set to true to enable tidying ACME accounts, orders and authorizations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_acme PkiSecretBackendConfigAutoTidy#tidy_acme}
	TidyAcme interface{} `field:"optional" json:"tidyAcme" yaml:"tidyAcme"`
	// Set to true to enable tidying up certificate metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_cert_metadata PkiSecretBackendConfigAutoTidy#tidy_cert_metadata}
	TidyCertMetadata interface{} `field:"optional" json:"tidyCertMetadata" yaml:"tidyCertMetadata"`
	// Set to true to enable tidying up the certificate store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_cert_store PkiSecretBackendConfigAutoTidy#tidy_cert_store}
	TidyCertStore interface{} `field:"optional" json:"tidyCertStore" yaml:"tidyCertStore"`
	// Set to true to enable tidying up the CMPv2 nonce store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_cmpv2_nonce_store PkiSecretBackendConfigAutoTidy#tidy_cmpv2_nonce_store}
	TidyCmpv2NonceStore interface{} `field:"optional" json:"tidyCmpv2NonceStore" yaml:"tidyCmpv2NonceStore"`
	// Set to true to enable tidying up the cross-cluster revoked certificate store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_cross_cluster_revoked_certs PkiSecretBackendConfigAutoTidy#tidy_cross_cluster_revoked_certs}
	TidyCrossClusterRevokedCerts interface{} `field:"optional" json:"tidyCrossClusterRevokedCerts" yaml:"tidyCrossClusterRevokedCerts"`
	// Set to true to automatically remove expired issuers past the issuer_safety_buffer.
	//
	// No keys will be removed as part of this operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_expired_issuers PkiSecretBackendConfigAutoTidy#tidy_expired_issuers}
	TidyExpiredIssuers interface{} `field:"optional" json:"tidyExpiredIssuers" yaml:"tidyExpiredIssuers"`
	// Set to true to move the legacy ca_bundle from /config/ca_bundle to /config/ca_bundle.bak.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_move_legacy_ca_bundle PkiSecretBackendConfigAutoTidy#tidy_move_legacy_ca_bundle}
	TidyMoveLegacyCaBundle interface{} `field:"optional" json:"tidyMoveLegacyCaBundle" yaml:"tidyMoveLegacyCaBundle"`
	// Set to true to remove stale revocation queue entries that haven't been confirmed by any active cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_revocation_queue PkiSecretBackendConfigAutoTidy#tidy_revocation_queue}
	TidyRevocationQueue interface{} `field:"optional" json:"tidyRevocationQueue" yaml:"tidyRevocationQueue"`
	// Set to true to validate issuer associations on revocation entries.
	//
	// This helps increase the performance of CRL building and OCSP responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_revoked_cert_issuer_associations PkiSecretBackendConfigAutoTidy#tidy_revoked_cert_issuer_associations}
	TidyRevokedCertIssuerAssociations interface{} `field:"optional" json:"tidyRevokedCertIssuerAssociations" yaml:"tidyRevokedCertIssuerAssociations"`
	// Set to true to remove all invalid and expired certificates from storage.
	//
	// A revoked storage entry is considered invalid if the entry is empty, or the value within the entry is empty. If a certificate is removed due to expiry, the entry will also be removed from the CRL, and the CRL will be rotated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy#tidy_revoked_certs PkiSecretBackendConfigAutoTidy#tidy_revoked_certs}
	TidyRevokedCerts interface{} `field:"optional" json:"tidyRevokedCerts" yaml:"tidyRevokedCerts"`
}

