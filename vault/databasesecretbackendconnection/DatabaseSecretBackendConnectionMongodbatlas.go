// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionMongodbatlas struct {
	// The Private Programmatic API Key used to connect with MongoDB Atlas API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#private_key DatabaseSecretBackendConnection#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// The Project ID the Database User should be created within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#project_id DatabaseSecretBackendConnection#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#public_key DatabaseSecretBackendConnection#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#username_template DatabaseSecretBackendConnection#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

