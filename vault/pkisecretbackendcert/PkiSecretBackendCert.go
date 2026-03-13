// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/pkisecretbackendcert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_cert vault_pki_secret_backend_cert}.
type PkiSecretBackendCert interface {
	cdktn.TerraformResource
	AltNames() *[]*string
	SetAltNames(val *[]*string)
	AltNamesInput() *[]*string
	AutoRenew() interface{}
	SetAutoRenew(val interface{})
	AutoRenewInput() interface{}
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	CaChain() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
	CertMetadata() *string
	SetCertMetadata(val *string)
	CertMetadataInput() *string
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	ExcludeCnFromSans() interface{}
	SetExcludeCnFromSans(val interface{})
	ExcludeCnFromSansInput() interface{}
	Expiration() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpSans() *[]*string
	SetIpSans(val *[]*string)
	IpSansInput() *[]*string
	IssuerRef() *string
	SetIssuerRef(val *string)
	IssuerRefInput() *string
	IssuingCa() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MinSecondsRemaining() *float64
	SetMinSecondsRemaining(val *float64)
	MinSecondsRemainingInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NotAfter() *string
	SetNotAfter(val *string)
	NotAfterInput() *string
	OtherSans() *[]*string
	SetOtherSans(val *[]*string)
	OtherSansInput() *[]*string
	PrivateKey() *string
	PrivateKeyFormat() *string
	SetPrivateKeyFormat(val *string)
	PrivateKeyFormatInput() *string
	PrivateKeyType() *string
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
	RemoveRootsFromChain() interface{}
	SetRemoveRootsFromChain(val interface{})
	RemoveRootsFromChainInput() interface{}
	RenewPending() cdktn.IResolvable
	Revoke() interface{}
	SetRevoke(val interface{})
	RevokeInput() interface{}
	RevokeWithKey() interface{}
	SetRevokeWithKey(val interface{})
	RevokeWithKeyInput() interface{}
	SerialNumber() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	UriSans() *[]*string
	SetUriSans(val *[]*string)
	UriSansInput() *[]*string
	UserIds() *[]*string
	SetUserIds(val *[]*string)
	UserIdsInput() *[]*string
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
	ResetAltNames()
	ResetAutoRenew()
	ResetCertMetadata()
	ResetExcludeCnFromSans()
	ResetFormat()
	ResetId()
	ResetIpSans()
	ResetIssuerRef()
	ResetMinSecondsRemaining()
	ResetNamespace()
	ResetNotAfter()
	ResetOtherSans()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateKeyFormat()
	ResetRemoveRootsFromChain()
	ResetRevoke()
	ResetRevokeWithKey()
	ResetTtl()
	ResetUriSans()
	ResetUserIds()
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

// The jsii proxy struct for PkiSecretBackendCert
type jsiiProxy_PkiSecretBackendCert struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendCert) AltNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) AltNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) AutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) AutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CaChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CertMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CertMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) ExcludeCnFromSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) ExcludeCnFromSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Expiration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IpSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IpSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IssuerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IssuerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) IssuingCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuingCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) MinSecondsRemaining() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSecondsRemaining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) MinSecondsRemainingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSecondsRemainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) NotAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) OtherSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) OtherSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) PrivateKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) PrivateKeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) PrivateKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RemoveRootsFromChain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeRootsFromChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RemoveRootsFromChainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeRootsFromChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RenewPending() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"renewPending",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Revoke() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RevokeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RevokeWithKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeWithKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) RevokeWithKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeWithKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) UriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) UriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) UserIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCert) UserIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_cert vault_pki_secret_backend_cert} Resource.
func NewPkiSecretBackendCert(scope constructs.Construct, id *string, config *PkiSecretBackendCertConfig) PkiSecretBackendCert {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendCert{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_cert vault_pki_secret_backend_cert} Resource.
func NewPkiSecretBackendCert_Override(p PkiSecretBackendCert, scope constructs.Construct, id *string, config *PkiSecretBackendCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetAltNames(val *[]*string) {
	if err := j.validateSetAltNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"altNames",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetAutoRenew(val interface{}) {
	if err := j.validateSetAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenew",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetCertMetadata(val *string) {
	if err := j.validateSetCertMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certMetadata",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetExcludeCnFromSans(val interface{}) {
	if err := j.validateSetExcludeCnFromSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCnFromSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetIpSans(val *[]*string) {
	if err := j.validateSetIpSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetIssuerRef(val *string) {
	if err := j.validateSetIssuerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerRef",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetMinSecondsRemaining(val *float64) {
	if err := j.validateSetMinSecondsRemainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSecondsRemaining",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetNotAfter(val *string) {
	if err := j.validateSetNotAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAfter",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetOtherSans(val *[]*string) {
	if err := j.validateSetOtherSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otherSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetPrivateKeyFormat(val *string) {
	if err := j.validateSetPrivateKeyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyFormat",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetRemoveRootsFromChain(val interface{}) {
	if err := j.validateSetRemoveRootsFromChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeRootsFromChain",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetRevoke(val interface{}) {
	if err := j.validateSetRevokeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revoke",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetRevokeWithKey(val interface{}) {
	if err := j.validateSetRevokeWithKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revokeWithKey",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetUriSans(val *[]*string) {
	if err := j.validateSetUriSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCert)SetUserIds(val *[]*string) {
	if err := j.validateSetUserIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIds",
		val,
	)
}

// Generates CDKTN code for importing a PkiSecretBackendCert resource upon running "cdktn plan <stack-name>".
func PkiSecretBackendCert_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendCert_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
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
func PkiSecretBackendCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendCert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendCert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiSecretBackendCert.PkiSecretBackendCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendCert) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendCert) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetAltNames() {
	_jsii_.InvokeVoid(
		p,
		"resetAltNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetAutoRenew() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoRenew",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetCertMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetCertMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetExcludeCnFromSans() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeCnFromSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetIpSans() {
	_jsii_.InvokeVoid(
		p,
		"resetIpSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetIssuerRef() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetMinSecondsRemaining() {
	_jsii_.InvokeVoid(
		p,
		"resetMinSecondsRemaining",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetNotAfter() {
	_jsii_.InvokeVoid(
		p,
		"resetNotAfter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetOtherSans() {
	_jsii_.InvokeVoid(
		p,
		"resetOtherSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetPrivateKeyFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKeyFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetRemoveRootsFromChain() {
	_jsii_.InvokeVoid(
		p,
		"resetRemoveRootsFromChain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetRevoke() {
	_jsii_.InvokeVoid(
		p,
		"resetRevoke",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetRevokeWithKey() {
	_jsii_.InvokeVoid(
		p,
		"resetRevokeWithKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetUriSans() {
	_jsii_.InvokeVoid(
		p,
		"resetUriSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) ResetUserIds() {
	_jsii_.InvokeVoid(
		p,
		"resetUserIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCert) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

