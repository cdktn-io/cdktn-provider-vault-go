// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quotaconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuotaConfigConfig struct {
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
	// Absolute paths exempt from all rate limit quotas, qualified from the root of the namespace hierarchy.
	//
	// This field is effectively root-managed; administrative namespaces can read returned values but cannot reliably manage them. Order is not significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/quota_config#absolute_rate_limit_exempt_paths QuotaConfig#absolute_rate_limit_exempt_paths}
	AbsoluteRateLimitExemptPaths *[]*string `field:"optional" json:"absoluteRateLimitExemptPaths" yaml:"absoluteRateLimitExemptPaths"`
	// Enables audit logging for requests rejected by rate limit quotas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/quota_config#enable_rate_limit_audit_logging QuotaConfig#enable_rate_limit_audit_logging}
	EnableRateLimitAuditLogging interface{} `field:"optional" json:"enableRateLimitAuditLogging" yaml:"enableRateLimitAuditLogging"`
	// Enables rate limit response headers on HTTP responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/quota_config#enable_rate_limit_response_headers QuotaConfig#enable_rate_limit_response_headers}
	EnableRateLimitResponseHeaders interface{} `field:"optional" json:"enableRateLimitResponseHeaders" yaml:"enableRateLimitResponseHeaders"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/quota_config#namespace QuotaConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Paths exempt from rate limit quotas relative to the current namespace context.
	//
	// This endpoint is only callable from the root or an administrative namespace, and exemption updates are effectively root-managed. Order is not significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/quota_config#rate_limit_exempt_paths QuotaConfig#rate_limit_exempt_paths}
	RateLimitExemptPaths *[]*string `field:"optional" json:"rateLimitExemptPaths" yaml:"rateLimitExemptPaths"`
}

