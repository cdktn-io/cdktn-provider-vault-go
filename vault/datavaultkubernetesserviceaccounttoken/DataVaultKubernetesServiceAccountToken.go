// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultkubernetesserviceaccounttoken

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/datavaultkubernetesserviceaccounttoken/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kubernetes_service_account_token vault_kubernetes_service_account_token}.
type DataVaultKubernetesServiceAccountToken interface {
	cdktn.TerraformDataSource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClusterRoleBinding() interface{}
	SetClusterRoleBinding(val interface{})
	ClusterRoleBindingInput() interface{}
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubernetesNamespace() *string
	SetKubernetesNamespace(val *string)
	KubernetesNamespaceInput() *string
	LeaseDuration() *float64
	LeaseId() *string
	LeaseRenewable() cdktn.IResolvable
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
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	ServiceAccountName() *string
	ServiceAccountNamespace() *string
	ServiceAccountToken() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
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
	ResetClusterRoleBinding()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTtl()
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

// The jsii proxy struct for DataVaultKubernetesServiceAccountToken
type jsiiProxy_DataVaultKubernetesServiceAccountToken struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ClusterRoleBinding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterRoleBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ClusterRoleBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterRoleBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) KubernetesNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) KubernetesNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) LeaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) LeaseRenewable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"leaseRenewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ServiceAccountNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) ServiceAccountToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kubernetes_service_account_token vault_kubernetes_service_account_token} Data Source.
func NewDataVaultKubernetesServiceAccountToken(scope constructs.Construct, id *string, config *DataVaultKubernetesServiceAccountTokenConfig) DataVaultKubernetesServiceAccountToken {
	_init_.Initialize()

	if err := validateNewDataVaultKubernetesServiceAccountTokenParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultKubernetesServiceAccountToken{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/data-sources/kubernetes_service_account_token vault_kubernetes_service_account_token} Data Source.
func NewDataVaultKubernetesServiceAccountToken_Override(d DataVaultKubernetesServiceAccountToken, scope constructs.Construct, id *string, config *DataVaultKubernetesServiceAccountTokenConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetClusterRoleBinding(val interface{}) {
	if err := j.validateSetClusterRoleBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterRoleBinding",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetKubernetesNamespace(val *string) {
	if err := j.validateSetKubernetesNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesNamespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_DataVaultKubernetesServiceAccountToken)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

// Generates CDKTN code for importing a DataVaultKubernetesServiceAccountToken resource upon running "cdktn plan <stack-name>".
func DataVaultKubernetesServiceAccountToken_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataVaultKubernetesServiceAccountToken_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
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
func DataVaultKubernetesServiceAccountToken_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesServiceAccountToken_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultKubernetesServiceAccountToken_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesServiceAccountToken_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataVaultKubernetesServiceAccountToken_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataVaultKubernetesServiceAccountToken_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultKubernetesServiceAccountToken_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.dataVaultKubernetesServiceAccountToken.DataVaultKubernetesServiceAccountToken",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ResetClusterRoleBinding() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterRoleBinding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultKubernetesServiceAccountToken) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

