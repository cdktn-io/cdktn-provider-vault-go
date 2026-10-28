// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultazureaccesscredentials

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/ephemeralvaultazureaccesscredentials/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/azure_access_credentials vault_azure_access_credentials}.
type EphemeralVaultAzureAccessCredentials interface {
	cdktn.TerraformEphemeralResource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientId() *string
	ClientSecret() *string
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
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	LeaseDuration() *float64
	LeaseId() *string
	LeaseRenewable() cdktn.IResolvable
	LeaseStartTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	MaxCredValidationSeconds() *float64
	SetMaxCredValidationSeconds(val *float64)
	MaxCredValidationSecondsInput() *float64
	Metadata() cdktn.StringMap
	MountId() *string
	SetMountId(val *string)
	MountIdInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NumSecondsBetweenTests() *float64
	SetNumSecondsBetweenTests(val *float64)
	NumSecondsBetweenTestsInput() *float64
	NumSequentialSuccesses() *float64
	SetNumSequentialSuccesses(val *float64)
	NumSequentialSuccessesInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RequestMetadata() *map[string]*string
	SetRequestMetadata(val *map[string]*string)
	RequestMetadataInput() *map[string]*string
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidateCreds() interface{}
	SetValidateCreds(val interface{})
	ValidateCredsInput() interface{}
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
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetEnvironment()
	ResetMaxCredValidationSeconds()
	ResetMountId()
	ResetNamespace()
	ResetNumSecondsBetweenTests()
	ResetNumSequentialSuccesses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestMetadata()
	ResetSubscriptionId()
	ResetTenantId()
	ResetValidateCreds()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this ephemeral resource to the terraform JSON output.
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

// The jsii proxy struct for EphemeralVaultAzureAccessCredentials
type jsiiProxy_EphemeralVaultAzureAccessCredentials struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) LeaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) LeaseRenewable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"leaseRenewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) LeaseStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) MaxCredValidationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCredValidationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) MaxCredValidationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCredValidationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) MountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) MountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) NumSecondsBetweenTests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSecondsBetweenTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) NumSecondsBetweenTestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSecondsBetweenTestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) NumSequentialSuccesses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSequentialSuccesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) NumSequentialSuccessesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSequentialSuccessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) RequestMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) RequestMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ValidateCreds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCreds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials) ValidateCredsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/azure_access_credentials vault_azure_access_credentials} Ephemeral Resource.
func NewEphemeralVaultAzureAccessCredentials(scope constructs.Construct, id *string, config *EphemeralVaultAzureAccessCredentialsConfig) EphemeralVaultAzureAccessCredentials {
	_init_.Initialize()

	if err := validateNewEphemeralVaultAzureAccessCredentialsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralVaultAzureAccessCredentials{}

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/azure_access_credentials vault_azure_access_credentials} Ephemeral Resource.
func NewEphemeralVaultAzureAccessCredentials_Override(e EphemeralVaultAzureAccessCredentials, scope constructs.Construct, id *string, config *EphemeralVaultAzureAccessCredentialsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetMaxCredValidationSeconds(val *float64) {
	if err := j.validateSetMaxCredValidationSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCredValidationSeconds",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetMountId(val *string) {
	if err := j.validateSetMountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountId",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetNumSecondsBetweenTests(val *float64) {
	if err := j.validateSetNumSecondsBetweenTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numSecondsBetweenTests",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetNumSequentialSuccesses(val *float64) {
	if err := j.validateSetNumSequentialSuccessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numSequentialSuccesses",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetRequestMetadata(val *map[string]*string) {
	if err := j.validateSetRequestMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMetadata",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultAzureAccessCredentials)SetValidateCreds(val interface{}) {
	if err := j.validateSetValidateCredsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateCreds",
		val,
	)
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
func EphemeralVaultAzureAccessCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultAzureAccessCredentials_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultAzureAccessCredentials_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultAzureAccessCredentials_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultAzureAccessCredentials_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultAzureAccessCredentials_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralVaultAzureAccessCredentials_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.ephemeralVaultAzureAccessCredentials.EphemeralVaultAzureAccessCredentials",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetMaxCredValidationSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxCredValidationSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetMountId() {
	_jsii_.InvokeVoid(
		e,
		"resetMountId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetNamespace() {
	_jsii_.InvokeVoid(
		e,
		"resetNamespace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetNumSecondsBetweenTests() {
	_jsii_.InvokeVoid(
		e,
		"resetNumSecondsBetweenTests",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetNumSequentialSuccesses() {
	_jsii_.InvokeVoid(
		e,
		"resetNumSequentialSuccesses",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetRequestMetadata() {
	_jsii_.InvokeVoid(
		e,
		"resetRequestMetadata",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ResetValidateCreds() {
	_jsii_.InvokeVoid(
		e,
		"resetValidateCreds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultAzureAccessCredentials) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

