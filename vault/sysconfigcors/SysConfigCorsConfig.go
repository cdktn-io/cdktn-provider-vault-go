// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sysconfigcors

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SysConfigCorsConfig struct {
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
	// Set of origins permitted to make cross-origin requests. Use `*` as the only value to allow all origins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/sys_config_cors#allowed_origins SysConfigCors#allowed_origins}
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Set of additional custom headers allowed on cross-origin requests.
	//
	// Vault automatically includes standard headers (Content-Type, X-Requested-With, X-Vault-AWS-IAM-Server-ID, X-Vault-MFA, X-Vault-No-Request-Forwarding, X-Vault-Wrap-Format, X-Vault-Wrap-TTL, X-Vault-Policy-Override, Authorization, X-Vault-Token), so only specify custom headers here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/sys_config_cors#allowed_headers SysConfigCors#allowed_headers}
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
}

