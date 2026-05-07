// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/ldapauthbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend vault_ldap_auth_backend}.
type LdapAuthBackend interface {
	cdktn.TerraformResource
	Accessor() *string
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AnonymousGroupSearch() interface{}
	SetAnonymousGroupSearch(val interface{})
	AnonymousGroupSearchInput() interface{}
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
	CaseSensitiveNames() interface{}
	SetCaseSensitiveNames(val interface{})
	CaseSensitiveNamesInput() interface{}
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
	DenyNullBind() interface{}
	SetDenyNullBind(val interface{})
	DenyNullBindInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DereferenceAliases() *string
	SetDereferenceAliases(val *string)
	DereferenceAliasesInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	DisableRemount() interface{}
	SetDisableRemount(val interface{})
	DisableRemountInput() interface{}
	Discoverdn() interface{}
	SetDiscoverdn(val interface{})
	DiscoverdnInput() interface{}
	EnableSamaccountnameLogin() interface{}
	SetEnableSamaccountnameLogin(val interface{})
	EnableSamaccountnameLoginInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groupattr() *string
	SetGroupattr(val *string)
	GroupattrInput() *string
	Groupdn() *string
	SetGroupdn(val *string)
	GroupdnInput() *string
	Groupfilter() *string
	SetGroupfilter(val *string)
	GroupfilterInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	MaxPageSize() *float64
	SetMaxPageSize(val *float64)
	MaxPageSizeInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	Starttls() interface{}
	SetStarttls(val interface{})
	StarttlsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsMaxVersion() *string
	SetTlsMaxVersion(val *string)
	TlsMaxVersionInput() *string
	TlsMinVersion() *string
	SetTlsMinVersion(val *string)
	TlsMinVersionInput() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	TokenExplicitMaxTtl() *float64
	SetTokenExplicitMaxTtl(val *float64)
	TokenExplicitMaxTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	TokenNoDefaultPolicy() interface{}
	SetTokenNoDefaultPolicy(val interface{})
	TokenNoDefaultPolicyInput() interface{}
	TokenNumUses() *float64
	SetTokenNumUses(val *float64)
	TokenNumUsesInput() *float64
	TokenPeriod() *float64
	SetTokenPeriod(val *float64)
	TokenPeriodInput() *float64
	TokenPolicies() *[]*string
	SetTokenPolicies(val *[]*string)
	TokenPoliciesInput() *[]*string
	TokenTtl() *float64
	SetTokenTtl(val *float64)
	TokenTtlInput() *float64
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
	Tune() LdapAuthBackendTuneList
	TuneInput() interface{}
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
	Userfilter() *string
	SetUserfilter(val *string)
	UserfilterInput() *string
	UsernameAsAlias() interface{}
	SetUsernameAsAlias(val interface{})
	UsernameAsAliasInput() interface{}
	UseTokenGroups() interface{}
	SetUseTokenGroups(val interface{})
	UseTokenGroupsInput() interface{}
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
	PutTune(value interface{})
	ResetAliasMetadata()
	ResetAnonymousGroupSearch()
	ResetBinddn()
	ResetBindpass()
	ResetBindpassWo()
	ResetBindpassWoVersion()
	ResetCaseSensitiveNames()
	ResetCertificate()
	ResetClientTlsCert()
	ResetClientTlsKey()
	ResetConnectionTimeout()
	ResetDenyNullBind()
	ResetDereferenceAliases()
	ResetDescription()
	ResetDisableAutomatedRotation()
	ResetDisableRemount()
	ResetDiscoverdn()
	ResetEnableSamaccountnameLogin()
	ResetGroupattr()
	ResetGroupdn()
	ResetGroupfilter()
	ResetId()
	ResetInsecureTls()
	ResetLocal()
	ResetMaxPageSize()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetRequestTimeout()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetStarttls()
	ResetTlsMaxVersion()
	ResetTlsMinVersion()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
	ResetTune()
	ResetUpndomain()
	ResetUserattr()
	ResetUserdn()
	ResetUserfilter()
	ResetUsernameAsAlias()
	ResetUseTokenGroups()
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

// The jsii proxy struct for LdapAuthBackend
type jsiiProxy_LdapAuthBackend struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LdapAuthBackend) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) AnonymousGroupSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) AnonymousGroupSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Binddn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BinddnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Bindpass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BindpassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BindpassWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BindpassWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BindpassWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) BindpassWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) CaseSensitiveNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) CaseSensitiveNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ClientTlsCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ClientTlsCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ClientTlsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ClientTlsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DenyNullBind() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DenyNullBindInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DereferenceAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DereferenceAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Discoverdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) DiscoverdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) EnableSamaccountnameLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamaccountnameLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) EnableSamaccountnameLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamaccountnameLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Groupattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) GroupattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Groupdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) GroupdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Groupfilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) GroupfilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) MaxPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) MaxPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Starttls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) StarttlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TlsMaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Tune() LdapAuthBackendTuneList {
	var returns LdapAuthBackendTuneList
	_jsii_.Get(
		j,
		"tune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) TuneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Upndomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UpndomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Userattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UserattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Userdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UserdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) Userfilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userfilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UserfilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userfilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UsernameAsAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameAsAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UsernameAsAliasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameAsAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UseTokenGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackend) UseTokenGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend vault_ldap_auth_backend} Resource.
func NewLdapAuthBackend(scope constructs.Construct, id *string, config *LdapAuthBackendConfig) LdapAuthBackend {
	_init_.Initialize()

	if err := validateNewLdapAuthBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LdapAuthBackend{}

	_jsii_.Create(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/ldap_auth_backend vault_ldap_auth_backend} Resource.
func NewLdapAuthBackend_Override(l LdapAuthBackend, scope constructs.Construct, id *string, config *LdapAuthBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetAnonymousGroupSearch(val interface{}) {
	if err := j.validateSetAnonymousGroupSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousGroupSearch",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetBinddn(val *string) {
	if err := j.validateSetBinddnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binddn",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetBindpass(val *string) {
	if err := j.validateSetBindpassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpass",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetBindpassWo(val *string) {
	if err := j.validateSetBindpassWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWo",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetBindpassWoVersion(val *float64) {
	if err := j.validateSetBindpassWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWoVersion",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetCaseSensitiveNames(val interface{}) {
	if err := j.validateSetCaseSensitiveNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitiveNames",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetClientTlsCert(val *string) {
	if err := j.validateSetClientTlsCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsCert",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetClientTlsKey(val *string) {
	if err := j.validateSetClientTlsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsKey",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDenyNullBind(val interface{}) {
	if err := j.validateSetDenyNullBindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyNullBind",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDereferenceAliases(val *string) {
	if err := j.validateSetDereferenceAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dereferenceAliases",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetDiscoverdn(val interface{}) {
	if err := j.validateSetDiscoverdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoverdn",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetEnableSamaccountnameLogin(val interface{}) {
	if err := j.validateSetEnableSamaccountnameLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSamaccountnameLogin",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetGroupattr(val *string) {
	if err := j.validateSetGroupattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupattr",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetGroupdn(val *string) {
	if err := j.validateSetGroupdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupdn",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetGroupfilter(val *string) {
	if err := j.validateSetGroupfilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupfilter",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetMaxPageSize(val *float64) {
	if err := j.validateSetMaxPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPageSize",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetRequestTimeout(val *float64) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetStarttls(val interface{}) {
	if err := j.validateSetStarttlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"starttls",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTlsMaxVersion(val *string) {
	if err := j.validateSetTlsMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMaxVersion",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUpndomain(val *string) {
	if err := j.validateSetUpndomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upndomain",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUserattr(val *string) {
	if err := j.validateSetUserattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userattr",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUserdn(val *string) {
	if err := j.validateSetUserdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userdn",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUserfilter(val *string) {
	if err := j.validateSetUserfilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userfilter",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUsernameAsAlias(val interface{}) {
	if err := j.validateSetUsernameAsAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAsAlias",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackend)SetUseTokenGroups(val interface{}) {
	if err := j.validateSetUseTokenGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTokenGroups",
		val,
	)
}

// Generates CDKTN code for importing a LdapAuthBackend resource upon running "cdktn plan <stack-name>".
func LdapAuthBackend_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLdapAuthBackend_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
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
func LdapAuthBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapAuthBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapAuthBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapAuthBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapAuthBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapAuthBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LdapAuthBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LdapAuthBackend) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LdapAuthBackend) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LdapAuthBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LdapAuthBackend) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LdapAuthBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LdapAuthBackend) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LdapAuthBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LdapAuthBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LdapAuthBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LdapAuthBackend) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LdapAuthBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LdapAuthBackend) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LdapAuthBackend) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LdapAuthBackend) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LdapAuthBackend) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LdapAuthBackend) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LdapAuthBackend) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LdapAuthBackend) PutTune(value interface{}) {
	if err := l.validatePutTuneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTune",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		l,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetAnonymousGroupSearch() {
	_jsii_.InvokeVoid(
		l,
		"resetAnonymousGroupSearch",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetBinddn() {
	_jsii_.InvokeVoid(
		l,
		"resetBinddn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetBindpass() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpass",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetBindpassWo() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpassWo",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetBindpassWoVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetBindpassWoVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetCaseSensitiveNames() {
	_jsii_.InvokeVoid(
		l,
		"resetCaseSensitiveNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetCertificate() {
	_jsii_.InvokeVoid(
		l,
		"resetCertificate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetClientTlsCert() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTlsCert",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetClientTlsKey() {
	_jsii_.InvokeVoid(
		l,
		"resetClientTlsKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDenyNullBind() {
	_jsii_.InvokeVoid(
		l,
		"resetDenyNullBind",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDereferenceAliases() {
	_jsii_.InvokeVoid(
		l,
		"resetDereferenceAliases",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetDiscoverdn() {
	_jsii_.InvokeVoid(
		l,
		"resetDiscoverdn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetEnableSamaccountnameLogin() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableSamaccountnameLogin",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetGroupattr() {
	_jsii_.InvokeVoid(
		l,
		"resetGroupattr",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetGroupdn() {
	_jsii_.InvokeVoid(
		l,
		"resetGroupdn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetGroupfilter() {
	_jsii_.InvokeVoid(
		l,
		"resetGroupfilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		l,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetLocal() {
	_jsii_.InvokeVoid(
		l,
		"resetLocal",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetMaxPageSize() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxPageSize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		l,
		"resetNamespace",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetPath() {
	_jsii_.InvokeVoid(
		l,
		"resetPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		l,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetStarttls() {
	_jsii_.InvokeVoid(
		l,
		"resetStarttls",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTlsMaxVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetTlsMaxVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTokenType() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetTune() {
	_jsii_.InvokeVoid(
		l,
		"resetTune",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUpndomain() {
	_jsii_.InvokeVoid(
		l,
		"resetUpndomain",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUserattr() {
	_jsii_.InvokeVoid(
		l,
		"resetUserattr",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUserdn() {
	_jsii_.InvokeVoid(
		l,
		"resetUserdn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUserfilter() {
	_jsii_.InvokeVoid(
		l,
		"resetUserfilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUsernameAsAlias() {
	_jsii_.InvokeVoid(
		l,
		"resetUsernameAsAlias",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) ResetUseTokenGroups() {
	_jsii_.InvokeVoid(
		l,
		"resetUseTokenGroups",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackend) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

