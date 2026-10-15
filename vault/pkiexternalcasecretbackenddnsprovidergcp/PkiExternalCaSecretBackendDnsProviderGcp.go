// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackenddnsprovidergcp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v18/pkiexternalcasecretbackenddnsprovidergcp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp vault_pki_external_ca_secret_backend_dns_provider_gcp}.
type PkiExternalCaSecretBackendDnsProviderGcp interface {
	cdktn.TerraformResource
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
	CreationDate() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	CredentialsWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetCredentialsWo(val *string)
	CredentialsWoInput() *string
	CredentialsWoVersion() *float64
	SetCredentialsWoVersion(val *float64)
	CredentialsWoVersionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Identifiers() *[]*string
	SetIdentifiers(val *[]*string)
	IdentifiersInput() *[]*string
	ImpersonateServiceAccount() *string
	SetImpersonateServiceAccount(val *string)
	ImpersonateServiceAccountInput() *string
	LastUpdatedDate() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Mount() *string
	SetMount(val *string)
	MountInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nameserver() *string
	SetNameserver(val *string)
	NameserverInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	ZoneName() *string
	SetZoneName(val *string)
	ZoneNameInput() *string
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetCredentialsWo()
	ResetCredentialsWoVersion()
	ResetImpersonateServiceAccount()
	ResetNameserver()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTtl()
	ResetZoneName()
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

// The jsii proxy struct for PkiExternalCaSecretBackendDnsProviderGcp
type jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CredentialsWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CredentialsWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CredentialsWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"credentialsWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) CredentialsWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"credentialsWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Identifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) IdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ImpersonateServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ImpersonateServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Mount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) MountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Nameserver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameserver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) NameserverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameserverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ZoneNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp vault_pki_external_ca_secret_backend_dns_provider_gcp} Resource.
func NewPkiExternalCaSecretBackendDnsProviderGcp(scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendDnsProviderGcpConfig) PkiExternalCaSecretBackendDnsProviderGcp {
	_init_.Initialize()

	if err := validateNewPkiExternalCaSecretBackendDnsProviderGcpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp vault_pki_external_ca_secret_backend_dns_provider_gcp} Resource.
func NewPkiExternalCaSecretBackendDnsProviderGcp_Override(p PkiExternalCaSecretBackendDnsProviderGcp, scope constructs.Construct, id *string, config *PkiExternalCaSecretBackendDnsProviderGcpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetCredentialsWo(val *string) {
	if err := j.validateSetCredentialsWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsWo",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetCredentialsWoVersion(val *float64) {
	if err := j.validateSetCredentialsWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsWoVersion",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetIdentifiers(val *[]*string) {
	if err := j.validateSetIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifiers",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetImpersonateServiceAccount(val *string) {
	if err := j.validateSetImpersonateServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impersonateServiceAccount",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetMount(val *string) {
	if err := j.validateSetMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mount",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetNameserver(val *string) {
	if err := j.validateSetNameserverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameserver",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp)SetZoneName(val *string) {
	if err := j.validateSetZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneName",
		val,
	)
}

// Generates CDKTN code for importing a PkiExternalCaSecretBackendDnsProviderGcp resource upon running "cdktn plan <stack-name>".
func PkiExternalCaSecretBackendDnsProviderGcp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendDnsProviderGcp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
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
func PkiExternalCaSecretBackendDnsProviderGcp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendDnsProviderGcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendDnsProviderGcp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendDnsProviderGcp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiExternalCaSecretBackendDnsProviderGcp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiExternalCaSecretBackendDnsProviderGcp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiExternalCaSecretBackendDnsProviderGcp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiExternalCaSecretBackendDnsProviderGcp.PkiExternalCaSecretBackendDnsProviderGcp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := p.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := p.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetCredentialsWo() {
	_jsii_.InvokeVoid(
		p,
		"resetCredentialsWo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetCredentialsWoVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetCredentialsWoVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetImpersonateServiceAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetImpersonateServiceAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetNameserver() {
	_jsii_.InvokeVoid(
		p,
		"resetNameserver",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetProject() {
	_jsii_.InvokeVoid(
		p,
		"resetProject",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ResetZoneName() {
	_jsii_.InvokeVoid(
		p,
		"resetZoneName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiExternalCaSecretBackendDnsProviderGcp) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

