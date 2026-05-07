// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncConfigConfig struct {
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
	// Disables the syncing process between Vault and external destinations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_config#disabled SecretsSyncConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_config#id SecretsSyncConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_config#namespace SecretsSyncConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Maximum number of pending sync operations allowed on the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_config#queue_capacity SecretsSyncConfig#queue_capacity}
	QueueCapacity *float64 `field:"optional" json:"queueCapacity" yaml:"queueCapacity"`
}

