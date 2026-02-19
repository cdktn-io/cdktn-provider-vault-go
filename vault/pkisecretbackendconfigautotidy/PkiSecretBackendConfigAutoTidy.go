// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigautotidy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/pkisecretbackendconfigautotidy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy vault_pki_secret_backend_config_auto_tidy}.
type PkiSecretBackendConfigAutoTidy interface {
	cdktn.TerraformResource
	AcmeAccountSafetyBuffer() *string
	SetAcmeAccountSafetyBuffer(val *string)
	AcmeAccountSafetyBufferInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	IntervalDuration() *string
	SetIntervalDuration(val *string)
	IntervalDurationInput() *string
	IssuerSafetyBuffer() *string
	SetIssuerSafetyBuffer(val *string)
	IssuerSafetyBufferInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintainStoredCertificateCounts() interface{}
	SetMaintainStoredCertificateCounts(val interface{})
	MaintainStoredCertificateCountsInput() interface{}
	MaxStartupBackoffDuration() *string
	SetMaxStartupBackoffDuration(val *string)
	MaxStartupBackoffDurationInput() *string
	MinStartupBackoffDuration() *string
	SetMinStartupBackoffDuration(val *string)
	MinStartupBackoffDurationInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PauseDuration() *string
	SetPauseDuration(val *string)
	PauseDurationInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublishStoredCertificateCountMetrics() interface{}
	SetPublishStoredCertificateCountMetrics(val interface{})
	PublishStoredCertificateCountMetricsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RevocationQueueSafetyBuffer() *string
	SetRevocationQueueSafetyBuffer(val *string)
	RevocationQueueSafetyBufferInput() *string
	SafetyBuffer() *string
	SetSafetyBuffer(val *string)
	SafetyBufferInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TidyAcme() interface{}
	SetTidyAcme(val interface{})
	TidyAcmeInput() interface{}
	TidyCertMetadata() interface{}
	SetTidyCertMetadata(val interface{})
	TidyCertMetadataInput() interface{}
	TidyCertStore() interface{}
	SetTidyCertStore(val interface{})
	TidyCertStoreInput() interface{}
	TidyCmpv2NonceStore() interface{}
	SetTidyCmpv2NonceStore(val interface{})
	TidyCmpv2NonceStoreInput() interface{}
	TidyCrossClusterRevokedCerts() interface{}
	SetTidyCrossClusterRevokedCerts(val interface{})
	TidyCrossClusterRevokedCertsInput() interface{}
	TidyExpiredIssuers() interface{}
	SetTidyExpiredIssuers(val interface{})
	TidyExpiredIssuersInput() interface{}
	TidyMoveLegacyCaBundle() interface{}
	SetTidyMoveLegacyCaBundle(val interface{})
	TidyMoveLegacyCaBundleInput() interface{}
	TidyRevocationQueue() interface{}
	SetTidyRevocationQueue(val interface{})
	TidyRevocationQueueInput() interface{}
	TidyRevokedCertIssuerAssociations() interface{}
	SetTidyRevokedCertIssuerAssociations(val interface{})
	TidyRevokedCertIssuerAssociationsInput() interface{}
	TidyRevokedCerts() interface{}
	SetTidyRevokedCerts(val interface{})
	TidyRevokedCertsInput() interface{}
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
	ResetAcmeAccountSafetyBuffer()
	ResetId()
	ResetIntervalDuration()
	ResetIssuerSafetyBuffer()
	ResetMaintainStoredCertificateCounts()
	ResetMaxStartupBackoffDuration()
	ResetMinStartupBackoffDuration()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPauseDuration()
	ResetPublishStoredCertificateCountMetrics()
	ResetRevocationQueueSafetyBuffer()
	ResetSafetyBuffer()
	ResetTidyAcme()
	ResetTidyCertMetadata()
	ResetTidyCertStore()
	ResetTidyCmpv2NonceStore()
	ResetTidyCrossClusterRevokedCerts()
	ResetTidyExpiredIssuers()
	ResetTidyMoveLegacyCaBundle()
	ResetTidyRevocationQueue()
	ResetTidyRevokedCertIssuerAssociations()
	ResetTidyRevokedCerts()
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
}

// The jsii proxy struct for PkiSecretBackendConfigAutoTidy
type jsiiProxy_PkiSecretBackendConfigAutoTidy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) AcmeAccountSafetyBuffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmeAccountSafetyBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) AcmeAccountSafetyBufferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmeAccountSafetyBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) IntervalDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) IntervalDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) IssuerSafetyBuffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerSafetyBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) IssuerSafetyBufferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerSafetyBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MaintainStoredCertificateCounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintainStoredCertificateCounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MaintainStoredCertificateCountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintainStoredCertificateCountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MaxStartupBackoffDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxStartupBackoffDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MaxStartupBackoffDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxStartupBackoffDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MinStartupBackoffDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minStartupBackoffDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) MinStartupBackoffDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minStartupBackoffDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) PauseDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) PauseDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pauseDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) PublishStoredCertificateCountMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishStoredCertificateCountMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) PublishStoredCertificateCountMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishStoredCertificateCountMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) RevocationQueueSafetyBuffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationQueueSafetyBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) RevocationQueueSafetyBufferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationQueueSafetyBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) SafetyBuffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safetyBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) SafetyBufferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"safetyBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyAcme() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyAcme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyAcmeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyAcmeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCertMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCertMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCertMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCertMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCertStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCertStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCertStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCertStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCmpv2NonceStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCmpv2NonceStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCmpv2NonceStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCmpv2NonceStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCrossClusterRevokedCerts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCrossClusterRevokedCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyCrossClusterRevokedCertsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyCrossClusterRevokedCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyExpiredIssuers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyExpiredIssuers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyExpiredIssuersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyExpiredIssuersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyMoveLegacyCaBundle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyMoveLegacyCaBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyMoveLegacyCaBundleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyMoveLegacyCaBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevocationQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevocationQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevocationQueueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevocationQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevokedCertIssuerAssociations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevokedCertIssuerAssociations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevokedCertIssuerAssociationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevokedCertIssuerAssociationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevokedCerts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevokedCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy) TidyRevokedCertsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tidyRevokedCertsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy vault_pki_secret_backend_config_auto_tidy} Resource.
func NewPkiSecretBackendConfigAutoTidy(scope constructs.Construct, id *string, config *PkiSecretBackendConfigAutoTidyConfig) PkiSecretBackendConfigAutoTidy {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendConfigAutoTidyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendConfigAutoTidy{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/pki_secret_backend_config_auto_tidy vault_pki_secret_backend_config_auto_tidy} Resource.
func NewPkiSecretBackendConfigAutoTidy_Override(p PkiSecretBackendConfigAutoTidy, scope constructs.Construct, id *string, config *PkiSecretBackendConfigAutoTidyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetAcmeAccountSafetyBuffer(val *string) {
	if err := j.validateSetAcmeAccountSafetyBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acmeAccountSafetyBuffer",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetIntervalDuration(val *string) {
	if err := j.validateSetIntervalDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetIssuerSafetyBuffer(val *string) {
	if err := j.validateSetIssuerSafetyBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerSafetyBuffer",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetMaintainStoredCertificateCounts(val interface{}) {
	if err := j.validateSetMaintainStoredCertificateCountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintainStoredCertificateCounts",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetMaxStartupBackoffDuration(val *string) {
	if err := j.validateSetMaxStartupBackoffDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStartupBackoffDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetMinStartupBackoffDuration(val *string) {
	if err := j.validateSetMinStartupBackoffDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minStartupBackoffDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetPauseDuration(val *string) {
	if err := j.validateSetPauseDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetPublishStoredCertificateCountMetrics(val interface{}) {
	if err := j.validateSetPublishStoredCertificateCountMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishStoredCertificateCountMetrics",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetRevocationQueueSafetyBuffer(val *string) {
	if err := j.validateSetRevocationQueueSafetyBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revocationQueueSafetyBuffer",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetSafetyBuffer(val *string) {
	if err := j.validateSetSafetyBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"safetyBuffer",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyAcme(val interface{}) {
	if err := j.validateSetTidyAcmeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyAcme",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyCertMetadata(val interface{}) {
	if err := j.validateSetTidyCertMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyCertMetadata",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyCertStore(val interface{}) {
	if err := j.validateSetTidyCertStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyCertStore",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyCmpv2NonceStore(val interface{}) {
	if err := j.validateSetTidyCmpv2NonceStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyCmpv2NonceStore",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyCrossClusterRevokedCerts(val interface{}) {
	if err := j.validateSetTidyCrossClusterRevokedCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyCrossClusterRevokedCerts",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyExpiredIssuers(val interface{}) {
	if err := j.validateSetTidyExpiredIssuersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyExpiredIssuers",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyMoveLegacyCaBundle(val interface{}) {
	if err := j.validateSetTidyMoveLegacyCaBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyMoveLegacyCaBundle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyRevocationQueue(val interface{}) {
	if err := j.validateSetTidyRevocationQueueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyRevocationQueue",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyRevokedCertIssuerAssociations(val interface{}) {
	if err := j.validateSetTidyRevokedCertIssuerAssociationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyRevokedCertIssuerAssociations",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigAutoTidy)SetTidyRevokedCerts(val interface{}) {
	if err := j.validateSetTidyRevokedCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tidyRevokedCerts",
		val,
	)
}

// Generates CDKTN code for importing a PkiSecretBackendConfigAutoTidy resource upon running "cdktn plan <stack-name>".
func PkiSecretBackendConfigAutoTidy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigAutoTidy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
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
func PkiSecretBackendConfigAutoTidy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigAutoTidy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendConfigAutoTidy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigAutoTidy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendConfigAutoTidy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigAutoTidy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendConfigAutoTidy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiSecretBackendConfigAutoTidy.PkiSecretBackendConfigAutoTidy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetAcmeAccountSafetyBuffer() {
	_jsii_.InvokeVoid(
		p,
		"resetAcmeAccountSafetyBuffer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetIntervalDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetIntervalDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetIssuerSafetyBuffer() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerSafetyBuffer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetMaintainStoredCertificateCounts() {
	_jsii_.InvokeVoid(
		p,
		"resetMaintainStoredCertificateCounts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetMaxStartupBackoffDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxStartupBackoffDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetMinStartupBackoffDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetMinStartupBackoffDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetPauseDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetPauseDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetPublishStoredCertificateCountMetrics() {
	_jsii_.InvokeVoid(
		p,
		"resetPublishStoredCertificateCountMetrics",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetRevocationQueueSafetyBuffer() {
	_jsii_.InvokeVoid(
		p,
		"resetRevocationQueueSafetyBuffer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetSafetyBuffer() {
	_jsii_.InvokeVoid(
		p,
		"resetSafetyBuffer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyAcme() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyAcme",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyCertMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyCertMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyCertStore() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyCertStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyCmpv2NonceStore() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyCmpv2NonceStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyCrossClusterRevokedCerts() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyCrossClusterRevokedCerts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyExpiredIssuers() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyExpiredIssuers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyMoveLegacyCaBundle() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyMoveLegacyCaBundle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyRevocationQueue() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyRevocationQueue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyRevokedCertIssuerAssociations() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyRevokedCertIssuerAssociations",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ResetTidyRevokedCerts() {
	_jsii_.InvokeVoid(
		p,
		"resetTidyRevokedCerts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigAutoTidy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

