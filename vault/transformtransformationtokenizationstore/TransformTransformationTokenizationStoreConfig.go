// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformtransformationtokenizationstore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransformTransformationTokenizationStoreConfig struct {
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
	// A database connection string with template slots for username and password that Vault will use for locating and connecting to a database.
	//
	// Each database driver type has a different syntax for its connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#connection_string TransformTransformationTokenizationStore#connection_string}
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Specifies the database driver to use, and thus which SQL database type.
	//
	// Currently the supported options are postgres, mysql, and mssql.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#driver TransformTransformationTokenizationStore#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
	// The name of the store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#name TransformTransformationTokenizationStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Password value to use to connect to database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#password TransformTransformationTokenizationStore#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#path TransformTransformationTokenizationStore#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Specifies the type of store, currently only "sql" is supported,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#type TransformTransformationTokenizationStore#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Username value to use to connect to database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#username TransformTransformationTokenizationStore#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#id TransformTransformationTokenizationStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum amount of time a connection can be open before closing it.
	//
	// 0 means no limit. Default is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#max_connection_lifetime TransformTransformationTokenizationStore#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// The maximum number of idle connections to the database at any given time. Default is 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#max_idle_connections TransformTransformationTokenizationStore#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// The maximum number of connections to the database at any given time. Default is 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#max_open_connections TransformTransformationTokenizationStore#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#namespace TransformTransformationTokenizationStore#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The schema within the database to expect tokenization state tables. Default is "public".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#schema TransformTransformationTokenizationStore#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// The types of transformations this store can support, currently only "tokenization" is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/transform_transformation_tokenization_store#supported_transformations TransformTransformationTokenizationStore#supported_transformations}
	SupportedTransformations *[]*string `field:"optional" json:"supportedTransformations" yaml:"supportedTransformations"`
}

