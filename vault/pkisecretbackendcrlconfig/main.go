// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendcrlconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		reflect.TypeOf((*PkiSecretBackendCrlConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoRebuild", GoGetter: "AutoRebuild"},
			_jsii_.MemberProperty{JsiiProperty: "autoRebuildGracePeriod", GoGetter: "AutoRebuildGracePeriod"},
			_jsii_.MemberProperty{JsiiProperty: "autoRebuildGracePeriodInput", GoGetter: "AutoRebuildGracePeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoRebuildInput", GoGetter: "AutoRebuildInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "crossClusterRevocation", GoGetter: "CrossClusterRevocation"},
			_jsii_.MemberProperty{JsiiProperty: "crossClusterRevocationInput", GoGetter: "CrossClusterRevocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "deltaRebuildInterval", GoGetter: "DeltaRebuildInterval"},
			_jsii_.MemberProperty{JsiiProperty: "deltaRebuildIntervalInput", GoGetter: "DeltaRebuildIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disable", GoGetter: "Disable"},
			_jsii_.MemberProperty{JsiiProperty: "disableInput", GoGetter: "DisableInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableDelta", GoGetter: "EnableDelta"},
			_jsii_.MemberProperty{JsiiProperty: "enableDeltaInput", GoGetter: "EnableDeltaInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiry", GoGetter: "Expiry"},
			_jsii_.MemberProperty{JsiiProperty: "expiryInput", GoGetter: "ExpiryInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxCrlEntries", GoGetter: "MaxCrlEntries"},
			_jsii_.MemberProperty{JsiiProperty: "maxCrlEntriesInput", GoGetter: "MaxCrlEntriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ocspDisable", GoGetter: "OcspDisable"},
			_jsii_.MemberProperty{JsiiProperty: "ocspDisableInput", GoGetter: "OcspDisableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ocspExpiry", GoGetter: "OcspExpiry"},
			_jsii_.MemberProperty{JsiiProperty: "ocspExpiryInput", GoGetter: "OcspExpiryInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRebuild", GoMethod: "ResetAutoRebuild"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRebuildGracePeriod", GoMethod: "ResetAutoRebuildGracePeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossClusterRevocation", GoMethod: "ResetCrossClusterRevocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeltaRebuildInterval", GoMethod: "ResetDeltaRebuildInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisable", GoMethod: "ResetDisable"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableDelta", GoMethod: "ResetEnableDelta"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpiry", GoMethod: "ResetExpiry"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxCrlEntries", GoMethod: "ResetMaxCrlEntries"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOcspDisable", GoMethod: "ResetOcspDisable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOcspExpiry", GoMethod: "ResetOcspExpiry"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnifiedCrl", GoMethod: "ResetUnifiedCrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnifiedCrlOnExistingPaths", GoMethod: "ResetUnifiedCrlOnExistingPaths"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedCrl", GoGetter: "UnifiedCrl"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedCrlInput", GoGetter: "UnifiedCrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedCrlOnExistingPaths", GoGetter: "UnifiedCrlOnExistingPaths"},
			_jsii_.MemberProperty{JsiiProperty: "unifiedCrlOnExistingPathsInput", GoGetter: "UnifiedCrlOnExistingPathsInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendCrlConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfigConfig",
		reflect.TypeOf((*PkiSecretBackendCrlConfigConfig)(nil)).Elem(),
	)
}
