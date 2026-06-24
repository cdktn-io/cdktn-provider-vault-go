// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultkvsecretslistv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultKvSecretsListV2Config struct {
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
	// Path where KV-V2 engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kv_secrets_list_v2#mount DataVaultKvSecretsListV2#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kv_secrets_list_v2#id DataVaultKvSecretsListV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Full named path of the secret.
	//
	// For a nested secret, the name is the nested path excluding the mount and data prefix. For example, for a secret at 'kvv2/data/foo/bar/baz', the name is 'foo/bar/baz'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kv_secrets_list_v2#name DataVaultKvSecretsListV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kv_secrets_list_v2#namespace DataVaultKvSecretsListV2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

