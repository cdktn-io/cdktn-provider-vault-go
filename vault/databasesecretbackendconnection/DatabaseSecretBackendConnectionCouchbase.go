// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionCouchbase struct {
	// A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#hosts DatabaseSecretBackendConnection#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Specifies the password corresponding to the given username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the username for Vault to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Required if `tls` is `true`.
	//
	// Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#base64_pem DatabaseSecretBackendConnection#base64_pem}
	Base64Pem *string `field:"optional" json:"base64Pem" yaml:"base64Pem"`
	// Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#bucket_name DatabaseSecretBackendConnection#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Specifies whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#insecure_tls DatabaseSecretBackendConnection#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// Specifies whether to use TLS when connecting to Couchbase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#tls DatabaseSecretBackendConnection#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#username_template DatabaseSecretBackendConnection#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

