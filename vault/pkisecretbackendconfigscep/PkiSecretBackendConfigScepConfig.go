// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigscep

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigScepConfig struct {
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
	// The PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#backend PkiSecretBackendConfigScep#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// List of allowed digest algorithms for SCEP requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#allowed_digest_algorithms PkiSecretBackendConfigScep#allowed_digest_algorithms}
	AllowedDigestAlgorithms *[]*string `field:"optional" json:"allowedDigestAlgorithms" yaml:"allowedDigestAlgorithms"`
	// List of allowed encryption algorithms for SCEP requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#allowed_encryption_algorithms PkiSecretBackendConfigScep#allowed_encryption_algorithms}
	AllowedEncryptionAlgorithms *[]*string `field:"optional" json:"allowedEncryptionAlgorithms" yaml:"allowedEncryptionAlgorithms"`
	// authenticators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#authenticators PkiSecretBackendConfigScep#authenticators}
	Authenticators *PkiSecretBackendConfigScepAuthenticators `field:"optional" json:"authenticators" yaml:"authenticators"`
	// Specifies the behavior for requests using the default SCEP label. Can be sign-verbatim or a role given by role:<role_name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#default_path_policy PkiSecretBackendConfigScep#default_path_policy}
	DefaultPathPolicy *string `field:"optional" json:"defaultPathPolicy" yaml:"defaultPathPolicy"`
	// Specifies whether SCEP is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#enabled PkiSecretBackendConfigScep#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// external_validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#external_validation PkiSecretBackendConfigScep#external_validation}
	ExternalValidation interface{} `field:"optional" json:"externalValidation" yaml:"externalValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#id PkiSecretBackendConfigScep#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The level of logging verbosity, affects only SCEP logs on this mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#log_level PkiSecretBackendConfigScep#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#namespace PkiSecretBackendConfigScep#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If true, only return the issuer CA, otherwise the entire CA certificate chain will be returned if available from the PKI mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_config_scep#restrict_ca_chain_to_issuer PkiSecretBackendConfigScep#restrict_ca_chain_to_issuer}
	RestrictCaChainToIssuer interface{} `field:"optional" json:"restrictCaChainToIssuer" yaml:"restrictCaChainToIssuer"`
}

