// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2",
		reflect.TypeOf((*KvSecretV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cas", GoGetter: "Cas"},
			_jsii_.MemberProperty{JsiiProperty: "casInput", GoGetter: "CasInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customMetadata", GoGetter: "CustomMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "customMetadataInput", GoGetter: "CustomMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "dataJson", GoGetter: "DataJson"},
			_jsii_.MemberProperty{JsiiProperty: "dataJsonInput", GoGetter: "DataJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataJsonWo", GoGetter: "DataJsonWo"},
			_jsii_.MemberProperty{JsiiProperty: "dataJsonWoInput", GoGetter: "DataJsonWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataJsonWoVersion", GoGetter: "DataJsonWoVersion"},
			_jsii_.MemberProperty{JsiiProperty: "dataJsonWoVersionInput", GoGetter: "DataJsonWoVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAllVersions", GoGetter: "DeleteAllVersions"},
			_jsii_.MemberProperty{JsiiProperty: "deleteAllVersionsInput", GoGetter: "DeleteAllVersionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableRead", GoGetter: "DisableRead"},
			_jsii_.MemberProperty{JsiiProperty: "disableReadInput", GoGetter: "DisableReadInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
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
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "optionsInput", GoGetter: "OptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomMetadata", GoMethod: "PutCustomMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCas", GoMethod: "ResetCas"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomMetadata", GoMethod: "ResetCustomMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataJson", GoMethod: "ResetDataJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataJsonWo", GoMethod: "ResetDataJsonWo"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataJsonWoVersion", GoMethod: "ResetDataJsonWoVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteAllVersions", GoMethod: "ResetDeleteAllVersions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRead", GoMethod: "ResetDisableRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptions", GoMethod: "ResetOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
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
			j := jsiiProxy_KvSecretV2{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2Config",
		reflect.TypeOf((*KvSecretV2Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2CustomMetadata",
		reflect.TypeOf((*KvSecretV2CustomMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2CustomMetadataOutputReference",
		reflect.TypeOf((*KvSecretV2CustomMetadataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "casRequired", GoGetter: "CasRequired"},
			_jsii_.MemberProperty{JsiiProperty: "casRequiredInput", GoGetter: "CasRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "dataInput", GoGetter: "DataInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteVersionAfter", GoGetter: "DeleteVersionAfter"},
			_jsii_.MemberProperty{JsiiProperty: "deleteVersionAfterInput", GoGetter: "DeleteVersionAfterInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxVersions", GoGetter: "MaxVersions"},
			_jsii_.MemberProperty{JsiiProperty: "maxVersionsInput", GoGetter: "MaxVersionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCasRequired", GoMethod: "ResetCasRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetData", GoMethod: "ResetData"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteVersionAfter", GoMethod: "ResetDeleteVersionAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxVersions", GoMethod: "ResetMaxVersions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KvSecretV2CustomMetadataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
