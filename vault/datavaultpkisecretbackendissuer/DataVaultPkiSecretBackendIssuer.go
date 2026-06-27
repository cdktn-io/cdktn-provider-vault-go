// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpkisecretbackendissuer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaultpkisecretbackendissuer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_issuer vault_pki_secret_backend_issuer}.
type DataVaultPkiSecretBackendIssuer interface {
	cdktn.TerraformDataSource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	CaChain() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
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
	DisableCriticalExtensionChecks() interface{}
	SetDisableCriticalExtensionChecks(val interface{})
	DisableCriticalExtensionChecksInput() interface{}
	DisableNameChecks() interface{}
	SetDisableNameChecks(val interface{})
	DisableNameChecksInput() interface{}
	DisableNameConstraintChecks() interface{}
	SetDisableNameConstraintChecks(val interface{})
	DisableNameConstraintChecksInput() interface{}
	DisablePathLengthChecks() interface{}
	SetDisablePathLengthChecks(val interface{})
	DisablePathLengthChecksInput() interface{}
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
	IssuerId() *string
	IssuerName() *string
	IssuerRef() *string
	SetIssuerRef(val *string)
	IssuerRefInput() *string
	KeyId() *string
	LeafNotAfterBehavior() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ManualChain() *[]*string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Usage() *string
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
	ResetDisableCriticalExtensionChecks()
	ResetDisableNameChecks()
	ResetDisableNameConstraintChecks()
	ResetDisablePathLengthChecks()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataVaultPkiSecretBackendIssuer
type jsiiProxy_DataVaultPkiSecretBackendIssuer struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) CaChain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"caChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableCriticalExtensionChecks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCriticalExtensionChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableCriticalExtensionChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCriticalExtensionChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableNameChecks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNameChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableNameChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNameChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableNameConstraintChecks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNameConstraintChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisableNameConstraintChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableNameConstraintChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisablePathLengthChecks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePathLengthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) DisablePathLengthChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePathLengthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) IssuerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) IssuerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) IssuerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) IssuerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) LeafNotAfterBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leafNotAfterBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) ManualChain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"manualChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer) Usage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_issuer vault_pki_secret_backend_issuer} Data Source.
func NewDataVaultPkiSecretBackendIssuer(scope constructs.Construct, id *string, config *DataVaultPkiSecretBackendIssuerConfig) DataVaultPkiSecretBackendIssuer {
	_init_.Initialize()

	if err := validateNewDataVaultPkiSecretBackendIssuerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultPkiSecretBackendIssuer{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/pki_secret_backend_issuer vault_pki_secret_backend_issuer} Data Source.
func NewDataVaultPkiSecretBackendIssuer_Override(d DataVaultPkiSecretBackendIssuer, scope constructs.Construct, id *string, config *DataVaultPkiSecretBackendIssuerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetDisableCriticalExtensionChecks(val interface{}) {
	if err := j.validateSetDisableCriticalExtensionChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCriticalExtensionChecks",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetDisableNameChecks(val interface{}) {
	if err := j.validateSetDisableNameChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableNameChecks",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetDisableNameConstraintChecks(val interface{}) {
	if err := j.validateSetDisableNameConstraintChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableNameConstraintChecks",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetDisablePathLengthChecks(val interface{}) {
	if err := j.validateSetDisablePathLengthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePathLengthChecks",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetIssuerRef(val *string) {
	if err := j.validateSetIssuerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerRef",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultPkiSecretBackendIssuer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultPkiSecretBackendIssuer resource upon running "cdktn plan <stack-name>".
func DataVaultPkiSecretBackendIssuer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultPkiSecretBackendIssuer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
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
func DataVaultPkiSecretBackendIssuer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultPkiSecretBackendIssuer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultPkiSecretBackendIssuer_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultPkiSecretBackendIssuer_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultPkiSecretBackendIssuer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultPkiSecretBackendIssuer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultPkiSecretBackendIssuer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultPkiSecretBackendIssuer.DataVaultPkiSecretBackendIssuer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetDisableCriticalExtensionChecks() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableCriticalExtensionChecks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetDisableNameChecks() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableNameChecks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetDisableNameConstraintChecks() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableNameConstraintChecks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetDisablePathLengthChecks() {
	_jsii_.InvokeVoid(
		d,
		"resetDisablePathLengthChecks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPkiSecretBackendIssuer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

