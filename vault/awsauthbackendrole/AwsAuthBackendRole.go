// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/awsauthbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role vault_aws_auth_backend_role}.
type AwsAuthBackendRole interface {
	cdktn.TerraformResource
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AllowInstanceMigration() interface{}
	SetAllowInstanceMigration(val interface{})
	AllowInstanceMigrationInput() interface{}
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BoundAccountIds() *[]*string
	SetBoundAccountIds(val *[]*string)
	BoundAccountIdsInput() *[]*string
	BoundAmiIds() *[]*string
	SetBoundAmiIds(val *[]*string)
	BoundAmiIdsInput() *[]*string
	BoundEc2InstanceIds() *[]*string
	SetBoundEc2InstanceIds(val *[]*string)
	BoundEc2InstanceIdsInput() *[]*string
	BoundIamInstanceProfileArns() *[]*string
	SetBoundIamInstanceProfileArns(val *[]*string)
	BoundIamInstanceProfileArnsInput() *[]*string
	BoundIamPrincipalArns() *[]*string
	SetBoundIamPrincipalArns(val *[]*string)
	BoundIamPrincipalArnsInput() *[]*string
	BoundIamRoleArns() *[]*string
	SetBoundIamRoleArns(val *[]*string)
	BoundIamRoleArnsInput() *[]*string
	BoundRegions() *[]*string
	SetBoundRegions(val *[]*string)
	BoundRegionsInput() *[]*string
	BoundSubnetIds() *[]*string
	SetBoundSubnetIds(val *[]*string)
	BoundSubnetIdsInput() *[]*string
	BoundVpcIds() *[]*string
	SetBoundVpcIds(val *[]*string)
	BoundVpcIdsInput() *[]*string
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
	DisallowReauthentication() interface{}
	SetDisallowReauthentication(val interface{})
	DisallowReauthenticationInput() interface{}
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
	InferredAwsRegion() *string
	SetInferredAwsRegion(val *string)
	InferredAwsRegionInput() *string
	InferredEntityType() *string
	SetInferredEntityType(val *string)
	InferredEntityTypeInput() *string
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResolveAwsUniqueIds() interface{}
	SetResolveAwsUniqueIds(val interface{})
	ResolveAwsUniqueIdsInput() interface{}
	Role() *string
	SetRole(val *string)
	RoleId() *string
	RoleInput() *string
	RoleTag() *string
	SetRoleTag(val *string)
	RoleTagInput() *string
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
	ResetAliasMetadata()
	ResetAllowInstanceMigration()
	ResetAuthType()
	ResetBackend()
	ResetBoundAccountIds()
	ResetBoundAmiIds()
	ResetBoundEc2InstanceIds()
	ResetBoundIamInstanceProfileArns()
	ResetBoundIamPrincipalArns()
	ResetBoundIamRoleArns()
	ResetBoundRegions()
	ResetBoundSubnetIds()
	ResetBoundVpcIds()
	ResetDisallowReauthentication()
	ResetId()
	ResetInferredAwsRegion()
	ResetInferredEntityType()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResolveAwsUniqueIds()
	ResetRoleTag()
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

// The jsii proxy struct for AwsAuthBackendRole
type jsiiProxy_AwsAuthBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAuthBackendRole) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AllowInstanceMigration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstanceMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AllowInstanceMigrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstanceMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAmiIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAmiIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAmiIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAmiIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundEc2InstanceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundEc2InstanceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundEc2InstanceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundEc2InstanceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamInstanceProfileArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamInstanceProfileArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamInstanceProfileArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamInstanceProfileArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamPrincipalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamPrincipalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamPrincipalArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamPrincipalArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundSubnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundVpcIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundVpcIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundVpcIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundVpcIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DisallowReauthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowReauthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DisallowReauthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowReauthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredAwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredAwsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredAwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredAwsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredEntityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredEntityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredEntityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredEntityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ResolveAwsUniqueIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveAwsUniqueIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ResolveAwsUniqueIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveAwsUniqueIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role vault_aws_auth_backend_role} Resource.
func NewAwsAuthBackendRole(scope constructs.Construct, id *string, config *AwsAuthBackendRoleConfig) AwsAuthBackendRole {
	_init_.Initialize()

	if err := validateNewAwsAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAuthBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_auth_backend_role vault_aws_auth_backend_role} Resource.
func NewAwsAuthBackendRole_Override(a AwsAuthBackendRole, scope constructs.Construct, id *string, config *AwsAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetAllowInstanceMigration(val interface{}) {
	if err := j.validateSetAllowInstanceMigrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInstanceMigration",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundAccountIds(val *[]*string) {
	if err := j.validateSetBoundAccountIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundAmiIds(val *[]*string) {
	if err := j.validateSetBoundAmiIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundAmiIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundEc2InstanceIds(val *[]*string) {
	if err := j.validateSetBoundEc2InstanceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundEc2InstanceIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundIamInstanceProfileArns(val *[]*string) {
	if err := j.validateSetBoundIamInstanceProfileArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundIamInstanceProfileArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundIamPrincipalArns(val *[]*string) {
	if err := j.validateSetBoundIamPrincipalArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundIamPrincipalArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundIamRoleArns(val *[]*string) {
	if err := j.validateSetBoundIamRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundIamRoleArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundRegions(val *[]*string) {
	if err := j.validateSetBoundRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundRegions",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundSubnetIds(val *[]*string) {
	if err := j.validateSetBoundSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundSubnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetBoundVpcIds(val *[]*string) {
	if err := j.validateSetBoundVpcIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundVpcIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetDisallowReauthentication(val interface{}) {
	if err := j.validateSetDisallowReauthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowReauthentication",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetInferredAwsRegion(val *string) {
	if err := j.validateSetInferredAwsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferredAwsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetInferredEntityType(val *string) {
	if err := j.validateSetInferredEntityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferredEntityType",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetResolveAwsUniqueIds(val interface{}) {
	if err := j.validateSetResolveAwsUniqueIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolveAwsUniqueIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetRoleTag(val *string) {
	if err := j.validateSetRoleTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleTag",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

// Generates CDKTN code for importing a AwsAuthBackendRole resource upon running "cdktn plan <stack-name>".
func AwsAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
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
func AwsAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.awsAuthBackendRole.AwsAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAuthBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAuthBackendRole) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetAllowInstanceMigration() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowInstanceMigration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundAmiIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundAmiIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundEc2InstanceIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundEc2InstanceIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamInstanceProfileArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamInstanceProfileArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamPrincipalArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamPrincipalArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamRoleArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamRoleArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundVpcIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundVpcIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetDisallowReauthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetDisallowReauthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetInferredAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetInferredAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetInferredEntityType() {
	_jsii_.InvokeVoid(
		a,
		"resetInferredEntityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetResolveAwsUniqueIds() {
	_jsii_.InvokeVoid(
		a,
		"resetResolveAwsUniqueIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetRoleTag() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

