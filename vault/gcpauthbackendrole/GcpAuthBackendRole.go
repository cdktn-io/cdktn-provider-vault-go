// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/gcpauthbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_auth_backend_role vault_gcp_auth_backend_role}.
type GcpAuthBackendRole interface {
	cdktn.TerraformResource
	AddGroupAliases() interface{}
	SetAddGroupAliases(val interface{})
	AddGroupAliasesInput() interface{}
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AllowGceInference() interface{}
	SetAllowGceInference(val interface{})
	AllowGceInferenceInput() interface{}
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BoundInstanceGroups() *[]*string
	SetBoundInstanceGroups(val *[]*string)
	BoundInstanceGroupsInput() *[]*string
	BoundLabels() *[]*string
	SetBoundLabels(val *[]*string)
	BoundLabelsInput() *[]*string
	BoundProjects() *[]*string
	SetBoundProjects(val *[]*string)
	BoundProjectsInput() *[]*string
	BoundRegions() *[]*string
	SetBoundRegions(val *[]*string)
	BoundRegionsInput() *[]*string
	BoundServiceAccounts() *[]*string
	SetBoundServiceAccounts(val *[]*string)
	BoundServiceAccountsInput() *[]*string
	BoundZones() *[]*string
	SetBoundZones(val *[]*string)
	BoundZonesInput() *[]*string
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
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxJwtExp() *string
	SetMaxJwtExp(val *string)
	MaxJwtExpInput() *string
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
	Role() *string
	SetRole(val *string)
	RoleId() *string
	SetRoleId(val *string)
	RoleIdInput() *string
	RoleInput() *string
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
	ResetAddGroupAliases()
	ResetAliasMetadata()
	ResetAllowGceInference()
	ResetBackend()
	ResetBoundInstanceGroups()
	ResetBoundLabels()
	ResetBoundProjects()
	ResetBoundRegions()
	ResetBoundServiceAccounts()
	ResetBoundZones()
	ResetId()
	ResetMaxJwtExp()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleId()
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

// The jsii proxy struct for GcpAuthBackendRole
type jsiiProxy_GcpAuthBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GcpAuthBackendRole) AddGroupAliases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addGroupAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) AddGroupAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addGroupAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) AllowGceInference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGceInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) AllowGceInferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGceInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundInstanceGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundInstanceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundInstanceGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundInstanceGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundLabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) BoundZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) MaxJwtExp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxJwtExp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) MaxJwtExpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxJwtExpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) RoleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendRole) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_auth_backend_role vault_gcp_auth_backend_role} Resource.
func NewGcpAuthBackendRole(scope constructs.Construct, id *string, config *GcpAuthBackendRoleConfig) GcpAuthBackendRole {
	_init_.Initialize()

	if err := validateNewGcpAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GcpAuthBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/gcp_auth_backend_role vault_gcp_auth_backend_role} Resource.
func NewGcpAuthBackendRole_Override(g GcpAuthBackendRole, scope constructs.Construct, id *string, config *GcpAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetAddGroupAliases(val interface{}) {
	if err := j.validateSetAddGroupAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addGroupAliases",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetAllowGceInference(val interface{}) {
	if err := j.validateSetAllowGceInferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGceInference",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundInstanceGroups(val *[]*string) {
	if err := j.validateSetBoundInstanceGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundInstanceGroups",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundLabels(val *[]*string) {
	if err := j.validateSetBoundLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundLabels",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundProjects(val *[]*string) {
	if err := j.validateSetBoundProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundProjects",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundRegions(val *[]*string) {
	if err := j.validateSetBoundRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundRegions",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundServiceAccounts(val *[]*string) {
	if err := j.validateSetBoundServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetBoundZones(val *[]*string) {
	if err := j.validateSetBoundZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundZones",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetMaxJwtExp(val *string) {
	if err := j.validateSetMaxJwtExpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxJwtExp",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetRoleId(val *string) {
	if err := j.validateSetRoleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleId",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendRole)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a GcpAuthBackendRole resource upon running "cdktn plan <stack-name>".
func GcpAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGcpAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
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
func GcpAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGcpAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GcpAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGcpAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GcpAuthBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGcpAuthBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GcpAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.gcpAuthBackendRole.GcpAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetAddGroupAliases() {
	_jsii_.InvokeVoid(
		g,
		"resetAddGroupAliases",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetAllowGceInference() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowGceInference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		g,
		"resetBackend",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundInstanceGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundInstanceGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundProjects() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundProjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetBoundZones() {
	_jsii_.InvokeVoid(
		g,
		"resetBoundZones",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetMaxJwtExp() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxJwtExp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetRoleId() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

