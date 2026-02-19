// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.provider.VaultProvider",
		reflect.TypeOf((*VaultProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addAddressToEnv", GoGetter: "AddAddressToEnv"},
			_jsii_.MemberProperty{JsiiProperty: "addAddressToEnvInput", GoGetter: "AddAddressToEnvInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "address", GoGetter: "Address"},
			_jsii_.MemberProperty{JsiiProperty: "addressInput", GoGetter: "AddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLogin", GoGetter: "AuthLogin"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginAws", GoGetter: "AuthLoginAws"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginAwsInput", GoGetter: "AuthLoginAwsInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginAzure", GoGetter: "AuthLoginAzure"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginAzureInput", GoGetter: "AuthLoginAzureInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginCert", GoGetter: "AuthLoginCert"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginCertInput", GoGetter: "AuthLoginCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginGcp", GoGetter: "AuthLoginGcp"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginGcpInput", GoGetter: "AuthLoginGcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginInput", GoGetter: "AuthLoginInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginJwt", GoGetter: "AuthLoginJwt"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginJwtInput", GoGetter: "AuthLoginJwtInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginKerberos", GoGetter: "AuthLoginKerberos"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginKerberosInput", GoGetter: "AuthLoginKerberosInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginOci", GoGetter: "AuthLoginOci"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginOciInput", GoGetter: "AuthLoginOciInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginOidc", GoGetter: "AuthLoginOidc"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginOidcInput", GoGetter: "AuthLoginOidcInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginRadius", GoGetter: "AuthLoginRadius"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginRadiusInput", GoGetter: "AuthLoginRadiusInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginTokenFile", GoGetter: "AuthLoginTokenFile"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginTokenFileInput", GoGetter: "AuthLoginTokenFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginUserpass", GoGetter: "AuthLoginUserpass"},
			_jsii_.MemberProperty{JsiiProperty: "authLoginUserpassInput", GoGetter: "AuthLoginUserpassInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertDir", GoGetter: "CaCertDir"},
			_jsii_.MemberProperty{JsiiProperty: "caCertDirInput", GoGetter: "CaCertDirInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertFile", GoGetter: "CaCertFile"},
			_jsii_.MemberProperty{JsiiProperty: "caCertFileInput", GoGetter: "CaCertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuth", GoGetter: "ClientAuth"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthInput", GoGetter: "ClientAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSeconds", GoGetter: "MaxLeaseTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlSecondsInput", GoGetter: "MaxLeaseTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesCcc", GoGetter: "MaxRetriesCcc"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesCccInput", GoGetter: "MaxRetriesCccInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddAddressToEnv", GoMethod: "ResetAddAddressToEnv"},
			_jsii_.MemberMethod{JsiiMethod: "resetAddress", GoMethod: "ResetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLogin", GoMethod: "ResetAuthLogin"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginAws", GoMethod: "ResetAuthLoginAws"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginAzure", GoMethod: "ResetAuthLoginAzure"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginCert", GoMethod: "ResetAuthLoginCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginGcp", GoMethod: "ResetAuthLoginGcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginJwt", GoMethod: "ResetAuthLoginJwt"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginKerberos", GoMethod: "ResetAuthLoginKerberos"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginOci", GoMethod: "ResetAuthLoginOci"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginOidc", GoMethod: "ResetAuthLoginOidc"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginRadius", GoMethod: "ResetAuthLoginRadius"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginTokenFile", GoMethod: "ResetAuthLoginTokenFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthLoginUserpass", GoMethod: "ResetAuthLoginUserpass"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertDir", GoMethod: "ResetCaCertDir"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaCertFile", GoMethod: "ResetCaCertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientAuth", GoMethod: "ResetClientAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtlSeconds", GoMethod: "ResetMaxLeaseTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetriesCcc", GoMethod: "ResetMaxRetriesCcc"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetNamespaceFromToken", GoMethod: "ResetSetNamespaceFromToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipChildToken", GoMethod: "ResetSkipChildToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGetVaultVersion", GoMethod: "ResetSkipGetVaultVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipTlsVerify", GoMethod: "ResetSkipTlsVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsServerName", GoMethod: "ResetTlsServerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenName", GoMethod: "ResetTokenName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVaultVersionOverride", GoMethod: "ResetVaultVersionOverride"},
			_jsii_.MemberProperty{JsiiProperty: "setNamespaceFromToken", GoGetter: "SetNamespaceFromToken"},
			_jsii_.MemberProperty{JsiiProperty: "setNamespaceFromTokenInput", GoGetter: "SetNamespaceFromTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipChildToken", GoGetter: "SkipChildToken"},
			_jsii_.MemberProperty{JsiiProperty: "skipChildTokenInput", GoGetter: "SkipChildTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipGetVaultVersion", GoGetter: "SkipGetVaultVersion"},
			_jsii_.MemberProperty{JsiiProperty: "skipGetVaultVersionInput", GoGetter: "SkipGetVaultVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipTlsVerify", GoGetter: "SkipTlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "skipTlsVerifyInput", GoGetter: "SkipTlsVerifyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerName", GoGetter: "TlsServerName"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerNameInput", GoGetter: "TlsServerNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenName", GoGetter: "TokenName"},
			_jsii_.MemberProperty{JsiiProperty: "tokenNameInput", GoGetter: "TokenNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vaultVersionOverride", GoGetter: "VaultVersionOverride"},
			_jsii_.MemberProperty{JsiiProperty: "vaultVersionOverrideInput", GoGetter: "VaultVersionOverrideInput"},
		},
		func() interface{} {
			j := jsiiProxy_VaultProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLogin",
		reflect.TypeOf((*VaultProviderAuthLogin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginAws",
		reflect.TypeOf((*VaultProviderAuthLoginAws)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginAzure",
		reflect.TypeOf((*VaultProviderAuthLoginAzure)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginCert",
		reflect.TypeOf((*VaultProviderAuthLoginCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginGcp",
		reflect.TypeOf((*VaultProviderAuthLoginGcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginJwt",
		reflect.TypeOf((*VaultProviderAuthLoginJwt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginKerberos",
		reflect.TypeOf((*VaultProviderAuthLoginKerberos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginOci",
		reflect.TypeOf((*VaultProviderAuthLoginOci)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginOidc",
		reflect.TypeOf((*VaultProviderAuthLoginOidc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginRadius",
		reflect.TypeOf((*VaultProviderAuthLoginRadius)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginTokenFile",
		reflect.TypeOf((*VaultProviderAuthLoginTokenFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderAuthLoginUserpass",
		reflect.TypeOf((*VaultProviderAuthLoginUserpass)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderClientAuth",
		reflect.TypeOf((*VaultProviderClientAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderConfig",
		reflect.TypeOf((*VaultProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.provider.VaultProviderHeaders",
		reflect.TypeOf((*VaultProviderHeaders)(nil)).Elem(),
	)
}
