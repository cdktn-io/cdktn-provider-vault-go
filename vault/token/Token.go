// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package token

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/token/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/token vault_token}.
type Token interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientToken() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	LeaseStarted() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Renewable() interface{}
	SetRenewable(val interface{})
	RenewableInput() interface{}
	RenewIncrement() *float64
	SetRenewIncrement(val *float64)
	RenewIncrementInput() *float64
	RenewMinLease() *float64
	SetRenewMinLease(val *float64)
	RenewMinLeaseInput() *float64
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	WrappedToken() *string
	WrappingAccessor() *string
	WrappingTtl() *string
	SetWrappingTtl(val *string)
	WrappingTtlInput() *string
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
	ResetDisplayName()
	ResetExplicitMaxTtl()
	ResetId()
	ResetMetadata()
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
	ResetRenewIncrement()
	ResetRenewMinLease()
	ResetRoleName()
	ResetTtl()
	ResetWrappingTtl()
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

// The jsii proxy struct for Token
type jsiiProxy_Token struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Token) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) ExplicitMaxTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) ExplicitMaxTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) LeaseStarted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseStarted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NoParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NoParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) NumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) PoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Renewable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RenewableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RenewIncrement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renewIncrement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RenewIncrementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renewIncrementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RenewMinLease() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renewMinLease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RenewMinLeaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renewMinLeaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) WrappedToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappedToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) WrappingAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) WrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Token) WrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/token vault_token} Resource.
func NewToken(scope constructs.Construct, id *string, config *TokenConfig) Token {
	_init_.Initialize()

	if err := validateNewTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Token{}

	_jsii_.Create(
		"@cdktn/provider-vault.token.Token",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/token vault_token} Resource.
func NewToken_Override(t Token, scope constructs.Construct, id *string, config *TokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.token.Token",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_Token)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Token)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Token)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Token)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_Token)SetExplicitMaxTtl(val *string) {
	if err := j.validateSetExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_Token)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Token)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Token)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Token)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Token)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_Token)SetNoDefaultPolicy(val interface{}) {
	if err := j.validateSetNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_Token)SetNoParent(val interface{}) {
	if err := j.validateSetNoParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noParent",
		val,
	)
}

func (j *jsiiProxy_Token)SetNumUses(val *float64) {
	if err := j.validateSetNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numUses",
		val,
	)
}

func (j *jsiiProxy_Token)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_Token)SetPolicies(val *[]*string) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_Token)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Token)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Token)SetRenewable(val interface{}) {
	if err := j.validateSetRenewableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewable",
		val,
	)
}

func (j *jsiiProxy_Token)SetRenewIncrement(val *float64) {
	if err := j.validateSetRenewIncrementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewIncrement",
		val,
	)
}

func (j *jsiiProxy_Token)SetRenewMinLease(val *float64) {
	if err := j.validateSetRenewMinLeaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewMinLease",
		val,
	)
}

func (j *jsiiProxy_Token)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_Token)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_Token)SetWrappingTtl(val *string) {
	if err := j.validateSetWrappingTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrappingTtl",
		val,
	)
}

// Generates CDKTN code for importing a Token resource upon running "cdktn plan <stack-name>".
func Token_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateToken_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.token.Token",
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
func Token_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.token.Token",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Token_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.token.Token",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Token_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateToken_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.token.Token",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Token_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.token.Token",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_Token) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_Token) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_Token) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_Token) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_Token) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_Token) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_Token) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_Token) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_Token) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_Token) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_Token) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_Token) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_Token) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_Token) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_Token) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_Token) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_Token) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_Token) ResetDisplayName() {
	_jsii_.InvokeVoid(
		t,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetExplicitMaxTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetNoDefaultPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetNoParent() {
	_jsii_.InvokeVoid(
		t,
		"resetNoParent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetNumUses() {
	_jsii_.InvokeVoid(
		t,
		"resetNumUses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetPolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetRenewable() {
	_jsii_.InvokeVoid(
		t,
		"resetRenewable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetRenewIncrement() {
	_jsii_.InvokeVoid(
		t,
		"resetRenewIncrement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetRenewMinLease() {
	_jsii_.InvokeVoid(
		t,
		"resetRenewMinLease",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetRoleName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoleName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) ResetWrappingTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetWrappingTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Token) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Token) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

