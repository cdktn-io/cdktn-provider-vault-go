// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oktaauthbackenduser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OktaAuthBackendUserAConfig struct {
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
	// Path to the Okta auth backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#path OktaAuthBackendUserA#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Name of the user within Okta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#username OktaAuthBackendUserA#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Groups within the Okta auth backend to associate with this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#groups OktaAuthBackendUserA#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#id OktaAuthBackendUserA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#namespace OktaAuthBackendUserA#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Policies to associate with this user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend_user#policies OktaAuthBackendUserA#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

