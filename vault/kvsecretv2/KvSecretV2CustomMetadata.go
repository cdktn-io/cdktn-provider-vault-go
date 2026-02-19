// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretv2


type KvSecretV2CustomMetadata struct {
	// If true, all keys will require the cas parameter to be set on all write requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/kv_secret_v2#cas_required KvSecretV2#cas_required}
	CasRequired interface{} `field:"optional" json:"casRequired" yaml:"casRequired"`
	// A map of arbitrary string to string valued user-provided metadata meant to describe the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/kv_secret_v2#data KvSecretV2#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// If set, specifies the length of time before a version is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/kv_secret_v2#delete_version_after KvSecretV2#delete_version_after}
	DeleteVersionAfter *float64 `field:"optional" json:"deleteVersionAfter" yaml:"deleteVersionAfter"`
	// The number of versions to keep per key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/kv_secret_v2#max_versions KvSecretV2#max_versions}
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
}

