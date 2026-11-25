// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformtransformationtokenizationimportedkey

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKey",
		reflect.TypeOf((*TransformTransformationTokenizationImportedKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRoles", GoGetter: "AllowedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRolesInput", GoGetter: "AllowedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowRotation", GoGetter: "AllowRotation"},
			_jsii_.MemberProperty{JsiiProperty: "allowRotationInput", GoGetter: "AllowRotationInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotatePeriod", GoGetter: "AutoRotatePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "autoRotatePeriodInput", GoGetter: "AutoRotatePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "ciphertextWo", GoGetter: "CiphertextWo"},
			_jsii_.MemberProperty{JsiiProperty: "ciphertextWoInput", GoGetter: "CiphertextWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "ciphertextWoVersion", GoGetter: "CiphertextWoVersion"},
			_jsii_.MemberProperty{JsiiProperty: "ciphertextWoVersionInput", GoGetter: "CiphertextWoVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "deletionAllowed", GoGetter: "DeletionAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "deletionAllowedInput", GoGetter: "DeletionAllowedInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "hashFunction", GoGetter: "HashFunction"},
			_jsii_.MemberProperty{JsiiProperty: "hashFunctionInput", GoGetter: "HashFunctionInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestVersion", GoGetter: "LatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mappingMode", GoGetter: "MappingMode"},
			_jsii_.MemberProperty{JsiiProperty: "mappingModeInput", GoGetter: "MappingModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minAvailableVersion", GoGetter: "MinAvailableVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minDecryptionVersion", GoGetter: "MinDecryptionVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minDecryptionVersionInput", GoGetter: "MinDecryptionVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "mount", GoGetter: "Mount"},
			_jsii_.MemberProperty{JsiiProperty: "mountInput", GoGetter: "MountInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedRoles", GoMethod: "ResetAllowedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRotation", GoMethod: "ResetAllowRotation"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRotatePeriod", GoMethod: "ResetAutoRotatePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetCiphertextWo", GoMethod: "ResetCiphertextWo"},
			_jsii_.MemberMethod{JsiiMethod: "resetCiphertextWoVersion", GoMethod: "ResetCiphertextWoVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeletionAllowed", GoMethod: "ResetDeletionAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetHashFunction", GoMethod: "ResetHashFunction"},
			_jsii_.MemberMethod{JsiiMethod: "resetMappingMode", GoMethod: "ResetMappingMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinDecryptionVersion", GoMethod: "ResetMinDecryptionVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStores", GoMethod: "ResetStores"},
			_jsii_.MemberProperty{JsiiProperty: "stores", GoGetter: "Stores"},
			_jsii_.MemberProperty{JsiiProperty: "storesInput", GoGetter: "StoresInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_TransformTransformationTokenizationImportedKey{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.transformTransformationTokenizationImportedKey.TransformTransformationTokenizationImportedKeyConfig",
		reflect.TypeOf((*TransformTransformationTokenizationImportedKeyConfig)(nil)).Elem(),
	)
}
