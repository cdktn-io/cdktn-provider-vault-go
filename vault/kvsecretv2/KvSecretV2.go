// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/kvsecretv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_v2 vault_kv_secret_v2}.
type KvSecretV2 interface {
	cdktn.TerraformResource
	Cas() *float64
	SetCas(val *float64)
	CasInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomMetadata() KvSecretV2CustomMetadataOutputReference
	CustomMetadataInput() *KvSecretV2CustomMetadata
	Data() cdktn.StringMap
	DataJson() *string
	SetDataJson(val *string)
	DataJsonInput() *string
	DataJsonWo() *string
	SetDataJsonWo(val *string)
	DataJsonWoInput() *string
	DataJsonWoVersion() *float64
	SetDataJsonWoVersion(val *float64)
	DataJsonWoVersionInput() *float64
	DeleteAllVersions() interface{}
	SetDeleteAllVersions(val interface{})
	DeleteAllVersionsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableRead() interface{}
	SetDisableRead(val interface{})
	DisableReadInput() interface{}
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
	Metadata() cdktn.StringMap
	Mount() *string
	SetMount(val *string)
	MountInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	Path() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCustomMetadata(value *KvSecretV2CustomMetadata)
	ResetCas()
	ResetCustomMetadata()
	ResetDataJson()
	ResetDataJsonWo()
	ResetDataJsonWoVersion()
	ResetDeleteAllVersions()
	ResetDisableRead()
	ResetId()
	ResetNamespace()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for KvSecretV2
type jsiiProxy_KvSecretV2 struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KvSecretV2) Cas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) CasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"casInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) CustomMetadata() KvSecretV2CustomMetadataOutputReference {
	var returns KvSecretV2CustomMetadataOutputReference
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) CustomMetadataInput() *KvSecretV2CustomMetadata {
	var returns *KvSecretV2CustomMetadata
	_jsii_.Get(
		j,
		"customMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Data() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJsonWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataJsonWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJsonWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataJsonWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJsonWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataJsonWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DataJsonWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataJsonWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DeleteAllVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DeleteAllVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAllVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DisableRead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) DisableReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_v2 vault_kv_secret_v2} Resource.
func NewKvSecretV2(scope constructs.Construct, id *string, config *KvSecretV2Config) KvSecretV2 {
	_init_.Initialize()

	if err := validateNewKvSecretV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KvSecretV2{}

	_jsii_.Create(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kv_secret_v2 vault_kv_secret_v2} Resource.
func NewKvSecretV2_Override(k KvSecretV2, scope constructs.Construct, id *string, config *KvSecretV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KvSecretV2)SetCas(val *float64) {
	if err := j.validateSetCasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cas",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDataJson(val *string) {
	if err := j.validateSetDataJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataJson",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDataJsonWo(val *string) {
	if err := j.validateSetDataJsonWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataJsonWo",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDataJsonWoVersion(val *float64) {
	if err := j.validateSetDataJsonWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataJsonWoVersion",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDeleteAllVersions(val interface{}) {
	if err := j.validateSetDeleteAllVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAllVersions",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetDisableRead(val interface{}) {
	if err := j.validateSetDisableReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRead",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a KvSecretV2 resource upon running "cdktn plan <stack-name>".
func KvSecretV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKvSecretV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
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
func KvSecretV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKvSecretV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KvSecretV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKvSecretV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KvSecretV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKvSecretV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KvSecretV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KvSecretV2) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KvSecretV2) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KvSecretV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KvSecretV2) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KvSecretV2) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KvSecretV2) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KvSecretV2) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KvSecretV2) PutCustomMetadata(value *KvSecretV2CustomMetadata) {
	if err := k.validatePutCustomMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCustomMetadata",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KvSecretV2) ResetCas() {
	_jsii_.InvokeVoid(
		k,
		"resetCas",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetCustomMetadata() {
	_jsii_.InvokeVoid(
		k,
		"resetCustomMetadata",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetDataJson() {
	_jsii_.InvokeVoid(
		k,
		"resetDataJson",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetDataJsonWo() {
	_jsii_.InvokeVoid(
		k,
		"resetDataJsonWo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetDataJsonWoVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetDataJsonWoVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetDeleteAllVersions() {
	_jsii_.InvokeVoid(
		k,
		"resetDeleteAllVersions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetDisableRead() {
	_jsii_.InvokeVoid(
		k,
		"resetDisableRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetOptions() {
	_jsii_.InvokeVoid(
		k,
		"resetOptions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

