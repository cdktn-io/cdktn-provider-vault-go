// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rotationpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RotationPolicyConfig struct {
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
	// Maximum retries per cycle for this rotation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/rotation_policy#max_retries_per_cycle RotationPolicy#max_retries_per_cycle}
	MaxRetriesPerCycle *float64 `field:"required" json:"maxRetriesPerCycle" yaml:"maxRetriesPerCycle"`
	// Maximum retry cycles for this rotation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/rotation_policy#max_retry_cycles RotationPolicy#max_retry_cycles}
	MaxRetryCycles *float64 `field:"required" json:"maxRetryCycles" yaml:"maxRetryCycles"`
	// Name of the rotation policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/rotation_policy#name RotationPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/rotation_policy#namespace RotationPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

