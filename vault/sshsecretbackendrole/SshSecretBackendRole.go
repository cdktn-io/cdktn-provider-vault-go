// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sshsecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/sshsecretbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ssh_secret_backend_role vault_ssh_secret_backend_role}.
type SshSecretBackendRole interface {
	cdktn.TerraformResource
	AlgorithmSigner() *string
	SetAlgorithmSigner(val *string)
	AlgorithmSignerInput() *string
	AllowBareDomains() interface{}
	SetAllowBareDomains(val interface{})
	AllowBareDomainsInput() interface{}
	AllowedCriticalOptions() *string
	SetAllowedCriticalOptions(val *string)
	AllowedCriticalOptionsInput() *string
	AllowedDomains() *string
	SetAllowedDomains(val *string)
	AllowedDomainsInput() *string
	AllowedDomainsTemplate() interface{}
	SetAllowedDomainsTemplate(val interface{})
	AllowedDomainsTemplateInput() interface{}
	AllowedExtensions() *string
	SetAllowedExtensions(val *string)
	AllowedExtensionsInput() *string
	AllowedUserKeyConfig() SshSecretBackendRoleAllowedUserKeyConfigList
	AllowedUserKeyConfigInput() interface{}
	AllowedUsers() *string
	SetAllowedUsers(val *string)
	AllowedUsersInput() *string
	AllowedUsersTemplate() interface{}
	SetAllowedUsersTemplate(val interface{})
	AllowedUsersTemplateInput() interface{}
	AllowEmptyPrincipals() interface{}
	SetAllowEmptyPrincipals(val interface{})
	AllowEmptyPrincipalsInput() interface{}
	AllowHostCertificates() interface{}
	SetAllowHostCertificates(val interface{})
	AllowHostCertificatesInput() interface{}
	AllowSubdomains() interface{}
	SetAllowSubdomains(val interface{})
	AllowSubdomainsInput() interface{}
	AllowUserCertificates() interface{}
	SetAllowUserCertificates(val interface{})
	AllowUserCertificatesInput() interface{}
	AllowUserKeyIds() interface{}
	SetAllowUserKeyIds(val interface{})
	AllowUserKeyIdsInput() interface{}
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CidrList() *string
	SetCidrList(val *string)
	CidrListInput() *string
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
	DefaultCriticalOptions() *map[string]*string
	SetDefaultCriticalOptions(val *map[string]*string)
	DefaultCriticalOptionsInput() *map[string]*string
	DefaultExtensions() *map[string]*string
	SetDefaultExtensions(val *map[string]*string)
	DefaultExtensionsInput() *map[string]*string
	DefaultExtensionsTemplate() interface{}
	SetDefaultExtensionsTemplate(val interface{})
	DefaultExtensionsTemplateInput() interface{}
	DefaultUser() *string
	SetDefaultUser(val *string)
	DefaultUserInput() *string
	DefaultUserTemplate() interface{}
	SetDefaultUserTemplate(val interface{})
	DefaultUserTemplateInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcludeCidrList() *[]*string
	SetExcludeCidrList(val *[]*string)
	ExcludeCidrListInput() *[]*string
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
	KeyIdFormat() *string
	SetKeyIdFormat(val *string)
	KeyIdFormatInput() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxTtl() *string
	SetMaxTtl(val *string)
	MaxTtlInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NotBeforeDuration() *string
	SetNotBeforeDuration(val *string)
	NotBeforeDurationInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
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
	PutAllowedUserKeyConfig(value interface{})
	ResetAlgorithmSigner()
	ResetAllowBareDomains()
	ResetAllowedCriticalOptions()
	ResetAllowedDomains()
	ResetAllowedDomainsTemplate()
	ResetAllowedExtensions()
	ResetAllowedUserKeyConfig()
	ResetAllowedUsers()
	ResetAllowedUsersTemplate()
	ResetAllowEmptyPrincipals()
	ResetAllowHostCertificates()
	ResetAllowSubdomains()
	ResetAllowUserCertificates()
	ResetAllowUserKeyIds()
	ResetCidrList()
	ResetDefaultCriticalOptions()
	ResetDefaultExtensions()
	ResetDefaultExtensionsTemplate()
	ResetDefaultUser()
	ResetDefaultUserTemplate()
	ResetExcludeCidrList()
	ResetId()
	ResetKeyIdFormat()
	ResetMaxTtl()
	ResetNamespace()
	ResetNotBeforeDuration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetTtl()
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

// The jsii proxy struct for SshSecretBackendRole
type jsiiProxy_SshSecretBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SshSecretBackendRole) AlgorithmSigner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmSigner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AlgorithmSignerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmSignerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowBareDomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowBareDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowBareDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowBareDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedCriticalOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedCriticalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedCriticalOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedCriticalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedDomains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedDomainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedDomainsTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedDomainsTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedDomainsTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedDomainsTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedExtensions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedExtensionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUserKeyConfig() SshSecretBackendRoleAllowedUserKeyConfigList {
	var returns SshSecretBackendRoleAllowedUserKeyConfigList
	_jsii_.Get(
		j,
		"allowedUserKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUserKeyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedUserKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUsers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUsersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUsersTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedUsersTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowedUsersTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedUsersTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowEmptyPrincipals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowEmptyPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowEmptyPrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowEmptyPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowHostCertificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHostCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowHostCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowHostCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowUserCertificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUserCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowUserCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUserCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowUserKeyIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUserKeyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) AllowUserKeyIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUserKeyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) CidrList() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) CidrListInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultCriticalOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultCriticalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultCriticalOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultCriticalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultExtensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultExtensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultExtensionsTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultExtensionsTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultExtensionsTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultExtensionsTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultUserTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultUserTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DefaultUserTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultUserTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) ExcludeCidrList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeCidrList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) ExcludeCidrListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeCidrListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) KeyIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) KeyIdFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) MaxTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) MaxTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) NotBeforeDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) NotBeforeDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SshSecretBackendRole) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ssh_secret_backend_role vault_ssh_secret_backend_role} Resource.
func NewSshSecretBackendRole(scope constructs.Construct, id *string, config *SshSecretBackendRoleConfig) SshSecretBackendRole {
	_init_.Initialize()

	if err := validateNewSshSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SshSecretBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ssh_secret_backend_role vault_ssh_secret_backend_role} Resource.
func NewSshSecretBackendRole_Override(s SshSecretBackendRole, scope constructs.Construct, id *string, config *SshSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAlgorithmSigner(val *string) {
	if err := j.validateSetAlgorithmSignerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithmSigner",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowBareDomains(val interface{}) {
	if err := j.validateSetAllowBareDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowBareDomains",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedCriticalOptions(val *string) {
	if err := j.validateSetAllowedCriticalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCriticalOptions",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedDomains(val *string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedDomainsTemplate(val interface{}) {
	if err := j.validateSetAllowedDomainsTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomainsTemplate",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedExtensions(val *string) {
	if err := j.validateSetAllowedExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExtensions",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedUsers(val *string) {
	if err := j.validateSetAllowedUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUsers",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowedUsersTemplate(val interface{}) {
	if err := j.validateSetAllowedUsersTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUsersTemplate",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowEmptyPrincipals(val interface{}) {
	if err := j.validateSetAllowEmptyPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowEmptyPrincipals",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowHostCertificates(val interface{}) {
	if err := j.validateSetAllowHostCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowHostCertificates",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowSubdomains(val interface{}) {
	if err := j.validateSetAllowSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSubdomains",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowUserCertificates(val interface{}) {
	if err := j.validateSetAllowUserCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUserCertificates",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetAllowUserKeyIds(val interface{}) {
	if err := j.validateSetAllowUserKeyIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUserKeyIds",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetCidrList(val *string) {
	if err := j.validateSetCidrListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrList",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDefaultCriticalOptions(val *map[string]*string) {
	if err := j.validateSetDefaultCriticalOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCriticalOptions",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDefaultExtensions(val *map[string]*string) {
	if err := j.validateSetDefaultExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExtensions",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDefaultExtensionsTemplate(val interface{}) {
	if err := j.validateSetDefaultExtensionsTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExtensionsTemplate",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDefaultUser(val *string) {
	if err := j.validateSetDefaultUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultUser",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDefaultUserTemplate(val interface{}) {
	if err := j.validateSetDefaultUserTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultUserTemplate",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetExcludeCidrList(val *[]*string) {
	if err := j.validateSetExcludeCidrListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCidrList",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetKeyIdFormat(val *string) {
	if err := j.validateSetKeyIdFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyIdFormat",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetMaxTtl(val *string) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetNotBeforeDuration(val *string) {
	if err := j.validateSetNotBeforeDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notBeforeDuration",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SshSecretBackendRole)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

// Generates CDKTN code for importing a SshSecretBackendRole resource upon running "cdktn plan <stack-name>".
func SshSecretBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSshSecretBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
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
func SshSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSshSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SshSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSshSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SshSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSshSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SshSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.sshSecretBackendRole.SshSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SshSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SshSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SshSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SshSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SshSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SshSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SshSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SshSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SshSecretBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SshSecretBackendRole) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) PutAllowedUserKeyConfig(value interface{}) {
	if err := s.validatePutAllowedUserKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAllowedUserKeyConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAlgorithmSigner() {
	_jsii_.InvokeVoid(
		s,
		"resetAlgorithmSigner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowBareDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowBareDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedCriticalOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedCriticalOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedDomainsTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedDomainsTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedExtensions() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedExtensions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedUserKeyConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedUserKeyConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedUsers() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedUsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowedUsersTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedUsersTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowEmptyPrincipals() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowEmptyPrincipals",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowHostCertificates() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowHostCertificates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowSubdomains() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowSubdomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowUserCertificates() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowUserCertificates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetAllowUserKeyIds() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowUserKeyIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetCidrList() {
	_jsii_.InvokeVoid(
		s,
		"resetCidrList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetDefaultCriticalOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultCriticalOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetDefaultExtensions() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultExtensions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetDefaultExtensionsTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultExtensionsTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetDefaultUser() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultUser",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetDefaultUserTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultUserTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetExcludeCidrList() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeCidrList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetKeyIdFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetKeyIdFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetNotBeforeDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetNotBeforeDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) ResetTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SshSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SshSecretBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

