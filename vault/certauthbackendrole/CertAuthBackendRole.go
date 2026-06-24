// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/certauthbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cert_auth_backend_role vault_cert_auth_backend_role}.
type CertAuthBackendRole interface {
	cdktn.TerraformResource
	AliasMetadata() *map[string]*string
	SetAliasMetadata(val *map[string]*string)
	AliasMetadataInput() *map[string]*string
	AllowedCommonNames() *[]*string
	SetAllowedCommonNames(val *[]*string)
	AllowedCommonNamesInput() *[]*string
	AllowedDnsSans() *[]*string
	SetAllowedDnsSans(val *[]*string)
	AllowedDnsSansInput() *[]*string
	AllowedEmailSans() *[]*string
	SetAllowedEmailSans(val *[]*string)
	AllowedEmailSansInput() *[]*string
	AllowedNames() *[]*string
	SetAllowedNames(val *[]*string)
	AllowedNamesInput() *[]*string
	AllowedOrganizationalUnits() *[]*string
	SetAllowedOrganizationalUnits(val *[]*string)
	AllowedOrganizationalUnitsInput() *[]*string
	AllowedUriSans() *[]*string
	SetAllowedUriSans(val *[]*string)
	AllowedUriSansInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	// The tree node.
	Node() constructs.Node
	OcspCaCertificates() *string
	SetOcspCaCertificates(val *string)
	OcspCaCertificatesInput() *string
	OcspEnabled() interface{}
	SetOcspEnabled(val interface{})
	OcspEnabledInput() interface{}
	OcspFailOpen() interface{}
	SetOcspFailOpen(val interface{})
	OcspFailOpenInput() interface{}
	OcspMaxRetries() *float64
	SetOcspMaxRetries(val *float64)
	OcspMaxRetriesInput() *float64
	OcspQueryAllServers() interface{}
	SetOcspQueryAllServers(val interface{})
	OcspQueryAllServersInput() interface{}
	OcspServersOverride() *[]*string
	SetOcspServersOverride(val *[]*string)
	OcspServersOverrideInput() *[]*string
	OcspThisUpdateMaxAge() *float64
	SetOcspThisUpdateMaxAge(val *float64)
	OcspThisUpdateMaxAgeInput() *float64
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
	RequiredExtensions() *[]*string
	SetRequiredExtensions(val *[]*string)
	RequiredExtensionsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	TokenExplicitMaxTtl() *float64
	SetTokenExplicitMaxTtl(val *float64)
	TokenExplicitMaxTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	TokenNoDefaultPolicy() interface{}
	SetTokenNoDefaultPolicy(val interface{})
	TokenNoDefaultPolicyInput() interface{}
	TokenNumUses() *float64
	SetTokenNumUses(val *float64)
	TokenNumUsesInput() *float64
	TokenPeriod() *float64
	SetTokenPeriod(val *float64)
	TokenPeriodInput() *float64
	TokenPolicies() *[]*string
	SetTokenPolicies(val *[]*string)
	TokenPoliciesInput() *[]*string
	TokenTtl() *float64
	SetTokenTtl(val *float64)
	TokenTtlInput() *float64
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	ResetAliasMetadata()
	ResetAllowedCommonNames()
	ResetAllowedDnsSans()
	ResetAllowedEmailSans()
	ResetAllowedNames()
	ResetAllowedOrganizationalUnits()
	ResetAllowedUriSans()
	ResetBackend()
	ResetDisplayName()
	ResetId()
	ResetNamespace()
	ResetOcspCaCertificates()
	ResetOcspEnabled()
	ResetOcspFailOpen()
	ResetOcspMaxRetries()
	ResetOcspQueryAllServers()
	ResetOcspServersOverride()
	ResetOcspThisUpdateMaxAge()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequiredExtensions()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
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

// The jsii proxy struct for CertAuthBackendRole
type jsiiProxy_CertAuthBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CertAuthBackendRole) AliasMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AliasMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliasMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedCommonNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCommonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedCommonNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCommonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedDnsSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDnsSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedDnsSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDnsSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedEmailSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedEmailSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEmailSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedOrganizationalUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationalUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedOrganizationalUnitsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrganizationalUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedUriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) AllowedUriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUriSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspCaCertificates() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspCaCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspCaCertificatesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspCaCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspFailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspFailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspMaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocspMaxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspMaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocspMaxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspQueryAllServers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspQueryAllServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspQueryAllServersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspQueryAllServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspServersOverride() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ocspServersOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspServersOverrideInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ocspServersOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspThisUpdateMaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocspThisUpdateMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) OcspThisUpdateMaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocspThisUpdateMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) RequiredExtensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) RequiredExtensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cert_auth_backend_role vault_cert_auth_backend_role} Resource.
func NewCertAuthBackendRole(scope constructs.Construct, id *string, config *CertAuthBackendRoleConfig) CertAuthBackendRole {
	_init_.Initialize()

	if err := validateNewCertAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertAuthBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/cert_auth_backend_role vault_cert_auth_backend_role} Resource.
func NewCertAuthBackendRole_Override(c CertAuthBackendRole, scope constructs.Construct, id *string, config *CertAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAliasMetadata(val *map[string]*string) {
	if err := j.validateSetAliasMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliasMetadata",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedCommonNames(val *[]*string) {
	if err := j.validateSetAllowedCommonNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCommonNames",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedDnsSans(val *[]*string) {
	if err := j.validateSetAllowedDnsSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDnsSans",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedEmailSans(val *[]*string) {
	if err := j.validateSetAllowedEmailSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEmailSans",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedNames(val *[]*string) {
	if err := j.validateSetAllowedNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedNames",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedOrganizationalUnits(val *[]*string) {
	if err := j.validateSetAllowedOrganizationalUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrganizationalUnits",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetAllowedUriSans(val *[]*string) {
	if err := j.validateSetAllowedUriSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUriSans",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspCaCertificates(val *string) {
	if err := j.validateSetOcspCaCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspCaCertificates",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspEnabled(val interface{}) {
	if err := j.validateSetOcspEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspEnabled",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspFailOpen(val interface{}) {
	if err := j.validateSetOcspFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspFailOpen",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspMaxRetries(val *float64) {
	if err := j.validateSetOcspMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspMaxRetries",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspQueryAllServers(val interface{}) {
	if err := j.validateSetOcspQueryAllServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspQueryAllServers",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspServersOverride(val *[]*string) {
	if err := j.validateSetOcspServersOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspServersOverride",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetOcspThisUpdateMaxAge(val *float64) {
	if err := j.validateSetOcspThisUpdateMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspThisUpdateMaxAge",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetRequiredExtensions(val *[]*string) {
	if err := j.validateSetRequiredExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredExtensions",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_CertAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

// Generates CDKTN code for importing a CertAuthBackendRole resource upon running "cdktn plan <stack-name>".
func CertAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCertAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
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
func CertAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CertAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CertAuthBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertAuthBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CertAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.certAuthBackendRole.CertAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAliasMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetAliasMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedCommonNames() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedCommonNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedDnsSans() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedDnsSans",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedEmailSans() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedEmailSans",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedNames() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedOrganizationalUnits() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedOrganizationalUnits",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetAllowedUriSans() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedUriSans",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		c,
		"resetBackend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetDisplayName() {
	_jsii_.InvokeVoid(
		c,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspCaCertificates() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspCaCertificates",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspFailOpen() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspFailOpen",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspMaxRetries() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspMaxRetries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspQueryAllServers() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspQueryAllServers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspServersOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspServersOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOcspThisUpdateMaxAge() {
	_jsii_.InvokeVoid(
		c,
		"resetOcspThisUpdateMaxAge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetRequiredExtensions() {
	_jsii_.InvokeVoid(
		c,
		"resetRequiredExtensions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		c,
		"resetTokenType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CertAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertAuthBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

