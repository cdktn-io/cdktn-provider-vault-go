// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultgcpauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaultgcpauthbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/gcp_auth_backend_role vault_gcp_auth_backend_role}.
type DataVaultGcpAuthBackendRole interface {
	cdktn.TerraformDataSource
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BoundInstanceGroups() *[]*string
	BoundLabels() *[]*string
	BoundProjects() *[]*string
	BoundRegions() *[]*string
	BoundServiceAccounts() *[]*string
	BoundZones() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	RawOverrides() interface{}
	RoleId() *string
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	TokenExplicitMaxTtl() *float64
	SetTokenExplicitMaxTtl(val *float64)
	TokenExplicitMaxTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	TokenNoDefaultPolicy() interface{}
	SetTokenNoDefaultPolicy(val interface{})
	TokenNoDefaultPolicyInput() interface{}
	TokenNumUses() *float64
	SetTokenNumUses(val *float64)
	TokenNumUsesInput() *float64
	TokenPeriod() *float64
	SetTokenPeriod(val *float64)
	TokenPeriodInput() *float64
	TokenPolicies() *[]*string
	SetTokenPolicies(val *[]*string)
	TokenPoliciesInput() *[]*string
	TokenTtl() *float64
	SetTokenTtl(val *float64)
	TokenTtlInput() *float64
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
	Type() *string
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
	ResetAliasMetadata()
	ResetBackend()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataVaultGcpAuthBackendRole
type jsiiProxy_DataVaultGcpAuthBackendRole struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundInstanceGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundInstanceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) BoundZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/gcp_auth_backend_role vault_gcp_auth_backend_role} Data Source.
func NewDataVaultGcpAuthBackendRole(scope constructs.Construct, id *string, config *DataVaultGcpAuthBackendRoleConfig) DataVaultGcpAuthBackendRole {
	_init_.Initialize()

	if err := validateNewDataVaultGcpAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultGcpAuthBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/gcp_auth_backend_role vault_gcp_auth_backend_role} Data Source.
func NewDataVaultGcpAuthBackendRole_Override(d DataVaultGcpAuthBackendRole, scope constructs.Construct, id *string, config *DataVaultGcpAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_DataVaultGcpAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultGcpAuthBackendRole resource upon running "cdktn plan <stack-name>".
func DataVaultGcpAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultGcpAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
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
func DataVaultGcpAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultGcpAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultGcpAuthBackendRole_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultGcpAuthBackendRole_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultGcpAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultGcpAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultGcpAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultGcpAuthBackendRole.DataVaultGcpAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		d,
		"resetBackend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		d,
		"resetTokenType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultGcpAuthBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

