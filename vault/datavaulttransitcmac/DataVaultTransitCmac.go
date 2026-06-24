// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitcmac

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaulttransitcmac/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac vault_transit_cmac}.
type DataVaultTransitCmac interface {
	cdktn.TerraformDataSource
	BatchInput() interface{}
	SetBatchInput(val interface{})
	BatchInputInput() interface{}
	BatchResults() interface{}
	SetBatchResults(val interface{})
	BatchResultsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Cmac() *string
	SetCmac(val *string)
	CmacInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Input() *string
	SetInput(val *string)
	InputInput() *string
	KeyVersion() *float64
	SetKeyVersion(val *float64)
	KeyVersionInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MacLength() *float64
	SetMacLength(val *float64)
	MacLengthInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UrlMacLength() *float64
	SetUrlMacLength(val *float64)
	UrlMacLengthInput() *float64
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
	ResetBatchInput()
	ResetBatchResults()
	ResetCmac()
	ResetId()
	ResetInput()
	ResetKeyVersion()
	ResetMacLength()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUrlMacLength()
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

// The jsii proxy struct for DataVaultTransitCmac
type jsiiProxy_DataVaultTransitCmac struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultTransitCmac) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) BatchInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) BatchResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) BatchResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Cmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) CmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) KeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) KeyVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) MacLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"macLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) MacLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"macLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) UrlMacLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"urlMacLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitCmac) UrlMacLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"urlMacLengthInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac vault_transit_cmac} Data Source.
func NewDataVaultTransitCmac(scope constructs.Construct, id *string, config *DataVaultTransitCmacConfig) DataVaultTransitCmac {
	_init_.Initialize()

	if err := validateNewDataVaultTransitCmacParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultTransitCmac{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_cmac vault_transit_cmac} Data Source.
func NewDataVaultTransitCmac_Override(d DataVaultTransitCmac, scope constructs.Construct, id *string, config *DataVaultTransitCmacConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetBatchInput(val interface{}) {
	if err := j.validateSetBatchInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchInput",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetBatchResults(val interface{}) {
	if err := j.validateSetBatchResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchResults",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetCmac(val *string) {
	if err := j.validateSetCmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmac",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetKeyVersion(val *float64) {
	if err := j.validateSetKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVersion",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetMacLength(val *float64) {
	if err := j.validateSetMacLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macLength",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitCmac)SetUrlMacLength(val *float64) {
	if err := j.validateSetUrlMacLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlMacLength",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultTransitCmac resource upon running "cdktn plan <stack-name>".
func DataVaultTransitCmac_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultTransitCmac_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
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
func DataVaultTransitCmac_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitCmac_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitCmac_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitCmac_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitCmac_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitCmac_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultTransitCmac_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultTransitCmac.DataVaultTransitCmac",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultTransitCmac) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultTransitCmac) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitCmac) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetBatchInput() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetBatchResults() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchResults",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetCmac() {
	_jsii_.InvokeVoid(
		d,
		"resetCmac",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetInput() {
	_jsii_.InvokeVoid(
		d,
		"resetInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetKeyVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetMacLength() {
	_jsii_.InvokeVoid(
		d,
		"resetMacLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) ResetUrlMacLength() {
	_jsii_.InvokeVoid(
		d,
		"resetUrlMacLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitCmac) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitCmac) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

