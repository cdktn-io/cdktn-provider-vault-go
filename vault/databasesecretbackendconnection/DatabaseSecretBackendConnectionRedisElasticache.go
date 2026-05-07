// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionRedisElasticache struct {
	// The configuration endpoint for the ElastiCache cluster to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_connection#url DatabaseSecretBackendConnection#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The AWS secret key id to use to talk to ElastiCache.
	//
	// If omitted the credentials chain provider is used instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The AWS region where the ElastiCache cluster is hosted.
	//
	// If omitted the plugin tries to infer the region from the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_connection#region DatabaseSecretBackendConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The AWS access key id to use to talk to ElastiCache.
	//
	// If omitted the credentials chain provider is used instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

