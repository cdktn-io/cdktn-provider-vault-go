// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spiffeauthbackendconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/spiffeauthbackendconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_auth_backend_config vault_spiffe_auth_backend_config}.
type SpiffeAuthBackendConfig interface {
	cdktn.TerraformResource
	Audience() *[]*string
	SetAudience(val *[]*string)
	AudienceInput() *[]*string
	Bundle() *string
	SetBundle(val *string)
	BundleInput() *string
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
	DeferBundleFetch() interface{}
	SetDeferBundleFetch(val interface{})
	DeferBundleFetchInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointRootCaTruststorePem() *string
	SetEndpointRootCaTruststorePem(val *string)
	EndpointRootCaTruststorePemInput() *string
	EndpointSpiffeId() *string
	SetEndpointSpiffeId(val *string)
	EndpointSpiffeIdInput() *string
	EndpointUrl() *string
	SetEndpointUrl(val *string)
	EndpointUrlInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Mount() *string
	SetMount(val *string)
	MountInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
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
	TrustDomain() *string
	SetTrustDomain(val *string)
	TrustDomainInput() *string
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
	ResetAudience()
	ResetBundle()
	ResetDeferBundleFetch()
	ResetEndpointRootCaTruststorePem()
	ResetEndpointSpiffeId()
	ResetEndpointUrl()
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

// The jsii proxy struct for SpiffeAuthBackendConfig
type jsiiProxy_SpiffeAuthBackendConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Audience() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) AudienceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Bundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) BundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) DeferBundleFetch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferBundleFetch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) DeferBundleFetchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferBundleFetchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointRootCaTruststorePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointRootCaTruststorePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointRootCaTruststorePemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointRootCaTruststorePemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointSpiffeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointSpiffeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointSpiffeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointSpiffeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) EndpointUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) TrustDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiffeAuthBackendConfig) TrustDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDomainInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_auth_backend_config vault_spiffe_auth_backend_config} Resource.
func NewSpiffeAuthBackendConfig(scope constructs.Construct, id *string, config *SpiffeAuthBackendConfigConfig) SpiffeAuthBackendConfig {
	_init_.Initialize()

	if err := validateNewSpiffeAuthBackendConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiffeAuthBackendConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_auth_backend_config vault_spiffe_auth_backend_config} Resource.
func NewSpiffeAuthBackendConfig_Override(s SpiffeAuthBackendConfig, scope constructs.Construct, id *string, config *SpiffeAuthBackendConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetAudience(val *[]*string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetBundle(val *string) {
	if err := j.validateSetBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bundle",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetDeferBundleFetch(val interface{}) {
	if err := j.validateSetDeferBundleFetchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferBundleFetch",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetEndpointRootCaTruststorePem(val *string) {
	if err := j.validateSetEndpointRootCaTruststorePemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointRootCaTruststorePem",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetEndpointSpiffeId(val *string) {
	if err := j.validateSetEndpointSpiffeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointSpiffeId",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetEndpointUrl(val *string) {
	if err := j.validateSetEndpointUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointUrl",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SpiffeAuthBackendConfig)SetTrustDomain(val *string) {
	if err := j.validateSetTrustDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDomain",
		val,
	)
}

// Generates CDKTN code for importing a SpiffeAuthBackendConfig resource upon running "cdktn plan <stack-name>".
func SpiffeAuthBackendConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSpiffeAuthBackendConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
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
func SpiffeAuthBackendConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpiffeAuthBackendConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpiffeAuthBackendConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpiffeAuthBackendConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SpiffeAuthBackendConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpiffeAuthBackendConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SpiffeAuthBackendConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.spiffeAuthBackendConfig.SpiffeAuthBackendConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetAudience() {
	_jsii_.InvokeVoid(
		s,
		"resetAudience",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetBundle() {
	_jsii_.InvokeVoid(
		s,
		"resetBundle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetDeferBundleFetch() {
	_jsii_.InvokeVoid(
		s,
		"resetDeferBundleFetch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetEndpointRootCaTruststorePem() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointRootCaTruststorePem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetEndpointSpiffeId() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointSpiffeId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetEndpointUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiffeAuthBackendConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

