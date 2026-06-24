// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigurls

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendConfigUrlsConfig struct {
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
	// The path of the PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#backend PkiSecretBackendConfigUrls#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Specifies the URL values for the CRL Distribution Points field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#crl_distribution_points PkiSecretBackendConfigUrls#crl_distribution_points}
	CrlDistributionPoints *[]*string `field:"optional" json:"crlDistributionPoints" yaml:"crlDistributionPoints"`
	// Specifies that templating of AIA fields is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#enable_templating PkiSecretBackendConfigUrls#enable_templating}
	EnableTemplating interface{} `field:"optional" json:"enableTemplating" yaml:"enableTemplating"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#id PkiSecretBackendConfigUrls#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the URL values for the Issuing Certificate field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#issuing_certificates PkiSecretBackendConfigUrls#issuing_certificates}
	IssuingCertificates *[]*string `field:"optional" json:"issuingCertificates" yaml:"issuingCertificates"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#namespace PkiSecretBackendConfigUrls#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the URL values for the OCSP Servers field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_urls#ocsp_servers PkiSecretBackendConfigUrls#ocsp_servers}
	OcspServers *[]*string `field:"optional" json:"ocspServers" yaml:"ocspServers"`
}

