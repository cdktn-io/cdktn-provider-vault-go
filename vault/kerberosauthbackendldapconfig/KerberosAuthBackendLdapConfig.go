// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kerberosauthbackendldapconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/kerberosauthbackendldapconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/resources/kerberos_auth_backend_ldap_config vault_kerberos_auth_backend_ldap_config}.
type KerberosAuthBackendLdapConfig interface {
	cdktn.TerraformResource
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AnonymousGroupSearch() interface{}
	SetAnonymousGroupSearch(val interface{})
	AnonymousGroupSearchInput() interface{}
	Binddn() *string
	SetBinddn(val *string)
	BinddnInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	BindpassWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
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
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientTlsCertWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientTlsCertWo(val *string)
	ClientTlsCertWoInput() *string
	ClientTlsCertWoVersion() *float64
	SetClientTlsCertWoVersion(val *float64)
	ClientTlsCertWoVersionInput() *float64
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientTlsKeyWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientTlsKeyWo(val *string)
	ClientTlsKeyWoInput() *string
	ClientTlsKeyWoVersion() *float64
	SetClientTlsKeyWoVersion(val *float64)
	ClientTlsKeyWoVersionInput() *float64
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
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxPageSize() *float64
	SetMaxPageSize(val *float64)
	MaxPageSizeInput() *float64
	Mount() *string
	SetMount(val *string)
	MountInput() *string
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
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAliasMetadata()
	ResetAnonymousGroupSearch()
	ResetBinddn()
	ResetBindpassWo()
	ResetBindpassWoVersion()
	ResetCaseSensitiveNames()
	ResetCertificate()
	ResetClientTlsCertWo()
	ResetClientTlsCertWoVersion()
	ResetClientTlsKeyWo()
	ResetClientTlsKeyWoVersion()
	ResetConnectionTimeout()
	ResetDenyNullBind()
	ResetDereferenceAliases()
	ResetDiscoverdn()
	ResetEnableSamaccountnameLogin()
	ResetGroupattr()
	ResetGroupdn()
	ResetGroupfilter()
	ResetInsecureTls()
	ResetMaxPageSize()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestTimeout()
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
	ResetUpndomain()
	ResetUrl()
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

// The jsii proxy struct for KerberosAuthBackendLdapConfig
type jsiiProxy_KerberosAuthBackendLdapConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) AnonymousGroupSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) AnonymousGroupSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Binddn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) BinddnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) BindpassWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) BindpassWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) BindpassWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) BindpassWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bindpassWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) CaseSensitiveNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) CaseSensitiveNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsCertWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCertWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsCertWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCertWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsCertWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTlsCertWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsCertWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTlsCertWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTlsKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ClientTlsKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTlsKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ConnectionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ConnectionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DenyNullBind() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DenyNullBindInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DereferenceAliases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DereferenceAliasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dereferenceAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Discoverdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) DiscoverdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) EnableSamaccountnameLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamaccountnameLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) EnableSamaccountnameLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamaccountnameLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Groupattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) GroupattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Groupdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) GroupdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Groupfilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) GroupfilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) MaxPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) MaxPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Starttls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) StarttlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TlsMaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Upndomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UpndomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Userattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UserattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Userdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UserdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) Userfilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userfilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UserfilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userfilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UsernameAsAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameAsAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UsernameAsAliasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usernameAsAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UseTokenGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig) UseTokenGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/resources/kerberos_auth_backend_ldap_config vault_kerberos_auth_backend_ldap_config} Resource.
func NewKerberosAuthBackendLdapConfig(scope constructs.Construct, id *string, config *KerberosAuthBackendLdapConfigConfig) KerberosAuthBackendLdapConfig {
	_init_.Initialize()

	if err := validateNewKerberosAuthBackendLdapConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KerberosAuthBackendLdapConfig{}

	_jsii_.Create(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.12.0/docs/resources/kerberos_auth_backend_ldap_config vault_kerberos_auth_backend_ldap_config} Resource.
func NewKerberosAuthBackendLdapConfig_Override(k KerberosAuthBackendLdapConfig, scope constructs.Construct, id *string, config *KerberosAuthBackendLdapConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetAnonymousGroupSearch(val interface{}) {
	if err := j.validateSetAnonymousGroupSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousGroupSearch",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetBinddn(val *string) {
	if err := j.validateSetBinddnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binddn",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetBindpassWo(val *string) {
	if err := j.validateSetBindpassWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWo",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetBindpassWoVersion(val *float64) {
	if err := j.validateSetBindpassWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpassWoVersion",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetCaseSensitiveNames(val interface{}) {
	if err := j.validateSetCaseSensitiveNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitiveNames",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetClientTlsCertWo(val *string) {
	if err := j.validateSetClientTlsCertWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsCertWo",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetClientTlsCertWoVersion(val *float64) {
	if err := j.validateSetClientTlsCertWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsCertWoVersion",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetClientTlsKeyWo(val *string) {
	if err := j.validateSetClientTlsKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsKeyWo",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetClientTlsKeyWoVersion(val *float64) {
	if err := j.validateSetClientTlsKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetConnectionTimeout(val *float64) {
	if err := j.validateSetConnectionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTimeout",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetDenyNullBind(val interface{}) {
	if err := j.validateSetDenyNullBindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyNullBind",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetDereferenceAliases(val *string) {
	if err := j.validateSetDereferenceAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dereferenceAliases",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetDiscoverdn(val interface{}) {
	if err := j.validateSetDiscoverdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoverdn",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetEnableSamaccountnameLogin(val interface{}) {
	if err := j.validateSetEnableSamaccountnameLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSamaccountnameLogin",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetGroupattr(val *string) {
	if err := j.validateSetGroupattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupattr",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetGroupdn(val *string) {
	if err := j.validateSetGroupdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupdn",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetGroupfilter(val *string) {
	if err := j.validateSetGroupfilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupfilter",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetMaxPageSize(val *float64) {
	if err := j.validateSetMaxPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPageSize",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetRequestTimeout(val *float64) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetStarttls(val interface{}) {
	if err := j.validateSetStarttlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"starttls",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTlsMaxVersion(val *string) {
	if err := j.validateSetTlsMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMaxVersion",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUpndomain(val *string) {
	if err := j.validateSetUpndomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upndomain",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUserattr(val *string) {
	if err := j.validateSetUserattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userattr",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUserdn(val *string) {
	if err := j.validateSetUserdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userdn",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUserfilter(val *string) {
	if err := j.validateSetUserfilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userfilter",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUsernameAsAlias(val interface{}) {
	if err := j.validateSetUsernameAsAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameAsAlias",
		val,
	)
}

func (j *jsiiProxy_KerberosAuthBackendLdapConfig)SetUseTokenGroups(val interface{}) {
	if err := j.validateSetUseTokenGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTokenGroups",
		val,
	)
}

// Generates CDKTN code for importing a KerberosAuthBackendLdapConfig resource upon running "cdktn plan <stack-name>".
func KerberosAuthBackendLdapConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKerberosAuthBackendLdapConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
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
func KerberosAuthBackendLdapConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKerberosAuthBackendLdapConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KerberosAuthBackendLdapConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKerberosAuthBackendLdapConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KerberosAuthBackendLdapConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKerberosAuthBackendLdapConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KerberosAuthBackendLdapConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.kerberosAuthBackendLdapConfig.KerberosAuthBackendLdapConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := k.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := k.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		k,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetAnonymousGroupSearch() {
	_jsii_.InvokeVoid(
		k,
		"resetAnonymousGroupSearch",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetBinddn() {
	_jsii_.InvokeVoid(
		k,
		"resetBinddn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetBindpassWo() {
	_jsii_.InvokeVoid(
		k,
		"resetBindpassWo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetBindpassWoVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetBindpassWoVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetCaseSensitiveNames() {
	_jsii_.InvokeVoid(
		k,
		"resetCaseSensitiveNames",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetCertificate() {
	_jsii_.InvokeVoid(
		k,
		"resetCertificate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetClientTlsCertWo() {
	_jsii_.InvokeVoid(
		k,
		"resetClientTlsCertWo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetClientTlsCertWoVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetClientTlsCertWoVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetClientTlsKeyWo() {
	_jsii_.InvokeVoid(
		k,
		"resetClientTlsKeyWo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetClientTlsKeyWoVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetClientTlsKeyWoVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetConnectionTimeout() {
	_jsii_.InvokeVoid(
		k,
		"resetConnectionTimeout",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetDenyNullBind() {
	_jsii_.InvokeVoid(
		k,
		"resetDenyNullBind",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetDereferenceAliases() {
	_jsii_.InvokeVoid(
		k,
		"resetDereferenceAliases",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetDiscoverdn() {
	_jsii_.InvokeVoid(
		k,
		"resetDiscoverdn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetEnableSamaccountnameLogin() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableSamaccountnameLogin",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetGroupattr() {
	_jsii_.InvokeVoid(
		k,
		"resetGroupattr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetGroupdn() {
	_jsii_.InvokeVoid(
		k,
		"resetGroupdn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetGroupfilter() {
	_jsii_.InvokeVoid(
		k,
		"resetGroupfilter",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		k,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetMaxPageSize() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxPageSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		k,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetStarttls() {
	_jsii_.InvokeVoid(
		k,
		"resetStarttls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTlsMaxVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsMaxVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetTokenType() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUpndomain() {
	_jsii_.InvokeVoid(
		k,
		"resetUpndomain",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUrl() {
	_jsii_.InvokeVoid(
		k,
		"resetUrl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUserattr() {
	_jsii_.InvokeVoid(
		k,
		"resetUserattr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUserdn() {
	_jsii_.InvokeVoid(
		k,
		"resetUserdn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUserfilter() {
	_jsii_.InvokeVoid(
		k,
		"resetUserfilter",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUsernameAsAlias() {
	_jsii_.InvokeVoid(
		k,
		"resetUsernameAsAlias",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ResetUseTokenGroups() {
	_jsii_.InvokeVoid(
		k,
		"resetUseTokenGroups",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KerberosAuthBackendLdapConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

