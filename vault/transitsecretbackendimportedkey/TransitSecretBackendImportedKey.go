// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transitsecretbackendimportedkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/transitsecretbackendimportedkey/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key vault_transit_secret_backend_imported_key}.
type TransitSecretBackendImportedKey interface {
	cdktn.TerraformResource
	AllowPlaintextBackup() interface{}
	SetAllowPlaintextBackup(val interface{})
	AllowPlaintextBackupInput() interface{}
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
	Context() *string
	SetContext(val *string)
	ContextInput() *string
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
	Derived() interface{}
	SetDerived(val interface{})
	DerivedInput() interface{}
	Exportable() interface{}
	SetExportable(val interface{})
	ExportableInput() interface{}
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
	ImportedKey() cdktn.IResolvable
	ImportedKeyAllowRotation() cdktn.IResolvable
	ImportVersionType() *string
	SetImportVersionType(val *string)
	ImportVersionTypeInput() *string
	Keys() cdktn.StringMap
	LatestVersion() *float64
	LatestVersionType() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MinAvailableVersion() *float64
	MinDecryptionVersion() *float64
	SetMinDecryptionVersion(val *float64)
	MinDecryptionVersionInput() *float64
	MinEncryptionVersion() *float64
	SetMinEncryptionVersion(val *float64)
	MinEncryptionVersionInput() *float64
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
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	SupportsDecryption() cdktn.IResolvable
	SupportsDerivation() cdktn.IResolvable
	SupportsEncryption() cdktn.IResolvable
	SupportsSigning() cdktn.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
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
	ResetAllowPlaintextBackup()
	ResetAllowRotation()
	ResetAutoRotatePeriod()
	ResetCiphertextWo()
	ResetCiphertextWoVersion()
	ResetContext()
	ResetDeletionAllowed()
	ResetDerived()
	ResetExportable()
	ResetHashFunction()
	ResetImportVersionType()
	ResetMinDecryptionVersion()
	ResetMinEncryptionVersion()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicKey()
	ResetVersion()
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

// The jsii proxy struct for TransitSecretBackendImportedKey
type jsiiProxy_TransitSecretBackendImportedKey struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AllowPlaintextBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPlaintextBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AllowPlaintextBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPlaintextBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AllowRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AllowRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AutoRotatePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) AutoRotatePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) CiphertextWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciphertextWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) CiphertextWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ciphertextWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) CiphertextWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciphertextWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) CiphertextWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ciphertextWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) DeletionAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) DeletionAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Derived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"derived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) DerivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"derivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Exportable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ExportableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) HashFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) HashFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ImportedKey() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"importedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ImportedKeyAllowRotation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"importedKeyAllowRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ImportVersionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importVersionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) ImportVersionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importVersionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Keys() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) LatestVersionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestVersionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MinAvailableVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAvailableVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MinDecryptionVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MinDecryptionVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MinEncryptionVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEncryptionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MinEncryptionVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEncryptionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) SupportsDecryption() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsDecryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) SupportsDerivation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsDerivation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) SupportsEncryption() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) SupportsSigning() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendImportedKey) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key vault_transit_secret_backend_imported_key} Resource.
func NewTransitSecretBackendImportedKey(scope constructs.Construct, id *string, config *TransitSecretBackendImportedKeyConfig) TransitSecretBackendImportedKey {
	_init_.Initialize()

	if err := validateNewTransitSecretBackendImportedKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitSecretBackendImportedKey{}

	_jsii_.Create(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key vault_transit_secret_backend_imported_key} Resource.
func NewTransitSecretBackendImportedKey_Override(t TransitSecretBackendImportedKey, scope constructs.Construct, id *string, config *TransitSecretBackendImportedKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetAllowPlaintextBackup(val interface{}) {
	if err := j.validateSetAllowPlaintextBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPlaintextBackup",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetAllowRotation(val interface{}) {
	if err := j.validateSetAllowRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowRotation",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetAutoRotatePeriod(val *float64) {
	if err := j.validateSetAutoRotatePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRotatePeriod",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetCiphertextWo(val *string) {
	if err := j.validateSetCiphertextWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphertextWo",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetCiphertextWoVersion(val *float64) {
	if err := j.validateSetCiphertextWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphertextWoVersion",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetDeletionAllowed(val interface{}) {
	if err := j.validateSetDeletionAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionAllowed",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetDerived(val interface{}) {
	if err := j.validateSetDerivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"derived",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetExportable(val interface{}) {
	if err := j.validateSetExportableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportable",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetHashFunction(val *string) {
	if err := j.validateSetHashFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashFunction",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetImportVersionType(val *string) {
	if err := j.validateSetImportVersionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importVersionType",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetMinDecryptionVersion(val *float64) {
	if err := j.validateSetMinDecryptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDecryptionVersion",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetMinEncryptionVersion(val *float64) {
	if err := j.validateSetMinEncryptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minEncryptionVersion",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendImportedKey)SetVersion(val *float64) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a TransitSecretBackendImportedKey resource upon running "cdktn plan <stack-name>".
func TransitSecretBackendImportedKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTransitSecretBackendImportedKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
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
func TransitSecretBackendImportedKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendImportedKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransitSecretBackendImportedKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendImportedKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransitSecretBackendImportedKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendImportedKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransitSecretBackendImportedKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.transitSecretBackendImportedKey.TransitSecretBackendImportedKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TransitSecretBackendImportedKey) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetAllowPlaintextBackup() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowPlaintextBackup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetAllowRotation() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowRotation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetAutoRotatePeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRotatePeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetCiphertextWo() {
	_jsii_.InvokeVoid(
		t,
		"resetCiphertextWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetCiphertextWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetCiphertextWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetContext() {
	_jsii_.InvokeVoid(
		t,
		"resetContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetDeletionAllowed() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionAllowed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetDerived() {
	_jsii_.InvokeVoid(
		t,
		"resetDerived",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetExportable() {
	_jsii_.InvokeVoid(
		t,
		"resetExportable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetHashFunction() {
	_jsii_.InvokeVoid(
		t,
		"resetHashFunction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetImportVersionType() {
	_jsii_.InvokeVoid(
		t,
		"resetImportVersionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetMinDecryptionVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinDecryptionVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetMinEncryptionVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinEncryptionVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetPublicKey() {
	_jsii_.InvokeVoid(
		t,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendImportedKey) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

