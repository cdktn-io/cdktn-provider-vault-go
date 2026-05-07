// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azureauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AzureAuthBackendConfigConfig struct {
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
	// The configured URL for the application registered in Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#resource AzureAuthBackendConfig#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The tenant id for the Azure Active Directory organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#tenant_id AzureAuthBackendConfig#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#backend AzureAuthBackendConfig#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#client_id AzureAuthBackendConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret for credentials to query the Azure APIs. Mutually exclusive with 'client_secret_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#client_secret AzureAuthBackendConfig#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The client secret for credentials to query the Azure APIs.
	//
	// This field is write-only and will never be stored in state. Mutually exclusive with 'client_secret'. Requires 'client_secret_wo_version' to trigger updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#client_secret_wo AzureAuthBackendConfig#client_secret_wo}
	ClientSecretWo *string `field:"optional" json:"clientSecretWo" yaml:"clientSecretWo"`
	// Version counter for the write-only client secret.
	//
	// Increment this value to trigger rotation of the client secret. Required when using 'client_secret_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#client_secret_wo_version AzureAuthBackendConfig#client_secret_wo_version}
	ClientSecretWoVersion *float64 `field:"optional" json:"clientSecretWoVersion" yaml:"clientSecretWoVersion"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#disable_automated_rotation AzureAuthBackendConfig#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#environment AzureAuthBackendConfig#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#id AzureAuthBackendConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#identity_token_audience AzureAuthBackendConfig#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The TTL of generated identity tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#identity_token_ttl AzureAuthBackendConfig#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Maximum number of retries for Azure API requests. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#max_retries AzureAuthBackendConfig#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The maximum delay in seconds between retries for Azure API requests. Defaults to 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#max_retry_delay AzureAuthBackendConfig#max_retry_delay}
	MaxRetryDelay *float64 `field:"optional" json:"maxRetryDelay" yaml:"maxRetryDelay"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#namespace AzureAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The initial delay in seconds between retries for Azure API requests. Defaults to 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#retry_delay AzureAuthBackendConfig#retry_delay}
	RetryDelay *float64 `field:"optional" json:"retryDelay" yaml:"retryDelay"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#rotation_period AzureAuthBackendConfig#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#rotation_schedule AzureAuthBackendConfig#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config#rotation_window AzureAuthBackendConfig#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
}

