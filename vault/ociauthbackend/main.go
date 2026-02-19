// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ociauthbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.ociAuthBackend.OciAuthBackend",
		reflect.TypeOf((*OciAuthBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessor", GoGetter: "Accessor"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableAutomatedRotation", GoGetter: "DisableAutomatedRotation"},
			_jsii_.MemberProperty{JsiiProperty: "disableAutomatedRotationInput", GoGetter: "DisableAutomatedRotationInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemount", GoGetter: "DisableRemount"},
			_jsii_.MemberProperty{JsiiProperty: "disableRemountInput", GoGetter: "DisableRemountInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "homeTenancyId", GoGetter: "HomeTenancyId"},
			_jsii_.MemberProperty{JsiiProperty: "homeTenancyIdInput", GoGetter: "HomeTenancyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTune", GoMethod: "PutTune"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableAutomatedRotation", GoMethod: "ResetDisableAutomatedRotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRemount", GoMethod: "ResetDisableRemount"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationPeriod", GoMethod: "ResetRotationPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationSchedule", GoMethod: "ResetRotationSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationWindow", GoMethod: "ResetRotationWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTune", GoMethod: "ResetTune"},
			_jsii_.MemberProperty{JsiiProperty: "rotationPeriod", GoGetter: "RotationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "rotationPeriodInput", GoGetter: "RotationPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationSchedule", GoGetter: "RotationSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "rotationScheduleInput", GoGetter: "RotationScheduleInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationWindow", GoGetter: "RotationWindow"},
			_jsii_.MemberProperty{JsiiProperty: "rotationWindowInput", GoGetter: "RotationWindowInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tune", GoGetter: "Tune"},
			_jsii_.MemberProperty{JsiiProperty: "tuneInput", GoGetter: "TuneInput"},
		},
		func() interface{} {
			j := jsiiProxy_OciAuthBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.ociAuthBackend.OciAuthBackendConfig",
		reflect.TypeOf((*OciAuthBackendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.ociAuthBackend.OciAuthBackendTune",
		reflect.TypeOf((*OciAuthBackendTune)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.ociAuthBackend.OciAuthBackendTuneList",
		reflect.TypeOf((*OciAuthBackendTuneList)(nil)).Elem(),
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
			j := jsiiProxy_OciAuthBackendTuneList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.ociAuthBackend.OciAuthBackendTuneOutputReference",
		reflect.TypeOf((*OciAuthBackendTuneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedResponseHeaders", GoGetter: "AllowedResponseHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "allowedResponseHeadersInput", GoGetter: "AllowedResponseHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeys", GoGetter: "AuditNonHmacRequestKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacRequestKeysInput", GoGetter: "AuditNonHmacRequestKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeys", GoGetter: "AuditNonHmacResponseKeys"},
			_jsii_.MemberProperty{JsiiProperty: "auditNonHmacResponseKeysInput", GoGetter: "AuditNonHmacResponseKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtl", GoGetter: "DefaultLeaseTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultLeaseTtlInput", GoGetter: "DefaultLeaseTtlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "listingVisibility", GoGetter: "ListingVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "listingVisibilityInput", GoGetter: "ListingVisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtl", GoGetter: "MaxLeaseTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxLeaseTtlInput", GoGetter: "MaxLeaseTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughRequestHeaders", GoGetter: "PassthroughRequestHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughRequestHeadersInput", GoGetter: "PassthroughRequestHeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedResponseHeaders", GoMethod: "ResetAllowedResponseHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacRequestKeys", GoMethod: "ResetAuditNonHmacRequestKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuditNonHmacResponseKeys", GoMethod: "ResetAuditNonHmacResponseKeys"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultLeaseTtl", GoMethod: "ResetDefaultLeaseTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetListingVisibility", GoMethod: "ResetListingVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLeaseTtl", GoMethod: "ResetMaxLeaseTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassthroughRequestHeaders", GoMethod: "ResetPassthroughRequestHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenType", GoMethod: "ResetTokenType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenType", GoGetter: "TokenType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTypeInput", GoGetter: "TokenTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_OciAuthBackendTuneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
