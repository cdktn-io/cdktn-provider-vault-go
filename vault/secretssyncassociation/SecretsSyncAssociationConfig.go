// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncAssociationConfig struct {
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
	// Specifies the mount where the secret is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#mount SecretsSyncAssociation#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#name SecretsSyncAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the name of the secret to synchronize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#secret_name SecretsSyncAssociation#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Type of sync destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#type SecretsSyncAssociation#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#id SecretsSyncAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_association#namespace SecretsSyncAssociation#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

