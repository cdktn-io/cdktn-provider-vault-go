// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package approleauthbackendrolesecretid

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/approleauthbackendrolesecretid/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id}.
type ApproleAuthBackendRoleSecretId interface {
	cdktn.TerraformResource
	Accessor() *string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CidrList() *[]*string
	SetCidrList(val *[]*string)
	CidrListInput() *[]*string
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
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Metadata() *string
	SetMetadata(val *string)
	MetadataInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NumUses() *float64
	SetNumUses(val *float64)
	NumUsesInput() *float64
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
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	SecretId() *string
	SetSecretId(val *string)
	SecretIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	WithWrappedAccessor() interface{}
	SetWithWrappedAccessor(val interface{})
	WithWrappedAccessorInput() interface{}
	WrappingAccessor() *string
	WrappingToken() *string
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
	ResetBackend()
	ResetCidrList()
	ResetId()
	ResetMetadata()
	ResetNamespace()
	ResetNumUses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecretId()
	ResetTokenBoundCidrs()
	ResetTtl()
	ResetWithWrappedAccessor()
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

// The jsii proxy struct for ApproleAuthBackendRoleSecretId
type jsiiProxy_ApproleAuthBackendRoleSecretId struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CidrList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CidrListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) NumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) NumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WithWrappedAccessor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withWrappedAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WithWrappedAccessorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withWrappedAccessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id} Resource.
func NewApproleAuthBackendRoleSecretId(scope constructs.Construct, id *string, config *ApproleAuthBackendRoleSecretIdConfig) ApproleAuthBackendRoleSecretId {
	_init_.Initialize()

	if err := validateNewApproleAuthBackendRoleSecretIdParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApproleAuthBackendRoleSecretId{}

	_jsii_.Create(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id} Resource.
func NewApproleAuthBackendRoleSecretId_Override(a ApproleAuthBackendRoleSecretId, scope constructs.Construct, id *string, config *ApproleAuthBackendRoleSecretIdConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetCidrList(val *[]*string) {
	if err := j.validateSetCidrListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrList",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetNumUses(val *float64) {
	if err := j.validateSetNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numUses",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetSecretId(val *string) {
	if err := j.validateSetSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetWithWrappedAccessor(val interface{}) {
	if err := j.validateSetWithWrappedAccessorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withWrappedAccessor",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId)SetWrappingTtl(val *string) {
	if err := j.validateSetWrappingTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrappingTtl",
		val,
	)
}

// Generates CDKTN code for importing a ApproleAuthBackendRoleSecretId resource upon running "cdktn plan <stack-name>".
func ApproleAuthBackendRoleSecretId_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApproleAuthBackendRoleSecretId_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
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
func ApproleAuthBackendRoleSecretId_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApproleAuthBackendRoleSecretId_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApproleAuthBackendRoleSecretId_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApproleAuthBackendRoleSecretId_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApproleAuthBackendRoleSecretId_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApproleAuthBackendRoleSecretId_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApproleAuthBackendRoleSecretId_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.approleAuthBackendRoleSecretId.ApproleAuthBackendRoleSecretId",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetCidrList() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetNumUses() {
	_jsii_.InvokeVoid(
		a,
		"resetNumUses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetSecretId() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetWithWrappedAccessor() {
	_jsii_.InvokeVoid(
		a,
		"resetWithWrappedAccessor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetWrappingTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetWrappingTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

