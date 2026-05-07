// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRole",
		reflect.TypeOf((*PkiExternalCaSecretBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acmeAccountName", GoGetter: "AcmeAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "acmeAccountNameInput", GoGetter: "AcmeAccountNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedChallengeTypes", GoGetter: "AllowedChallengeTypes"},
			_jsii_.MemberProperty{JsiiProperty: "allowedChallengeTypesInput", GoGetter: "AllowedChallengeTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainOptions", GoGetter: "AllowedDomainOptions"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainOptionsInput", GoGetter: "AllowedDomainOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomains", GoGetter: "AllowedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowedDomainsInput", GoGetter: "AllowedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "csrGenerateKeyType", GoGetter: "CsrGenerateKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "csrGenerateKeyTypeInput", GoGetter: "CsrGenerateKeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "csrIdentifierPopulation", GoGetter: "CsrIdentifierPopulation"},
			_jsii_.MemberProperty{JsiiProperty: "csrIdentifierPopulationInput", GoGetter: "CsrIdentifierPopulationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "force", GoGetter: "Force"},
			_jsii_.MemberProperty{JsiiProperty: "forceInput", GoGetter: "ForceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdateDate", GoGetter: "LastUpdateDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedChallengeTypes", GoMethod: "ResetAllowedChallengeTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDomainOptions", GoMethod: "ResetAllowedDomainOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedDomains", GoMethod: "ResetAllowedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetCsrGenerateKeyType", GoMethod: "ResetCsrGenerateKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetCsrIdentifierPopulation", GoMethod: "ResetCsrIdentifierPopulation"},
			_jsii_.MemberMethod{JsiiMethod: "resetForce", GoMethod: "ResetForce"},
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
			j := jsiiProxy_PkiExternalCaSecretBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendRole.PkiExternalCaSecretBackendRoleConfig",
		reflect.TypeOf((*PkiExternalCaSecretBackendRoleConfig)(nil)).Elem(),
	)
}
