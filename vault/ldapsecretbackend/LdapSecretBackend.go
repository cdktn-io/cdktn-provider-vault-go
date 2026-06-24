// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapsecretbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/ldapsecretbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend vault_ldap_secret_backend}.
type LdapSecretBackend interface {
	cdktn.TerraformResource
	Accessor() *string
	AllowedManagedKeys() *[]*string
	SetAllowedManagedKeys(val *[]*string)
	AllowedManagedKeysInput() *[]*string
	AllowedResponseHeaders() *[]*string
	SetAllowedResponseHeaders(val *[]*string)
	AllowedResponseHeadersInput() *[]*string
	AuditNonHmacRequestKeys() *[]*string
	SetAuditNonHmacRequestKeys(val *[]*string)
	AuditNonHmacRequestKeysInput() *[]*string
	AuditNonHmacResponseKeys() *[]*string
	SetAuditNonHmacResponseKeys(val *[]*string)
	AuditNonHmacResponseKeysInput() *[]*string
	Binddn() *string
	SetBinddn(val *string)
	BinddnInput() *string
	Bindpass() *string
	SetBindpass(val *string)
	BindpassInput() *string
	BindpassWo() *string
	SetBindpassWo(val *string)
	BindpassWoInput() *string
	BindpassWoVersion() *float64
	SetBindpassWoVersion(val *float64)
	BindpassWoVersionInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ClientTlsCert() *string
	SetClientTlsCert(val *string)
	ClientTlsCertInput() *string
	ClientTlsKey() *string
	SetClientTlsKey(val *string)
	ClientTlsKeyInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionTimeout() *float64
	SetConnectionTimeout(val *float64)
	ConnectionTimeoutInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CredentialType() *string
	SetCredentialType(val *string)
	CredentialTypeInput() *string
	DefaultLeaseTtlSeconds() *float64
	SetDefaultLeaseTtlSeconds(val *float64)
	DefaultLeaseTtlSecondsInput() *float64
	DelegatedAuthAccessors() *[]*string
	SetDelegatedAuthAccessors(val *[]*string)
	DelegatedAuthAccessorsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	DisableRemount() interface{}
	SetDisableRemount(val interface{})
	DisableRemountInput() interface{}
	ExternalEntropyAccess() interface{}
	SetExternalEntropyAccess(val interface{})
	ExternalEntropyAccessInput() interface{}
	ForceNoCache() interface{}
	SetForceNoCache(val interface{})
	ForceNoCacheInput() interface{}
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
	IdentityTokenKey() *string
	SetIdentityTokenKey(val *string)
	IdentityTokenKeyInput() *string
	IdInput() *string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListingVisibility() *string
	SetListingVisibility(val *string)
	ListingVisibilityInput() *string
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	PassthroughRequestHeaders() *[]*string
	SetPassthroughRequestHeaders(val *[]*string)
	PassthroughRequestHeadersInput() *[]*string
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PluginVersion() *string
	SetPluginVersion(val *string)
	PluginVersionInput() *string
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
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
	RotationPeriod() *float64
	SetRotationPeriod(val *float64)
	RotationPeriodInput() *float64
	RotationSchedule() *string
	SetRotationSchedule(val *string)
	RotationScheduleInput() *string
	RotationWindow() *float64
	SetRotationWindow(val *float64)
	RotationWindowInput() *float64
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	SealWrap() interface{}
	SetSealWrap(val interface{})
	SealWrapInput() interface{}
	SelfManaged() interface{}
	SetSelfManaged(val interface{})
	SelfManagedInput() interface{}
	SkipStaticRoleImportRotation() interface{}
	SetSkipStaticRoleImportRotation(val interface{})
	SkipStaticRoleImportRotationInput() interface{}
	Starttls() interface{}
	SetStarttls(val interface{})
	StarttlsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Upndomain() *string
	SetUpndomain(val *string)
	UpndomainInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Userattr() *string
	SetUserattr(val *string)
	UserattrInput() *string
	Userdn() *string
	SetUserdn(val *string)
	UserdnInput() *string
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
	ResetAllowedManagedKeys()
	ResetAllowedResponseHeaders()
	ResetAuditNonHmacRequestKeys()
	ResetAuditNonHmacResponseKeys()
	ResetBindpass()
	ResetBindpassWo()
	ResetBindpassWoVersion()
	ResetCertificate()
	ResetClientTlsCert()
	ResetClientTlsKey()
	ResetConnectionTimeout()
	ResetCredentialType()
	ResetDefaultLeaseTtlSeconds()
	ResetDelegatedAuthAccessors()
	ResetDescription()
	ResetDisableAutomatedRotation()
	ResetDisableRemount()
	ResetExternalEntropyAccess()
	ResetForceNoCache()
	ResetId()
	ResetIdentityTokenKey()
	ResetInsecureTls()
	ResetListingVisibility()
	ResetLocal()
	ResetMaxLeaseTtlSeconds()
	ResetNamespace()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassthroughRequestHeaders()
	ResetPasswordPolicy()
	ResetPath()
	ResetPluginVersion()
	ResetRequestTimeout()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetSchema()
	ResetSealWrap()
	ResetSelfManaged()
	ResetSkipStaticRoleImportRotation()
	ResetStarttls()
	ResetUpndomain()
	ResetUrl()
	ResetUserattr()
	ResetUserdn()
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

// The jsii proxy struct for LdapSecretBackend
type jsiiProxy_LdapSecretBackend struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LdapSecretBackend) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AllowedManagedKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AllowedManagedKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AllowedResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AllowedResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Binddn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BinddnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Bindpass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BindpassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BindpassWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BindpassWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BindpassWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) BindpassWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ClientTlsCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ClientTlsCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ClientTlsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ClientTlsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) CredentialType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) CredentialTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DefaultLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DefaultLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DelegatedAuthAccessors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DelegatedAuthAccessorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ExternalEntropyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ExternalEntropyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ForceNoCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ForceNoCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) IdentityTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) IdentityTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ListingVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) ListingVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PassthroughRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PassthroughRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PluginVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) PluginVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SealWrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SealWrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SelfManaged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SelfManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SkipStaticRoleImportRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) SkipStaticRoleImportRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Starttls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) StarttlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Upndomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) UpndomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Userattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) UserattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) Userdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapSecretBackend) UserdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend vault_ldap_secret_backend} Resource.
func NewLdapSecretBackend(scope constructs.Construct, id *string, config *LdapSecretBackendConfig) LdapSecretBackend {
	_init_.Initialize()

	if err := validateNewLdapSecretBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LdapSecretBackend{}

	_jsii_.Create(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ldap_secret_backend vault_ldap_secret_backend} Resource.
func NewLdapSecretBackend_Override(l LdapSecretBackend, scope constructs.Construct, id *string, config *LdapSecretBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetAllowedManagedKeys(val *[]*string) {
	if err := j.validateSetAllowedManagedKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedManagedKeys",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetAllowedResponseHeaders(val *[]*string) {
	if err := j.validateSetAllowedResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetAuditNonHmacRequestKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacRequestKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetAuditNonHmacResponseKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacResponseKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetBinddn(val *string) {
	if err := j.validateSetBinddnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binddn",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetBindpass(val *string) {
	if err := j.validateSetBindpassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpass",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetBindpassWo(val *string) {
	if err := j.validateSetBindpassWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWo",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetBindpassWoVersion(val *float64) {
	if err := j.validateSetBindpassWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWoVersion",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetClientTlsCert(val *string) {
	if err := j.validateSetClientTlsCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsCert",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetClientTlsKey(val *string) {
	if err := j.validateSetClientTlsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsKey",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetCredentialType(val *string) {
	if err := j.validateSetCredentialTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialType",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDefaultLeaseTtlSeconds(val *float64) {
	if err := j.validateSetDefaultLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDelegatedAuthAccessors(val *[]*string) {
	if err := j.validateSetDelegatedAuthAccessorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedAuthAccessors",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetExternalEntropyAccess(val interface{}) {
	if err := j.validateSetExternalEntropyAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEntropyAccess",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetForceNoCache(val interface{}) {
	if err := j.validateSetForceNoCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNoCache",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetIdentityTokenKey(val *string) {
	if err := j.validateSetIdentityTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKey",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetListingVisibility(val *string) {
	if err := j.validateSetListingVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingVisibility",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetMaxLeaseTtlSeconds(val *float64) {
	if err := j.validateSetMaxLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetPassthroughRequestHeaders(val *[]*string) {
	if err := j.validateSetPassthroughRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetPluginVersion(val *string) {
	if err := j.validateSetPluginVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginVersion",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetRequestTimeout(val *float64) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetSealWrap(val interface{}) {
	if err := j.validateSetSealWrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sealWrap",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetSelfManaged(val interface{}) {
	if err := j.validateSetSelfManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManaged",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetSkipStaticRoleImportRotation(val interface{}) {
	if err := j.validateSetSkipStaticRoleImportRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipStaticRoleImportRotation",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetStarttls(val interface{}) {
	if err := j.validateSetStarttlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"starttls",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetUpndomain(val *string) {
	if err := j.validateSetUpndomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upndomain",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetUserattr(val *string) {
	if err := j.validateSetUserattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userattr",
		val,
	)
}

func (j *jsiiProxy_LdapSecretBackend)SetUserdn(val *string) {
	if err := j.validateSetUserdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userdn",
		val,
	)
}

// Generates CDKTN code for importing a LdapSecretBackend resource upon running "cdktn plan <stack-name>".
func LdapSecretBackend_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLdapSecretBackend_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
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
func LdapSecretBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapSecretBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapSecretBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapSecretBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapSecretBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapSecretBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LdapSecretBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.ldapSecretBackend.LdapSecretBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LdapSecretBackend) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LdapSecretBackend) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LdapSecretBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LdapSecretBackend) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LdapSecretBackend) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LdapSecretBackend) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LdapSecretBackend) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetAllowedManagedKeys() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedManagedKeys",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetAllowedResponseHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedResponseHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		l,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		l,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetBindpass() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpass",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetBindpassWo() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpassWo",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetBindpassWoVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpassWoVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetCertificate() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetClientTlsCert() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTlsCert",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetClientTlsKey() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTlsKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetCredentialType() {
	_jsii_.InvokeVoid(
		l,
		"resetCredentialType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetDefaultLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetDelegatedAuthAccessors() {
	_jsii_.InvokeVoid(
		l,
		"resetDelegatedAuthAccessors",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetExternalEntropyAccess() {
	_jsii_.InvokeVoid(
		l,
		"resetExternalEntropyAccess",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetForceNoCache() {
	_jsii_.InvokeVoid(
		l,
		"resetForceNoCache",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetIdentityTokenKey() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentityTokenKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		l,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetListingVisibility() {
	_jsii_.InvokeVoid(
		l,
		"resetListingVisibility",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetLocal() {
	_jsii_.InvokeVoid(
		l,
		"resetLocal",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		l,
		"resetNamespace",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetOptions() {
	_jsii_.InvokeVoid(
		l,
		"resetOptions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetPassthroughRequestHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetPassthroughRequestHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetPath() {
	_jsii_.InvokeVoid(
		l,
		"resetPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetPluginVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPluginVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetSchema() {
	_jsii_.InvokeVoid(
		l,
		"resetSchema",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetSealWrap() {
	_jsii_.InvokeVoid(
		l,
		"resetSealWrap",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetSelfManaged() {
	_jsii_.InvokeVoid(
		l,
		"resetSelfManaged",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetSkipStaticRoleImportRotation() {
	_jsii_.InvokeVoid(
		l,
		"resetSkipStaticRoleImportRotation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetStarttls() {
	_jsii_.InvokeVoid(
		l,
		"resetStarttls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetUpndomain() {
	_jsii_.InvokeVoid(
		l,
		"resetUpndomain",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetUserattr() {
	_jsii_.InvokeVoid(
		l,
		"resetUserattr",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) ResetUserdn() {
	_jsii_.InvokeVoid(
		l,
		"resetUserdn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapSecretBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapSecretBackend) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

