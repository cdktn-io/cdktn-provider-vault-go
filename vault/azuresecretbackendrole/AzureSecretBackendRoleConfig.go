// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AzureSecretBackendRoleConfig struct {
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
	// Name of the role to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#role AzureSecretBackendRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Application Object ID for an existing service principal that will be used instead of creating dynamic service principals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#application_object_id AzureSecretBackendRole#application_object_id}
	ApplicationObjectId *string `field:"optional" json:"applicationObjectId" yaml:"applicationObjectId"`
	// azure_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#azure_groups AzureSecretBackendRole#azure_groups}
	AzureGroups interface{} `field:"optional" json:"azureGroups" yaml:"azureGroups"`
	// azure_roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#azure_roles AzureSecretBackendRole#azure_roles}
	AzureRoles interface{} `field:"optional" json:"azureRoles" yaml:"azureRoles"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#backend AzureSecretBackendRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#description AzureSecretBackendRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the explicit maximum lifetime of the lease and service principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#explicit_max_ttl AzureSecretBackendRole#explicit_max_ttl}
	ExplicitMaxTtl *string `field:"optional" json:"explicitMaxTtl" yaml:"explicitMaxTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#id AzureSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum TTL for service principals generated using this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#max_ttl AzureSecretBackendRole#max_ttl}
	MaxTtl *string `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#namespace AzureSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Indicates whether the applications and service principals created by Vault will be permanently deleted when the corresponding leases expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#permanently_delete AzureSecretBackendRole#permanently_delete}
	PermanentlyDelete interface{} `field:"optional" json:"permanentlyDelete" yaml:"permanentlyDelete"`
	// If true, persists the created service principal and application for the lifetime of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#persist_app AzureSecretBackendRole#persist_app}
	PersistApp interface{} `field:"optional" json:"persistApp" yaml:"persistApp"`
	// Specifies the security principal types that are allowed to sign in to the application.
	//
	// Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#sign_in_audience AzureSecretBackendRole#sign_in_audience}
	SignInAudience *string `field:"optional" json:"signInAudience" yaml:"signInAudience"`
	// Comma-separated strings of Azure tags to attach to an application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#tags AzureSecretBackendRole#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the default TTL for service principals generated using this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/azure_secret_backend_role#ttl AzureSecretBackendRole#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

