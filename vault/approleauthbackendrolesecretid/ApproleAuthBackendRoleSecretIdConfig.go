// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package approleauthbackendrolesecretid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApproleAuthBackendRoleSecretIdConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#role_name ApproleAuthBackendRoleSecretId#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#backend ApproleAuthBackendRoleSecretId#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// List of CIDR blocks that can log in using the SecretID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#cidr_list ApproleAuthBackendRoleSecretId#cidr_list}
	CidrList *[]*string `field:"optional" json:"cidrList" yaml:"cidrList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#id ApproleAuthBackendRoleSecretId#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// JSON-encoded secret data to write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#metadata ApproleAuthBackendRoleSecretId#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#namespace ApproleAuthBackendRoleSecretId#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The number of uses for the secret-id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#num_uses ApproleAuthBackendRoleSecretId#num_uses}
	NumUses *float64 `field:"optional" json:"numUses" yaml:"numUses"`
	// The SecretID to be managed. If not specified, Vault auto-generates one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#secret_id ApproleAuthBackendRoleSecretId#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#token_bound_cidrs ApproleAuthBackendRoleSecretId#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// The TTL duration of the SecretID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#ttl ApproleAuthBackendRoleSecretId#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Use the wrapped secret-id accessor as the id of this resource.
	//
	// If false, a fresh secret-id will be regenerated whenever the wrapping token is expired or invalidated through unwrapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#with_wrapped_accessor ApproleAuthBackendRoleSecretId#with_wrapped_accessor}
	WithWrappedAccessor interface{} `field:"optional" json:"withWrappedAccessor" yaml:"withWrappedAccessor"`
	// The TTL duration of the wrapped SecretID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id#wrapping_ttl ApproleAuthBackendRoleSecretId#wrapping_ttl}
	WrappingTtl *string `field:"optional" json:"wrappingTtl" yaml:"wrappingTtl"`
}

