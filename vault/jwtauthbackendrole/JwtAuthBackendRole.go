// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package jwtauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jwtauthbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/jwt_auth_backend_role vault_jwt_auth_backend_role}.
type JwtAuthBackendRole interface {
	cdktn.TerraformResource
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AllowedRedirectUris() *[]*string
	SetAllowedRedirectUris(val *[]*string)
	AllowedRedirectUrisInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BoundAudiences() *[]*string
	SetBoundAudiences(val *[]*string)
	BoundAudiencesInput() *[]*string
	BoundClaims() *map[string]*string
	SetBoundClaims(val *map[string]*string)
	BoundClaimsInput() *map[string]*string
	BoundClaimsType() *string
	SetBoundClaimsType(val *string)
	BoundClaimsTypeInput() *string
	BoundSubject() *string
	SetBoundSubject(val *string)
	BoundSubjectInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClaimMappings() *map[string]*string
	SetClaimMappings(val *map[string]*string)
	ClaimMappingsInput() *map[string]*string
	ClockSkewLeeway() *float64
	SetClockSkewLeeway(val *float64)
	ClockSkewLeewayInput() *float64
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
	DisableBoundClaimsParsing() interface{}
	SetDisableBoundClaimsParsing(val interface{})
	DisableBoundClaimsParsingInput() interface{}
	ExpirationLeeway() *float64
	SetExpirationLeeway(val *float64)
	ExpirationLeewayInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsClaim() *string
	SetGroupsClaim(val *string)
	GroupsClaimInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NotBeforeLeeway() *float64
	SetNotBeforeLeeway(val *float64)
	NotBeforeLeewayInput() *float64
	OidcScopes() *[]*string
	SetOidcScopes(val *[]*string)
	OidcScopesInput() *[]*string
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
	RoleType() *string
	SetRoleType(val *string)
	RoleTypeInput() *string
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
	UserClaim() *string
	SetUserClaim(val *string)
	UserClaimInput() *string
	UserClaimJsonPointer() interface{}
	SetUserClaimJsonPointer(val interface{})
	UserClaimJsonPointerInput() interface{}
	VerboseOidcLogging() interface{}
	SetVerboseOidcLogging(val interface{})
	VerboseOidcLoggingInput() interface{}
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
	ResetAllowedRedirectUris()
	ResetBackend()
	ResetBoundAudiences()
	ResetBoundClaims()
	ResetBoundClaimsType()
	ResetBoundSubject()
	ResetClaimMappings()
	ResetClockSkewLeeway()
	ResetDisableBoundClaimsParsing()
	ResetExpirationLeeway()
	ResetGroupsClaim()
	ResetId()
	ResetMaxAge()
	ResetNamespace()
	ResetNotBeforeLeeway()
	ResetOidcScopes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleType()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
	ResetUserClaimJsonPointer()
	ResetVerboseOidcLogging()
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

// The jsii proxy struct for JwtAuthBackendRole
type jsiiProxy_JwtAuthBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_JwtAuthBackendRole) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) AllowedRedirectUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) AllowedRedirectUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRedirectUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundClaims() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"boundClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundClaimsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"boundClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundClaimsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundClaimsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundClaimsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundClaimsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) BoundSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ClaimMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ClaimMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ClockSkewLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockSkewLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ClockSkewLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockSkewLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) DisableBoundClaimsParsing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBoundClaimsParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) DisableBoundClaimsParsingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableBoundClaimsParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ExpirationLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ExpirationLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GroupsClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GroupsClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) NotBeforeLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notBeforeLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) NotBeforeLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"notBeforeLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) OidcScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) OidcScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) RoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) RoleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) UserClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) UserClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) UserClaimJsonPointer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userClaimJsonPointer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) UserClaimJsonPointerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userClaimJsonPointerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) VerboseOidcLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verboseOidcLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) VerboseOidcLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verboseOidcLoggingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/jwt_auth_backend_role vault_jwt_auth_backend_role} Resource.
func NewJwtAuthBackendRole(scope constructs.Construct, id *string, config *JwtAuthBackendRoleConfig) JwtAuthBackendRole {
	_init_.Initialize()

	if err := validateNewJwtAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_JwtAuthBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/jwt_auth_backend_role vault_jwt_auth_backend_role} Resource.
func NewJwtAuthBackendRole_Override(j JwtAuthBackendRole, scope constructs.Construct, id *string, config *JwtAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetAllowedRedirectUris(val *[]*string) {
	if err := j.validateSetAllowedRedirectUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRedirectUris",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetBoundAudiences(val *[]*string) {
	if err := j.validateSetBoundAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundAudiences",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetBoundClaims(val *map[string]*string) {
	if err := j.validateSetBoundClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundClaims",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetBoundClaimsType(val *string) {
	if err := j.validateSetBoundClaimsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundClaimsType",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetBoundSubject(val *string) {
	if err := j.validateSetBoundSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundSubject",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetClaimMappings(val *map[string]*string) {
	if err := j.validateSetClaimMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimMappings",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetClockSkewLeeway(val *float64) {
	if err := j.validateSetClockSkewLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clockSkewLeeway",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetDisableBoundClaimsParsing(val interface{}) {
	if err := j.validateSetDisableBoundClaimsParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableBoundClaimsParsing",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetExpirationLeeway(val *float64) {
	if err := j.validateSetExpirationLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationLeeway",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetGroupsClaim(val *string) {
	if err := j.validateSetGroupsClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsClaim",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetNotBeforeLeeway(val *float64) {
	if err := j.validateSetNotBeforeLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notBeforeLeeway",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetOidcScopes(val *[]*string) {
	if err := j.validateSetOidcScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcScopes",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetRoleType(val *string) {
	if err := j.validateSetRoleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleType",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetUserClaim(val *string) {
	if err := j.validateSetUserClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userClaim",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetUserClaimJsonPointer(val interface{}) {
	if err := j.validateSetUserClaimJsonPointerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userClaimJsonPointer",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendRole)SetVerboseOidcLogging(val interface{}) {
	if err := j.validateSetVerboseOidcLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verboseOidcLogging",
		val,
	)
}

// Generates CDKTN code for importing a JwtAuthBackendRole resource upon running "cdktn plan <stack-name>".
func JwtAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateJwtAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
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
func JwtAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JwtAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JwtAuthBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func JwtAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.jwtAuthBackendRole.JwtAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) AddMoveTarget(moveTarget *string) {
	if err := j.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := j.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) MoveFromId(id *string) {
	if err := j.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveFromId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := j.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) MoveToId(id *string) {
	if err := j.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"moveToId",
		[]interface{}{id},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		j,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetAllowedRedirectUris() {
	_jsii_.InvokeVoid(
		j,
		"resetAllowedRedirectUris",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		j,
		"resetBackend",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetBoundAudiences() {
	_jsii_.InvokeVoid(
		j,
		"resetBoundAudiences",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetBoundClaims() {
	_jsii_.InvokeVoid(
		j,
		"resetBoundClaims",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetBoundClaimsType() {
	_jsii_.InvokeVoid(
		j,
		"resetBoundClaimsType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetBoundSubject() {
	_jsii_.InvokeVoid(
		j,
		"resetBoundSubject",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetClaimMappings() {
	_jsii_.InvokeVoid(
		j,
		"resetClaimMappings",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetClockSkewLeeway() {
	_jsii_.InvokeVoid(
		j,
		"resetClockSkewLeeway",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetDisableBoundClaimsParsing() {
	_jsii_.InvokeVoid(
		j,
		"resetDisableBoundClaimsParsing",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetExpirationLeeway() {
	_jsii_.InvokeVoid(
		j,
		"resetExpirationLeeway",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		j,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		j,
		"resetId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetMaxAge() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		j,
		"resetNamespace",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetNotBeforeLeeway() {
	_jsii_.InvokeVoid(
		j,
		"resetNotBeforeLeeway",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetOidcScopes() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcScopes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetRoleType() {
	_jsii_.InvokeVoid(
		j,
		"resetRoleType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetUserClaimJsonPointer() {
	_jsii_.InvokeVoid(
		j,
		"resetUserClaimJsonPointer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) ResetVerboseOidcLogging() {
	_jsii_.InvokeVoid(
		j,
		"resetVerboseOidcLogging",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		j,
		"with",
		args,
		&returns,
	)

	return returns
}

