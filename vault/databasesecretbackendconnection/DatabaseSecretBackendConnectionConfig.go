// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretBackendConnectionConfig struct {
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
	// Unique name of the Vault mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#backend DatabaseSecretBackendConnection#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#name DatabaseSecretBackendConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#allowed_roles DatabaseSecretBackendConnection#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// cassandra block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#cassandra DatabaseSecretBackendConnection#cassandra}
	Cassandra *DatabaseSecretBackendConnectionCassandra `field:"optional" json:"cassandra" yaml:"cassandra"`
	// couchbase block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#couchbase DatabaseSecretBackendConnection#couchbase}
	Couchbase *DatabaseSecretBackendConnectionCouchbase `field:"optional" json:"couchbase" yaml:"couchbase"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#data DatabaseSecretBackendConnection#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#disable_automated_rotation DatabaseSecretBackendConnection#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#elasticsearch DatabaseSecretBackendConnection#elasticsearch}
	Elasticsearch *DatabaseSecretBackendConnectionElasticsearch `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// hana block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#hana DatabaseSecretBackendConnection#hana}
	Hana *DatabaseSecretBackendConnectionHana `field:"optional" json:"hana" yaml:"hana"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#id DatabaseSecretBackendConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// influxdb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#influxdb DatabaseSecretBackendConnection#influxdb}
	Influxdb *DatabaseSecretBackendConnectionInfluxdb `field:"optional" json:"influxdb" yaml:"influxdb"`
	// mongodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mongodb DatabaseSecretBackendConnection#mongodb}
	Mongodb *DatabaseSecretBackendConnectionMongodb `field:"optional" json:"mongodb" yaml:"mongodb"`
	// mongodbatlas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mongodbatlas DatabaseSecretBackendConnection#mongodbatlas}
	Mongodbatlas *DatabaseSecretBackendConnectionMongodbatlas `field:"optional" json:"mongodbatlas" yaml:"mongodbatlas"`
	// mssql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mssql DatabaseSecretBackendConnection#mssql}
	Mssql *DatabaseSecretBackendConnectionMssql `field:"optional" json:"mssql" yaml:"mssql"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mysql DatabaseSecretBackendConnection#mysql}
	Mysql *DatabaseSecretBackendConnectionMysql `field:"optional" json:"mysql" yaml:"mysql"`
	// mysql_aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mysql_aurora DatabaseSecretBackendConnection#mysql_aurora}
	MysqlAurora *DatabaseSecretBackendConnectionMysqlAurora `field:"optional" json:"mysqlAurora" yaml:"mysqlAurora"`
	// mysql_legacy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mysql_legacy DatabaseSecretBackendConnection#mysql_legacy}
	MysqlLegacy *DatabaseSecretBackendConnectionMysqlLegacy `field:"optional" json:"mysqlLegacy" yaml:"mysqlLegacy"`
	// mysql_rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#mysql_rds DatabaseSecretBackendConnection#mysql_rds}
	MysqlRds *DatabaseSecretBackendConnectionMysqlRds `field:"optional" json:"mysqlRds" yaml:"mysqlRds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#namespace DatabaseSecretBackendConnection#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#oracle DatabaseSecretBackendConnection#oracle}
	Oracle *DatabaseSecretBackendConnectionOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// The name of the password policy to use when generating passwords for this database.
	//
	// If not specified, this will use a default policy defined as: 20 characters with at least 1 uppercase, 1 lowercase, 1 number, and 1 dash character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#password_policy DatabaseSecretBackendConnection#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Specifies the name of the plugin to use for this connection.
	//
	// Must be prefixed with the name of one of the supported database engine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#plugin_name DatabaseSecretBackendConnection#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// Specifies the semantic version of the plugin to use for this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#plugin_version DatabaseSecretBackendConnection#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#postgresql DatabaseSecretBackendConnection#postgresql}
	Postgresql *DatabaseSecretBackendConnectionPostgresql `field:"optional" json:"postgresql" yaml:"postgresql"`
	// redis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#redis DatabaseSecretBackendConnection#redis}
	Redis *DatabaseSecretBackendConnectionRedis `field:"optional" json:"redis" yaml:"redis"`
	// redis_elasticache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#redis_elasticache DatabaseSecretBackendConnection#redis_elasticache}
	RedisElasticache *DatabaseSecretBackendConnectionRedisElasticache `field:"optional" json:"redisElasticache" yaml:"redisElasticache"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#redshift DatabaseSecretBackendConnection#redshift}
	Redshift *DatabaseSecretBackendConnectionRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// A list of database statements to be executed to rotate the root user's credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#root_rotation_statements DatabaseSecretBackendConnection#root_rotation_statements}
	RootRotationStatements *[]*string `field:"optional" json:"rootRotationStatements" yaml:"rootRotationStatements"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#rotation_period DatabaseSecretBackendConnection#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#rotation_schedule DatabaseSecretBackendConnection#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#rotation_window DatabaseSecretBackendConnection#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Specifies if a given static account's password should be rotated on creation of the static roles associated with this database config.
	//
	// This is can be overridden at the role-level by the static role's skip_import_rotation field. The default is false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#skip_static_role_import_rotation DatabaseSecretBackendConnection#skip_static_role_import_rotation}
	SkipStaticRoleImportRotation interface{} `field:"optional" json:"skipStaticRoleImportRotation" yaml:"skipStaticRoleImportRotation"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#snowflake DatabaseSecretBackendConnection#snowflake}
	Snowflake *DatabaseSecretBackendConnectionSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/database_secret_backend_connection#verify_connection DatabaseSecretBackendConnection#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

