// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapauthbackenduser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LdapAuthBackendUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#username LdapAuthBackendUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#backend LdapAuthBackendUser#backend}.
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#groups LdapAuthBackendUser#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#id LdapAuthBackendUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#namespace LdapAuthBackendUser#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend_user#policies LdapAuthBackendUser#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

