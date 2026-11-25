// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformtransformationtokenizationimportedkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/transformtransformationtokenizationimportedkey/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_tokenization_imported_key vault_transform_transformation_tokenization_imported_key}.
type TransformTransformationTokenizationImportedKey interface {
	cdktn.TerraformResource
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
	AllowRotation() interface{}
	SetAllowRotation(val interface{})
	AllowRotationInput() interface{}
	AutoRotatePeriod() *float64
	SetAutoRotatePeriod(val *float64)
	AutoRotatePeriodInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	CiphertextWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetCiphertextWo(val *string)
	CiphertextWoInput() *string
	CiphertextWoVersion() *float64
	SetCiphertextWoVersion(val *float64)
	CiphertextWoVersionInput() *float64
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
	DeletionAllowed() interface{}
	SetDeletionAllowed(val interface{})
	DeletionAllowedInput() interface{}
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
	HashFunction() *string
	SetHashFunction(val *string)
	HashFunctionInput() *string
	LatestVersion() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MappingMode() *string
	SetMappingMode(val *string)
	MappingModeInput() *string
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	MinAvailableVersion() *float64
	MinDecryptionVersion() *float64
	SetMinDecryptionVersion(val *float64)
	MinDecryptionVersionInput() *float64
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
	Stores() *[]*string
	SetStores(val *[]*string)
	StoresInput() *[]*string
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	ResetAllowedRoles()
	ResetAllowRotation()
	ResetAutoRotatePeriod()
	ResetCiphertextWo()
	ResetCiphertextWoVersion()
	ResetDeletionAllowed()
	ResetHashFunction()
	ResetMappingMode()
	ResetMaxTtl()
	ResetMinDecryptionVersion()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStores()
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

// The jsii proxy struct for TransformTransformationTokenizationImportedKey
type jsiiProxy_TransformTransformationTokenizationImportedKey struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AllowRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AllowRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AutoRotatePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) AutoRotatePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) CiphertextWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciphertextWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) CiphertextWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciphertextWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) CiphertextWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciphertextWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) CiphertextWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciphertextWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) DeletionAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) DeletionAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) HashFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) HashFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MappingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MappingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MinAvailableVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAvailableVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MinDecryptionVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MinDecryptionVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) Stores() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) StoresInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_tokenization_imported_key vault_transform_transformation_tokenization_imported_key} Resource.
func NewTransformTransformationTokenizationImportedKey(scope constructs.Construct, id *string, config *TransformTransformationTokenizationImportedKeyConfig) TransformTransformationTokenizationImportedKey {
	_init_.Initialize()

	if err := validateNewTransformTransformationTokenizationImportedKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransformTransformationTokenizationImportedKey{}

	_jsii_.Create(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_tokenization_imported_key vault_transform_transformation_tokenization_imported_key} Resource.
func NewTransformTransformationTokenizationImportedKey_Override(t TransformTransformationTokenizationImportedKey, scope constructs.Construct, id *string, config *TransformTransformationTokenizationImportedKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetAllowRotation(val interface{}) {
	if err := j.validateSetAllowRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowRotation",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetAutoRotatePeriod(val *float64) {
	if err := j.validateSetAutoRotatePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRotatePeriod",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetCiphertextWo(val *string) {
	if err := j.validateSetCiphertextWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphertextWo",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetCiphertextWoVersion(val *float64) {
	if err := j.validateSetCiphertextWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphertextWoVersion",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetDeletionAllowed(val interface{}) {
	if err := j.validateSetDeletionAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionAllowed",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetHashFunction(val *string) {
	if err := j.validateSetHashFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashFunction",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetMappingMode(val *string) {
	if err := j.validateSetMappingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingMode",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetMinDecryptionVersion(val *float64) {
	if err := j.validateSetMinDecryptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDecryptionVersion",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TransformTransformationTokenizationImportedKey)SetStores(val *[]*string) {
	if err := j.validateSetStoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stores",
		val,
	)
}

// Generates CDKTN code for importing a TransformTransformationTokenizationImportedKey resource upon running "cdktn plan <stack-name>".
func TransformTransformationTokenizationImportedKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTransformTransformationTokenizationImportedKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
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
func TransformTransformationTokenizationImportedKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransformTransformationTokenizationImportedKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransformTransformationTokenizationImportedKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransformTransformationTokenizationImportedKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransformTransformationTokenizationImportedKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransformTransformationTokenizationImportedKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransformTransformationTokenizationImportedKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetAllowRotation() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowRotation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetAutoRotatePeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRotatePeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetCiphertextWo() {
	_jsii_.InvokeVoid(
		t,
		"resetCiphertextWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetCiphertextWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetCiphertextWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetDeletionAllowed() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionAllowed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetHashFunction() {
	_jsii_.InvokeVoid(
		t,
		"resetHashFunction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetMappingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetMappingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetMinDecryptionVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinDecryptionVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ResetStores() {
	_jsii_.InvokeVoid(
		t,
		"resetStores",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransformTransformationTokenizationImportedKey) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

