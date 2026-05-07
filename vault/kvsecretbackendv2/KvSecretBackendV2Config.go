// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretbackendv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KvSecretBackendV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#mount KvSecretBackendV2#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// If true, all keys will require the cas parameter to be set on all write requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#cas_required KvSecretBackendV2#cas_required}
	CasRequired interface{} `field:"optional" json:"casRequired" yaml:"casRequired"`
	// If set, specifies the length of time before a version is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#delete_version_after KvSecretBackendV2#delete_version_after}
	DeleteVersionAfter *float64 `field:"optional" json:"deleteVersionAfter" yaml:"deleteVersionAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#id KvSecretBackendV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of versions to keep per key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#max_versions KvSecretBackendV2#max_versions}
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_backend_v2#namespace KvSecretBackendV2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

