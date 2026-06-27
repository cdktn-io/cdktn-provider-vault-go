// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quotaleasecount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuotaLeaseCountConfig struct {
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
	// The maximum number of leases to be allowed by the quota rule. The max_leases must be positive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#max_leases QuotaLeaseCount#max_leases}
	MaxLeases *float64 `field:"required" json:"maxLeases" yaml:"maxLeases"`
	// The name of the quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#name QuotaLeaseCount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#id QuotaLeaseCount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set to true on a quota where path is set to a namespace, the same quota will be cumulatively applied to all child namespace.
	//
	// The inheritable parameter cannot be set to true if the path does not specify a namespace. Only the quotas associated with the root namespace are inheritable by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#inheritable QuotaLeaseCount#inheritable}
	Inheritable interface{} `field:"optional" json:"inheritable" yaml:"inheritable"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#namespace QuotaLeaseCount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Path of the mount or namespace to apply the quota. A blank path configures a global lease count quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#path QuotaLeaseCount#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If set on a quota where path is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/quota_lease_count#role QuotaLeaseCount#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

