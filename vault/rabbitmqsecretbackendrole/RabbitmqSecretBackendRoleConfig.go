// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rabbitmqsecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RabbitmqSecretBackendRoleConfig struct {
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
	// The path of the Rabbitmq Secret Backend the role belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#backend RabbitmqSecretBackendRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Unique name for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#name RabbitmqSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#id RabbitmqSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#namespace RabbitmqSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies a comma-separated RabbitMQ management tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#tags RabbitmqSecretBackendRole#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// vhost block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#vhost RabbitmqSecretBackendRole#vhost}
	Vhost interface{} `field:"optional" json:"vhost" yaml:"vhost"`
	// vhost_topic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/rabbitmq_secret_backend_role#vhost_topic RabbitmqSecretBackendRole#vhost_topic}
	VhostTopic interface{} `field:"optional" json:"vhostTopic" yaml:"vhostTopic"`
}

