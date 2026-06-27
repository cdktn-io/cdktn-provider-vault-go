// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oktaauthbackend


type OktaAuthBackendTune struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#allowed_response_headers OktaAuthBackend#allowed_response_headers}.
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#audit_non_hmac_request_keys OktaAuthBackend#audit_non_hmac_request_keys}.
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#audit_non_hmac_response_keys OktaAuthBackend#audit_non_hmac_response_keys}.
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#default_lease_ttl OktaAuthBackend#default_lease_ttl}.
	DefaultLeaseTtl *string `field:"optional" json:"defaultLeaseTtl" yaml:"defaultLeaseTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#listing_visibility OktaAuthBackend#listing_visibility}.
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#max_lease_ttl OktaAuthBackend#max_lease_ttl}.
	MaxLeaseTtl *string `field:"optional" json:"maxLeaseTtl" yaml:"maxLeaseTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#passthrough_request_headers OktaAuthBackend#passthrough_request_headers}.
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_type OktaAuthBackend#token_type}.
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

