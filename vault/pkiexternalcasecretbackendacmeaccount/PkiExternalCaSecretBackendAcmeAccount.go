// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendacmeaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/pkiexternalcasecretbackendacmeaccount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account vault_pki_external_ca_secret_backend_acme_account}.
type PkiExternalCaSecretBackendAcmeAccount interface {
	cdktn.TerraformResource
	ActiveKeyVersion() *float64
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
	DirectoryUrl() *string
	SetDirectoryUrl(val *string)
	DirectoryUrlInput() *string
	EabKey() *string
	SetEabKey(val *string)
	EabKeyInput() *string
	EabKid() *string
	SetEabKid(val *string)
	EabKidInput() *string
	EmailContacts() *[]*string
	SetEmailContacts(val *[]*string)
	EmailContactsInput() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
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
	TrustedCa() *string
	SetTrustedCa(val *string)
	TrustedCaInput() *string
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
	ResetEabKey()
	ResetEabKid()
	ResetKeyType()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTrustedCa()
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

// The jsii proxy struct for PkiExternalCaSecretBackendAcmeAccount
type jsiiProxy_PkiExternalCaSecretBackendAcmeAccount struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ActiveKeyVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) DirectoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) DirectoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EabKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eabKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EabKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eabKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EabKid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eabKid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EabKidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eabKidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EmailContacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) EmailContactsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) TrustedCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) TrustedCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedCaInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account vault_pki_external_ca_secret_backend_acme_account} Resource.
func NewPkiExternalCaSecretBackendAcmeAccount(scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendAcmeAccountConfig) PkiExternalCaSecretBackendAcmeAccount {
	_init_.Initialize()

	if err := validateNewPkiExternalCaSecretBackendAcmeAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiExternalCaSecretBackendAcmeAccount{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account vault_pki_external_ca_secret_backend_acme_account} Resource.
func NewPkiExternalCaSecretBackendAcmeAccount_Override(p PkiExternalCaSecretBackendAcmeAccount, scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendAcmeAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetDirectoryUrl(val *string) {
	if err := j.validateSetDirectoryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryUrl",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetEabKey(val *string) {
	if err := j.validateSetEabKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eabKey",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetEabKid(val *string) {
	if err := j.validateSetEabKidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eabKid",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetEmailContacts(val *[]*string) {
	if err := j.validateSetEmailContactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailContacts",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount)SetTrustedCa(val *string) {
	if err := j.validateSetTrustedCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedCa",
		val,
	)
}

// Generates CDKTN code for importing a PkiExternalCaSecretBackendAcmeAccount resource upon running "cdktn plan <stack-name>".
func PkiExternalCaSecretBackendAcmeAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendAcmeAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
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
func PkiExternalCaSecretBackendAcmeAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendAcmeAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendAcmeAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendAcmeAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendAcmeAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendAcmeAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiExternalCaSecretBackendAcmeAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendAcmeAccount.PkiExternalCaSecretBackendAcmeAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetEabKey() {
	_jsii_.InvokeVoid(
		p,
		"resetEabKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetEabKid() {
	_jsii_.InvokeVoid(
		p,
		"resetEabKid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetKeyType() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ResetTrustedCa() {
	_jsii_.InvokeVoid(
		p,
		"resetTrustedCa",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendAcmeAccount) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

