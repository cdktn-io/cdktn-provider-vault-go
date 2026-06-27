// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretsMountConfig struct {
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
	// Where the secret backend will be mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#path DatabaseSecretsMount#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#allowed_managed_keys DatabaseSecretsMount#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#allowed_response_headers DatabaseSecretsMount#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#audit_non_hmac_request_keys DatabaseSecretsMount#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#audit_non_hmac_response_keys DatabaseSecretsMount#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// cassandra block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#cassandra DatabaseSecretsMount#cassandra}
	Cassandra interface{} `field:"optional" json:"cassandra" yaml:"cassandra"`
	// couchbase block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#couchbase DatabaseSecretsMount#couchbase}
	Couchbase interface{} `field:"optional" json:"couchbase" yaml:"couchbase"`
	// Default lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#default_lease_ttl_seconds DatabaseSecretsMount#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#delegated_auth_accessors DatabaseSecretsMount#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#description DatabaseSecretsMount#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#elasticsearch DatabaseSecretsMount#elasticsearch}
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#external_entropy_access DatabaseSecretsMount#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#force_no_cache DatabaseSecretsMount#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// hana block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#hana DatabaseSecretsMount#hana}
	Hana interface{} `field:"optional" json:"hana" yaml:"hana"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#id DatabaseSecretsMount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key to use for signing plugin workload identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#identity_token_key DatabaseSecretsMount#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// influxdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#influxdb DatabaseSecretsMount#influxdb}
	Influxdb interface{} `field:"optional" json:"influxdb" yaml:"influxdb"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#listing_visibility DatabaseSecretsMount#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#local DatabaseSecretsMount#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#max_lease_ttl_seconds DatabaseSecretsMount#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// mongodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mongodb DatabaseSecretsMount#mongodb}
	Mongodb interface{} `field:"optional" json:"mongodb" yaml:"mongodb"`
	// mongodbatlas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mongodbatlas DatabaseSecretsMount#mongodbatlas}
	Mongodbatlas interface{} `field:"optional" json:"mongodbatlas" yaml:"mongodbatlas"`
	// mssql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mssql DatabaseSecretsMount#mssql}
	Mssql interface{} `field:"optional" json:"mssql" yaml:"mssql"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mysql DatabaseSecretsMount#mysql}
	Mysql interface{} `field:"optional" json:"mysql" yaml:"mysql"`
	// mysql_aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mysql_aurora DatabaseSecretsMount#mysql_aurora}
	MysqlAurora interface{} `field:"optional" json:"mysqlAurora" yaml:"mysqlAurora"`
	// mysql_legacy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mysql_legacy DatabaseSecretsMount#mysql_legacy}
	MysqlLegacy interface{} `field:"optional" json:"mysqlLegacy" yaml:"mysqlLegacy"`
	// mysql_rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#mysql_rds DatabaseSecretsMount#mysql_rds}
	MysqlRds interface{} `field:"optional" json:"mysqlRds" yaml:"mysqlRds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#namespace DatabaseSecretsMount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#options DatabaseSecretsMount#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#oracle DatabaseSecretsMount#oracle}
	Oracle interface{} `field:"optional" json:"oracle" yaml:"oracle"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#passthrough_request_headers DatabaseSecretsMount#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#plugin_version DatabaseSecretsMount#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#postgresql DatabaseSecretsMount#postgresql}
	Postgresql interface{} `field:"optional" json:"postgresql" yaml:"postgresql"`
	// redis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#redis DatabaseSecretsMount#redis}
	Redis interface{} `field:"optional" json:"redis" yaml:"redis"`
	// redis_elasticache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#redis_elasticache DatabaseSecretsMount#redis_elasticache}
	RedisElasticache interface{} `field:"optional" json:"redisElasticache" yaml:"redisElasticache"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#redshift DatabaseSecretsMount#redshift}
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#seal_wrap DatabaseSecretsMount#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#snowflake DatabaseSecretsMount#snowflake}
	Snowflake interface{} `field:"optional" json:"snowflake" yaml:"snowflake"`
}

