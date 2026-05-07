// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/kubernetessecretbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role}.
type KubernetesSecretBackendRole interface {
	cdktn.TerraformResource
	AllowedKubernetesNamespaces() *[]*string
	SetAllowedKubernetesNamespaces(val *[]*string)
	AllowedKubernetesNamespaceSelector() *string
	SetAllowedKubernetesNamespaceSelector(val *string)
	AllowedKubernetesNamespaceSelectorInput() *string
	AllowedKubernetesNamespacesInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExtraAnnotations() *map[string]*string
	SetExtraAnnotations(val *map[string]*string)
	ExtraAnnotationsInput() *map[string]*string
	ExtraLabels() *map[string]*string
	SetExtraLabels(val *map[string]*string)
	ExtraLabelsInput() *map[string]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedRoleRules() *string
	SetGeneratedRoleRules(val *string)
	GeneratedRoleRulesInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubernetesRoleName() *string
	SetKubernetesRoleName(val *string)
	KubernetesRoleNameInput() *string
	KubernetesRoleType() *string
	SetKubernetesRoleType(val *string)
	KubernetesRoleTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NameTemplate() *string
	SetNameTemplate(val *string)
	NameTemplateInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenDefaultAudiences() *[]*string
	SetTokenDefaultAudiences(val *[]*string)
	TokenDefaultAudiencesInput() *[]*string
	TokenDefaultTtl() *float64
	SetTokenDefaultTtl(val *float64)
	TokenDefaultTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAllowedKubernetesNamespaces()
	ResetAllowedKubernetesNamespaceSelector()
	ResetExtraAnnotations()
	ResetExtraLabels()
	ResetGeneratedRoleRules()
	ResetId()
	ResetKubernetesRoleName()
	ResetKubernetesRoleType()
	ResetNamespace()
	ResetNameTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceAccountName()
	ResetTokenDefaultAudiences()
	ResetTokenDefaultTtl()
	ResetTokenMaxTtl()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for KubernetesSecretBackendRole
type jsiiProxy_KubernetesSecretBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespaceSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespaceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespaceSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespaceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraAnnotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraAnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraAnnotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) GeneratedRoleRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedRoleRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) GeneratedRoleRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedRoleRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultAudiences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDefaultAudiences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultAudiencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDefaultAudiencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDefaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDefaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role} Resource.
func NewKubernetesSecretBackendRole(scope constructs.Construct, id *string, config *KubernetesSecretBackendRoleConfig) KubernetesSecretBackendRole {
	_init_.Initialize()

	if err := validateNewKubernetesSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesSecretBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role} Resource.
func NewKubernetesSecretBackendRole_Override(k KubernetesSecretBackendRole, scope constructs.Construct, id *string, config *KubernetesSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetAllowedKubernetesNamespaces(val *[]*string) {
	if err := j.validateSetAllowedKubernetesNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedKubernetesNamespaces",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetAllowedKubernetesNamespaceSelector(val *string) {
	if err := j.validateSetAllowedKubernetesNamespaceSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedKubernetesNamespaceSelector",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetExtraAnnotations(val *map[string]*string) {
	if err := j.validateSetExtraAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraAnnotations",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetExtraLabels(val *map[string]*string) {
	if err := j.validateSetExtraLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetGeneratedRoleRules(val *string) {
	if err := j.validateSetGeneratedRoleRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generatedRoleRules",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetKubernetesRoleName(val *string) {
	if err := j.validateSetKubernetesRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesRoleName",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetKubernetesRoleType(val *string) {
	if err := j.validateSetKubernetesRoleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesRoleType",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetNameTemplate(val *string) {
	if err := j.validateSetNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameTemplate",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetTokenDefaultAudiences(val *[]*string) {
	if err := j.validateSetTokenDefaultAudiencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDefaultAudiences",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetTokenDefaultTtl(val *float64) {
	if err := j.validateSetTokenDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDefaultTtl",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

// Generates CDKTN code for importing a KubernetesSecretBackendRole resource upon running "cdktn plan <stack-name>".
func KubernetesSecretBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func KubernetesSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetAllowedKubernetesNamespaces() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedKubernetesNamespaces",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetAllowedKubernetesNamespaceSelector() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedKubernetesNamespaceSelector",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetExtraAnnotations() {
	_jsii_.InvokeVoid(
		k,
		"resetExtraAnnotations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetExtraLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetExtraLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetGeneratedRoleRules() {
	_jsii_.InvokeVoid(
		k,
		"resetGeneratedRoleRules",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetKubernetesRoleName() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesRoleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetKubernetesRoleType() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesRoleType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetNameTemplate() {
	_jsii_.InvokeVoid(
		k,
		"resetNameTemplate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetTokenDefaultAudiences() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenDefaultAudiences",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetTokenDefaultTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenDefaultTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

