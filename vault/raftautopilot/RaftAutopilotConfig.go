// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package raftautopilot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RaftAutopilotConfig struct {
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
	// Specifies whether to remove dead server nodes periodically or when a new server joins.
	//
	// This requires that min-quorum is also set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#cleanup_dead_servers RaftAutopilot#cleanup_dead_servers}
	CleanupDeadServers interface{} `field:"optional" json:"cleanupDeadServers" yaml:"cleanupDeadServers"`
	// Limit the amount of time a server can go without leader contact before being considered failed.
	//
	// This only takes effect when cleanup_dead_servers is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#dead_server_last_contact_threshold RaftAutopilot#dead_server_last_contact_threshold}
	DeadServerLastContactThreshold *string `field:"optional" json:"deadServerLastContactThreshold" yaml:"deadServerLastContactThreshold"`
	// Disables automatically upgrading Vault using autopilot. (Enterprise-only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#disable_upgrade_migration RaftAutopilot#disable_upgrade_migration}
	DisableUpgradeMigration interface{} `field:"optional" json:"disableUpgradeMigration" yaml:"disableUpgradeMigration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#id RaftAutopilot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Limit the amount of time a server can go without leader contact before being considered unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#last_contact_threshold RaftAutopilot#last_contact_threshold}
	LastContactThreshold *string `field:"optional" json:"lastContactThreshold" yaml:"lastContactThreshold"`
	// Maximum number of log entries in the Raft log that a server can be behind its leader before being considered unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#max_trailing_logs RaftAutopilot#max_trailing_logs}
	MaxTrailingLogs *float64 `field:"optional" json:"maxTrailingLogs" yaml:"maxTrailingLogs"`
	// Minimum number of servers allowed in a cluster before autopilot can prune dead servers.
	//
	// This should at least be 3. Applicable only for voting nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#min_quorum RaftAutopilot#min_quorum}
	MinQuorum *float64 `field:"optional" json:"minQuorum" yaml:"minQuorum"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#namespace RaftAutopilot#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Minimum amount of time a server must be stable in the 'healthy' state before being added to the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/raft_autopilot#server_stabilization_time RaftAutopilot#server_stabilization_time}
	ServerStabilizationTime *string `field:"optional" json:"serverStabilizationTime" yaml:"serverStabilizationTime"`
}

