// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CertAuthBackendRoleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#certificate CertAuthBackendRole#certificate}.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#name CertAuthBackendRole#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#alias_metadata CertAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_common_names CertAuthBackendRole#allowed_common_names}.
	AllowedCommonNames *[]*string `field:"optional" json:"allowedCommonNames" yaml:"allowedCommonNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_dns_sans CertAuthBackendRole#allowed_dns_sans}.
	AllowedDnsSans *[]*string `field:"optional" json:"allowedDnsSans" yaml:"allowedDnsSans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_email_sans CertAuthBackendRole#allowed_email_sans}.
	AllowedEmailSans *[]*string `field:"optional" json:"allowedEmailSans" yaml:"allowedEmailSans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_names CertAuthBackendRole#allowed_names}.
	AllowedNames *[]*string `field:"optional" json:"allowedNames" yaml:"allowedNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_organizational_units CertAuthBackendRole#allowed_organizational_units}.
	AllowedOrganizationalUnits *[]*string `field:"optional" json:"allowedOrganizationalUnits" yaml:"allowedOrganizationalUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#allowed_uri_sans CertAuthBackendRole#allowed_uri_sans}.
	AllowedUriSans *[]*string `field:"optional" json:"allowedUriSans" yaml:"allowedUriSans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#backend CertAuthBackendRole#backend}.
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#display_name CertAuthBackendRole#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#id CertAuthBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#namespace CertAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Any additional CA certificates needed to verify OCSP responses. Provided as base64 encoded PEM data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_ca_certificates CertAuthBackendRole#ocsp_ca_certificates}
	OcspCaCertificates *string `field:"optional" json:"ocspCaCertificates" yaml:"ocspCaCertificates"`
	// If enabled, validate certificates' revocation status using OCSP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_enabled CertAuthBackendRole#ocsp_enabled}
	OcspEnabled interface{} `field:"optional" json:"ocspEnabled" yaml:"ocspEnabled"`
	// If true and an OCSP response cannot be fetched or is of an unknown status, the login will proceed as if the certificate has not been revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_fail_open CertAuthBackendRole#ocsp_fail_open}
	OcspFailOpen interface{} `field:"optional" json:"ocspFailOpen" yaml:"ocspFailOpen"`
	// The number of retries to attempt when connecting to an OCSP server.
	//
	// Defaults to 4 retries. Must be a non-negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_max_retries CertAuthBackendRole#ocsp_max_retries}
	OcspMaxRetries *float64 `field:"optional" json:"ocspMaxRetries" yaml:"ocspMaxRetries"`
	// If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_query_all_servers CertAuthBackendRole#ocsp_query_all_servers}
	OcspQueryAllServers interface{} `field:"optional" json:"ocspQueryAllServers" yaml:"ocspQueryAllServers"`
	// A comma-separated list of OCSP server addresses.
	//
	// If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_servers_override CertAuthBackendRole#ocsp_servers_override}
	OcspServersOverride *[]*string `field:"optional" json:"ocspServersOverride" yaml:"ocspServersOverride"`
	// The maximum age in seconds of the 'thisUpdate' field in an OCSP response before it is considered too old.
	//
	// Defaults to 0 (disabled). Must be a non-negative value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#ocsp_this_update_max_age CertAuthBackendRole#ocsp_this_update_max_age}
	OcspThisUpdateMaxAge *float64 `field:"optional" json:"ocspThisUpdateMaxAge" yaml:"ocspThisUpdateMaxAge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#required_extensions CertAuthBackendRole#required_extensions}.
	RequiredExtensions *[]*string `field:"optional" json:"requiredExtensions" yaml:"requiredExtensions"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_bound_cidrs CertAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_explicit_max_ttl CertAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_max_ttl CertAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_no_default_policy CertAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_num_uses CertAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_period CertAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_policies CertAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_ttl CertAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/cert_auth_backend_role#token_type CertAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

