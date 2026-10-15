// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kerberosauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KerberosAuthBackendConfigConfig struct {
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
	// Base64-encoded keytab file content (write-only). Must contain an entry matching service_account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#keytab_wo KerberosAuthBackendConfig#keytab_wo}
	KeytabWo *string `field:"required" json:"keytabWo" yaml:"keytabWo"`
	// Version identifier for keytab updates. Increment this value to trigger a keytab update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#keytab_wo_version KerberosAuthBackendConfig#keytab_wo_version}
	KeytabWoVersion *float64 `field:"required" json:"keytabWoVersion" yaml:"keytabWoVersion"`
	// Path where the Kerberos auth method is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#mount KerberosAuthBackendConfig#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The Kerberos service account associated with the keytab entry (e.g., 'vault_svc').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#service_account KerberosAuthBackendConfig#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// Adds group aliases during authentication. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#add_group_aliases KerberosAuthBackendConfig#add_group_aliases}
	AddGroupAliases interface{} `field:"optional" json:"addGroupAliases" yaml:"addGroupAliases"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#namespace KerberosAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Removes instance names from Kerberos service principal names. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/kerberos_auth_backend_config#remove_instance_name KerberosAuthBackendConfig#remove_instance_name}
	RemoveInstanceName interface{} `field:"optional" json:"removeInstanceName" yaml:"removeInstanceName"`
}

