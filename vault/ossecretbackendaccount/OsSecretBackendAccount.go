// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ossecretbackendaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/ossecretbackendaccount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account vault_os_secret_backend_account}.
type OsSecretBackendAccount interface {
	cdktn.TerraformResource
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
	CustomMetadata() *map[string]*string
	SetCustomMetadata(val *map[string]*string)
	CustomMetadataInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	LastVaultRotation() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Mount() *string
	SetMount(val *string)
	MountInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NextVaultRotation() *string
	// The tree node.
	Node() constructs.Node
	ParentAccountRef() *string
	SetParentAccountRef(val *string)
	ParentAccountRefInput() *string
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
	PasswordWo() *string
	SetPasswordWo(val *string)
	PasswordWoInput() *string
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
	RotationPeriod() *float64
	SetRotationPeriod(val *float64)
	RotationPeriodInput() *float64
	RotationSchedule() *string
	SetRotationSchedule(val *string)
	RotationScheduleInput() *string
	RotationWindow() *float64
	SetRotationWindow(val *float64)
	RotationWindowInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VerifyConnection() interface{}
	SetVerifyConnection(val interface{})
	VerifyConnectionInput() interface{}
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
	ResetCustomMetadata()
	ResetDisableAutomatedRotation()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentAccountRef()
	ResetPasswordPolicy()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetVerifyConnection()
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

// The jsii proxy struct for OsSecretBackendAccount
type jsiiProxy_OsSecretBackendAccount struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OsSecretBackendAccount) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) CustomMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) CustomMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) LastVaultRotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVaultRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) NextVaultRotation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextVaultRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) ParentAccountRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentAccountRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) ParentAccountRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentAccountRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsSecretBackendAccount) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account vault_os_secret_backend_account} Resource.
func NewOsSecretBackendAccount(scope constructs.Construct, id *string, config *OsSecretBackendAccountConfig) OsSecretBackendAccount {
	_init_.Initialize()

	if err := validateNewOsSecretBackendAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsSecretBackendAccount{}

	_jsii_.Create(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/os_secret_backend_account vault_os_secret_backend_account} Resource.
func NewOsSecretBackendAccount_Override(o OsSecretBackendAccount, scope constructs.Construct, id *string, config *OsSecretBackendAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetCustomMetadata(val *map[string]*string) {
	if err := j.validateSetCustomMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMetadata",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetParentAccountRef(val *string) {
	if err := j.validateSetParentAccountRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentAccountRef",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_OsSecretBackendAccount)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
		val,
	)
}

// Generates CDKTN code for importing a OsSecretBackendAccount resource upon running "cdktn plan <stack-name>".
func OsSecretBackendAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOsSecretBackendAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
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
func OsSecretBackendAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsSecretBackendAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsSecretBackendAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsSecretBackendAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsSecretBackendAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsSecretBackendAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OsSecretBackendAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.osSecretBackendAccount.OsSecretBackendAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetCustomMetadata() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomMetadata",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		o,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetNamespace() {
	_jsii_.InvokeVoid(
		o,
		"resetNamespace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetParentAccountRef() {
	_jsii_.InvokeVoid(
		o,
		"resetParentAccountRef",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		o,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		o,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		o,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsSecretBackendAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsSecretBackendAccount) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		o,
		"with",
		args,
		&returns,
	)

	return returns
}

