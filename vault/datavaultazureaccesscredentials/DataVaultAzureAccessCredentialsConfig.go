// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultazureaccesscredentials

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultAzureAccessCredentialsConfig struct {
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
	// Azure Secret Backend to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#backend DataVaultAzureAccessCredentials#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Azure Secret Role to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#role DataVaultAzureAccessCredentials#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The Azure environment to use during credential validation. Defaults to the Azure Public Cloud. Some possible values: AzurePublicCloud, AzureUSGovernmentCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#environment DataVaultAzureAccessCredentials#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#id DataVaultAzureAccessCredentials#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If 'validate_creds' is true, the number of seconds after which to give up validating credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#max_cred_validation_seconds DataVaultAzureAccessCredentials#max_cred_validation_seconds}
	MaxCredValidationSeconds *float64 `field:"optional" json:"maxCredValidationSeconds" yaml:"maxCredValidationSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#namespace DataVaultAzureAccessCredentials#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If 'validate_creds' is true, the number of seconds to wait between each test of generated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#num_seconds_between_tests DataVaultAzureAccessCredentials#num_seconds_between_tests}
	NumSecondsBetweenTests *float64 `field:"optional" json:"numSecondsBetweenTests" yaml:"numSecondsBetweenTests"`
	// If 'validate_creds' is true, the number of sequential successes required to validate generated credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#num_sequential_successes DataVaultAzureAccessCredentials#num_sequential_successes}
	NumSequentialSuccesses *float64 `field:"optional" json:"numSequentialSuccesses" yaml:"numSequentialSuccesses"`
	// The subscription ID to use during credential validation. Defaults to the subscription ID configured in the Vault backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#subscription_id DataVaultAzureAccessCredentials#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The tenant ID to use during credential validation. Defaults to the tenant ID configured in the Vault backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#tenant_id DataVaultAzureAccessCredentials#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Whether generated credentials should be validated before being returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/azure_access_credentials#validate_creds DataVaultAzureAccessCredentials#validate_creds}
	ValidateCreds interface{} `field:"optional" json:"validateCreds" yaml:"validateCreds"`
}

