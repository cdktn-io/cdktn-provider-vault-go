// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultgcpoauth2accesstoken

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.ephemeralVaultGcpOauth2AccessToken.EphemeralVaultGcpOauth2AccessToken",
		reflect.TypeOf((*EphemeralVaultGcpOauth2AccessToken)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "impersonatedAccount", GoGetter: "ImpersonatedAccount"},
			_jsii_.MemberProperty{JsiiProperty: "impersonatedAccountInput", GoGetter: "ImpersonatedAccountInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "leaseDuration", GoGetter: "LeaseDuration"},
			_jsii_.MemberProperty{JsiiProperty: "leaseId", GoGetter: "LeaseId"},
			_jsii_.MemberProperty{JsiiProperty: "leaseRenewable", GoGetter: "LeaseRenewable"},
			_jsii_.MemberProperty{JsiiProperty: "leaseStartTime", GoGetter: "LeaseStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "mount", GoGetter: "Mount"},
			_jsii_.MemberProperty{JsiiProperty: "mountId", GoGetter: "MountId"},
			_jsii_.MemberProperty{JsiiProperty: "mountIdInput", GoGetter: "MountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "mountInput", GoGetter: "MountInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetImpersonatedAccount", GoMethod: "ResetImpersonatedAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountId", GoMethod: "ResetMountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleset", GoMethod: "ResetRoleset"},
			_jsii_.MemberMethod{JsiiMethod: "resetStaticAccount", GoMethod: "ResetStaticAccount"},
			_jsii_.MemberProperty{JsiiProperty: "roleset", GoGetter: "Roleset"},
			_jsii_.MemberProperty{JsiiProperty: "rolesetInput", GoGetter: "RolesetInput"},
			_jsii_.MemberProperty{JsiiProperty: "staticAccount", GoGetter: "StaticAccount"},
			_jsii_.MemberProperty{JsiiProperty: "staticAccountInput", GoGetter: "StaticAccountInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenTtl", GoGetter: "TokenTtl"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_EphemeralVaultGcpOauth2AccessToken{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformEphemeralResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.ephemeralVaultGcpOauth2AccessToken.EphemeralVaultGcpOauth2AccessTokenConfig",
		reflect.TypeOf((*EphemeralVaultGcpOauth2AccessTokenConfig)(nil)).Elem(),
	)
}
