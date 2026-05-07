// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmipsecretlistener

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KmipSecretListenerConfig struct {
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
	// Host:port address to listen on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#address KmipSecretListener#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Name of the CA to use to generate the server certificate and verify client certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#ca KmipSecretListener#ca}
	Ca *string `field:"required" json:"ca" yaml:"ca"`
	// Unique name for the listener.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#name KmipSecretListener#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path where KMIP backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#path KmipSecretListener#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Names of additional TLS CAs to use to verify client certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#additional_client_cas KmipSecretListener#additional_client_cas}
	AdditionalClientCas *[]*string `field:"optional" json:"additionalClientCas" yaml:"additionalClientCas"`
	// Use the legacy unnamed CA for verifying client certificates as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#also_use_legacy_ca KmipSecretListener#also_use_legacy_ca}
	AlsoUseLegacyCa interface{} `field:"optional" json:"alsoUseLegacyCa" yaml:"alsoUseLegacyCa"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#namespace KmipSecretListener#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// DNS SANs to include in listener certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#server_hostnames KmipSecretListener#server_hostnames}
	ServerHostnames *[]*string `field:"optional" json:"serverHostnames" yaml:"serverHostnames"`
	// IP SANs to include in listener certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#server_ips KmipSecretListener#server_ips}
	ServerIps *[]*string `field:"optional" json:"serverIps" yaml:"serverIps"`
	// TLS cipher suites to allow (does not apply to tls13+).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#tls_cipher_suites KmipSecretListener#tls_cipher_suites}
	TlsCipherSuites *string `field:"optional" json:"tlsCipherSuites" yaml:"tlsCipherSuites"`
	// Maximum TLS version to accept (tls12 or tls13).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#tls_max_version KmipSecretListener#tls_max_version}
	TlsMaxVersion *string `field:"optional" json:"tlsMaxVersion" yaml:"tlsMaxVersion"`
	// Minimum TLS version to accept (tls12 or tls13).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kmip_secret_listener#tls_min_version KmipSecretListener#tls_min_version}
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

