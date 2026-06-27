// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpkisecretbackendconfigcmpv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultPkiSecretBackendConfigCmpv2Config struct {
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
	// Path where PKI engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_config_cmpv2#backend DataVaultPkiSecretBackendConfigCmpv2#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// A comma-separated list of validations not to perform on CMPv2 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_config_cmpv2#disabled_validations DataVaultPkiSecretBackendConfigCmpv2#disabled_validations}
	DisabledValidations *[]*string `field:"optional" json:"disabledValidations" yaml:"disabledValidations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_config_cmpv2#id DataVaultPkiSecretBackendConfigCmpv2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_config_cmpv2#namespace DataVaultPkiSecretBackendConfigCmpv2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

