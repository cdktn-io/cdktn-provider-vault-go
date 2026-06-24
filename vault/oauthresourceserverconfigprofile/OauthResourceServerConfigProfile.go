// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthresourceserverconfigprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/oauthresourceserverconfigprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/oauth_resource_server_config_profile vault_oauth_resource_server_config_profile}.
type OauthResourceServerConfigProfile interface {
	cdktn.TerraformResource
	Audiences() *[]*string
	SetAudiences(val *[]*string)
	AudiencesInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IssuerId() *string
	SetIssuerId(val *string)
	IssuerIdInput() *string
	JwksCaPem() *string
	SetJwksCaPem(val *string)
	JwksCaPemInput() *string
	JwksUri() *string
	SetJwksUri(val *string)
	JwksUriInput() *string
	JwtType() *string
	SetJwtType(val *string)
	JwtTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NoDefaultPolicy() interface{}
	SetNoDefaultPolicy(val interface{})
	NoDefaultPolicyInput() interface{}
	OptionalAuthorizationDetails() interface{}
	SetOptionalAuthorizationDetails(val interface{})
	OptionalAuthorizationDetailsInput() interface{}
	ProfileName() *string
	SetProfileName(val *string)
	ProfileNameInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicKeys() OauthResourceServerConfigProfilePublicKeysList
	PublicKeysInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SupportedAlgorithms() *[]*string
	SetSupportedAlgorithms(val *[]*string)
	SupportedAlgorithmsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseJwks() interface{}
	SetUseJwks(val interface{})
	UseJwksInput() interface{}
	UserClaim() *string
	SetUserClaim(val *string)
	UserClaimInput() *string
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
	PutPublicKeys(value interface{})
	ResetAudiences()
	ResetClockSkewLeeway()
	ResetEnabled()
	ResetJwksCaPem()
	ResetJwksUri()
	ResetJwtType()
	ResetNamespace()
	ResetNoDefaultPolicy()
	ResetOptionalAuthorizationDetails()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicKeys()
	ResetSupportedAlgorithms()
	ResetUseJwks()
	ResetUserClaim()
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

// The jsii proxy struct for OauthResourceServerConfigProfile
type jsiiProxy_OauthResourceServerConfigProfile struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Audiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) AudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ClockSkewLeeway() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockSkewLeeway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ClockSkewLeewayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clockSkewLeewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) IssuerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) IssuerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwksCaPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCaPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwksCaPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCaPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwksUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwksUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwtType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) JwtTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) NoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) NoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) OptionalAuthorizationDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionalAuthorizationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) OptionalAuthorizationDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionalAuthorizationDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) ProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) PublicKeys() OauthResourceServerConfigProfilePublicKeysList {
	var returns OauthResourceServerConfigProfilePublicKeysList
	_jsii_.Get(
		j,
		"publicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) PublicKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) SupportedAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) SupportedAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) UseJwks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useJwks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) UseJwksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useJwksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) UserClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthResourceServerConfigProfile) UserClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaimInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/oauth_resource_server_config_profile vault_oauth_resource_server_config_profile} Resource.
func NewOauthResourceServerConfigProfile(scope constructs.Construct, id *string, config *OauthResourceServerConfigProfileConfig) OauthResourceServerConfigProfile {
	_init_.Initialize()

	if err := validateNewOauthResourceServerConfigProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthResourceServerConfigProfile{}

	_jsii_.Create(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/oauth_resource_server_config_profile vault_oauth_resource_server_config_profile} Resource.
func NewOauthResourceServerConfigProfile_Override(o OauthResourceServerConfigProfile, scope constructs.Construct, id *string, config *OauthResourceServerConfigProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetAudiences(val *[]*string) {
	if err := j.validateSetAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audiences",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetClockSkewLeeway(val *float64) {
	if err := j.validateSetClockSkewLeewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clockSkewLeeway",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetIssuerId(val *string) {
	if err := j.validateSetIssuerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerId",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetJwksCaPem(val *string) {
	if err := j.validateSetJwksCaPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksCaPem",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetJwksUri(val *string) {
	if err := j.validateSetJwksUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUri",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetJwtType(val *string) {
	if err := j.validateSetJwtTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtType",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetNoDefaultPolicy(val interface{}) {
	if err := j.validateSetNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetOptionalAuthorizationDetails(val interface{}) {
	if err := j.validateSetOptionalAuthorizationDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionalAuthorizationDetails",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetProfileName(val *string) {
	if err := j.validateSetProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileName",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetSupportedAlgorithms(val *[]*string) {
	if err := j.validateSetSupportedAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedAlgorithms",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetUseJwks(val interface{}) {
	if err := j.validateSetUseJwksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useJwks",
		val,
	)
}

func (j *jsiiProxy_OauthResourceServerConfigProfile)SetUserClaim(val *string) {
	if err := j.validateSetUserClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userClaim",
		val,
	)
}

// Generates CDKTN code for importing a OauthResourceServerConfigProfile resource upon running "cdktn plan <stack-name>".
func OauthResourceServerConfigProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateOauthResourceServerConfigProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
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
func OauthResourceServerConfigProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthResourceServerConfigProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthResourceServerConfigProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthResourceServerConfigProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OauthResourceServerConfigProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOauthResourceServerConfigProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OauthResourceServerConfigProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_OauthResourceServerConfigProfile) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) PutPublicKeys(value interface{}) {
	if err := o.validatePutPublicKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPublicKeys",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetAudiences() {
	_jsii_.InvokeVoid(
		o,
		"resetAudiences",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetClockSkewLeeway() {
	_jsii_.InvokeVoid(
		o,
		"resetClockSkewLeeway",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetJwksCaPem() {
	_jsii_.InvokeVoid(
		o,
		"resetJwksCaPem",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetJwksUri() {
	_jsii_.InvokeVoid(
		o,
		"resetJwksUri",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetJwtType() {
	_jsii_.InvokeVoid(
		o,
		"resetJwtType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetNamespace() {
	_jsii_.InvokeVoid(
		o,
		"resetNamespace",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetNoDefaultPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetOptionalAuthorizationDetails() {
	_jsii_.InvokeVoid(
		o,
		"resetOptionalAuthorizationDetails",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetPublicKeys() {
	_jsii_.InvokeVoid(
		o,
		"resetPublicKeys",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetSupportedAlgorithms() {
	_jsii_.InvokeVoid(
		o,
		"resetSupportedAlgorithms",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetUseJwks() {
	_jsii_.InvokeVoid(
		o,
		"resetUseJwks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ResetUserClaim() {
	_jsii_.InvokeVoid(
		o,
		"resetUserClaim",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthResourceServerConfigProfile) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

