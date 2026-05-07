// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmipsecretcagenerated

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KmipSecretCaGeneratedConfig struct {
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
	// CA key bits. Valid values depend on key_type: For rsa: 2048, 3072, 4096. For ec: 224, 256, 384, 521.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#key_bits KmipSecretCaGenerated#key_bits}
	KeyBits *float64 `field:"required" json:"keyBits" yaml:"keyBits"`
	// CA key type (rsa or ec).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#key_type KmipSecretCaGenerated#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Name to identify the CA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#name KmipSecretCaGenerated#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path where KMIP backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#path KmipSecretCaGenerated#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#namespace KmipSecretCaGenerated#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// CA TTL in seconds. Defaults to 365 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_ca_generated#ttl KmipSecretCaGenerated#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

