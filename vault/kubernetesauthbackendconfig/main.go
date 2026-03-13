// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesauthbackendconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfig",
		reflect.TypeOf((*KubernetesAuthBackendConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableIssValidation", GoGetter: "DisableIssValidation"},
			_jsii_.MemberProperty{JsiiProperty: "disableIssValidationInput", GoGetter: "DisableIssValidationInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableLocalCaJwt", GoGetter: "DisableLocalCaJwt"},
			_jsii_.MemberProperty{JsiiProperty: "disableLocalCaJwtInput", GoGetter: "DisableLocalCaJwtInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuer", GoGetter: "Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "issuerInput", GoGetter: "IssuerInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesCaCert", GoGetter: "KubernetesCaCert"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesCaCertInput", GoGetter: "KubernetesCaCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesHost", GoGetter: "KubernetesHost"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesHostInput", GoGetter: "KubernetesHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pemKeys", GoGetter: "PemKeys"},
			_jsii_.MemberProperty{JsiiProperty: "pemKeysInput", GoGetter: "PemKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackend", GoMethod: "ResetBackend"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableIssValidation", GoMethod: "ResetDisableIssValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableLocalCaJwt", GoMethod: "ResetDisableLocalCaJwt"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuer", GoMethod: "ResetIssuer"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesCaCert", GoMethod: "ResetKubernetesCaCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPemKeys", GoMethod: "ResetPemKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenReviewerJwt", GoMethod: "ResetTokenReviewerJwt"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenReviewerJwtWo", GoMethod: "ResetTokenReviewerJwtWo"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenReviewerJwtWoVersion", GoMethod: "ResetTokenReviewerJwtWoVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAnnotationsAsAliasMetadata", GoMethod: "ResetUseAnnotationsAsAliasMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwt", GoGetter: "TokenReviewerJwt"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwtInput", GoGetter: "TokenReviewerJwtInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwtWo", GoGetter: "TokenReviewerJwtWo"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwtWoInput", GoGetter: "TokenReviewerJwtWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwtWoVersion", GoGetter: "TokenReviewerJwtWoVersion"},
			_jsii_.MemberProperty{JsiiProperty: "tokenReviewerJwtWoVersionInput", GoGetter: "TokenReviewerJwtWoVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useAnnotationsAsAliasMetadata", GoGetter: "UseAnnotationsAsAliasMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "useAnnotationsAsAliasMetadataInput", GoGetter: "UseAnnotationsAsAliasMetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesAuthBackendConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.kubernetesAuthBackendConfig.KubernetesAuthBackendConfigConfig",
		reflect.TypeOf((*KubernetesAuthBackendConfigConfig)(nil)).Elem(),
	)
}
