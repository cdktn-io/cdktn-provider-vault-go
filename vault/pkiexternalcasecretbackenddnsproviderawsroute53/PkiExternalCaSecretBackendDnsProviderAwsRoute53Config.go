// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackenddnsproviderawsroute53

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiExternalCaSecretBackendDnsProviderAwsRoute53Config struct {
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
	// List of domain identifiers this provider handles. Supports wildcard patterns with leftmost `*` (e.g. `*.example.com`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#identifiers PkiExternalCaSecretBackendDnsProviderAwsRoute53#identifiers}
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#mount PkiExternalCaSecretBackendDnsProviderAwsRoute53#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the DNS provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#name PkiExternalCaSecretBackendDnsProviderAwsRoute53#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS access key ID for Route53 API access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#access_key_id PkiExternalCaSecretBackendDnsProviderAwsRoute53#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// AWS IAM role ARN to assume for Route53 operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#assume_role_arn PkiExternalCaSecretBackendDnsProviderAwsRoute53#assume_role_arn}
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID for AWS STS AssumeRole.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#external_id PkiExternalCaSecretBackendDnsProviderAwsRoute53#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// AWS Route53 hosted zone ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#hosted_zone_id PkiExternalCaSecretBackendDnsProviderAwsRoute53#hosted_zone_id}
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Address of a DNS nameserver (`host` or `host:port`) to use when verifying DNS-01 challenge propagation instead of the domain's primary nameserver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#nameserver PkiExternalCaSecretBackendDnsProviderAwsRoute53#nameserver}
	Nameserver *string `field:"optional" json:"nameserver" yaml:"nameserver"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#namespace PkiExternalCaSecretBackendDnsProviderAwsRoute53#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// AWS region for Route53 operations. Defaults to `us-east-1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#region PkiExternalCaSecretBackendDnsProviderAwsRoute53#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// AWS secret access key for Route53 API access. Write-only — not returned by Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#secret_access_key_wo PkiExternalCaSecretBackendDnsProviderAwsRoute53#secret_access_key_wo}
	SecretAccessKeyWo *string `field:"optional" json:"secretAccessKeyWo" yaml:"secretAccessKeyWo"`
	// Version counter for the write-only `secret_access_key` field.
	//
	// Increment this value to trigger an update to the secret access key in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#secret_access_key_wo_version PkiExternalCaSecretBackendDnsProviderAwsRoute53#secret_access_key_wo_version}
	SecretAccessKeyWoVersion *float64 `field:"optional" json:"secretAccessKeyWoVersion" yaml:"secretAccessKeyWoVersion"`
	// TTL for DNS TXT records used in DNS-01 challenges in seconds. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_aws_route53#ttl PkiExternalCaSecretBackendDnsProviderAwsRoute53#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

