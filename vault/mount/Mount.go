// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/mount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/mount vault_mount}.
type Mount interface {
	cdktn.TerraformResource
	Accessor() *string
	AllowedManagedKeys() *[]*string
	SetAllowedManagedKeys(val *[]*string)
	AllowedManagedKeysInput() *[]*string
	AllowedResponseHeaders() *[]*string
	SetAllowedResponseHeaders(val *[]*string)
	AllowedResponseHeadersInput() *[]*string
	AuditNonHmacRequestKeys() *[]*string
	SetAuditNonHmacRequestKeys(val *[]*string)
	AuditNonHmacRequestKeysInput() *[]*string
	AuditNonHmacResponseKeys() *[]*string
	SetAuditNonHmacResponseKeys(val *[]*string)
	AuditNonHmacResponseKeysInput() *[]*string
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
	DefaultLeaseTtlSeconds() *float64
	SetDefaultLeaseTtlSeconds(val *float64)
	DefaultLeaseTtlSecondsInput() *float64
	DelegatedAuthAccessors() *[]*string
	SetDelegatedAuthAccessors(val *[]*string)
	DelegatedAuthAccessorsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExternalEntropyAccess() interface{}
	SetExternalEntropyAccess(val interface{})
	ExternalEntropyAccessInput() interface{}
	ForceNoCache() interface{}
	SetForceNoCache(val interface{})
	ForceNoCacheInput() interface{}
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
	IdentityTokenKey() *string
	SetIdentityTokenKey(val *string)
	IdentityTokenKeyInput() *string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListingVisibility() *string
	SetListingVisibility(val *string)
	ListingVisibilityInput() *string
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	PassthroughRequestHeaders() *[]*string
	SetPassthroughRequestHeaders(val *[]*string)
	PassthroughRequestHeadersInput() *[]*string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PluginVersion() *string
	SetPluginVersion(val *string)
	PluginVersionInput() *string
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
	SealWrap() interface{}
	SetSealWrap(val interface{})
	SealWrapInput() interface{}
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
	ResetAllowedManagedKeys()
	ResetAllowedResponseHeaders()
	ResetAuditNonHmacRequestKeys()
	ResetAuditNonHmacResponseKeys()
	ResetDefaultLeaseTtlSeconds()
	ResetDelegatedAuthAccessors()
	ResetDescription()
	ResetExternalEntropyAccess()
	ResetForceNoCache()
	ResetId()
	ResetIdentityTokenKey()
	ResetListingVisibility()
	ResetLocal()
	ResetMaxLeaseTtlSeconds()
	ResetNamespace()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassthroughRequestHeaders()
	ResetPluginVersion()
	ResetSealWrap()
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

// The jsii proxy struct for Mount
type jsiiProxy_Mount struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Mount) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AllowedManagedKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AllowedManagedKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AllowedResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AllowedResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DefaultLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DefaultLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DelegatedAuthAccessors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DelegatedAuthAccessorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ExternalEntropyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ExternalEntropyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ForceNoCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ForceNoCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) IdentityTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) IdentityTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ListingVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) ListingVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) PassthroughRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) PassthroughRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) PluginVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) PluginVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) SealWrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) SealWrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mount) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/mount vault_mount} Resource.
func NewMount(scope constructs.Construct, id *string, config *MountConfig) Mount {
	_init_.Initialize()

	if err := validateNewMountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mount{}

	_jsii_.Create(
		"@cdktn/provider-vault.mount.Mount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/mount vault_mount} Resource.
func NewMount_Override(m Mount, scope constructs.Construct, id *string, config *MountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.mount.Mount",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Mount)SetAllowedManagedKeys(val *[]*string) {
	if err := j.validateSetAllowedManagedKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedManagedKeys",
		val,
	)
}

func (j *jsiiProxy_Mount)SetAllowedResponseHeaders(val *[]*string) {
	if err := j.validateSetAllowedResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_Mount)SetAuditNonHmacRequestKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacRequestKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_Mount)SetAuditNonHmacResponseKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacResponseKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_Mount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Mount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Mount)SetDefaultLeaseTtlSeconds(val *float64) {
	if err := j.validateSetDefaultLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_Mount)SetDelegatedAuthAccessors(val *[]*string) {
	if err := j.validateSetDelegatedAuthAccessorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedAuthAccessors",
		val,
	)
}

func (j *jsiiProxy_Mount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Mount)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Mount)SetExternalEntropyAccess(val interface{}) {
	if err := j.validateSetExternalEntropyAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEntropyAccess",
		val,
	)
}

func (j *jsiiProxy_Mount)SetForceNoCache(val interface{}) {
	if err := j.validateSetForceNoCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNoCache",
		val,
	)
}

func (j *jsiiProxy_Mount)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Mount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Mount)SetIdentityTokenKey(val *string) {
	if err := j.validateSetIdentityTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKey",
		val,
	)
}

func (j *jsiiProxy_Mount)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Mount)SetListingVisibility(val *string) {
	if err := j.validateSetListingVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingVisibility",
		val,
	)
}

func (j *jsiiProxy_Mount)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_Mount)SetMaxLeaseTtlSeconds(val *float64) {
	if err := j.validateSetMaxLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_Mount)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_Mount)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_Mount)SetPassthroughRequestHeaders(val *[]*string) {
	if err := j.validateSetPassthroughRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_Mount)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_Mount)SetPluginVersion(val *string) {
	if err := j.validateSetPluginVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginVersion",
		val,
	)
}

func (j *jsiiProxy_Mount)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Mount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Mount)SetSealWrap(val interface{}) {
	if err := j.validateSetSealWrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sealWrap",
		val,
	)
}

func (j *jsiiProxy_Mount)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a Mount resource upon running "cdktn plan <stack-name>".
func Mount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.mount.Mount",
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
func Mount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.mount.Mount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Mount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.mount.Mount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Mount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.mount.Mount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Mount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.mount.Mount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Mount) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Mount) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Mount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Mount) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Mount) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Mount) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Mount) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Mount) ResetAllowedManagedKeys() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowedManagedKeys",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetAllowedResponseHeaders() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowedResponseHeaders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		m,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		m,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetDefaultLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetDelegatedAuthAccessors() {
	_jsii_.InvokeVoid(
		m,
		"resetDelegatedAuthAccessors",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetExternalEntropyAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetExternalEntropyAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetForceNoCache() {
	_jsii_.InvokeVoid(
		m,
		"resetForceNoCache",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetIdentityTokenKey() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentityTokenKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetListingVisibility() {
	_jsii_.InvokeVoid(
		m,
		"resetListingVisibility",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetLocal() {
	_jsii_.InvokeVoid(
		m,
		"resetLocal",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetNamespace() {
	_jsii_.InvokeVoid(
		m,
		"resetNamespace",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetPassthroughRequestHeaders() {
	_jsii_.InvokeVoid(
		m,
		"resetPassthroughRequestHeaders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetPluginVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) ResetSealWrap() {
	_jsii_.InvokeVoid(
		m,
		"resetSealWrap",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Mount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mount) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

