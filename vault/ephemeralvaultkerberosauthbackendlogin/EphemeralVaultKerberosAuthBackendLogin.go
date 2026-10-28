// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultkerberosauthbackendlogin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v19/ephemeralvaultkerberosauthbackendlogin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/kerberos_auth_backend_login vault_kerberos_auth_backend_login}.
type EphemeralVaultKerberosAuthBackendLogin interface {
	cdktn.TerraformEphemeralResource
	Accessor() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientToken() *string
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
	DisableFastNegotiation() interface{}
	SetDisableFastNegotiation(val interface{})
	DisableFastNegotiationInput() interface{}
	EntityId() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IdentityPolicies() *[]*string
	KeytabPath() *string
	SetKeytabPath(val *string)
	KeytabPathInput() *string
	Krb5ConfPath() *string
	SetKrb5ConfPath(val *string)
	Krb5ConfPathInput() *string
	LeaseDuration() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	Metadata() cdktn.StringMap
	Mount() *string
	SetMount(val *string)
	MountId() *string
	SetMountId(val *string)
	MountIdInput() *string
	MountInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Orphan() cdktn.IResolvable
	Policies() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
	RemoveInstanceName() interface{}
	SetRemoveInstanceName(val interface{})
	RemoveInstanceNameInput() interface{}
	Renewable() cdktn.IResolvable
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenPolicies() *[]*string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetDisableFastNegotiation()
	ResetMountId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoveInstanceName()
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

// The jsii proxy struct for EphemeralVaultKerberosAuthBackendLogin
type jsiiProxy_EphemeralVaultKerberosAuthBackendLogin struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) DisableFastNegotiation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFastNegotiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) DisableFastNegotiationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFastNegotiationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) IdentityPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) KeytabPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keytabPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) KeytabPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keytabPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Krb5ConfPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Krb5ConfPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) MountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) MountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Orphan() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"orphan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) RemoveInstanceName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) RemoveInstanceNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Renewable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"renewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/kerberos_auth_backend_login vault_kerberos_auth_backend_login} Ephemeral Resource.
func NewEphemeralVaultKerberosAuthBackendLogin(scope constructs.Construct, id *string, config *EphemeralVaultKerberosAuthBackendLoginConfig) EphemeralVaultKerberosAuthBackendLogin {
	_init_.Initialize()

	if err := validateNewEphemeralVaultKerberosAuthBackendLoginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralVaultKerberosAuthBackendLogin{}

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/kerberos_auth_backend_login vault_kerberos_auth_backend_login} Ephemeral Resource.
func NewEphemeralVaultKerberosAuthBackendLogin_Override(e EphemeralVaultKerberosAuthBackendLogin, scope constructs.Construct, id *string, config *EphemeralVaultKerberosAuthBackendLoginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetDisableFastNegotiation(val interface{}) {
	if err := j.validateSetDisableFastNegotiationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableFastNegotiation",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetKeytabPath(val *string) {
	if err := j.validateSetKeytabPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keytabPath",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetKrb5ConfPath(val *string) {
	if err := j.validateSetKrb5ConfPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5ConfPath",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetMountId(val *string) {
	if err := j.validateSetMountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mountId",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetRemoveInstanceName(val interface{}) {
	if err := j.validateSetRemoveInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInstanceName",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
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
func EphemeralVaultKerberosAuthBackendLogin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultKerberosAuthBackendLogin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultKerberosAuthBackendLogin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultKerberosAuthBackendLogin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralVaultKerberosAuthBackendLogin_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralVaultKerberosAuthBackendLogin_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralVaultKerberosAuthBackendLogin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.ephemeralVaultKerberosAuthBackendLogin.EphemeralVaultKerberosAuthBackendLogin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ResetDisableFastNegotiation() {
	_jsii_.InvokeVoid(
		e,
		"resetDisableFastNegotiation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ResetMountId() {
	_jsii_.InvokeVoid(
		e,
		"resetMountId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ResetNamespace() {
	_jsii_.InvokeVoid(
		e,
		"resetNamespace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ResetRemoveInstanceName() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoveInstanceName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralVaultKerberosAuthBackendLogin) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

