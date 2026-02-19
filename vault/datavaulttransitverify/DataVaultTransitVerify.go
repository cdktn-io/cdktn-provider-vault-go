// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitverify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/datavaulttransitverify/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/data-sources/transit_verify vault_transit_verify}.
type DataVaultTransitVerify interface {
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
	Hmac() *string
	SetHmac(val *string)
	HmacInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MacLength() *float64
	SetMacLength(val *float64)
	MacLengthInput() *float64
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
	Valid() interface{}
	SetValid(val interface{})
	ValidInput() interface{}
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
	ResetContext()
	ResetHashAlgorithm()
	ResetHmac()
	ResetId()
	ResetInput()
	ResetMacLength()
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
	ResetValid()
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
}

// The jsii proxy struct for DataVaultTransitVerify
type jsiiProxy_DataVaultTransitVerify struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultTransitVerify) BatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) BatchInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) BatchResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) BatchResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Cmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) CmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) HashAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) HashAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Hmac() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hmac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) HmacInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hmacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) MacLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"macLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) MacLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"macLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) MarshalingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marshalingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) MarshalingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marshalingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Prehashed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prehashed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) PrehashedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prehashedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Reference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) ReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SaltLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saltLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SaltLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saltLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Signature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SignatureContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SignatureContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) SignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) Valid() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultTransitVerify) ValidInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/data-sources/transit_verify vault_transit_verify} Data Source.
func NewDataVaultTransitVerify(scope constructs.Construct, id *string, config *DataVaultTransitVerifyConfig) DataVaultTransitVerify {
	_init_.Initialize()

	if err := validateNewDataVaultTransitVerifyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultTransitVerify{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/data-sources/transit_verify vault_transit_verify} Data Source.
func NewDataVaultTransitVerify_Override(d DataVaultTransitVerify, scope constructs.Construct, id *string, config *DataVaultTransitVerifyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetBatchInput(val interface{}) {
	if err := j.validateSetBatchInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchInput",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetBatchResults(val interface{}) {
	if err := j.validateSetBatchResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchResults",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetCmac(val *string) {
	if err := j.validateSetCmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmac",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetHashAlgorithm(val *string) {
	if err := j.validateSetHashAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetHmac(val *string) {
	if err := j.validateSetHmacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hmac",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetMacLength(val *float64) {
	if err := j.validateSetMacLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macLength",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetMarshalingAlgorithm(val *string) {
	if err := j.validateSetMarshalingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marshalingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetPrehashed(val interface{}) {
	if err := j.validateSetPrehashedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prehashed",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetReference(val *string) {
	if err := j.validateSetReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reference",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetSaltLength(val *string) {
	if err := j.validateSetSaltLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saltLength",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetSignature(val *string) {
	if err := j.validateSetSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signature",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetSignatureAlgorithm(val *string) {
	if err := j.validateSetSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetSignatureContext(val *string) {
	if err := j.validateSetSignatureContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureContext",
		val,
	)
}

func (j *jsiiProxy_DataVaultTransitVerify)SetValid(val interface{}) {
	if err := j.validateSetValidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valid",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultTransitVerify resource upon running "cdktn plan <stack-name>".
func DataVaultTransitVerify_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultTransitVerify_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
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
func DataVaultTransitVerify_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitVerify_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitVerify_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitVerify_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultTransitVerify_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultTransitVerify_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultTransitVerify_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultTransitVerify.DataVaultTransitVerify",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultTransitVerify) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultTransitVerify) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultTransitVerify) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetBatchInput() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetBatchResults() {
	_jsii_.InvokeVoid(
		d,
		"resetBatchResults",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetCmac() {
	_jsii_.InvokeVoid(
		d,
		"resetCmac",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetContext() {
	_jsii_.InvokeVoid(
		d,
		"resetContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetHashAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetHashAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetHmac() {
	_jsii_.InvokeVoid(
		d,
		"resetHmac",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetInput() {
	_jsii_.InvokeVoid(
		d,
		"resetInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetMacLength() {
	_jsii_.InvokeVoid(
		d,
		"resetMacLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetMarshalingAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetMarshalingAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetPrehashed() {
	_jsii_.InvokeVoid(
		d,
		"resetPrehashed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetReference() {
	_jsii_.InvokeVoid(
		d,
		"resetReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetSaltLength() {
	_jsii_.InvokeVoid(
		d,
		"resetSaltLength",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetSignature() {
	_jsii_.InvokeVoid(
		d,
		"resetSignature",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		d,
		"resetSignatureAlgorithm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetSignatureContext() {
	_jsii_.InvokeVoid(
		d,
		"resetSignatureContext",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) ResetValid() {
	_jsii_.InvokeVoid(
		d,
		"resetValid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultTransitVerify) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultTransitVerify) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

