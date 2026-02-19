// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncgcpdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/secretssyncgcpdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gcp_destination vault_secrets_sync_gcp_destination}.
type SecretsSyncGcpDestination interface {
	cdktn.TerraformResource
	AllowedIpv4Addresses() *[]*string
	SetAllowedIpv4Addresses(val *[]*string)
	AllowedIpv4AddressesInput() *[]*string
	AllowedIpv6Addresses() *[]*string
	SetAllowedIpv6Addresses(val *[]*string)
	AllowedIpv6AddressesInput() *[]*string
	AllowedPorts() *[]*float64
	SetAllowedPorts(val *[]*float64)
	AllowedPortsInput() *[]*float64
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
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	CustomTags() *map[string]*string
	SetCustomTags(val *map[string]*string)
	CustomTagsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableStrictNetworking() interface{}
	SetDisableStrictNetworking(val interface{})
	DisableStrictNetworkingInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalKmsKey() *string
	SetGlobalKmsKey(val *string)
	GlobalKmsKeyInput() *string
	Granularity() *string
	SetGranularity(val *string)
	GranularityInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocationalKmsKeys() *map[string]*string
	SetLocationalKmsKeys(val *map[string]*string)
	LocationalKmsKeysInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	ReplicationLocations() *[]*string
	SetReplicationLocations(val *[]*string)
	ReplicationLocationsInput() *[]*string
	SecretNameTemplate() *string
	SetSecretNameTemplate(val *string)
	SecretNameTemplateInput() *string
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
	ResetAllowedIpv4Addresses()
	ResetAllowedIpv6Addresses()
	ResetAllowedPorts()
	ResetCredentials()
	ResetCustomTags()
	ResetDisableStrictNetworking()
	ResetGlobalKmsKey()
	ResetGranularity()
	ResetId()
	ResetLocationalKmsKeys()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectId()
	ResetReplicationLocations()
	ResetSecretNameTemplate()
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

// The jsii proxy struct for SecretsSyncGcpDestination
type jsiiProxy_SecretsSyncGcpDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedIpv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedIpv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedIpv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedIpv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) AllowedPortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) DisableStrictNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) DisableStrictNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) GlobalKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) GlobalKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) LocationalKmsKeys() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"locationalKmsKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) LocationalKmsKeysInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"locationalKmsKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ReplicationLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) ReplicationLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) SecretNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) SecretNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncGcpDestination) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gcp_destination vault_secrets_sync_gcp_destination} Resource.
func NewSecretsSyncGcpDestination(scope constructs.Construct, id *string, config *SecretsSyncGcpDestinationConfig) SecretsSyncGcpDestination {
	_init_.Initialize()

	if err := validateNewSecretsSyncGcpDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsSyncGcpDestination{}

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/secrets_sync_gcp_destination vault_secrets_sync_gcp_destination} Resource.
func NewSecretsSyncGcpDestination_Override(s SecretsSyncGcpDestination, scope constructs.Construct, id *string, config *SecretsSyncGcpDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetAllowedIpv4Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetAllowedIpv6Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetAllowedPorts(val *[]*float64) {
	if err := j.validateSetAllowedPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPorts",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetCredentials(val *string) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetDisableStrictNetworking(val interface{}) {
	if err := j.validateSetDisableStrictNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableStrictNetworking",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetGlobalKmsKey(val *string) {
	if err := j.validateSetGlobalKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalKmsKey",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetLocationalKmsKeys(val *map[string]*string) {
	if err := j.validateSetLocationalKmsKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationalKmsKeys",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetReplicationLocations(val *[]*string) {
	if err := j.validateSetReplicationLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationLocations",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncGcpDestination)SetSecretNameTemplate(val *string) {
	if err := j.validateSetSecretNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretNameTemplate",
		val,
	)
}

// Generates CDKTN code for importing a SecretsSyncGcpDestination resource upon running "cdktn plan <stack-name>".
func SecretsSyncGcpDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSecretsSyncGcpDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
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
func SecretsSyncGcpDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGcpDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncGcpDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGcpDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncGcpDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncGcpDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecretsSyncGcpDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.secretsSyncGcpDestination.SecretsSyncGcpDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncGcpDestination) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetAllowedIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetAllowedIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetAllowedPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetCredentials() {
	_jsii_.InvokeVoid(
		s,
		"resetCredentials",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetCustomTags() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetDisableStrictNetworking() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableStrictNetworking",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetGlobalKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetGlobalKmsKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetGranularity() {
	_jsii_.InvokeVoid(
		s,
		"resetGranularity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetLocationalKmsKeys() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationalKmsKeys",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetProjectId() {
	_jsii_.InvokeVoid(
		s,
		"resetProjectId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetReplicationLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ResetSecretNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncGcpDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncGcpDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

