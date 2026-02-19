// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/awsauthbackendclient/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_client vault_aws_auth_backend_client}.
type AwsAuthBackendClient interface {
	cdktn.TerraformResource
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AllowedStsHeaderValues() *[]*string
	SetAllowedStsHeaderValues(val *[]*string)
	AllowedStsHeaderValuesInput() *[]*string
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
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	Ec2Endpoint() *string
	SetEc2Endpoint(val *string)
	Ec2EndpointInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamEndpoint() *string
	SetIamEndpoint(val *string)
	IamEndpointInput() *string
	IamServerIdHeaderValue() *string
	SetIamServerIdHeaderValue(val *string)
	IamServerIdHeaderValueInput() *string
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
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	RotationPeriod() *float64
	SetRotationPeriod(val *float64)
	RotationPeriodInput() *float64
	RotationSchedule() *string
	SetRotationSchedule(val *string)
	RotationScheduleInput() *string
	RotationWindow() *float64
	SetRotationWindow(val *float64)
	RotationWindowInput() *float64
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SecretKeyWo() *string
	SetSecretKeyWo(val *string)
	SecretKeyWoInput() *string
	SecretKeyWoVersion() *float64
	SetSecretKeyWoVersion(val *float64)
	SecretKeyWoVersionInput() *float64
	StsEndpoint() *string
	SetStsEndpoint(val *string)
	StsEndpointInput() *string
	StsRegion() *string
	SetStsRegion(val *string)
	StsRegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseStsRegionFromClient() interface{}
	SetUseStsRegionFromClient(val interface{})
	UseStsRegionFromClientInput() interface{}
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
	ResetAccessKey()
	ResetAllowedStsHeaderValues()
	ResetBackend()
	ResetDisableAutomatedRotation()
	ResetEc2Endpoint()
	ResetIamEndpoint()
	ResetIamServerIdHeaderValue()
	ResetId()
	ResetIdentityTokenAudience()
	ResetIdentityTokenTtl()
	ResetMaxRetries()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoleArn()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetSecretKey()
	ResetSecretKeyWo()
	ResetSecretKeyWoVersion()
	ResetStsEndpoint()
	ResetStsRegion()
	ResetUseStsRegionFromClient()
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

// The jsii proxy struct for AwsAuthBackendClient
type jsiiProxy_AwsAuthBackendClient struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAuthBackendClient) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) AllowedStsHeaderValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedStsHeaderValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) AllowedStsHeaderValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedStsHeaderValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Ec2Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Ec2EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IamEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IamEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IamServerIdHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServerIdHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IamServerIdHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServerIdHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IdentityTokenAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IdentityTokenAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IdentityTokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IdentityTokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"identityTokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secretKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) SecretKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secretKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) StsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) StsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) StsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) StsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) UseStsRegionFromClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStsRegionFromClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendClient) UseStsRegionFromClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStsRegionFromClientInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_client vault_aws_auth_backend_client} Resource.
func NewAwsAuthBackendClient(scope constructs.Construct, id *string, config *AwsAuthBackendClientConfig) AwsAuthBackendClient {
	_init_.Initialize()

	if err := validateNewAwsAuthBackendClientParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAuthBackendClient{}

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_client vault_aws_auth_backend_client} Resource.
func NewAwsAuthBackendClient_Override(a AwsAuthBackendClient, scope constructs.Construct, id *string, config *AwsAuthBackendClientConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetAllowedStsHeaderValues(val *[]*string) {
	if err := j.validateSetAllowedStsHeaderValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedStsHeaderValues",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetEc2Endpoint(val *string) {
	if err := j.validateSetEc2EndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2Endpoint",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetIamEndpoint(val *string) {
	if err := j.validateSetIamEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetIamServerIdHeaderValue(val *string) {
	if err := j.validateSetIamServerIdHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamServerIdHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetIdentityTokenAudience(val *string) {
	if err := j.validateSetIdentityTokenAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenAudience",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetIdentityTokenTtl(val *float64) {
	if err := j.validateSetIdentityTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetSecretKeyWo(val *string) {
	if err := j.validateSetSecretKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKeyWo",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetSecretKeyWoVersion(val *float64) {
	if err := j.validateSetSecretKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetStsEndpoint(val *string) {
	if err := j.validateSetStsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetStsRegion(val *string) {
	if err := j.validateSetStsRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendClient)SetUseStsRegionFromClient(val interface{}) {
	if err := j.validateSetUseStsRegionFromClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStsRegionFromClient",
		val,
	)
}

// Generates CDKTN code for importing a AwsAuthBackendClient resource upon running "cdktn plan <stack-name>".
func AwsAuthBackendClient_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAuthBackendClient_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
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
func AwsAuthBackendClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendClient_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendClient_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendClient_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendClient_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAuthBackendClient_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.awsAuthBackendClient.AwsAuthBackendClient",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAuthBackendClient) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAuthBackendClient) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAuthBackendClient) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetAllowedStsHeaderValues() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedStsHeaderValues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetEc2Endpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2Endpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetIamEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetIamEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetIamServerIdHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetIamServerIdHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetIdentityTokenAudience() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityTokenAudience",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetIdentityTokenTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityTokenTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetSecretKeyWo() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKeyWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetSecretKeyWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKeyWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetStsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetStsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetStsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetStsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) ResetUseStsRegionFromClient() {
	_jsii_.InvokeVoid(
		a,
		"resetUseStsRegionFromClient",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendClient) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendClient) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

