// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/pkiexternalcasecretbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_role vault_pki_external_ca_secret_backend_role}.
type PkiExternalCaSecretBackendRole interface {
	cdktn.TerraformResource
	AcmeAccountName() *string
	SetAcmeAccountName(val *string)
	AcmeAccountNameInput() *string
	AllowedChallengeTypes() *[]*string
	SetAllowedChallengeTypes(val *[]*string)
	AllowedChallengeTypesInput() *[]*string
	AllowedDomainOptions() *[]*string
	SetAllowedDomainOptions(val *[]*string)
	AllowedDomainOptionsInput() *[]*string
	AllowedDomains() *[]*string
	SetAllowedDomains(val *[]*string)
	AllowedDomainsInput() *[]*string
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
	CreationDate() *string
	CsrGenerateKeyType() *string
	SetCsrGenerateKeyType(val *string)
	CsrGenerateKeyTypeInput() *string
	CsrIdentifierPopulation() *string
	SetCsrIdentifierPopulation(val *string)
	CsrIdentifierPopulationInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Force() interface{}
	SetForce(val interface{})
	ForceInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	LastUpdateDate() *string
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
	ResetAllowedChallengeTypes()
	ResetAllowedDomainOptions()
	ResetAllowedDomains()
	ResetCsrGenerateKeyType()
	ResetCsrIdentifierPopulation()
	ResetForce()
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

// The jsii proxy struct for PkiExternalCaSecretBackendRole
type jsiiProxy_PkiExternalCaSecretBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AcmeAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmeAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AcmeAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmeAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedChallengeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedChallengeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedChallengeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedChallengeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedDomainOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedDomainOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) AllowedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CsrGenerateKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrGenerateKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CsrGenerateKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrGenerateKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CsrIdentifierPopulation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrIdentifierPopulation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) CsrIdentifierPopulationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrIdentifierPopulationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Force() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) ForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) LastUpdateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_role vault_pki_external_ca_secret_backend_role} Resource.
func NewPkiExternalCaSecretBackendRole(scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendRoleConfig) PkiExternalCaSecretBackendRole {
	_init_.Initialize()

	if err := validateNewPkiExternalCaSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiExternalCaSecretBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_role vault_pki_external_ca_secret_backend_role} Resource.
func NewPkiExternalCaSecretBackendRole_Override(p PkiExternalCaSecretBackendRole, scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetAcmeAccountName(val *string) {
	if err := j.validateSetAcmeAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmeAccountName",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetAllowedChallengeTypes(val *[]*string) {
	if err := j.validateSetAllowedChallengeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedChallengeTypes",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetAllowedDomainOptions(val *[]*string) {
	if err := j.validateSetAllowedDomainOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomainOptions",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetAllowedDomains(val *[]*string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetCsrGenerateKeyType(val *string) {
	if err := j.validateSetCsrGenerateKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csrGenerateKeyType",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetCsrIdentifierPopulation(val *string) {
	if err := j.validateSetCsrIdentifierPopulationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csrIdentifierPopulation",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetForce(val interface{}) {
	if err := j.validateSetForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"force",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a PkiExternalCaSecretBackendRole resource upon running "cdktn plan <stack-name>".
func PkiExternalCaSecretBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
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
func PkiExternalCaSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiExternalCaSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetAllowedChallengeTypes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedChallengeTypes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetAllowedDomainOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDomainOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetCsrGenerateKeyType() {
	_jsii_.InvokeVoid(
		p,
		"resetCsrGenerateKeyType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetCsrIdentifierPopulation() {
	_jsii_.InvokeVoid(
		p,
		"resetCsrIdentifierPopulation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetForce() {
	_jsii_.InvokeVoid(
		p,
		"resetForce",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

