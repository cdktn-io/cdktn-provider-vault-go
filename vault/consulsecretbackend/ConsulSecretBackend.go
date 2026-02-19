// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package consulsecretbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/consulsecretbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/consul_secret_backend vault_consul_secret_backend}.
type ConsulSecretBackend interface {
	cdktn.TerraformResource
	Accessor() *string
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
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
	Bootstrap() interface{}
	SetBootstrap(val interface{})
	BootstrapInput() interface{}
	CaCert() *string
	SetCaCert(val *string)
	CaCertInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCert() *string
	SetClientCert(val *string)
	ClientCertInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
	ClientKeyWo() *string
	SetClientKeyWo(val *string)
	ClientKeyWoInput() *string
	ClientKeyWoVersion() *float64
	SetClientKeyWoVersion(val *float64)
	ClientKeyWoVersionInput() *float64
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
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	SealWrap() interface{}
	SetSealWrap(val interface{})
	SealWrapInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	TokenWo() *string
	SetTokenWo(val *string)
	TokenWoInput() *string
	TokenWoVersion() *float64
	SetTokenWoVersion(val *float64)
	TokenWoVersionInput() *float64
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
	ResetBootstrap()
	ResetCaCert()
	ResetClientCert()
	ResetClientKey()
	ResetClientKeyWo()
	ResetClientKeyWoVersion()
	ResetDefaultLeaseTtlSeconds()
	ResetDelegatedAuthAccessors()
	ResetDescription()
	ResetDisableRemount()
	ResetExternalEntropyAccess()
	ResetForceNoCache()
	ResetId()
	ResetIdentityTokenKey()
	ResetListingVisibility()
	ResetLocal()
	ResetMaxLeaseTtlSeconds()
	ResetNamespace()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassthroughRequestHeaders()
	ResetPath()
	ResetPluginVersion()
	ResetScheme()
	ResetSealWrap()
	ResetToken()
	ResetTokenWo()
	ResetTokenWoVersion()
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

// The jsii proxy struct for ConsulSecretBackend
type jsiiProxy_ConsulSecretBackend struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ConsulSecretBackend) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AllowedManagedKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AllowedManagedKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedManagedKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AllowedResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AllowedResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Bootstrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) BootstrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) CaCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) CaCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKeyWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKeyWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKeyWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeyWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ClientKeyWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeyWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DefaultLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DefaultLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DelegatedAuthAccessors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DelegatedAuthAccessorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"delegatedAuthAccessorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ExternalEntropyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ExternalEntropyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ForceNoCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ForceNoCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNoCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) IdentityTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) IdentityTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ListingVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) ListingVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) PassthroughRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) PassthroughRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) PluginVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) PluginVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) SealWrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) SealWrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TokenWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TokenWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TokenWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulSecretBackend) TokenWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenWoVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/consul_secret_backend vault_consul_secret_backend} Resource.
func NewConsulSecretBackend(scope constructs.Construct, id *string, config *ConsulSecretBackendConfig) ConsulSecretBackend {
	_init_.Initialize()

	if err := validateNewConsulSecretBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConsulSecretBackend{}

	_jsii_.Create(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/consul_secret_backend vault_consul_secret_backend} Resource.
func NewConsulSecretBackend_Override(c ConsulSecretBackend, scope constructs.Construct, id *string, config *ConsulSecretBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetAllowedManagedKeys(val *[]*string) {
	if err := j.validateSetAllowedManagedKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedManagedKeys",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetAllowedResponseHeaders(val *[]*string) {
	if err := j.validateSetAllowedResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetAuditNonHmacRequestKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacRequestKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetAuditNonHmacResponseKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacResponseKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetBootstrap(val interface{}) {
	if err := j.validateSetBootstrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrap",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetCaCert(val *string) {
	if err := j.validateSetCaCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCert",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetClientCert(val *string) {
	if err := j.validateSetClientCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCert",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetClientKeyWo(val *string) {
	if err := j.validateSetClientKeyWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeyWo",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetClientKeyWoVersion(val *float64) {
	if err := j.validateSetClientKeyWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeyWoVersion",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetDefaultLeaseTtlSeconds(val *float64) {
	if err := j.validateSetDefaultLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetDelegatedAuthAccessors(val *[]*string) {
	if err := j.validateSetDelegatedAuthAccessorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delegatedAuthAccessors",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetExternalEntropyAccess(val interface{}) {
	if err := j.validateSetExternalEntropyAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEntropyAccess",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetForceNoCache(val interface{}) {
	if err := j.validateSetForceNoCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNoCache",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetIdentityTokenKey(val *string) {
	if err := j.validateSetIdentityTokenKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityTokenKey",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetListingVisibility(val *string) {
	if err := j.validateSetListingVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingVisibility",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetMaxLeaseTtlSeconds(val *float64) {
	if err := j.validateSetMaxLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetPassthroughRequestHeaders(val *[]*string) {
	if err := j.validateSetPassthroughRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetPluginVersion(val *string) {
	if err := j.validateSetPluginVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginVersion",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetSealWrap(val interface{}) {
	if err := j.validateSetSealWrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sealWrap",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetTokenWo(val *string) {
	if err := j.validateSetTokenWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenWo",
		val,
	)
}

func (j *jsiiProxy_ConsulSecretBackend)SetTokenWoVersion(val *float64) {
	if err := j.validateSetTokenWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenWoVersion",
		val,
	)
}

// Generates CDKTN code for importing a ConsulSecretBackend resource upon running "cdktn plan <stack-name>".
func ConsulSecretBackend_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateConsulSecretBackend_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
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
func ConsulSecretBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulSecretBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulSecretBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulSecretBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ConsulSecretBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConsulSecretBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConsulSecretBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.consulSecretBackend.ConsulSecretBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetAllowedManagedKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedManagedKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetAllowedResponseHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedResponseHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetBootstrap() {
	_jsii_.InvokeVoid(
		c,
		"resetBootstrap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetCaCert() {
	_jsii_.InvokeVoid(
		c,
		"resetCaCert",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetClientCert() {
	_jsii_.InvokeVoid(
		c,
		"resetClientCert",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetClientKey() {
	_jsii_.InvokeVoid(
		c,
		"resetClientKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetClientKeyWo() {
	_jsii_.InvokeVoid(
		c,
		"resetClientKeyWo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetClientKeyWoVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetClientKeyWoVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetDefaultLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetDelegatedAuthAccessors() {
	_jsii_.InvokeVoid(
		c,
		"resetDelegatedAuthAccessors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetExternalEntropyAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetExternalEntropyAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetForceNoCache() {
	_jsii_.InvokeVoid(
		c,
		"resetForceNoCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetIdentityTokenKey() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentityTokenKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetListingVisibility() {
	_jsii_.InvokeVoid(
		c,
		"resetListingVisibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetLocal() {
	_jsii_.InvokeVoid(
		c,
		"resetLocal",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetPassthroughRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetPassthroughRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetPluginVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetPluginVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetSealWrap() {
	_jsii_.InvokeVoid(
		c,
		"resetSealWrap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetToken() {
	_jsii_.InvokeVoid(
		c,
		"resetToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetTokenWo() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenWo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) ResetTokenWoVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenWoVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConsulSecretBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConsulSecretBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

