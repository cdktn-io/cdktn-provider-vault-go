// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package raftsnapshotagentconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/raftsnapshotagentconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config vault_raft_snapshot_agent_config}.
type RaftSnapshotAgentConfig interface {
	cdktn.TerraformResource
	AutoloadEnabled() interface{}
	SetAutoloadEnabled(val interface{})
	AutoloadEnabledInput() interface{}
	AwsAccessKeyId() *string
	SetAwsAccessKeyId(val *string)
	AwsAccessKeyIdInput() *string
	AwsS3Bucket() *string
	SetAwsS3Bucket(val *string)
	AwsS3BucketInput() *string
	AwsS3DisableTls() interface{}
	SetAwsS3DisableTls(val interface{})
	AwsS3DisableTlsInput() interface{}
	AwsS3EnableKms() interface{}
	SetAwsS3EnableKms(val interface{})
	AwsS3EnableKmsInput() interface{}
	AwsS3Endpoint() *string
	SetAwsS3Endpoint(val *string)
	AwsS3EndpointInput() *string
	AwsS3ForcePathStyle() interface{}
	SetAwsS3ForcePathStyle(val interface{})
	AwsS3ForcePathStyleInput() interface{}
	AwsS3KmsKey() *string
	SetAwsS3KmsKey(val *string)
	AwsS3KmsKeyInput() *string
	AwsS3Region() *string
	SetAwsS3Region(val *string)
	AwsS3RegionInput() *string
	AwsS3ServerSideEncryption() interface{}
	SetAwsS3ServerSideEncryption(val interface{})
	AwsS3ServerSideEncryptionInput() interface{}
	AwsSecretAccessKey() *string
	SetAwsSecretAccessKey(val *string)
	AwsSecretAccessKeyInput() *string
	AwsSessionToken() *string
	SetAwsSessionToken(val *string)
	AwsSessionTokenInput() *string
	AzureAccountKey() *string
	SetAzureAccountKey(val *string)
	AzureAccountKeyInput() *string
	AzureAccountName() *string
	SetAzureAccountName(val *string)
	AzureAccountNameInput() *string
	AzureAuthMode() *string
	SetAzureAuthMode(val *string)
	AzureAuthModeInput() *string
	AzureBlobEnvironment() *string
	SetAzureBlobEnvironment(val *string)
	AzureBlobEnvironmentInput() *string
	AzureClientId() *string
	SetAzureClientId(val *string)
	AzureClientIdInput() *string
	AzureContainerName() *string
	SetAzureContainerName(val *string)
	AzureContainerNameInput() *string
	AzureEndpoint() *string
	SetAzureEndpoint(val *string)
	AzureEndpointInput() *string
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
	FilePrefix() *string
	SetFilePrefix(val *string)
	FilePrefixInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleDisableTls() interface{}
	SetGoogleDisableTls(val interface{})
	GoogleDisableTlsInput() interface{}
	GoogleEndpoint() *string
	SetGoogleEndpoint(val *string)
	GoogleEndpointInput() *string
	GoogleGcsBucket() *string
	SetGoogleGcsBucket(val *string)
	GoogleGcsBucketInput() *string
	GoogleServiceAccountKey() *string
	SetGoogleServiceAccountKey(val *string)
	GoogleServiceAccountKeyInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntervalSeconds() *float64
	SetIntervalSeconds(val *float64)
	IntervalSecondsInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalMaxSpace() *float64
	SetLocalMaxSpace(val *float64)
	LocalMaxSpaceInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PathPrefix() *string
	SetPathPrefix(val *string)
	PathPrefixInput() *string
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
	Retain() *float64
	SetRetain(val *float64)
	RetainInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	ResetAutoloadEnabled()
	ResetAwsAccessKeyId()
	ResetAwsS3Bucket()
	ResetAwsS3DisableTls()
	ResetAwsS3EnableKms()
	ResetAwsS3Endpoint()
	ResetAwsS3ForcePathStyle()
	ResetAwsS3KmsKey()
	ResetAwsS3Region()
	ResetAwsS3ServerSideEncryption()
	ResetAwsSecretAccessKey()
	ResetAwsSessionToken()
	ResetAzureAccountKey()
	ResetAzureAccountName()
	ResetAzureAuthMode()
	ResetAzureBlobEnvironment()
	ResetAzureClientId()
	ResetAzureContainerName()
	ResetAzureEndpoint()
	ResetFilePrefix()
	ResetGoogleDisableTls()
	ResetGoogleEndpoint()
	ResetGoogleGcsBucket()
	ResetGoogleServiceAccountKey()
	ResetId()
	ResetLocalMaxSpace()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetain()
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

// The jsii proxy struct for RaftSnapshotAgentConfig
type jsiiProxy_RaftSnapshotAgentConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AutoloadEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoloadEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AutoloadEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoloadEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3DisableTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3DisableTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3DisableTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3DisableTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3EnableKms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3EnableKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3EnableKmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3EnableKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3Endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3EndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3ForcePathStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3ForcePathStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3ForcePathStyleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3ForcePathStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3KmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3KmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsS3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3ServerSideEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3ServerSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsS3ServerSideEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsS3ServerSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsSessionToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSessionToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AwsSessionTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSessionTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureAuthModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureBlobEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureBlobEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureBlobEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) AzureEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) FilePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) FilePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleDisableTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleDisableTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleDisableTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleDisableTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleGcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleGcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleGcsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleGcsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleServiceAccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) GoogleServiceAccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleServiceAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) LocalMaxSpace() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localMaxSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) LocalMaxSpaceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localMaxSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) PathPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) PathPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) Retain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) RetainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RaftSnapshotAgentConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config vault_raft_snapshot_agent_config} Resource.
func NewRaftSnapshotAgentConfig(scope constructs.Construct, id *string, config *RaftSnapshotAgentConfigConfig) RaftSnapshotAgentConfig {
	_init_.Initialize()

	if err := validateNewRaftSnapshotAgentConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RaftSnapshotAgentConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config vault_raft_snapshot_agent_config} Resource.
func NewRaftSnapshotAgentConfig_Override(r RaftSnapshotAgentConfig, scope constructs.Construct, id *string, config *RaftSnapshotAgentConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAutoloadEnabled(val interface{}) {
	if err := j.validateSetAutoloadEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoloadEnabled",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsAccessKeyId(val *string) {
	if err := j.validateSetAwsAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3Bucket(val *string) {
	if err := j.validateSetAwsS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3Bucket",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3DisableTls(val interface{}) {
	if err := j.validateSetAwsS3DisableTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3DisableTls",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3EnableKms(val interface{}) {
	if err := j.validateSetAwsS3EnableKmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3EnableKms",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3Endpoint(val *string) {
	if err := j.validateSetAwsS3EndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3Endpoint",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3ForcePathStyle(val interface{}) {
	if err := j.validateSetAwsS3ForcePathStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3ForcePathStyle",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3KmsKey(val *string) {
	if err := j.validateSetAwsS3KmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3KmsKey",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3Region(val *string) {
	if err := j.validateSetAwsS3RegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3Region",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsS3ServerSideEncryption(val interface{}) {
	if err := j.validateSetAwsS3ServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsS3ServerSideEncryption",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsSecretAccessKey(val *string) {
	if err := j.validateSetAwsSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAwsSessionToken(val *string) {
	if err := j.validateSetAwsSessionTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsSessionToken",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureAccountKey(val *string) {
	if err := j.validateSetAzureAccountKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAccountKey",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureAccountName(val *string) {
	if err := j.validateSetAzureAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAccountName",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureAuthMode(val *string) {
	if err := j.validateSetAzureAuthModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureAuthMode",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureBlobEnvironment(val *string) {
	if err := j.validateSetAzureBlobEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureBlobEnvironment",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureClientId(val *string) {
	if err := j.validateSetAzureClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureClientId",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureContainerName(val *string) {
	if err := j.validateSetAzureContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureContainerName",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetAzureEndpoint(val *string) {
	if err := j.validateSetAzureEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureEndpoint",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetFilePrefix(val *string) {
	if err := j.validateSetFilePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePrefix",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetGoogleDisableTls(val interface{}) {
	if err := j.validateSetGoogleDisableTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleDisableTls",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetGoogleEndpoint(val *string) {
	if err := j.validateSetGoogleEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleEndpoint",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetGoogleGcsBucket(val *string) {
	if err := j.validateSetGoogleGcsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleGcsBucket",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetGoogleServiceAccountKey(val *string) {
	if err := j.validateSetGoogleServiceAccountKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleServiceAccountKey",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetLocalMaxSpace(val *float64) {
	if err := j.validateSetLocalMaxSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localMaxSpace",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetPathPrefix(val *string) {
	if err := j.validateSetPathPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathPrefix",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetRetain(val *float64) {
	if err := j.validateSetRetainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retain",
		val,
	)
}

func (j *jsiiProxy_RaftSnapshotAgentConfig)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

// Generates CDKTN code for importing a RaftSnapshotAgentConfig resource upon running "cdktn plan <stack-name>".
func RaftSnapshotAgentConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRaftSnapshotAgentConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
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
func RaftSnapshotAgentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftSnapshotAgentConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RaftSnapshotAgentConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftSnapshotAgentConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RaftSnapshotAgentConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRaftSnapshotAgentConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RaftSnapshotAgentConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.raftSnapshotAgentConfig.RaftSnapshotAgentConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAutoloadEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoloadEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsAccessKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsAccessKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3Bucket() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3Bucket",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3DisableTls() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3DisableTls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3EnableKms() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3EnableKms",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3Endpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3Endpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3ForcePathStyle() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3ForcePathStyle",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3KmsKey() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3KmsKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3Region() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3Region",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsS3ServerSideEncryption() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsS3ServerSideEncryption",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsSecretAccessKey() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsSecretAccessKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAwsSessionToken() {
	_jsii_.InvokeVoid(
		r,
		"resetAwsSessionToken",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureAccountKey() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureAccountKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureAccountName() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureAccountName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureAuthMode() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureAuthMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureBlobEnvironment() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureBlobEnvironment",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureClientId() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureClientId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureContainerName() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureContainerName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetAzureEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetAzureEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetFilePrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetFilePrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetGoogleDisableTls() {
	_jsii_.InvokeVoid(
		r,
		"resetGoogleDisableTls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetGoogleEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetGoogleEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetGoogleGcsBucket() {
	_jsii_.InvokeVoid(
		r,
		"resetGoogleGcsBucket",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetGoogleServiceAccountKey() {
	_jsii_.InvokeVoid(
		r,
		"resetGoogleServiceAccountKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetLocalMaxSpace() {
	_jsii_.InvokeVoid(
		r,
		"resetLocalMaxSpace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ResetRetain() {
	_jsii_.InvokeVoid(
		r,
		"resetRetain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RaftSnapshotAgentConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

