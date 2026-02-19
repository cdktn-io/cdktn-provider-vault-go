// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretbackendrole

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		reflect.TypeOf((*KubernetesSecretBackendRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedKubernetesNamespaces", GoGetter: "AllowedKubernetesNamespaces"},
			_jsii_.MemberProperty{JsiiProperty: "allowedKubernetesNamespaceSelector", GoGetter: "AllowedKubernetesNamespaceSelector"},
			_jsii_.MemberProperty{JsiiProperty: "allowedKubernetesNamespaceSelectorInput", GoGetter: "AllowedKubernetesNamespaceSelectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedKubernetesNamespacesInput", GoGetter: "AllowedKubernetesNamespacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "extraAnnotations", GoGetter: "ExtraAnnotations"},
			_jsii_.MemberProperty{JsiiProperty: "extraAnnotationsInput", GoGetter: "ExtraAnnotationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "extraLabels", GoGetter: "ExtraLabels"},
			_jsii_.MemberProperty{JsiiProperty: "extraLabelsInput", GoGetter: "ExtraLabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "generatedRoleRules", GoGetter: "GeneratedRoleRules"},
			_jsii_.MemberProperty{JsiiProperty: "generatedRoleRulesInput", GoGetter: "GeneratedRoleRulesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kubernetesRoleName", GoGetter: "KubernetesRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesRoleNameInput", GoGetter: "KubernetesRoleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesRoleType", GoGetter: "KubernetesRoleType"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesRoleTypeInput", GoGetter: "KubernetesRoleTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "nameTemplate", GoGetter: "NameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "nameTemplateInput", GoGetter: "NameTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedKubernetesNamespaces", GoMethod: "ResetAllowedKubernetesNamespaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedKubernetesNamespaceSelector", GoMethod: "ResetAllowedKubernetesNamespaceSelector"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtraAnnotations", GoMethod: "ResetExtraAnnotations"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtraLabels", GoMethod: "ResetExtraLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeneratedRoleRules", GoMethod: "ResetGeneratedRoleRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesRoleName", GoMethod: "ResetKubernetesRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetKubernetesRoleType", GoMethod: "ResetKubernetesRoleType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameTemplate", GoMethod: "ResetNameTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountName", GoMethod: "ResetServiceAccountName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenDefaultTtl", GoMethod: "ResetTokenDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenMaxTtl", GoMethod: "ResetTokenMaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountNameInput", GoGetter: "ServiceAccountNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tokenDefaultTtl", GoGetter: "TokenDefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenDefaultTtlInput", GoGetter: "TokenDefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenMaxTtl", GoGetter: "TokenMaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenMaxTtlInput", GoGetter: "TokenMaxTtlInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesSecretBackendRole{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRoleConfig",
		reflect.TypeOf((*KubernetesSecretBackendRoleConfig)(nil)).Elem(),
	)
}
