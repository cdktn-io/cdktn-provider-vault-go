// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package samlauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/samlauthbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend vault_saml_auth_backend}.
type SamlAuthBackend interface {
	cdktn.TerraformResource
	AcsUrls() *[]*string
	SetAcsUrls(val *[]*string)
	AcsUrlsInput() *[]*string
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
	DefaultRole() *string
	SetDefaultRole(val *string)
	DefaultRoleInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableRemount() interface{}
	SetDisableRemount(val interface{})
	DisableRemountInput() interface{}
	EntityId() *string
	SetEntityId(val *string)
	EntityIdInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdpCert() *string
	SetIdpCert(val *string)
	IdpCertInput() *string
	IdpEntityId() *string
	SetIdpEntityId(val *string)
	IdpEntityIdInput() *string
	IdpMetadataUrl() *string
	SetIdpMetadataUrl(val *string)
	IdpMetadataUrlInput() *string
	IdpSsoUrl() *string
	SetIdpSsoUrl(val *string)
	IdpSsoUrlInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tune() SamlAuthBackendTuneList
	TuneInput() interface{}
	ValidateAssertionSignature() interface{}
	SetValidateAssertionSignature(val interface{})
	ValidateAssertionSignatureInput() interface{}
	ValidateResponseSignature() interface{}
	SetValidateResponseSignature(val interface{})
	ValidateResponseSignatureInput() interface{}
	VerboseLogging() interface{}
	SetVerboseLogging(val interface{})
	VerboseLoggingInput() interface{}
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
	PutTune(value interface{})
	ResetDefaultRole()
	ResetDisableRemount()
	ResetId()
	ResetIdpCert()
	ResetIdpEntityId()
	ResetIdpMetadataUrl()
	ResetIdpSsoUrl()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetTune()
	ResetValidateAssertionSignature()
	ResetValidateResponseSignature()
	ResetVerboseLogging()
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

// The jsii proxy struct for SamlAuthBackend
type jsiiProxy_SamlAuthBackend struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SamlAuthBackend) AcsUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) AcsUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) EntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpEntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpEntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpEntityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpMetadataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpMetadataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpSsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpSsoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) IdpSsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpSsoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) Tune() SamlAuthBackendTuneList {
	var returns SamlAuthBackendTuneList
	_jsii_.Get(
		j,
		"tune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) TuneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ValidateAssertionSignature() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateAssertionSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ValidateAssertionSignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateAssertionSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ValidateResponseSignature() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateResponseSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) ValidateResponseSignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateResponseSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) VerboseLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verboseLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlAuthBackend) VerboseLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verboseLoggingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend vault_saml_auth_backend} Resource.
func NewSamlAuthBackend(scope constructs.Construct, id *string, config *SamlAuthBackendConfig) SamlAuthBackend {
	_init_.Initialize()

	if err := validateNewSamlAuthBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlAuthBackend{}

	_jsii_.Create(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/saml_auth_backend vault_saml_auth_backend} Resource.
func NewSamlAuthBackend_Override(s SamlAuthBackend, scope constructs.Construct, id *string, config *SamlAuthBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetAcsUrls(val *[]*string) {
	if err := j.validateSetAcsUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acsUrls",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetEntityId(val *string) {
	if err := j.validateSetEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityId",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetIdpCert(val *string) {
	if err := j.validateSetIdpCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpCert",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetIdpEntityId(val *string) {
	if err := j.validateSetIdpEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpEntityId",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetIdpMetadataUrl(val *string) {
	if err := j.validateSetIdpMetadataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpMetadataUrl",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetIdpSsoUrl(val *string) {
	if err := j.validateSetIdpSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpSsoUrl",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetValidateAssertionSignature(val interface{}) {
	if err := j.validateSetValidateAssertionSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateAssertionSignature",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetValidateResponseSignature(val interface{}) {
	if err := j.validateSetValidateResponseSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateResponseSignature",
		val,
	)
}

func (j *jsiiProxy_SamlAuthBackend)SetVerboseLogging(val interface{}) {
	if err := j.validateSetVerboseLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verboseLogging",
		val,
	)
}

// Generates CDKTN code for importing a SamlAuthBackend resource upon running "cdktn plan <stack-name>".
func SamlAuthBackend_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSamlAuthBackend_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
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
func SamlAuthBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlAuthBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlAuthBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlAuthBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SamlAuthBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSamlAuthBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SamlAuthBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.samlAuthBackend.SamlAuthBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SamlAuthBackend) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SamlAuthBackend) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SamlAuthBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SamlAuthBackend) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlAuthBackend) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SamlAuthBackend) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SamlAuthBackend) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SamlAuthBackend) PutTune(value interface{}) {
	if err := s.validatePutTuneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTune",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetIdpCert() {
	_jsii_.InvokeVoid(
		s,
		"resetIdpCert",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetIdpEntityId() {
	_jsii_.InvokeVoid(
		s,
		"resetIdpEntityId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetIdpMetadataUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetIdpMetadataUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetIdpSsoUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetIdpSsoUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		s,
		"resetNamespace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetPath() {
	_jsii_.InvokeVoid(
		s,
		"resetPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetTune() {
	_jsii_.InvokeVoid(
		s,
		"resetTune",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetValidateAssertionSignature() {
	_jsii_.InvokeVoid(
		s,
		"resetValidateAssertionSignature",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetValidateResponseSignature() {
	_jsii_.InvokeVoid(
		s,
		"resetValidateResponseSignature",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) ResetVerboseLogging() {
	_jsii_.InvokeVoid(
		s,
		"resetVerboseLogging",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SamlAuthBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlAuthBackend) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

