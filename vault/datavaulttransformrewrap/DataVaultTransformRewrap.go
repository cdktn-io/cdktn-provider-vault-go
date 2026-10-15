// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransformrewrap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/datavaulttransformrewrap/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap vault_transform_rewrap}.
type DataVaultTransformRewrap interface {
	cdktn.TerraformDataSource
	BatchInput() interface{}
	SetBatchInput(val interface{})
	BatchInputInput() interface{}
	BatchResults() cdktn.StringMapList
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DecodeTransformation() *string
	SetDecodeTransformation(val *string)
	DecodeTransformationInput() *string
	DecodeTweak() *string
	SetDecodeTweak(val *string)
	DecodeTweakInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncodedValue() *string
	SetEncodedValue(val *string)
	EncodedValueInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Transformation() *string
	SetTransformation(val *string)
	TransformationInput() *string
	Tweak() *string
	SetTweak(val *string)
	TweakInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetBatchInput()
	ResetDecodeTransformation()
	ResetDecodeTweak()
	ResetEncodedValue()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTransformation()
	ResetTweak()
	ResetValue()
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

// The jsii proxy struct for DataVaultTransformRewrap
type jsiiProxy_DataVaultTransformRewrap struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultTransformRewrap) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) BatchInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) BatchResults() cdktn.StringMapList {
	var returns cdktn.StringMapList
	_jsii_.Get(
		j,
		"batchResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) DecodeTransformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decodeTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) DecodeTransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decodeTransformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) DecodeTweak() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decodeTweak",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) DecodeTweakInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decodeTweakInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) EncodedValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) EncodedValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Transformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) TransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Tweak() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tweak",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) TweakInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tweakInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransformRewrap) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap vault_transform_rewrap} Data Source.
func NewDataVaultTransformRewrap(scope constructs.Construct, id *string, config *DataVaultTransformRewrapConfig) DataVaultTransformRewrap {
	_init_.Initialize()

	if err := validateNewDataVaultTransformRewrapParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultTransformRewrap{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/data-sources/transform_rewrap vault_transform_rewrap} Data Source.
func NewDataVaultTransformRewrap_Override(d DataVaultTransformRewrap, scope constructs.Construct, id *string, config *DataVaultTransformRewrapConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetBatchInput(val interface{}) {
	if err := j.validateSetBatchInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchInput",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetDecodeTransformation(val *string) {
	if err := j.validateSetDecodeTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decodeTransformation",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetDecodeTweak(val *string) {
	if err := j.validateSetDecodeTweakParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decodeTweak",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetEncodedValue(val *string) {
	if err := j.validateSetEncodedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodedValue",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetTransformation(val *string) {
	if err := j.validateSetTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformation",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetTweak(val *string) {
	if err := j.validateSetTweakParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tweak",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransformRewrap)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultTransformRewrap resource upon running "cdktn plan <stack-name>".
func DataVaultTransformRewrap_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultTransformRewrap_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
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
func DataVaultTransformRewrap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransformRewrap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransformRewrap_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransformRewrap_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransformRewrap_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransformRewrap_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultTransformRewrap_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultTransformRewrap.DataVaultTransformRewrap",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultTransformRewrap) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultTransformRewrap) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransformRewrap) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetBatchInput() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetDecodeTransformation() {
	_jsii_.InvokeVoid(
		d,
		"resetDecodeTransformation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetDecodeTweak() {
	_jsii_.InvokeVoid(
		d,
		"resetDecodeTweak",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetEncodedValue() {
	_jsii_.InvokeVoid(
		d,
		"resetEncodedValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetTransformation() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetTweak() {
	_jsii_.InvokeVoid(
		d,
		"resetTweak",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) ResetValue() {
	_jsii_.InvokeVoid(
		d,
		"resetValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransformRewrap) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransformRewrap) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

