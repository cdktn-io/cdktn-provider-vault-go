// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesauthbackendconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/kubernetesauthbackendconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config}.
type KubernetesAuthBackendConfig interface {
	cdktn.TerraformResource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
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
	TokenReviewerJwt() *string
	SetTokenReviewerJwt(val *string)
	TokenReviewerJwtInput() *string
	TokenReviewerJwtWo() *string
	SetTokenReviewerJwtWo(val *string)
	TokenReviewerJwtWoInput() *string
	TokenReviewerJwtWoVersion() *float64
	SetTokenReviewerJwtWoVersion(val *float64)
	TokenReviewerJwtWoVersionInput() *float64
	UseAnnotationsAsAliasMetadata() interface{}
	SetUseAnnotationsAsAliasMetadata(val interface{})
	UseAnnotationsAsAliasMetadataInput() interface{}
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
	ResetBackend()
	ResetDisableIssValidation()
	ResetDisableLocalCaJwt()
	ResetId()
	ResetIssuer()
	ResetKubernetesCaCert()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPemKeys()
	ResetTokenReviewerJwt()
	ResetTokenReviewerJwtWo()
	ResetTokenReviewerJwtWoVersion()
	ResetUseAnnotationsAsAliasMetadata()
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

// The jsii proxy struct for KubernetesAuthBackendConfig
type jsiiProxy_KubernetesAuthBackendConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) DisableIssValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIssValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) DisableIssValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableIssValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) DisableLocalCaJwt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalCaJwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) DisableLocalCaJwtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalCaJwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) KubernetesCaCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) KubernetesCaCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesCaCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) KubernetesHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) KubernetesHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) PemKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pemKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) PemKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pemKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenReviewerJwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenReviewerJwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwtWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenReviewerJwtWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwtWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenReviewerJwtWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwtWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenReviewerJwtWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) TokenReviewerJwtWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenReviewerJwtWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) UseAnnotationsAsAliasMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAnnotationsAsAliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAuthBackendConfig) UseAnnotationsAsAliasMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAnnotationsAsAliasMetadataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config} Resource.
func NewKubernetesAuthBackendConfig(scope constructs.Construct, id *string, config *KubernetesAuthBackendConfigConfig) KubernetesAuthBackendConfig {
	_init_.Initialize()

	if err := validateNewKubernetesAuthBackendConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesAuthBackendConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config vault_kubernetes_auth_backend_config} Resource.
func NewKubernetesAuthBackendConfig_Override(k KubernetesAuthBackendConfig, scope constructs.Construct, id *string, config *KubernetesAuthBackendConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetDisableIssValidation(val interface{}) {
	if err := j.validateSetDisableIssValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableIssValidation",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetDisableLocalCaJwt(val interface{}) {
	if err := j.validateSetDisableLocalCaJwtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableLocalCaJwt",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetKubernetesCaCert(val *string) {
	if err := j.validateSetKubernetesCaCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesCaCert",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetKubernetesHost(val *string) {
	if err := j.validateSetKubernetesHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesHost",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetPemKeys(val *[]*string) {
	if err := j.validateSetPemKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemKeys",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetTokenReviewerJwt(val *string) {
	if err := j.validateSetTokenReviewerJwtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenReviewerJwt",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetTokenReviewerJwtWo(val *string) {
	if err := j.validateSetTokenReviewerJwtWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenReviewerJwtWo",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetTokenReviewerJwtWoVersion(val *float64) {
	if err := j.validateSetTokenReviewerJwtWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenReviewerJwtWoVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesAuthBackendConfig)SetUseAnnotationsAsAliasMetadata(val interface{}) {
	if err := j.validateSetUseAnnotationsAsAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAnnotationsAsAliasMetadata",
		val,
	)
}

// Generates CDKTN code for importing a KubernetesAuthBackendConfig resource upon running "cdktn plan <stack-name>".
func KubernetesAuthBackendConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesAuthBackendConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
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
func KubernetesAuthBackendConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAuthBackendConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesAuthBackendConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAuthBackendConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesAuthBackendConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAuthBackendConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesAuthBackendConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAuthBackendConfig) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetBackend() {
	_jsii_.InvokeVoid(
		k,
		"resetBackend",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetDisableIssValidation() {
	_jsii_.InvokeVoid(
		k,
		"resetDisableIssValidation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetDisableLocalCaJwt() {
	_jsii_.InvokeVoid(
		k,
		"resetDisableLocalCaJwt",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetIssuer() {
	_jsii_.InvokeVoid(
		k,
		"resetIssuer",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetKubernetesCaCert() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesCaCert",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetPemKeys() {
	_jsii_.InvokeVoid(
		k,
		"resetPemKeys",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetTokenReviewerJwt() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenReviewerJwt",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetTokenReviewerJwtWo() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenReviewerJwtWo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetTokenReviewerJwtWoVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenReviewerJwtWoVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ResetUseAnnotationsAsAliasMetadata() {
	_jsii_.InvokeVoid(
		k,
		"resetUseAnnotationsAsAliasMetadata",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAuthBackendConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

