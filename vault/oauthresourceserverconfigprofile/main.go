// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthresourceserverconfigprofile

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfile",
		reflect.TypeOf((*OauthResourceServerConfigProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "audiences", GoGetter: "Audiences"},
			_jsii_.MemberProperty{JsiiProperty: "audiencesInput", GoGetter: "AudiencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clockSkewLeeway", GoGetter: "ClockSkewLeeway"},
			_jsii_.MemberProperty{JsiiProperty: "clockSkewLeewayInput", GoGetter: "ClockSkewLeewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerId", GoGetter: "IssuerId"},
			_jsii_.MemberProperty{JsiiProperty: "issuerIdInput", GoGetter: "IssuerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwksCaPem", GoGetter: "JwksCaPem"},
			_jsii_.MemberProperty{JsiiProperty: "jwksCaPemInput", GoGetter: "JwksCaPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUri", GoGetter: "JwksUri"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUriInput", GoGetter: "JwksUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwtType", GoGetter: "JwtType"},
			_jsii_.MemberProperty{JsiiProperty: "jwtTypeInput", GoGetter: "JwtTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "noDefaultPolicy", GoGetter: "NoDefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "noDefaultPolicyInput", GoGetter: "NoDefaultPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "optionalAuthorizationDetails", GoGetter: "OptionalAuthorizationDetails"},
			_jsii_.MemberProperty{JsiiProperty: "optionalAuthorizationDetailsInput", GoGetter: "OptionalAuthorizationDetailsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "profileName", GoGetter: "ProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "profileNameInput", GoGetter: "ProfileNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeys", GoGetter: "PublicKeys"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeysInput", GoGetter: "PublicKeysInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPublicKeys", GoMethod: "PutPublicKeys"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudiences", GoMethod: "ResetAudiences"},
			_jsii_.MemberMethod{JsiiMethod: "resetClockSkewLeeway", GoMethod: "ResetClockSkewLeeway"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwksCaPem", GoMethod: "ResetJwksCaPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwksUri", GoMethod: "ResetJwksUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetJwtType", GoMethod: "ResetJwtType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoDefaultPolicy", GoMethod: "ResetNoDefaultPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalAuthorizationDetails", GoMethod: "ResetOptionalAuthorizationDetails"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicKeys", GoMethod: "ResetPublicKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportedAlgorithms", GoMethod: "ResetSupportedAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseJwks", GoMethod: "ResetUseJwks"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserClaim", GoMethod: "ResetUserClaim"},
			_jsii_.MemberProperty{JsiiProperty: "supportedAlgorithms", GoGetter: "SupportedAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "supportedAlgorithmsInput", GoGetter: "SupportedAlgorithmsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useJwks", GoGetter: "UseJwks"},
			_jsii_.MemberProperty{JsiiProperty: "useJwksInput", GoGetter: "UseJwksInput"},
			_jsii_.MemberProperty{JsiiProperty: "userClaim", GoGetter: "UserClaim"},
			_jsii_.MemberProperty{JsiiProperty: "userClaimInput", GoGetter: "UserClaimInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OauthResourceServerConfigProfile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfileConfig",
		reflect.TypeOf((*OauthResourceServerConfigProfileConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfilePublicKeys",
		reflect.TypeOf((*OauthResourceServerConfigProfilePublicKeys)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfilePublicKeysList",
		reflect.TypeOf((*OauthResourceServerConfigProfilePublicKeysList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_OauthResourceServerConfigProfilePublicKeysList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.oauthResourceServerConfigProfile.OauthResourceServerConfigProfilePublicKeysOutputReference",
		reflect.TypeOf((*OauthResourceServerConfigProfilePublicKeysOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "keyIdInput", GoGetter: "KeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pem", GoGetter: "Pem"},
			_jsii_.MemberProperty{JsiiProperty: "pemInput", GoGetter: "PemInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OauthResourceServerConfigProfilePublicKeysOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
