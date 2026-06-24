// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package radiusauthbackenduser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RadiusAuthBackendUserConfig struct {
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
	// Path to the RADIUS auth mount where the user will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/radius_auth_backend_user#mount RadiusAuthBackendUser#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The username to register with the RADIUS auth backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/radius_auth_backend_user#username RadiusAuthBackendUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/radius_auth_backend_user#namespace RadiusAuthBackendUser#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A set of Vault policies to associate with this user.
	//
	// If not set, only the `default` policy will be applicable to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/radius_auth_backend_user#policies RadiusAuthBackendUser#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

