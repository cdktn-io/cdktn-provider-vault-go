// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/provider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs vault}.
type VaultProvider interface {
	cdktn.TerraformProvider
	AddAddressToEnv() *string
	SetAddAddressToEnv(val *string)
	AddAddressToEnvInput() *string
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthLogin() interface{}
	SetAuthLogin(val interface{})
	AuthLoginAws() interface{}
	SetAuthLoginAws(val interface{})
	AuthLoginAwsInput() interface{}
	AuthLoginAzure() interface{}
	SetAuthLoginAzure(val interface{})
	AuthLoginAzureInput() interface{}
	AuthLoginCert() interface{}
	SetAuthLoginCert(val interface{})
	AuthLoginCertInput() interface{}
	AuthLoginGcp() interface{}
	SetAuthLoginGcp(val interface{})
	AuthLoginGcpInput() interface{}
	AuthLoginInput() interface{}
	AuthLoginJwt() interface{}
	SetAuthLoginJwt(val interface{})
	AuthLoginJwtInput() interface{}
	AuthLoginKerberos() interface{}
	SetAuthLoginKerberos(val interface{})
	AuthLoginKerberosInput() interface{}
	AuthLoginOci() interface{}
	SetAuthLoginOci(val interface{})
	AuthLoginOciInput() interface{}
	AuthLoginOidc() interface{}
	SetAuthLoginOidc(val interface{})
	AuthLoginOidcInput() interface{}
	AuthLoginRadius() interface{}
	SetAuthLoginRadius(val interface{})
	AuthLoginRadiusInput() interface{}
	AuthLoginTokenFile() interface{}
	SetAuthLoginTokenFile(val interface{})
	AuthLoginTokenFileInput() interface{}
	AuthLoginUserpass() interface{}
	SetAuthLoginUserpass(val interface{})
	AuthLoginUserpassInput() interface{}
	CaCertDir() *string
	SetCaCertDir(val *string)
	CaCertDirInput() *string
	CaCertFile() *string
	SetCaCertFile(val *string)
	CaCertFileInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientAuth() interface{}
	SetClientAuth(val interface{})
	ClientAuthInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Headers() interface{}
	SetHeaders(val interface{})
	HeadersInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesCcc() *float64
	SetMaxRetriesCcc(val *float64)
	MaxRetriesCccInput() *float64
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	SetNamespaceFromToken() interface{}
	SetSetNamespaceFromToken(val interface{})
	SetNamespaceFromTokenInput() interface{}
	SkipChildToken() interface{}
	SetSkipChildToken(val interface{})
	SkipChildTokenInput() interface{}
	SkipGetVaultVersion() interface{}
	SetSkipGetVaultVersion(val interface{})
	SkipGetVaultVersionInput() interface{}
	SkipTlsVerify() interface{}
	SetSkipTlsVerify(val interface{})
	SkipTlsVerifyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TlsServerName() *string
	SetTlsServerName(val *string)
	TlsServerNameInput() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	TokenName() *string
	SetTokenName(val *string)
	TokenNameInput() *string
	VaultVersionOverride() *string
	SetVaultVersionOverride(val *string)
	VaultVersionOverrideInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAddAddressToEnv()
	ResetAddress()
	ResetAlias()
	ResetAuthLogin()
	ResetAuthLoginAws()
	ResetAuthLoginAzure()
	ResetAuthLoginCert()
	ResetAuthLoginGcp()
	ResetAuthLoginJwt()
	ResetAuthLoginKerberos()
	ResetAuthLoginOci()
	ResetAuthLoginOidc()
	ResetAuthLoginRadius()
	ResetAuthLoginTokenFile()
	ResetAuthLoginUserpass()
	ResetCaCertDir()
	ResetCaCertFile()
	ResetClientAuth()
	ResetHeaders()
	ResetMaxLeaseTtlSeconds()
	ResetMaxRetries()
	ResetMaxRetriesCcc()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSetNamespaceFromToken()
	ResetSkipChildToken()
	ResetSkipGetVaultVersion()
	ResetSkipTlsVerify()
	ResetTlsServerName()
	ResetToken()
	ResetTokenName()
	ResetVaultVersionOverride()
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

// The jsii proxy struct for VaultProvider
type jsiiProxy_VaultProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_VaultProvider) AddAddressToEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addAddressToEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AddAddressToEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addAddressToEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginAws() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginAws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginAwsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginAwsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginAzure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginAzure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginAzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginAzureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginCert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginCertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginGcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginGcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginGcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginGcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginJwt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginJwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginJwtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginJwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginKerberos() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginKerberos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginKerberosInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginKerberosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginOci() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginOciInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginOciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginOidc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginOidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginRadius() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginRadius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginRadiusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginRadiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginTokenFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginTokenFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginTokenFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginTokenFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginUserpass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginUserpass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginUserpassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginUserpassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ClientAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ClientAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Headers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesCcc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesCcc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesCccInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesCccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SetNamespaceFromToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setNamespaceFromToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SetNamespaceFromTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setNamespaceFromTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipChildToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipChildToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipChildTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipChildTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipGetVaultVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGetVaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipGetVaultVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGetVaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TlsServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TlsServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) VaultVersionOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultVersionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) VaultVersionOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultVersionOverrideInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs vault} Resource.
func NewVaultProvider(scope constructs.Construct, id *string, config *VaultProviderConfig) VaultProvider {
	_init_.Initialize()

	if err := validateNewVaultProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultProvider{}

	_jsii_.Create(
		"@cdktn/provider-vault.provider.VaultProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs vault} Resource.
func NewVaultProvider_Override(v VaultProvider, scope constructs.Construct, id *string, config *VaultProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.provider.VaultProvider",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultProvider)SetAddAddressToEnv(val *string) {
	_jsii_.Set(
		j,
		"addAddressToEnv",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLogin(val interface{}) {
	if err := j.validateSetAuthLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLogin",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginAws(val interface{}) {
	if err := j.validateSetAuthLoginAwsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginAws",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginAzure(val interface{}) {
	if err := j.validateSetAuthLoginAzureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginAzure",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginCert(val interface{}) {
	if err := j.validateSetAuthLoginCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginCert",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginGcp(val interface{}) {
	if err := j.validateSetAuthLoginGcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginGcp",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginJwt(val interface{}) {
	if err := j.validateSetAuthLoginJwtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginJwt",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginKerberos(val interface{}) {
	if err := j.validateSetAuthLoginKerberosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginKerberos",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginOci(val interface{}) {
	if err := j.validateSetAuthLoginOciParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginOci",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginOidc(val interface{}) {
	if err := j.validateSetAuthLoginOidcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginOidc",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginRadius(val interface{}) {
	if err := j.validateSetAuthLoginRadiusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginRadius",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginTokenFile(val interface{}) {
	if err := j.validateSetAuthLoginTokenFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginTokenFile",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLoginUserpass(val interface{}) {
	if err := j.validateSetAuthLoginUserpassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLoginUserpass",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetCaCertDir(val *string) {
	_jsii_.Set(
		j,
		"caCertDir",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetCaCertFile(val *string) {
	_jsii_.Set(
		j,
		"caCertFile",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetClientAuth(val interface{}) {
	if err := j.validateSetClientAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuth",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetHeaders(val interface{}) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxLeaseTtlSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxRetriesCcc(val *float64) {
	_jsii_.Set(
		j,
		"maxRetriesCcc",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSetNamespaceFromToken(val interface{}) {
	if err := j.validateSetSetNamespaceFromTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setNamespaceFromToken",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSkipChildToken(val interface{}) {
	if err := j.validateSetSkipChildTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipChildToken",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSkipGetVaultVersion(val interface{}) {
	if err := j.validateSetSkipGetVaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGetVaultVersion",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSkipTlsVerify(val interface{}) {
	if err := j.validateSetSkipTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipTlsVerify",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetTlsServerName(val *string) {
	_jsii_.Set(
		j,
		"tlsServerName",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetTokenName(val *string) {
	_jsii_.Set(
		j,
		"tokenName",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetVaultVersionOverride(val *string) {
	_jsii_.Set(
		j,
		"vaultVersionOverride",
		val,
	)
}

// Generates CDKTN code for importing a VaultProvider resource upon running "cdktn plan <stack-name>".
func VaultProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVaultProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.provider.VaultProvider",
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
func VaultProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.provider.VaultProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.provider.VaultProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VaultProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.provider.VaultProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.provider.VaultProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultProvider) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultProvider) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultProvider) ResetAddAddressToEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetAddAddressToEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		v,
		"resetAlias",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLogin() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLogin",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginAws() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginAws",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginAzure() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginAzure",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginCert() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginCert",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginGcp() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginGcp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginJwt() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginJwt",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginKerberos() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginKerberos",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginOci() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginOci",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginOidc() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginOidc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginRadius() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginRadius",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginTokenFile() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginTokenFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLoginUserpass() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLoginUserpass",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetCaCertDir() {
	_jsii_.InvokeVoid(
		v,
		"resetCaCertDir",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetCaCertFile() {
	_jsii_.InvokeVoid(
		v,
		"resetCaCertFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetClientAuth() {
	_jsii_.InvokeVoid(
		v,
		"resetClientAuth",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetHeaders() {
	_jsii_.InvokeVoid(
		v,
		"resetHeaders",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxRetriesCcc() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxRetriesCcc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetNamespace() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespace",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSetNamespaceFromToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSetNamespaceFromToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSkipChildToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSkipChildToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSkipGetVaultVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetSkipGetVaultVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSkipTlsVerify() {
	_jsii_.InvokeVoid(
		v,
		"resetSkipTlsVerify",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetTlsServerName() {
	_jsii_.InvokeVoid(
		v,
		"resetTlsServerName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetToken() {
	_jsii_.InvokeVoid(
		v,
		"resetToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetTokenName() {
	_jsii_.InvokeVoid(
		v,
		"resetTokenName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetVaultVersionOverride() {
	_jsii_.InvokeVoid(
		v,
		"resetVaultVersionOverride",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

