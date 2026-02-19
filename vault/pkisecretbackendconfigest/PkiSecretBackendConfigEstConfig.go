// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigest

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigEstConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#backend PkiSecretBackendConfigEst#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#audit_fields PkiSecretBackendConfigEst#audit_fields}
	AuditFields *[]*string `field:"optional" json:"auditFields" yaml:"auditFields"`
	// authenticators block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#authenticators PkiSecretBackendConfigEst#authenticators}
	Authenticators *PkiSecretBackendConfigEstAuthenticators `field:"optional" json:"authenticators" yaml:"authenticators"`
	// If set, this mount will register the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#default_mount PkiSecretBackendConfigEst#default_mount}
	DefaultMount interface{} `field:"optional" json:"defaultMount" yaml:"defaultMount"`
	// Required to be set if default_mount is enabled.
	//
	// Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:<role_name>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#default_path_policy PkiSecretBackendConfigEst#default_path_policy}
	DefaultPathPolicy *string `field:"optional" json:"defaultPathPolicy" yaml:"defaultPathPolicy"`
	// Specifies whether EST is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#enabled PkiSecretBackendConfigEst#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If set, parse out fields from the provided CSR making them available for Sentinel policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#enable_sentinel_parsing PkiSecretBackendConfigEst#enable_sentinel_parsing}
	EnableSentinelParsing interface{} `field:"optional" json:"enableSentinelParsing" yaml:"enableSentinelParsing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#id PkiSecretBackendConfigEst#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Configures a pairing of an EST label with the redirected behavior for requests hitting that role.
	//
	// The path policy can be sign-verbatim or a role given by role:<role_name>. Labels must be unique across Vault cluster, and will register .well-known/est/<label> URL paths
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#label_to_path_policy PkiSecretBackendConfigEst#label_to_path_policy}
	LabelToPathPolicy *map[string]*string `field:"optional" json:"labelToPathPolicy" yaml:"labelToPathPolicy"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_est#namespace PkiSecretBackendConfigEst#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

