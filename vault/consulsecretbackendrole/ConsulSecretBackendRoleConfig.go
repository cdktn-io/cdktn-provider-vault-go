// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package consulsecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConsulSecretBackendRoleConfig struct {
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
	// The name of an existing role against which to create this Consul credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#name ConsulSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the Consul Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#backend ConsulSecretBackendRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// The Consul namespace that the token will be created in. Applicable for Vault 1.10+ and Consul 1.7+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#consul_namespace ConsulSecretBackendRole#consul_namespace}
	ConsulNamespace *string `field:"optional" json:"consulNamespace" yaml:"consulNamespace"`
	// List of Consul policies to associate with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#consul_policies ConsulSecretBackendRole#consul_policies}
	ConsulPolicies *[]*string `field:"optional" json:"consulPolicies" yaml:"consulPolicies"`
	// Set of Consul roles to attach to the token. Applicable for Vault 1.10+ with Consul 1.5+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#consul_roles ConsulSecretBackendRole#consul_roles}
	ConsulRoles *[]*string `field:"optional" json:"consulRoles" yaml:"consulRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#id ConsulSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates that the token should not be replicated globally and instead be local to the current datacenter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#local ConsulSecretBackendRole#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum TTL for leases associated with this role, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#max_ttl ConsulSecretBackendRole#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#namespace ConsulSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set of Consul node identities to attach to 				the token. Applicable for Vault 1.11+ with Consul 1.8+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#node_identities ConsulSecretBackendRole#node_identities}
	NodeIdentities *[]*string `field:"optional" json:"nodeIdentities" yaml:"nodeIdentities"`
	// The Consul admin partition that the token will be created in. Applicable for Vault 1.10+ and Consul 1.11+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#partition ConsulSecretBackendRole#partition}
	Partition *string `field:"optional" json:"partition" yaml:"partition"`
	// List of Consul policies to associate with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#policies ConsulSecretBackendRole#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Set of Consul service identities to attach to 				the token. Applicable for Vault 1.11+ with Consul 1.5+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#service_identities ConsulSecretBackendRole#service_identities}
	ServiceIdentities *[]*string `field:"optional" json:"serviceIdentities" yaml:"serviceIdentities"`
	// Specifies the TTL for this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend_role#ttl ConsulSecretBackendRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

