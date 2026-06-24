// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultsshsecretbackendsign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaultsshsecretbackendsign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/ssh_secret_backend_sign vault_ssh_secret_backend_sign}.
type DataVaultSshSecretBackendSign interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertType() *string
	SetCertType(val *string)
	CertTypeInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CriticalOptions() *map[string]*string
	SetCriticalOptions(val *map[string]*string)
	CriticalOptionsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Extensions() *map[string]*string
	SetExtensions(val *map[string]*string)
	ExtensionsInput() *map[string]*string
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
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
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
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	SerialNumber() *string
	SignedKey() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	ValidPrincipals() *string
	SetValidPrincipals(val *string)
	ValidPrincipalsInput() *string
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
	ResetCertType()
	ResetCriticalOptions()
	ResetExtensions()
	ResetId()
	ResetKeyId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTtl()
	ResetValidPrincipals()
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

// The jsii proxy struct for DataVaultSshSecretBackendSign
type jsiiProxy_DataVaultSshSecretBackendSign struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) CertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) CertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) CriticalOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"criticalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) CriticalOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"criticalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Extensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) ExtensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) SignedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) ValidPrincipals() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign) ValidPrincipalsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validPrincipalsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/ssh_secret_backend_sign vault_ssh_secret_backend_sign} Data Source.
func NewDataVaultSshSecretBackendSign(scope constructs.Construct, id *string, config *DataVaultSshSecretBackendSignConfig) DataVaultSshSecretBackendSign {
	_init_.Initialize()

	if err := validateNewDataVaultSshSecretBackendSignParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultSshSecretBackendSign{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/ssh_secret_backend_sign vault_ssh_secret_backend_sign} Data Source.
func NewDataVaultSshSecretBackendSign_Override(d DataVaultSshSecretBackendSign, scope constructs.Construct, id *string, config *DataVaultSshSecretBackendSignConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetCertType(val *string) {
	if err := j.validateSetCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certType",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetCriticalOptions(val *map[string]*string) {
	if err := j.validateSetCriticalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticalOptions",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetExtensions(val *map[string]*string) {
	if err := j.validateSetExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensions",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetKeyId(val *string) {
	if err := j.validateSetKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetPublicKey(val *string) {
	if err := j.validateSetPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_DataVaultSshSecretBackendSign)SetValidPrincipals(val *string) {
	if err := j.validateSetValidPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validPrincipals",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultSshSecretBackendSign resource upon running "cdktn plan <stack-name>".
func DataVaultSshSecretBackendSign_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultSshSecretBackendSign_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
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
func DataVaultSshSecretBackendSign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultSshSecretBackendSign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultSshSecretBackendSign_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultSshSecretBackendSign_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultSshSecretBackendSign_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultSshSecretBackendSign_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultSshSecretBackendSign_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultSshSecretBackendSign.DataVaultSshSecretBackendSign",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultSshSecretBackendSign) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetCertType() {
	_jsii_.InvokeVoid(
		d,
		"resetCertType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetCriticalOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetCriticalOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetExtensions() {
	_jsii_.InvokeVoid(
		d,
		"resetExtensions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ResetValidPrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetValidPrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultSshSecretBackendSign) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

