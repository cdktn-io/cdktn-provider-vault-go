// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaulttoken

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/ephemeralvaulttoken/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token vault_token}.
type EphemeralVaultToken interface {
	cdktn.TerraformEphemeralResource
	Accessor() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientToken() *string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EntityAlias() *string
	SetEntityAlias(val *string)
	EntityAliasInput() *string
	EntityId() *string
	ExplicitMaxTtl() *string
	SetExplicitMaxTtl(val *string)
	ExplicitMaxTtlInput() *string
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
	LeaseDuration() *float64
	LeaseId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	MountId() *string
	SetMountId(val *string)
	MountIdInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NoDefaultPolicy() interface{}
	SetNoDefaultPolicy(val interface{})
	NoDefaultPolicyInput() interface{}
	NoParent() interface{}
	SetNoParent(val interface{})
	NoParentInput() interface{}
	NumUses() *float64
	SetNumUses(val *float64)
	NumUsesInput() *float64
	Orphan() cdktn.IResolvable
	Period() *string
	SetPeriod(val *string)
	PeriodInput() *string
	Policies() *[]*string
	SetPolicies(val *[]*string)
	PoliciesInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Renewable() interface{}
	SetRenewable(val interface{})
	RenewableInput() interface{}
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenPolicies() *[]*string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WrappedToken() *string
	WrappingAccessor() *string
	WrappingTtl() *string
	SetWrappingTtl(val *string)
	WrappingTtlInput() *string
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
	ResetDisplayName()
	ResetEntityAlias()
	ResetExplicitMaxTtl()
	ResetId()
	ResetMetadata()
	ResetMountId()
	ResetNamespace()
	ResetNoDefaultPolicy()
	ResetNoParent()
	ResetNumUses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetPolicies()
	ResetRenewable()
	ResetRoleName()
	ResetTtl()
	ResetType()
	ResetWrappingTtl()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this ephemeral resource to the terraform JSON output.
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

// The jsii proxy struct for EphemeralVaultToken
type jsiiProxy_EphemeralVaultToken struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralVaultToken) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) EntityAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) EntityAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) ExplicitMaxTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) ExplicitMaxTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) LeaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) MountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) MountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NoParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NoParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) NumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Orphan() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"orphan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) PoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Renewable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) RenewableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) WrappedToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappedToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) WrappingAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) WrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultToken) WrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token vault_token} Ephemeral Resource.
func NewEphemeralVaultToken(scope constructs.Construct, id *string, config *EphemeralVaultTokenConfig) EphemeralVaultToken {
	_init_.Initialize()

	if err := validateNewEphemeralVaultTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralVaultToken{}

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/token vault_token} Ephemeral Resource.
func NewEphemeralVaultToken_Override(e EphemeralVaultToken, scope constructs.Construct, id *string, config *EphemeralVaultTokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetEntityAlias(val *string) {
	if err := j.validateSetEntityAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityAlias",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetExplicitMaxTtl(val *string) {
	if err := j.validateSetExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetMountId(val *string) {
	if err := j.validateSetMountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountId",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetNoDefaultPolicy(val interface{}) {
	if err := j.validateSetNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetNoParent(val interface{}) {
	if err := j.validateSetNoParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noParent",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetNumUses(val *float64) {
	if err := j.validateSetNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numUses",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetPolicies(val *[]*string) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetRenewable(val interface{}) {
	if err := j.validateSetRenewableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewable",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultToken)SetWrappingTtl(val *string) {
	if err := j.validateSetWrappingTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrappingTtl",
		val,
	)
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
func EphemeralVaultToken_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultToken_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultToken_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultToken_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralVaultToken_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.ephemeralVaultToken.EphemeralVaultToken",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralVaultToken) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralVaultToken) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetDisplayName() {
	_jsii_.InvokeVoid(
		e,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetEntityAlias() {
	_jsii_.InvokeVoid(
		e,
		"resetEntityAlias",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		e,
		"resetExplicitMaxTtl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetMetadata() {
	_jsii_.InvokeVoid(
		e,
		"resetMetadata",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetMountId() {
	_jsii_.InvokeVoid(
		e,
		"resetMountId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetNamespace() {
	_jsii_.InvokeVoid(
		e,
		"resetNamespace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetNoDefaultPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetNoParent() {
	_jsii_.InvokeVoid(
		e,
		"resetNoParent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetNumUses() {
	_jsii_.InvokeVoid(
		e,
		"resetNumUses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetPeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetPeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetPolicies() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetRenewable() {
	_jsii_.InvokeVoid(
		e,
		"resetRenewable",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetRoleName() {
	_jsii_.InvokeVoid(
		e,
		"resetRoleName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetTtl() {
	_jsii_.InvokeVoid(
		e,
		"resetTtl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) ResetWrappingTtl() {
	_jsii_.InvokeVoid(
		e,
		"resetWrappingTtl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultToken) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultToken) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

