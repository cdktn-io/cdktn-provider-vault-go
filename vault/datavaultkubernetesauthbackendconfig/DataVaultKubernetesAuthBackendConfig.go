// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultkubernetesauthbackendconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaultkubernetesauthbackendconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config}.
type DataVaultKubernetesAuthBackendConfig interface {
	cdktn.TerraformDataSource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DisableIssValidation() interface{}
	SetDisableIssValidation(val interface{})
	DisableIssValidationInput() interface{}
	DisableLocalCaJwt() interface{}
	SetDisableLocalCaJwt(val interface{})
	DisableLocalCaJwtInput() interface{}
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	KubernetesCaCert() *string
	SetKubernetesCaCert(val *string)
	KubernetesCaCertInput() *string
	KubernetesHost() *string
	SetKubernetesHost(val *string)
	KubernetesHostInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PemKeys() *[]*string
	SetPemKeys(val *[]*string)
	PemKeysInput() *[]*string
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
	UseAnnotationsAsAliasMetadata() interface{}
	SetUseAnnotationsAsAliasMetadata(val interface{})
	UseAnnotationsAsAliasMetadataInput() interface{}
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
	ResetBackend()
	ResetDisableIssValidation()
	ResetDisableLocalCaJwt()
	ResetId()
	ResetIssuer()
	ResetKubernetesCaCert()
	ResetKubernetesHost()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPemKeys()
	ResetUseAnnotationsAsAliasMetadata()
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

// The jsii proxy struct for DataVaultKubernetesAuthBackendConfig
type jsiiProxy_DataVaultKubernetesAuthBackendConfig struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) DisableIssValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIssValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) DisableIssValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIssValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) DisableLocalCaJwt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalCaJwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) DisableLocalCaJwtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalCaJwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) KubernetesCaCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) KubernetesCaCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) KubernetesHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) KubernetesHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) PemKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pemKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) PemKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pemKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) UseAnnotationsAsAliasMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAnnotationsAsAliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig) UseAnnotationsAsAliasMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAnnotationsAsAliasMetadataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config} Data Source.
func NewDataVaultKubernetesAuthBackendConfig(scope constructs.Construct, id *string, config *DataVaultKubernetesAuthBackendConfigConfig) DataVaultKubernetesAuthBackendConfig {
	_init_.Initialize()

	if err := validateNewDataVaultKubernetesAuthBackendConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultKubernetesAuthBackendConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config} Data Source.
func NewDataVaultKubernetesAuthBackendConfig_Override(d DataVaultKubernetesAuthBackendConfig, scope constructs.Construct, id *string, config *DataVaultKubernetesAuthBackendConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetDisableIssValidation(val interface{}) {
	if err := j.validateSetDisableIssValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableIssValidation",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetDisableLocalCaJwt(val interface{}) {
	if err := j.validateSetDisableLocalCaJwtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableLocalCaJwt",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetKubernetesCaCert(val *string) {
	if err := j.validateSetKubernetesCaCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesCaCert",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetKubernetesHost(val *string) {
	if err := j.validateSetKubernetesHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesHost",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetPemKeys(val *[]*string) {
	if err := j.validateSetPemKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemKeys",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesAuthBackendConfig)SetUseAnnotationsAsAliasMetadata(val interface{}) {
	if err := j.validateSetUseAnnotationsAsAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAnnotationsAsAliasMetadata",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultKubernetesAuthBackendConfig resource upon running "cdktn plan <stack-name>".
func DataVaultKubernetesAuthBackendConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultKubernetesAuthBackendConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
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
func DataVaultKubernetesAuthBackendConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesAuthBackendConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultKubernetesAuthBackendConfig_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesAuthBackendConfig_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultKubernetesAuthBackendConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesAuthBackendConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultKubernetesAuthBackendConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultKubernetesAuthBackendConfig.DataVaultKubernetesAuthBackendConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetBackend() {
	_jsii_.InvokeVoid(
		d,
		"resetBackend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetDisableIssValidation() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableIssValidation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetDisableLocalCaJwt() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableLocalCaJwt",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetIssuer() {
	_jsii_.InvokeVoid(
		d,
		"resetIssuer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetKubernetesCaCert() {
	_jsii_.InvokeVoid(
		d,
		"resetKubernetesCaCert",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetKubernetesHost() {
	_jsii_.InvokeVoid(
		d,
		"resetKubernetesHost",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetPemKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetPemKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ResetUseAnnotationsAsAliasMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetUseAnnotationsAsAliasMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesAuthBackendConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

