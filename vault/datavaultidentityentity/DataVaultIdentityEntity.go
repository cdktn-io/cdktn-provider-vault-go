// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultidentityentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/datavaultidentityentity/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/identity_entity vault_identity_entity}.
type DataVaultIdentityEntity interface {
	cdktn.TerraformDataSource
	Aliases() DataVaultIdentityEntityAliasesList
	AliasId() *string
	SetAliasId(val *string)
	AliasIdInput() *string
	AliasMountAccessor() *string
	SetAliasMountAccessor(val *string)
	AliasMountAccessorInput() *string
	AliasName() *string
	SetAliasName(val *string)
	AliasNameInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	DataJson() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectGroupIds() *[]*string
	Disabled() cdktn.IResolvable
	EntityId() *string
	SetEntityId(val *string)
	EntityIdInput() *string
	EntityName() *string
	SetEntityName(val *string)
	EntityNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupIds() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InheritedGroupIds() *[]*string
	LastUpdateTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MergedEntityIds() *[]*string
	Metadata() cdktn.StringMap
	Namespace() *string
	SetNamespace(val *string)
	NamespaceId() *string
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Policies() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAliasId()
	ResetAliasMountAccessor()
	ResetAliasName()
	ResetEntityId()
	ResetEntityName()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataVaultIdentityEntity
type jsiiProxy_DataVaultIdentityEntity struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultIdentityEntity) Aliases() DataVaultIdentityEntityAliasesList {
	var returns DataVaultIdentityEntityAliasesList
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasMountAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasMountAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasMountAccessorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasMountAccessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) AliasNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) DataJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) DirectGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"directGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Disabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) EntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) EntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) EntityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) GroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) InheritedGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inheritedGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) LastUpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) MergedEntityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mergedEntityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultIdentityEntity) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/identity_entity vault_identity_entity} Data Source.
func NewDataVaultIdentityEntity(scope constructs.Construct, id *string, config *DataVaultIdentityEntityConfig) DataVaultIdentityEntity {
	_init_.Initialize()

	if err := validateNewDataVaultIdentityEntityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultIdentityEntity{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/identity_entity vault_identity_entity} Data Source.
func NewDataVaultIdentityEntity_Override(d DataVaultIdentityEntity, scope constructs.Construct, id *string, config *DataVaultIdentityEntityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetAliasId(val *string) {
	if err := j.validateSetAliasIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasId",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetAliasMountAccessor(val *string) {
	if err := j.validateSetAliasMountAccessorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMountAccessor",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetAliasName(val *string) {
	if err := j.validateSetAliasNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasName",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetEntityId(val *string) {
	if err := j.validateSetEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityId",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetEntityName(val *string) {
	if err := j.validateSetEntityNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityName",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultIdentityEntity)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultIdentityEntity resource upon running "cdktn plan <stack-name>".
func DataVaultIdentityEntity_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultIdentityEntity_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataVaultIdentityEntity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultIdentityEntity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultIdentityEntity_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultIdentityEntity_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultIdentityEntity_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultIdentityEntity_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultIdentityEntity_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultIdentityEntity.DataVaultIdentityEntity",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetAliasId() {
	_jsii_.InvokeVoid(
		d,
		"resetAliasId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetAliasMountAccessor() {
	_jsii_.InvokeVoid(
		d,
		"resetAliasMountAccessor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetAliasName() {
	_jsii_.InvokeVoid(
		d,
		"resetAliasName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetEntityId() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetEntityName() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultIdentityEntity) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultIdentityEntity) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

