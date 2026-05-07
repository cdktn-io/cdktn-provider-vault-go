// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigacme

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.pkiSecretBackendConfigAcme.PkiSecretBackendConfigAcme",
		reflect.TypeOf((*PkiSecretBackendConfigAcme)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedIssuers", GoGetter: "AllowedIssuers"},
			_jsii_.MemberProperty{JsiiProperty: "allowedIssuersInput", GoGetter: "AllowedIssuersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRoles", GoGetter: "AllowedRoles"},
			_jsii_.MemberProperty{JsiiProperty: "allowedRolesInput", GoGetter: "AllowedRolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowRoleExtKeyUsage", GoGetter: "AllowRoleExtKeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "allowRoleExtKeyUsageInput", GoGetter: "AllowRoleExtKeyUsageInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "challengeExcludedIpRanges", GoGetter: "ChallengeExcludedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "challengeExcludedIpRangesInput", GoGetter: "ChallengeExcludedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "challengePermittedIpRanges", GoGetter: "ChallengePermittedIpRanges"},
			_jsii_.MemberProperty{JsiiProperty: "challengePermittedIpRangesInput", GoGetter: "ChallengePermittedIpRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDirectoryPolicy", GoGetter: "DefaultDirectoryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "defaultDirectoryPolicyInput", GoGetter: "DefaultDirectoryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dnsResolver", GoGetter: "DnsResolver"},
			_jsii_.MemberProperty{JsiiProperty: "dnsResolverInput", GoGetter: "DnsResolverInput"},
			_jsii_.MemberProperty{JsiiProperty: "eabPolicy", GoGetter: "EabPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "eabPolicyInput", GoGetter: "EabPolicyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedIssuers", GoMethod: "ResetAllowedIssuers"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedRoles", GoMethod: "ResetAllowedRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRoleExtKeyUsage", GoMethod: "ResetAllowRoleExtKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetChallengeExcludedIpRanges", GoMethod: "ResetChallengeExcludedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetChallengePermittedIpRanges", GoMethod: "ResetChallengePermittedIpRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultDirectoryPolicy", GoMethod: "ResetDefaultDirectoryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsResolver", GoMethod: "ResetDnsResolver"},
			_jsii_.MemberMethod{JsiiMethod: "resetEabPolicy", GoMethod: "ResetEabPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
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
			j := jsiiProxy_PkiSecretBackendConfigAcme{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.pkiSecretBackendConfigAcme.PkiSecretBackendConfigAcmeConfig",
		reflect.TypeOf((*PkiSecretBackendConfigAcmeConfig)(nil)).Elem(),
	)
}
