// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cfauthbackendconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.cfAuthBackendConfig.CfAuthBackendConfig",
		reflect.TypeOf((*CfAuthBackendConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cfApiAddr", GoGetter: "CfApiAddr"},
			_jsii_.MemberProperty{JsiiProperty: "cfApiAddrInput", GoGetter: "CfApiAddrInput"},
			_jsii_.MemberProperty{JsiiProperty: "cfApiTrustedCertificates", GoGetter: "CfApiTrustedCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "cfApiTrustedCertificatesInput", GoGetter: "CfApiTrustedCertificatesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cfPasswordWo", GoGetter: "CfPasswordWo"},
			_jsii_.MemberProperty{JsiiProperty: "cfPasswordWoInput", GoGetter: "CfPasswordWoInput"},
			_jsii_.MemberProperty{JsiiProperty: "cfPasswordWoVersion", GoGetter: "CfPasswordWoVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfPasswordWoVersionInput", GoGetter: "CfPasswordWoVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cfTimeout", GoGetter: "CfTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "cfTimeoutInput", GoGetter: "CfTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "cfUsername", GoGetter: "CfUsername"},
			_jsii_.MemberProperty{JsiiProperty: "cfUsernameInput", GoGetter: "CfUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "identityCaCertificates", GoGetter: "IdentityCaCertificates"},
			_jsii_.MemberProperty{JsiiProperty: "identityCaCertificatesInput", GoGetter: "IdentityCaCertificatesInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loginMaxSecondsNotAfter", GoGetter: "LoginMaxSecondsNotAfter"},
			_jsii_.MemberProperty{JsiiProperty: "loginMaxSecondsNotAfterInput", GoGetter: "LoginMaxSecondsNotAfterInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginMaxSecondsNotBefore", GoGetter: "LoginMaxSecondsNotBefore"},
			_jsii_.MemberProperty{JsiiProperty: "loginMaxSecondsNotBeforeInput", GoGetter: "LoginMaxSecondsNotBeforeInput"},
			_jsii_.MemberProperty{JsiiProperty: "mount", GoGetter: "Mount"},
			_jsii_.MemberProperty{JsiiProperty: "mountInput", GoGetter: "MountInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCfApiTrustedCertificates", GoMethod: "ResetCfApiTrustedCertificates"},
			_jsii_.MemberMethod{JsiiMethod: "resetCfTimeout", GoMethod: "ResetCfTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMaxSecondsNotAfter", GoMethod: "ResetLoginMaxSecondsNotAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMaxSecondsNotBefore", GoMethod: "ResetLoginMaxSecondsNotBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
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
			j := jsiiProxy_CfAuthBackendConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.cfAuthBackendConfig.CfAuthBackendConfigConfig",
		reflect.TypeOf((*CfAuthBackendConfigConfig)(nil)).Elem(),
	)
}
