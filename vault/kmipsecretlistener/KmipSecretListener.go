// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kmipsecretlistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/kmipsecretlistener/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kmip_secret_listener vault_kmip_secret_listener}.
type KmipSecretListener interface {
	cdktn.TerraformResource
	AdditionalClientCas() *[]*string
	SetAdditionalClientCas(val *[]*string)
	AdditionalClientCasInput() *[]*string
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	AlsoUseLegacyCa() interface{}
	SetAlsoUseLegacyCa(val interface{})
	AlsoUseLegacyCaInput() interface{}
	Ca() *string
	SetCa(val *string)
	CaInput() *string
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ServerHostnames() *[]*string
	SetServerHostnames(val *[]*string)
	ServerHostnamesInput() *[]*string
	ServerIps() *[]*string
	SetServerIps(val *[]*string)
	ServerIpsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsCipherSuites() *string
	SetTlsCipherSuites(val *string)
	TlsCipherSuitesInput() *string
	TlsMaxVersion() *string
	SetTlsMaxVersion(val *string)
	TlsMaxVersionInput() *string
	TlsMinVersion() *string
	SetTlsMinVersion(val *string)
	TlsMinVersionInput() *string
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
	ResetAdditionalClientCas()
	ResetAlsoUseLegacyCa()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerHostnames()
	ResetServerIps()
	ResetTlsCipherSuites()
	ResetTlsMaxVersion()
	ResetTlsMinVersion()
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

// The jsii proxy struct for KmipSecretListener
type jsiiProxy_KmipSecretListener struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KmipSecretListener) AdditionalClientCas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalClientCas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) AdditionalClientCasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalClientCasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) AlsoUseLegacyCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alsoUseLegacyCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) AlsoUseLegacyCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alsoUseLegacyCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Ca() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ca",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) CaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ServerHostnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ServerHostnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ServerIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) ServerIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsCipherSuites() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCipherSuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsCipherSuitesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCipherSuitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsMaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretListener) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kmip_secret_listener vault_kmip_secret_listener} Resource.
func NewKmipSecretListener(scope constructs.Construct, id *string, config *KmipSecretListenerConfig) KmipSecretListener {
	_init_.Initialize()

	if err := validateNewKmipSecretListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmipSecretListener{}

	_jsii_.Create(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/kmip_secret_listener vault_kmip_secret_listener} Resource.
func NewKmipSecretListener_Override(k KmipSecretListener, scope constructs.Construct, id *string, config *KmipSecretListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetAdditionalClientCas(val *[]*string) {
	if err := j.validateSetAdditionalClientCasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalClientCas",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetAlsoUseLegacyCa(val interface{}) {
	if err := j.validateSetAlsoUseLegacyCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alsoUseLegacyCa",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetCa(val *string) {
	if err := j.validateSetCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ca",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetServerHostnames(val *[]*string) {
	if err := j.validateSetServerHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverHostnames",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetServerIps(val *[]*string) {
	if err := j.validateSetServerIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverIps",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetTlsCipherSuites(val *string) {
	if err := j.validateSetTlsCipherSuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCipherSuites",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetTlsMaxVersion(val *string) {
	if err := j.validateSetTlsMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMaxVersion",
		val,
	)
}

func (j *jsiiProxy_KmipSecretListener)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
		val,
	)
}

// Generates CDKTN code for importing a KmipSecretListener resource upon running "cdktn plan <stack-name>".
func KmipSecretListener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKmipSecretListener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
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
func KmipSecretListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretListener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretListener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretListener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretListener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KmipSecretListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.kmipSecretListener.KmipSecretListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmipSecretListener) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KmipSecretListener) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KmipSecretListener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KmipSecretListener) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmipSecretListener) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KmipSecretListener) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmipSecretListener) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetAdditionalClientCas() {
	_jsii_.InvokeVoid(
		k,
		"resetAdditionalClientCas",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetAlsoUseLegacyCa() {
	_jsii_.InvokeVoid(
		k,
		"resetAlsoUseLegacyCa",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetServerHostnames() {
	_jsii_.InvokeVoid(
		k,
		"resetServerHostnames",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetServerIps() {
	_jsii_.InvokeVoid(
		k,
		"resetServerIps",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetTlsCipherSuites() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsCipherSuites",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetTlsMaxVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsMaxVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretListener) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

