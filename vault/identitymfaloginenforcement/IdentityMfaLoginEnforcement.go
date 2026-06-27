// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfaloginenforcement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/identitymfaloginenforcement/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_mfa_login_enforcement vault_identity_mfa_login_enforcement}.
type IdentityMfaLoginEnforcement interface {
	cdktn.TerraformResource
	AuthMethodAccessors() *[]*string
	SetAuthMethodAccessors(val *[]*string)
	AuthMethodAccessorsInput() *[]*string
	AuthMethodTypes() *[]*string
	SetAuthMethodTypes(val *[]*string)
	AuthMethodTypesInput() *[]*string
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
	IdentityEntityIds() *[]*string
	SetIdentityEntityIds(val *[]*string)
	IdentityEntityIdsInput() *[]*string
	IdentityGroupIds() *[]*string
	SetIdentityGroupIds(val *[]*string)
	IdentityGroupIdsInput() *[]*string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MfaMethodIds() *[]*string
	SetMfaMethodIds(val *[]*string)
	MfaMethodIdsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceId() *string
	NamespaceInput() *string
	NamespacePath() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Uuid() *string
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
	ResetAuthMethodAccessors()
	ResetAuthMethodTypes()
	ResetId()
	ResetIdentityEntityIds()
	ResetIdentityGroupIds()
	ResetNamespace()
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

// The jsii proxy struct for IdentityMfaLoginEnforcement
type jsiiProxy_IdentityMfaLoginEnforcement struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) AuthMethodAccessors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authMethodAccessors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) AuthMethodAccessorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authMethodAccessorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) AuthMethodTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authMethodTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) AuthMethodTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authMethodTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) IdentityEntityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityEntityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) IdentityEntityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityEntityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) IdentityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) IdentityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) MfaMethodIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mfaMethodIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) MfaMethodIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mfaMethodIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) NamespacePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_mfa_login_enforcement vault_identity_mfa_login_enforcement} Resource.
func NewIdentityMfaLoginEnforcement(scope constructs.Construct, id *string, config *IdentityMfaLoginEnforcementConfig) IdentityMfaLoginEnforcement {
	_init_.Initialize()

	if err := validateNewIdentityMfaLoginEnforcementParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityMfaLoginEnforcement{}

	_jsii_.Create(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_mfa_login_enforcement vault_identity_mfa_login_enforcement} Resource.
func NewIdentityMfaLoginEnforcement_Override(i IdentityMfaLoginEnforcement, scope constructs.Construct, id *string, config *IdentityMfaLoginEnforcementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetAuthMethodAccessors(val *[]*string) {
	if err := j.validateSetAuthMethodAccessorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethodAccessors",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetAuthMethodTypes(val *[]*string) {
	if err := j.validateSetAuthMethodTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authMethodTypes",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetIdentityEntityIds(val *[]*string) {
	if err := j.validateSetIdentityEntityIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityEntityIds",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetIdentityGroupIds(val *[]*string) {
	if err := j.validateSetIdentityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityGroupIds",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetMfaMethodIds(val *[]*string) {
	if err := j.validateSetMfaMethodIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mfaMethodIds",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaLoginEnforcement)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a IdentityMfaLoginEnforcement resource upon running "cdktn plan <stack-name>".
func IdentityMfaLoginEnforcement_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIdentityMfaLoginEnforcement_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
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
func IdentityMfaLoginEnforcement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaLoginEnforcement_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaLoginEnforcement_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaLoginEnforcement_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaLoginEnforcement_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaLoginEnforcement_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityMfaLoginEnforcement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.identityMfaLoginEnforcement.IdentityMfaLoginEnforcement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetAuthMethodAccessors() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthMethodAccessors",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetAuthMethodTypes() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthMethodTypes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetIdentityEntityIds() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentityEntityIds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetIdentityGroupIds() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentityGroupIds",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetNamespace() {
	_jsii_.InvokeVoid(
		i,
		"resetNamespace",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaLoginEnforcement) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

