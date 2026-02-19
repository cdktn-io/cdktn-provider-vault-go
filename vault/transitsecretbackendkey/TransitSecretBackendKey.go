// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transitsecretbackendkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/transitsecretbackendkey/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/transit_secret_backend_key vault_transit_secret_backend_key}.
type TransitSecretBackendKey interface {
	cdktn.TerraformResource
	AllowPlaintextBackup() interface{}
	SetAllowPlaintextBackup(val interface{})
	AllowPlaintextBackupInput() interface{}
	AutoRotatePeriod() *float64
	SetAutoRotatePeriod(val *float64)
	AutoRotatePeriodInput() *float64
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	ConvergentEncryption() interface{}
	SetConvergentEncryption(val interface{})
	ConvergentEncryptionInput() interface{}
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
	HybridKeyTypeEc() *string
	SetHybridKeyTypeEc(val *string)
	HybridKeyTypeEcInput() *string
	HybridKeyTypePqc() *string
	SetHybridKeyTypePqc(val *string)
	HybridKeyTypePqcInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Keys() cdktn.StringMapList
	KeySize() *float64
	SetKeySize(val *float64)
	KeySizeInput() *float64
	LatestVersion() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ManagedKeyId() *string
	SetManagedKeyId(val *string)
	ManagedKeyIdInput() *string
	ManagedKeyName() *string
	SetManagedKeyName(val *string)
	ManagedKeyNameInput() *string
	MinAvailableVersion() *float64
	MinDecryptionVersion() *float64
	SetMinDecryptionVersion(val *float64)
	MinDecryptionVersionInput() *float64
	MinEncryptionVersion() *float64
	SetMinEncryptionVersion(val *float64)
	MinEncryptionVersionInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	ParameterSet() *string
	SetParameterSet(val *string)
	ParameterSetInput() *string
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
	ResetAllowPlaintextBackup()
	ResetAutoRotatePeriod()
	ResetContext()
	ResetConvergentEncryption()
	ResetDeletionAllowed()
	ResetDerived()
	ResetExportable()
	ResetHybridKeyTypeEc()
	ResetHybridKeyTypePqc()
	ResetId()
	ResetKeySize()
	ResetManagedKeyId()
	ResetManagedKeyName()
	ResetMinDecryptionVersion()
	ResetMinEncryptionVersion()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterSet()
	ResetType()
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
}

// The jsii proxy struct for TransitSecretBackendKey
type jsiiProxy_TransitSecretBackendKey struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TransitSecretBackendKey) AllowPlaintextBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPlaintextBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) AllowPlaintextBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPlaintextBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) AutoRotatePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) AutoRotatePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRotatePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ConvergentEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convergentEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ConvergentEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"convergentEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) DeletionAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) DeletionAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Derived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"derived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) DerivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"derivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Exportable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ExportableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) HybridKeyTypeEc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridKeyTypeEc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) HybridKeyTypeEcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridKeyTypeEcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) HybridKeyTypePqc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridKeyTypePqc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) HybridKeyTypePqcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridKeyTypePqcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Keys() cdktn.StringMapList {
	var returns cdktn.StringMapList
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) KeySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) KeySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) LatestVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ManagedKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ManagedKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) MinAvailableVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAvailableVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) MinDecryptionVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) MinDecryptionVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDecryptionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) MinEncryptionVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEncryptionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) MinEncryptionVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEncryptionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ParameterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) ParameterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) SupportsDecryption() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsDecryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) SupportsDerivation() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsDerivation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) SupportsEncryption() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) SupportsSigning() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"supportsSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransitSecretBackendKey) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/transit_secret_backend_key vault_transit_secret_backend_key} Resource.
func NewTransitSecretBackendKey(scope constructs.Construct, id *string, config *TransitSecretBackendKeyConfig) TransitSecretBackendKey {
	_init_.Initialize()

	if err := validateNewTransitSecretBackendKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransitSecretBackendKey{}

	_jsii_.Create(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/transit_secret_backend_key vault_transit_secret_backend_key} Resource.
func NewTransitSecretBackendKey_Override(t TransitSecretBackendKey, scope constructs.Construct, id *string, config *TransitSecretBackendKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetAllowPlaintextBackup(val interface{}) {
	if err := j.validateSetAllowPlaintextBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPlaintextBackup",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetAutoRotatePeriod(val *float64) {
	if err := j.validateSetAutoRotatePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRotatePeriod",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetConvergentEncryption(val interface{}) {
	if err := j.validateSetConvergentEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"convergentEncryption",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetDeletionAllowed(val interface{}) {
	if err := j.validateSetDeletionAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionAllowed",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetDerived(val interface{}) {
	if err := j.validateSetDerivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"derived",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetExportable(val interface{}) {
	if err := j.validateSetExportableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportable",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetHybridKeyTypeEc(val *string) {
	if err := j.validateSetHybridKeyTypeEcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridKeyTypeEc",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetHybridKeyTypePqc(val *string) {
	if err := j.validateSetHybridKeyTypePqcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridKeyTypePqc",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetKeySize(val *float64) {
	if err := j.validateSetKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySize",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetManagedKeyId(val *string) {
	if err := j.validateSetManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedKeyId",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetManagedKeyName(val *string) {
	if err := j.validateSetManagedKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedKeyName",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetMinDecryptionVersion(val *float64) {
	if err := j.validateSetMinDecryptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDecryptionVersion",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetMinEncryptionVersion(val *float64) {
	if err := j.validateSetMinEncryptionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minEncryptionVersion",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetParameterSet(val *string) {
	if err := j.validateSetParameterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterSet",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TransitSecretBackendKey)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a TransitSecretBackendKey resource upon running "cdktn plan <stack-name>".
func TransitSecretBackendKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTransitSecretBackendKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
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
func TransitSecretBackendKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransitSecretBackendKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TransitSecretBackendKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTransitSecretBackendKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TransitSecretBackendKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.transitSecretBackendKey.TransitSecretBackendKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TransitSecretBackendKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TransitSecretBackendKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TransitSecretBackendKey) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetAllowPlaintextBackup() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowPlaintextBackup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetAutoRotatePeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRotatePeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetContext() {
	_jsii_.InvokeVoid(
		t,
		"resetContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetConvergentEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetConvergentEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetDeletionAllowed() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionAllowed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetDerived() {
	_jsii_.InvokeVoid(
		t,
		"resetDerived",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetExportable() {
	_jsii_.InvokeVoid(
		t,
		"resetExportable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetHybridKeyTypeEc() {
	_jsii_.InvokeVoid(
		t,
		"resetHybridKeyTypeEc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetHybridKeyTypePqc() {
	_jsii_.InvokeVoid(
		t,
		"resetHybridKeyTypePqc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetKeySize() {
	_jsii_.InvokeVoid(
		t,
		"resetKeySize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetManagedKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetManagedKeyName() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedKeyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetMinDecryptionVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinDecryptionVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetMinEncryptionVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMinEncryptionVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetParameterSet() {
	_jsii_.InvokeVoid(
		t,
		"resetParameterSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransitSecretBackendKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransitSecretBackendKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

