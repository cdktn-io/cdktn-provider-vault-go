// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncawsdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/secretssyncawsdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination vault_secrets_sync_aws_destination}.
type SecretsSyncAwsDestination interface {
	cdktn.TerraformResource
	AccessKeyId() *string
	SetAccessKeyId(val *string)
	AccessKeyIdInput() *string
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
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SecretAccessKey() *string
	SetSecretAccessKey(val *string)
	SecretAccessKeyInput() *string
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
	ResetAccessKeyId()
	ResetAllowedIpv4Addresses()
	ResetAllowedIpv6Addresses()
	ResetAllowedPorts()
	ResetCustomTags()
	ResetDisableStrictNetworking()
	ResetExternalId()
	ResetGranularity()
	ResetId()
	ResetIdentityTokenAudienceWo()
	ResetIdentityTokenAudienceWoVersion()
	ResetIdentityTokenKeyWo()
	ResetIdentityTokenKeyWoVersion()
	ResetIdentityTokenTtl()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetRoleArn()
	ResetSecretAccessKey()
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

// The jsii proxy struct for SecretsSyncAwsDestination
type jsiiProxy_SecretsSyncAwsDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedIpv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedIpv4AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv4AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedIpv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedIpv6AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpv6AddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedPorts() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) AllowedPortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) CustomTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) CustomTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) DisableStrictNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) DisableStrictNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableStrictNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenAudienceWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenAudienceWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenAudienceWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenAudienceWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenAudienceWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenAudienceWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdentityTokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) SecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) SecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) SecretNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) SecretNameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsSyncAwsDestination) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination vault_secrets_sync_aws_destination} Resource.
func NewSecretsSyncAwsDestination(scope constructs.Construct, id *string, config *SecretsSyncAwsDestinationConfig) SecretsSyncAwsDestination {
	_init_.Initialize()

	if err := validateNewSecretsSyncAwsDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsSyncAwsDestination{}

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination vault_secrets_sync_aws_destination} Resource.
func NewSecretsSyncAwsDestination_Override(s SecretsSyncAwsDestination, scope constructs.Construct, id *string, config *SecretsSyncAwsDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetAccessKeyId(val *string) {
	if err := j.validateSetAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeyId",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetAllowedIpv4Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv4AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv4Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetAllowedIpv6Addresses(val *[]*string) {
	if err := j.validateSetAllowedIpv6AddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpv6Addresses",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetAllowedPorts(val *[]*float64) {
	if err := j.validateSetAllowedPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPorts",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetCustomTags(val *map[string]*string) {
	if err := j.validateSetCustomTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTags",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetDisableStrictNetworking(val interface{}) {
	if err := j.validateSetDisableStrictNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableStrictNetworking",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetExternalId(val *string) {
	if err := j.validateSetExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetIdentityTokenAudienceWo(val *string) {
	if err := j.validateSetIdentityTokenAudienceWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudienceWo",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetIdentityTokenAudienceWoVersion(val *float64) {
	if err := j.validateSetIdentityTokenAudienceWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudienceWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetIdentityTokenKeyWo(val *string) {
	if err := j.validateSetIdentityTokenKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKeyWo",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetIdentityTokenKeyWoVersion(val *float64) {
	if err := j.validateSetIdentityTokenKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetIdentityTokenTtl(val *float64) {
	if err := j.validateSetIdentityTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenTtl",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetSecretAccessKey(val *string) {
	if err := j.validateSetSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKey",
		val,
	)
}

func (j *jsiiProxy_SecretsSyncAwsDestination)SetSecretNameTemplate(val *string) {
	if err := j.validateSetSecretNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretNameTemplate",
		val,
	)
}

// Generates CDKTN code for importing a SecretsSyncAwsDestination resource upon running "cdktn plan <stack-name>".
func SecretsSyncAwsDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSecretsSyncAwsDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
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
func SecretsSyncAwsDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAwsDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncAwsDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAwsDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecretsSyncAwsDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsSyncAwsDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecretsSyncAwsDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.secretsSyncAwsDestination.SecretsSyncAwsDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecretsSyncAwsDestination) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetAccessKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetAllowedIpv4Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv4Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetAllowedIpv6Addresses() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedIpv6Addresses",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetAllowedPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetCustomTags() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetDisableStrictNetworking() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableStrictNetworking",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetExternalId() {
	_jsii_.InvokeVoid(
		s,
		"resetExternalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetGranularity() {
	_jsii_.InvokeVoid(
		s,
		"resetGranularity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetIdentityTokenAudienceWo() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenAudienceWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetIdentityTokenAudienceWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenAudienceWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetIdentityTokenKeyWo() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenKeyWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetIdentityTokenKeyWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenKeyWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetIdentityTokenTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityTokenTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetRoleArn() {
	_jsii_.InvokeVoid(
		s,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetSecretAccessKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretAccessKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ResetSecretNameTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSecretNameTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecretsSyncAwsDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsSyncAwsDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

