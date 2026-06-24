// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigcmpv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigCmpv2Config struct {
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
	// The PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#backend PkiSecretBackendConfigCmpv2#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#audit_fields PkiSecretBackendConfigCmpv2#audit_fields}
	AuditFields *[]*string `field:"optional" json:"auditFields" yaml:"auditFields"`
	// authenticators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#authenticators PkiSecretBackendConfigCmpv2#authenticators}
	Authenticators *PkiSecretBackendConfigCmpv2Authenticators `field:"optional" json:"authenticators" yaml:"authenticators"`
	// Can be sign-verbatim or a role given by role:<role_name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#default_path_policy PkiSecretBackendConfigCmpv2#default_path_policy}
	DefaultPathPolicy *string `field:"optional" json:"defaultPathPolicy" yaml:"defaultPathPolicy"`
	// A comma-separated list of validations not to perform on CMPv2 messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#disabled_validations PkiSecretBackendConfigCmpv2#disabled_validations}
	DisabledValidations *[]*string `field:"optional" json:"disabledValidations" yaml:"disabledValidations"`
	// Specifies whether CMPv2 is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#enabled PkiSecretBackendConfigCmpv2#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If set, parse out fields from the provided CSR making them available for Sentinel policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#enable_sentinel_parsing PkiSecretBackendConfigCmpv2#enable_sentinel_parsing}
	EnableSentinelParsing interface{} `field:"optional" json:"enableSentinelParsing" yaml:"enableSentinelParsing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#id PkiSecretBackendConfigCmpv2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_cmpv2#namespace PkiSecretBackendConfigCmpv2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

