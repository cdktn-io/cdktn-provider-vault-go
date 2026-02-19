// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncghdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/secretssyncghdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gh_destination vault_secrets_sync_gh_destination}.
type SecretsSyncGhDestination interface {
	cdktn.TerraformResource
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	AllowedIpv4Addresses() *[]*string
	SetAllowedIpv4Addresses(val *[]*string)
	AllowedIpv4AddressesInput() *[]*string
	AllowedIpv6Addresses() *[]*string
	SetAllowedIpv6Addresses(val *[]*string)
	AllowedIpv6AddressesInput() *[]*string
	AllowedPorts() *[]*float64
	SetAllowedPorts(val *[]*float64)
	AllowedPortsInput() *[]*float64
	AppName() *string
	SetAppName(val *string)
	AppNameInput() *string
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
	DisableStrictNetworking() interface{}
	SetDisableStrictNetworking(val interface{})
	DisableStrictNetworkingInput() interface{}
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	EnvironmentNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Granularity() *string
	SetGranularity(val *string)
	GranularityInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallationId() *float64
	SetInstallationId(val *float64)
	InstallationIdInput() *float64
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
	RepositoryName() *string
	SetRepositoryName(val *string)
	RepositoryNameInput() *string
	RepositoryOwner() *string
	SetRepositoryOwner(val *string)
	RepositoryOwnerInput() *string
	SecretNameTemplate() *string
	SetSecretNameTemplate(val *string)
	SecretNameTemplateInput() *string
	SecretsLocation() *string
	SetSecretsLocation(val *string)
	SecretsLocationInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
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
	ResetAccessToken()
	ResetAllowedIpv4Addresses()
	ResetAllowedIpv6Addresses()
	ResetAllowedPorts()
	ResetAppName()
	ResetDisableStrictNetworking()
	ResetEnvironmentName()
	ResetGranularity()
	ResetId()
	ResetInstallationId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRepositoryName()
	ResetRepositoryOwner()
	ResetSecretNameTemplate()
	ResetSecretsLocation()
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
}

// The jsii proxy struct for SecretsSyncGhDestination
type jsiiProxy_SecretsSyncGhDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SecretsSyncGhDestination) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedIpv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedIpv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedIpv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedIpv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AllowedPortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) DisableStrictNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) DisableStrictNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) EnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) InstallationId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"installationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) InstallationIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"installationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) RepositoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) RepositoryOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) RepositoryOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) SecretNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) SecretNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) SecretsLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) SecretsLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGhDestination) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gh_destination vault_secrets_sync_gh_destination} Resource.
func NewSecretsSyncGhDestination(scope constructs.Construct, id *string, config *SecretsSyncGhDestinationConfig) SecretsSyncGhDestination {
	_init_.Initialize()

	if err := validateNewSecretsSyncGhDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsSyncGhDestination{}

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gh_destination vault_secrets_sync_gh_destination} Resource.
func NewSecretsSyncGhDestination_Override(s SecretsSyncGhDestination, scope constructs.Construct, id *string, config *SecretsSyncGhDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetAllowedIpv4Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetAllowedIpv6Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetAllowedPorts(val *[]*float64) {
	if err := j.validateSetAllowedPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPorts",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetDisableStrictNetworking(val interface{}) {
	if err := j.validateSetDisableStrictNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableStrictNetworking",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetEnvironmentName(val *string) {
	if err := j.validateSetEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetInstallationId(val *float64) {
	if err := j.validateSetInstallationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installationId",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetRepositoryName(val *string) {
	if err := j.validateSetRepositoryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryName",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetRepositoryOwner(val *string) {
	if err := j.validateSetRepositoryOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryOwner",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetSecretNameTemplate(val *string) {
	if err := j.validateSetSecretNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretNameTemplate",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGhDestination)SetSecretsLocation(val *string) {
	if err := j.validateSetSecretsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsLocation",
		val,
	)
}

// Generates CDKTN code for importing a SecretsSyncGhDestination resource upon running "cdktn plan <stack-name>".
func SecretsSyncGhDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSecretsSyncGhDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
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
func SecretsSyncGhDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGhDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncGhDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGhDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncGhDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGhDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecretsSyncGhDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.secretsSyncGhDestination.SecretsSyncGhDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsSyncGhDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsSyncGhDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncGhDestination) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetAccessToken() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetAllowedIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetAllowedIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetAllowedPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetAppName() {
	_jsii_.InvokeVoid(
		s,
		"resetAppName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetDisableStrictNetworking() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableStrictNetworking",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetEnvironmentName() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironmentName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetGranularity() {
	_jsii_.InvokeVoid(
		s,
		"resetGranularity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetInstallationId() {
	_jsii_.InvokeVoid(
		s,
		"resetInstallationId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetRepositoryName() {
	_jsii_.InvokeVoid(
		s,
		"resetRepositoryName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetRepositoryOwner() {
	_jsii_.InvokeVoid(
		s,
		"resetRepositoryOwner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetSecretNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) ResetSecretsLocation() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretsLocation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGhDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGhDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

