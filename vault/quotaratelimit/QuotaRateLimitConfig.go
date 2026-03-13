// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quotaratelimit

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuotaRateLimitConfig struct {
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
	// The name of the quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#name QuotaRateLimit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The maximum number of requests at any given second to be allowed by the quota rule.
	//
	// The rate must be positive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#rate QuotaRateLimit#rate}
	Rate *float64 `field:"required" json:"rate" yaml:"rate"`
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#block_interval QuotaRateLimit#block_interval}
	BlockInterval *float64 `field:"optional" json:"blockInterval" yaml:"blockInterval"`
	// Attribute used to group requests for rate limiting.
	//
	// Limits are enforced independently for each group. Valid group_by modes are: 1) "ip" that groups requests by their source IP address (group_by defaults to ip if unset); 2) "none" that groups all requests that match the rate limit quota rule together; 3) "entity_then_ip" that groups requests by their entity ID for authenticated requests that carry one, or by their IP for unauthenticated requests (or requests whose authentication is not connected to an entity); and 4) "entity_then_none" which also groups requests by their entity ID when available, but the rest is all grouped together (i.e. unauthenticated or with authentication not connected to an entity).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#group_by QuotaRateLimit#group_by}
	GroupBy *string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#id QuotaRateLimit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set to true on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace.
	//
	// The inheritable parameter cannot be set to true if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#inheritable QuotaRateLimit#inheritable}
	Inheritable interface{} `field:"optional" json:"inheritable" yaml:"inheritable"`
	// The duration in seconds to enforce rate limiting for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#interval QuotaRateLimit#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#namespace QuotaRateLimit#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#path QuotaRateLimit#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#role QuotaRateLimit#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Only available when using the "entity_then_ip" or "entity_then_none" group_by modes.
	//
	// This is the rate limit applied to the requests that fall under the "ip" or "none" groupings, while the authenticated requests that contain an entity ID are subject to the "rate" field instead. Defaults to the same value as "rate".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/quota_rate_limit#secondary_rate QuotaRateLimit#secondary_rate}
	SecondaryRate *float64 `field:"optional" json:"secondaryRate" yaml:"secondaryRate"`
}

