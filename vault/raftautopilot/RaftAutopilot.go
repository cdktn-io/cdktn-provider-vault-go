// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package raftautopilot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/raftautopilot/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/raft_autopilot vault_raft_autopilot}.
type RaftAutopilot interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CleanupDeadServers() interface{}
	SetCleanupDeadServers(val interface{})
	CleanupDeadServersInput() interface{}
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
	DeadServerLastContactThreshold() *string
	SetDeadServerLastContactThreshold(val *string)
	DeadServerLastContactThresholdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableUpgradeMigration() interface{}
	SetDisableUpgradeMigration(val interface{})
	DisableUpgradeMigrationInput() interface{}
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
	LastContactThreshold() *string
	SetLastContactThreshold(val *string)
	LastContactThresholdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxTrailingLogs() *float64
	SetMaxTrailingLogs(val *float64)
	MaxTrailingLogsInput() *float64
	MinQuorum() *float64
	SetMinQuorum(val *float64)
	MinQuorumInput() *float64
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
	ServerStabilizationTime() *string
	SetServerStabilizationTime(val *string)
	ServerStabilizationTimeInput() *string
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
	ResetCleanupDeadServers()
	ResetDeadServerLastContactThreshold()
	ResetDisableUpgradeMigration()
	ResetId()
	ResetLastContactThreshold()
	ResetMaxTrailingLogs()
	ResetMinQuorum()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerStabilizationTime()
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

// The jsii proxy struct for RaftAutopilot
type jsiiProxy_RaftAutopilot struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RaftAutopilot) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) CleanupDeadServers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupDeadServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) CleanupDeadServersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupDeadServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) DeadServerLastContactThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadServerLastContactThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) DeadServerLastContactThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadServerLastContactThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) DisableUpgradeMigration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpgradeMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) DisableUpgradeMigrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableUpgradeMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) LastContactThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastContactThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) LastContactThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastContactThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) MaxTrailingLogs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrailingLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) MaxTrailingLogsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrailingLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) MinQuorum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minQuorum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) MinQuorumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minQuorumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) ServerStabilizationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverStabilizationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) ServerStabilizationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverStabilizationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftAutopilot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/raft_autopilot vault_raft_autopilot} Resource.
func NewRaftAutopilot(scope constructs.Construct, id *string, config *RaftAutopilotConfig) RaftAutopilot {
	_init_.Initialize()

	if err := validateNewRaftAutopilotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RaftAutopilot{}

	_jsii_.Create(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/raft_autopilot vault_raft_autopilot} Resource.
func NewRaftAutopilot_Override(r RaftAutopilot, scope constructs.Construct, id *string, config *RaftAutopilotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetCleanupDeadServers(val interface{}) {
	if err := j.validateSetCleanupDeadServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupDeadServers",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetDeadServerLastContactThreshold(val *string) {
	if err := j.validateSetDeadServerLastContactThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadServerLastContactThreshold",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetDisableUpgradeMigration(val interface{}) {
	if err := j.validateSetDisableUpgradeMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableUpgradeMigration",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetLastContactThreshold(val *string) {
	if err := j.validateSetLastContactThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastContactThreshold",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetMaxTrailingLogs(val *float64) {
	if err := j.validateSetMaxTrailingLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTrailingLogs",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetMinQuorum(val *float64) {
	if err := j.validateSetMinQuorumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minQuorum",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RaftAutopilot)SetServerStabilizationTime(val *string) {
	if err := j.validateSetServerStabilizationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverStabilizationTime",
		val,
	)
}

// Generates CDKTN code for importing a RaftAutopilot resource upon running "cdktn plan <stack-name>".
func RaftAutopilot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRaftAutopilot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
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
func RaftAutopilot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftAutopilot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RaftAutopilot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftAutopilot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RaftAutopilot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftAutopilot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RaftAutopilot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.raftAutopilot.RaftAutopilot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RaftAutopilot) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RaftAutopilot) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RaftAutopilot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RaftAutopilot) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RaftAutopilot) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RaftAutopilot) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RaftAutopilot) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetCleanupDeadServers() {
	_jsii_.InvokeVoid(
		r,
		"resetCleanupDeadServers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetDeadServerLastContactThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetDeadServerLastContactThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetDisableUpgradeMigration() {
	_jsii_.InvokeVoid(
		r,
		"resetDisableUpgradeMigration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetLastContactThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetLastContactThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetMaxTrailingLogs() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxTrailingLogs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetMinQuorum() {
	_jsii_.InvokeVoid(
		r,
		"resetMinQuorum",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) ResetServerStabilizationTime() {
	_jsii_.InvokeVoid(
		r,
		"resetServerStabilizationTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftAutopilot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftAutopilot) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

