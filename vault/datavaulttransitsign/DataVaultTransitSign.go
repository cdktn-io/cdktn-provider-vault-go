// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitsign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaulttransitsign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_sign vault_transit_sign}.
type DataVaultTransitSign interface {
	cdktn.TerraformDataSource
	BatchInput() interface{}
	SetBatchInput(val interface{})
	BatchInputInput() interface{}
	BatchResults() interface{}
	SetBatchResults(val interface{})
	BatchResultsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
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
	HashAlgorithm() *string
	SetHashAlgorithm(val *string)
	HashAlgorithmInput() *string
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
	MarshalingAlgorithm() *string
	SetMarshalingAlgorithm(val *string)
	MarshalingAlgorithmInput() *string
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
	Prehashed() interface{}
	SetPrehashed(val interface{})
	PrehashedInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Reference() *string
	SetReference(val *string)
	ReferenceInput() *string
	SaltLength() *string
	SetSaltLength(val *string)
	SaltLengthInput() *string
	Signature() *string
	SetSignature(val *string)
	SignatureAlgorithm() *string
	SetSignatureAlgorithm(val *string)
	SignatureAlgorithmInput() *string
	SignatureContext() *string
	SetSignatureContext(val *string)
	SignatureContextInput() *string
	SignatureInput() *string
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
	ResetBatchInput()
	ResetBatchResults()
	ResetContext()
	ResetHashAlgorithm()
	ResetId()
	ResetInput()
	ResetKeyVersion()
	ResetMarshalingAlgorithm()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrehashed()
	ResetReference()
	ResetSaltLength()
	ResetSignature()
	ResetSignatureAlgorithm()
	ResetSignatureContext()
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

// The jsii proxy struct for DataVaultTransitSign
type jsiiProxy_DataVaultTransitSign struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultTransitSign) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) BatchInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) BatchResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) BatchResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) HashAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) KeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) KeyVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) MarshalingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marshalingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) MarshalingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marshalingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Prehashed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prehashed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) PrehashedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prehashedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Reference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) ReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SaltLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saltLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SaltLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saltLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) Signature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SignatureContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SignatureContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) SignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitSign) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_sign vault_transit_sign} Data Source.
func NewDataVaultTransitSign(scope constructs.Construct, id *string, config *DataVaultTransitSignConfig) DataVaultTransitSign {
	_init_.Initialize()

	if err := validateNewDataVaultTransitSignParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultTransitSign{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/transit_sign vault_transit_sign} Data Source.
func NewDataVaultTransitSign_Override(d DataVaultTransitSign, scope constructs.Construct, id *string, config *DataVaultTransitSignConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetBatchInput(val interface{}) {
	if err := j.validateSetBatchInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchInput",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetBatchResults(val interface{}) {
	if err := j.validateSetBatchResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchResults",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetHashAlgorithm(val *string) {
	if err := j.validateSetHashAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetKeyVersion(val *float64) {
	if err := j.validateSetKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVersion",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetMarshalingAlgorithm(val *string) {
	if err := j.validateSetMarshalingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marshalingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetPrehashed(val interface{}) {
	if err := j.validateSetPrehashedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prehashed",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetReference(val *string) {
	if err := j.validateSetReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reference",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetSaltLength(val *string) {
	if err := j.validateSetSaltLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saltLength",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetSignature(val *string) {
	if err := j.validateSetSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signature",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetSignatureAlgorithm(val *string) {
	if err := j.validateSetSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitSign)SetSignatureContext(val *string) {
	if err := j.validateSetSignatureContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureContext",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultTransitSign resource upon running "cdktn plan <stack-name>".
func DataVaultTransitSign_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultTransitSign_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
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
func DataVaultTransitSign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitSign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitSign_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitSign_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitSign_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitSign_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultTransitSign_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultTransitSign.DataVaultTransitSign",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultTransitSign) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultTransitSign) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitSign) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultTransitSign) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultTransitSign) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultTransitSign) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultTransitSign) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultTransitSign) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultTransitSign) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultTransitSign) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitSign) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetBatchInput() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetBatchResults() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchResults",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetContext() {
	_jsii_.InvokeVoid(
		d,
		"resetContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetHashAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetHashAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetInput() {
	_jsii_.InvokeVoid(
		d,
		"resetInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetKeyVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetMarshalingAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetMarshalingAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetPrehashed() {
	_jsii_.InvokeVoid(
		d,
		"resetPrehashed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetReference() {
	_jsii_.InvokeVoid(
		d,
		"resetReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetSaltLength() {
	_jsii_.InvokeVoid(
		d,
		"resetSaltLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetSignature() {
	_jsii_.InvokeVoid(
		d,
		"resetSignature",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetSignatureAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) ResetSignatureContext() {
	_jsii_.InvokeVoid(
		d,
		"resetSignatureContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitSign) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitSign) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

