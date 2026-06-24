// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongodbatlassecretrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MongodbatlasSecretRoleConfig struct {
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
	// Path where MongoDB Atlas secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#mount MongodbatlasSecretRole#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#name MongodbatlasSecretRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of roles that the API Key needs to have.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#roles MongodbatlasSecretRole#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Whitelist entry in CIDR notation to be added for the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#cidr_blocks MongodbatlasSecretRole#cidr_blocks}
	CidrBlocks *[]*string `field:"optional" json:"cidrBlocks" yaml:"cidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#id MongodbatlasSecretRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IP address to be added to the whitelist for the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#ip_addresses MongodbatlasSecretRole#ip_addresses}
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// The maximum allowed lifetime of credentials issued using this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#max_ttl MongodbatlasSecretRole#max_ttl}
	MaxTtl *string `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#namespace MongodbatlasSecretRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// ID for the organization to which the target API Key belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#organization_id MongodbatlasSecretRole#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// ID for the project to which the target API Key belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#project_id MongodbatlasSecretRole#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Roles assigned when an org API key is assigned to a project API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#project_roles MongodbatlasSecretRole#project_roles}
	ProjectRoles *[]*string `field:"optional" json:"projectRoles" yaml:"projectRoles"`
	// Duration in seconds after which the issued credential should expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/mongodbatlas_secret_role#ttl MongodbatlasSecretRole#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

