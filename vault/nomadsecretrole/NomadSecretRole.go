// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nomadsecretrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/nomadsecretrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/nomad_secret_role vault_nomad_secret_role}.
type NomadSecretRole interface {
	cdktn.TerraformResource
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
	Global() interface{}
	SetGlobal(val interface{})
	GlobalInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
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
	ResetGlobal()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicies()
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

// The jsii proxy struct for NomadSecretRole
type jsiiProxy_NomadSecretRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_NomadSecretRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Global() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"global",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) GlobalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) PoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NomadSecretRole) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/nomad_secret_role vault_nomad_secret_role} Resource.
func NewNomadSecretRole(scope constructs.Construct, id *string, config *NomadSecretRoleConfig) NomadSecretRole {
	_init_.Initialize()

	if err := validateNewNomadSecretRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NomadSecretRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/nomad_secret_role vault_nomad_secret_role} Resource.
func NewNomadSecretRole_Override(n NomadSecretRole, scope constructs.Construct, id *string, config *NomadSecretRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetGlobal(val interface{}) {
	if err := j.validateSetGlobalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"global",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetPolicies(val *[]*string) {
	if err := j.validateSetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_NomadSecretRole)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a NomadSecretRole resource upon running "cdktn plan <stack-name>".
func NomadSecretRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateNomadSecretRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
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
func NomadSecretRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadSecretRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NomadSecretRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadSecretRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NomadSecretRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNomadSecretRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NomadSecretRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.nomadSecretRole.NomadSecretRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NomadSecretRole) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NomadSecretRole) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NomadSecretRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NomadSecretRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NomadSecretRole) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NomadSecretRole) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NomadSecretRole) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetGlobal() {
	_jsii_.InvokeVoid(
		n,
		"resetGlobal",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		n,
		"resetNamespace",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetPolicies() {
	_jsii_.InvokeVoid(
		n,
		"resetPolicies",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NomadSecretRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NomadSecretRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		n,
		"with",
		args,
		&returns,
	)

	return returns
}

