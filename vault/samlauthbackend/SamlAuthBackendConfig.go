// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package samlauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SamlAuthBackendConfig struct {
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
	// The well-formatted URLs of your Assertion Consumer Service (ACS) that should receive a response from the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#acs_urls SamlAuthBackend#acs_urls}
	AcsUrls *[]*string `field:"required" json:"acsUrls" yaml:"acsUrls"`
	// The entity ID of the SAML authentication service provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#entity_id SamlAuthBackend#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The role to use if no role is provided during login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#default_role SamlAuthBackend#default_role}
	DefaultRole *string `field:"optional" json:"defaultRole" yaml:"defaultRole"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#disable_remount SamlAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#id SamlAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The PEM encoded certificate of the identity provider. Mutually exclusive with 'idp_metadata_url'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#idp_cert SamlAuthBackend#idp_cert}
	IdpCert *string `field:"optional" json:"idpCert" yaml:"idpCert"`
	// The entity ID of the identity provider. Mutually exclusive with 'idp_metadata_url'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#idp_entity_id SamlAuthBackend#idp_entity_id}
	IdpEntityId *string `field:"optional" json:"idpEntityId" yaml:"idpEntityId"`
	// The metadata URL of the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#idp_metadata_url SamlAuthBackend#idp_metadata_url}
	IdpMetadataUrl *string `field:"optional" json:"idpMetadataUrl" yaml:"idpMetadataUrl"`
	// The SSO URL of the identity provider. Mutually exclusive with 'idp_metadata_url'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#idp_sso_url SamlAuthBackend#idp_sso_url}
	IdpSsoUrl *string `field:"optional" json:"idpSsoUrl" yaml:"idpSsoUrl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#namespace SamlAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#path SamlAuthBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#tune SamlAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
	// Whether to validate the assertion signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#validate_assertion_signature SamlAuthBackend#validate_assertion_signature}
	ValidateAssertionSignature interface{} `field:"optional" json:"validateAssertionSignature" yaml:"validateAssertionSignature"`
	// Whether to validate the response signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#validate_response_signature SamlAuthBackend#validate_response_signature}
	ValidateResponseSignature interface{} `field:"optional" json:"validateResponseSignature" yaml:"validateResponseSignature"`
	// Log additional, potentially sensitive information during the SAML exchange according to the current logging level. Not recommended for production.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend#verbose_logging SamlAuthBackend#verbose_logging}
	VerboseLogging interface{} `field:"optional" json:"verboseLogging" yaml:"verboseLogging"`
}

