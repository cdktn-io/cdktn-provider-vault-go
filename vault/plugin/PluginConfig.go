// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package plugin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PluginConfig struct {
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
	// Name of the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#name Plugin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of plugin; one of "auth", "secret", or "database".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#type Plugin#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of additional arguments to pass to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#args Plugin#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Command to execute the plugin, relative to the plugin_directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#command Plugin#command}
	Command *string `field:"optional" json:"command" yaml:"command"`
	// List of additional environment variables to run the plugin with in KEY=VALUE form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#env Plugin#env}
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#id Plugin#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// OCI image to run.
	//
	// If specified, setting command, args, and env will update the container's entrypoint, args, and environment variables (append-only) respectively.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#oci_image Plugin#oci_image}
	OciImage *string `field:"optional" json:"ociImage" yaml:"ociImage"`
	// Vault plugin runtime to use if oci_image is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#runtime Plugin#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// SHA256 sum of the plugin binary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#sha256 Plugin#sha256}
	Sha256 *string `field:"optional" json:"sha256" yaml:"sha256"`
	// Semantic version of the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin#version Plugin#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

