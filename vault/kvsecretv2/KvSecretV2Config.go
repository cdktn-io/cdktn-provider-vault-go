// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KvSecretV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#mount KvSecretV2#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Full name of the secret.
	//
	// For a nested secret, the name is the nested path excluding the mount and data prefix. For example, for a secret at 'kvv2/data/foo/bar/baz', the name is 'foo/bar/baz'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#name KvSecretV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// This flag is required if cas_required is set to true on either the secret or the engine's config.
	//
	// In order for a write to be successful, cas must be set to the current version of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#cas KvSecretV2#cas}
	Cas *float64 `field:"optional" json:"cas" yaml:"cas"`
	// custom_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#custom_metadata KvSecretV2#custom_metadata}
	CustomMetadata *KvSecretV2CustomMetadata `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// JSON-encoded secret data to write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#data_json KvSecretV2#data_json}
	DataJson *string `field:"optional" json:"dataJson" yaml:"dataJson"`
	// Write-Only JSON-encoded secret data to write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#data_json_wo KvSecretV2#data_json_wo}
	DataJsonWo *string `field:"optional" json:"dataJsonWo" yaml:"dataJsonWo"`
	// Version counter for write-only secret data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#data_json_wo_version KvSecretV2#data_json_wo_version}
	DataJsonWoVersion *float64 `field:"optional" json:"dataJsonWoVersion" yaml:"dataJsonWoVersion"`
	// If set to true, permanently deletes all versions for the specified key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#delete_all_versions KvSecretV2#delete_all_versions}
	DeleteAllVersions interface{} `field:"optional" json:"deleteAllVersions" yaml:"deleteAllVersions"`
	// If set to true, disables reading secret from Vault; note: drift won't be detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#disable_read KvSecretV2#disable_read}
	DisableRead interface{} `field:"optional" json:"disableRead" yaml:"disableRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#id KvSecretV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#namespace KvSecretV2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// An object that holds option settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kv_secret_v2#options KvSecretV2#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
}

