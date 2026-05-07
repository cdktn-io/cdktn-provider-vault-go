// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azureauthbackendconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/azureauthbackendconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config vault_azure_auth_backend_config}.
type AzureAuthBackendConfig interface {
	cdktn.TerraformResource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	ClientSecretWo() *string
	SetClientSecretWo(val *string)
	ClientSecretWoInput() *string
	ClientSecretWoVersion() *float64
	SetClientSecretWoVersion(val *float64)
	ClientSecretWoVersionInput() *float64
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
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
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
	Id() *string
	SetId(val *string)
	IdentityTokenAudience() *string
	SetIdentityTokenAudience(val *string)
	IdentityTokenAudienceInput() *string
	IdentityTokenTtl() *float64
	SetIdentityTokenTtl(val *float64)
	IdentityTokenTtlInput() *float64
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MaxRetryDelay() *float64
	SetMaxRetryDelay(val *float64)
	MaxRetryDelayInput() *float64
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	RetryDelay() *float64
	SetRetryDelay(val *float64)
	RetryDelayInput() *float64
	RotationPeriod() *float64
	SetRotationPeriod(val *float64)
	RotationPeriodInput() *float64
	RotationSchedule() *string
	SetRotationSchedule(val *string)
	RotationScheduleInput() *string
	RotationWindow() *float64
	SetRotationWindow(val *float64)
	RotationWindowInput() *float64
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetClientId()
	ResetClientSecret()
	ResetClientSecretWo()
	ResetClientSecretWoVersion()
	ResetDisableAutomatedRotation()
	ResetEnvironment()
	ResetId()
	ResetIdentityTokenAudience()
	ResetIdentityTokenTtl()
	ResetMaxRetries()
	ResetMaxRetryDelay()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetryDelay()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
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

// The jsii proxy struct for AzureAuthBackendConfig
type jsiiProxy_AzureAuthBackendConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AzureAuthBackendConfig) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecretWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSecretWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ClientSecretWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientSecretWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) IdentityTokenAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) IdentityTokenAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) IdentityTokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) IdentityTokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) MaxRetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) MaxRetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RetryDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RetryDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureAuthBackendConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config vault_azure_auth_backend_config} Resource.
func NewAzureAuthBackendConfig(scope constructs.Construct, id *string, config *AzureAuthBackendConfigConfig) AzureAuthBackendConfig {
	_init_.Initialize()

	if err := validateNewAzureAuthBackendConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzureAuthBackendConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_auth_backend_config vault_azure_auth_backend_config} Resource.
func NewAzureAuthBackendConfig_Override(a AzureAuthBackendConfig, scope constructs.Construct, id *string, config *AzureAuthBackendConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetClientSecretWoVersion(val *float64) {
	if err := j.validateSetClientSecretWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWoVersion",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetIdentityTokenAudience(val *string) {
	if err := j.validateSetIdentityTokenAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudience",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetIdentityTokenTtl(val *float64) {
	if err := j.validateSetIdentityTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenTtl",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetMaxRetryDelay(val *float64) {
	if err := j.validateSetMaxRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryDelay",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetRetryDelay(val *float64) {
	if err := j.validateSetRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDelay",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_AzureAuthBackendConfig)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTN code for importing a AzureAuthBackendConfig resource upon running "cdktn plan <stack-name>".
func AzureAuthBackendConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAzureAuthBackendConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
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
func AzureAuthBackendConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureAuthBackendConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzureAuthBackendConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureAuthBackendConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzureAuthBackendConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzureAuthBackendConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzureAuthBackendConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.azureAuthBackendConfig.AzureAuthBackendConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetClientSecretWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetIdentityTokenAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityTokenAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetIdentityTokenTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityTokenTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetMaxRetryDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetryDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetRetryDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureAuthBackendConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureAuthBackendConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

