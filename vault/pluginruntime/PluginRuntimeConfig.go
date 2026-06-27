// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pluginruntime

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PluginRuntimeConfig struct {
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
	// The name of the plugin runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#name PluginRuntime#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the plugin runtime type. Currently only `container` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#type PluginRuntime#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies the parent cgroup to set for each container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#cgroup_parent PluginRuntime#cgroup_parent}
	CgroupParent *string `field:"optional" json:"cgroupParent" yaml:"cgroupParent"`
	// Specifies CPU limit to set per container in billionths of a CPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#cpu_nanos PluginRuntime#cpu_nanos}
	CpuNanos *float64 `field:"optional" json:"cpuNanos" yaml:"cpuNanos"`
	// Specifies memory limit to set per container in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#memory_bytes PluginRuntime#memory_bytes}
	MemoryBytes *float64 `field:"optional" json:"memoryBytes" yaml:"memoryBytes"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#namespace PluginRuntime#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies OCI-compliant container runtime to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#oci_runtime PluginRuntime#oci_runtime}
	OciRuntime *string `field:"optional" json:"ociRuntime" yaml:"ociRuntime"`
	// Whether the container runtime is running as a non-privileged user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/plugin_runtime#rootless PluginRuntime#rootless}
	Rootless interface{} `field:"optional" json:"rootless" yaml:"rootless"`
}

