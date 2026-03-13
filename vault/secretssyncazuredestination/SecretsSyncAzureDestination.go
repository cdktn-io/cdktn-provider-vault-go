// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncazuredestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/secretssyncazuredestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_azure_destination vault_secrets_sync_azure_destination}.
type SecretsSyncAzureDestination interface {
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
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	Cloud() *string
	SetCloud(val *string)
	CloudInput() *string
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
	Granularity() *string
	SetGranularity(val *string)
	GranularityInput() *string
	Id() *string
	SetId(val *string)
	IdentityTokenAudienceWo() *string
	SetIdentityTokenAudienceWo(val *string)
	IdentityTokenAudienceWoInput() *string
	IdentityTokenAudienceWoVersion() *float64
	SetIdentityTokenAudienceWoVersion(val *float64)
	IdentityTokenAudienceWoVersionInput() *float64
	IdentityTokenKeyWo() *string
	SetIdentityTokenKeyWo(val *string)
	IdentityTokenKeyWoInput() *string
	IdentityTokenKeyWoVersion() *float64
	SetIdentityTokenKeyWoVersion(val *float64)
	IdentityTokenKeyWoVersionInput() *float64
	IdentityTokenTtl() *float64
	SetIdentityTokenTtl(val *float64)
	IdentityTokenTtlInput() *float64
	IdInput() *string
	KeyVaultUri() *string
	SetKeyVaultUri(val *string)
	KeyVaultUriInput() *string
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
	SecretNameTemplate() *string
	SetSecretNameTemplate(val *string)
	SecretNameTemplateInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetClientId()
	ResetClientSecret()
	ResetCloud()
	ResetCustomTags()
	ResetDisableStrictNetworking()
	ResetGranularity()
	ResetId()
	ResetIdentityTokenAudienceWo()
	ResetIdentityTokenAudienceWoVersion()
	ResetIdentityTokenKeyWo()
	ResetIdentityTokenKeyWoVersion()
	ResetIdentityTokenTtl()
	ResetKeyVaultUri()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecretNameTemplate()
	ResetTenantId()
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

// The jsii proxy struct for SecretsSyncAzureDestination
type jsiiProxy_SecretsSyncAzureDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedIpv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedIpv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedIpv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedIpv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) AllowedPortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Cloud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) CloudInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) DisableStrictNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) DisableStrictNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenAudienceWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenAudienceWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenAudienceWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenAudienceWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenAudienceWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenAudienceWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdentityTokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) KeyVaultUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) KeyVaultUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) SecretNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) SecretNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAzureDestination) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_azure_destination vault_secrets_sync_azure_destination} Resource.
func NewSecretsSyncAzureDestination(scope constructs.Construct, id *string, config *SecretsSyncAzureDestinationConfig) SecretsSyncAzureDestination {
	_init_.Initialize()

	if err := validateNewSecretsSyncAzureDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsSyncAzureDestination{}

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_azure_destination vault_secrets_sync_azure_destination} Resource.
func NewSecretsSyncAzureDestination_Override(s SecretsSyncAzureDestination, scope constructs.Construct, id *string, config *SecretsSyncAzureDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetAllowedIpv4Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetAllowedIpv6Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetAllowedPorts(val *[]*float64) {
	if err := j.validateSetAllowedPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPorts",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetCloud(val *string) {
	if err := j.validateSetCloudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloud",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetDisableStrictNetworking(val interface{}) {
	if err := j.validateSetDisableStrictNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableStrictNetworking",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetIdentityTokenAudienceWo(val *string) {
	if err := j.validateSetIdentityTokenAudienceWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudienceWo",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetIdentityTokenAudienceWoVersion(val *float64) {
	if err := j.validateSetIdentityTokenAudienceWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudienceWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetIdentityTokenKeyWo(val *string) {
	if err := j.validateSetIdentityTokenKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKeyWo",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetIdentityTokenKeyWoVersion(val *float64) {
	if err := j.validateSetIdentityTokenKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetIdentityTokenTtl(val *float64) {
	if err := j.validateSetIdentityTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenTtl",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetKeyVaultUri(val *string) {
	if err := j.validateSetKeyVaultUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultUri",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetSecretNameTemplate(val *string) {
	if err := j.validateSetSecretNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretNameTemplate",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAzureDestination)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTN code for importing a SecretsSyncAzureDestination resource upon running "cdktn plan <stack-name>".
func SecretsSyncAzureDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSecretsSyncAzureDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
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
func SecretsSyncAzureDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAzureDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncAzureDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAzureDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncAzureDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAzureDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecretsSyncAzureDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.secretsSyncAzureDestination.SecretsSyncAzureDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncAzureDestination) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetAllowedIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetAllowedIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetAllowedPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetClientSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetCloud() {
	_jsii_.InvokeVoid(
		s,
		"resetCloud",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetCustomTags() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetDisableStrictNetworking() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableStrictNetworking",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetGranularity() {
	_jsii_.InvokeVoid(
		s,
		"resetGranularity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetIdentityTokenAudienceWo() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenAudienceWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetIdentityTokenAudienceWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenAudienceWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetIdentityTokenKeyWo() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenKeyWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetIdentityTokenKeyWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenKeyWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetIdentityTokenTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetKeyVaultUri() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyVaultUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetSecretNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ResetTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAzureDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAzureDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

