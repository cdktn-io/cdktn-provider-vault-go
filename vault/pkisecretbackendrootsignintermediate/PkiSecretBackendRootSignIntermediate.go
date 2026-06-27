// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrootsignintermediate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/pkisecretbackendrootsignintermediate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_root_sign_intermediate vault_pki_secret_backend_root_sign_intermediate}.
type PkiSecretBackendRootSignIntermediate interface {
	cdktn.TerraformResource
	AltNames() *[]*string
	SetAltNames(val *[]*string)
	AltNamesInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	CaChain() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
	CertificateBundle() *string
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
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	Csr() *string
	SetCsr(val *string)
	CsrInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcludeCnFromSans() interface{}
	SetExcludeCnFromSans(val interface{})
	ExcludeCnFromSansInput() interface{}
	ExcludedDnsDomains() *[]*string
	SetExcludedDnsDomains(val *[]*string)
	ExcludedDnsDomainsInput() *[]*string
	ExcludedEmailAddresses() *[]*string
	SetExcludedEmailAddresses(val *[]*string)
	ExcludedEmailAddressesInput() *[]*string
	ExcludedIpRanges() *[]*string
	SetExcludedIpRanges(val *[]*string)
	ExcludedIpRangesInput() *[]*string
	ExcludedUriDomains() *[]*string
	SetExcludedUriDomains(val *[]*string)
	ExcludedUriDomainsInput() *[]*string
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
	KeyUsage() *[]*string
	SetKeyUsage(val *[]*string)
	KeyUsageInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	MaxPathLength() *float64
	SetMaxPathLength(val *float64)
	MaxPathLengthInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NotAfter() *string
	SetNotAfter(val *string)
	NotAfterInput() *string
	NotBeforeDuration() *string
	SetNotBeforeDuration(val *string)
	NotBeforeDurationInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	OtherSans() *[]*string
	SetOtherSans(val *[]*string)
	OtherSansInput() *[]*string
	Ou() *string
	SetOu(val *string)
	OuInput() *string
	PermittedDnsDomains() *[]*string
	SetPermittedDnsDomains(val *[]*string)
	PermittedDnsDomainsInput() *[]*string
	PermittedEmailAddresses() *[]*string
	SetPermittedEmailAddresses(val *[]*string)
	PermittedEmailAddressesInput() *[]*string
	PermittedIpRanges() *[]*string
	SetPermittedIpRanges(val *[]*string)
	PermittedIpRangesInput() *[]*string
	PermittedUriDomains() *[]*string
	SetPermittedUriDomains(val *[]*string)
	PermittedUriDomainsInput() *[]*string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	Province() *string
	SetProvince(val *string)
	ProvinceInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Revoke() interface{}
	SetRevoke(val interface{})
	RevokeInput() interface{}
	SerialNumber() *string
	SignatureBits() *float64
	SetSignatureBits(val *float64)
	SignatureBitsInput() *float64
	Skid() *string
	SetSkid(val *string)
	SkidInput() *string
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
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
	UseCsrValues() interface{}
	SetUseCsrValues(val interface{})
	UseCsrValuesInput() interface{}
	UsePss() interface{}
	SetUsePss(val interface{})
	UsePssInput() interface{}
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
	ResetCountry()
	ResetExcludeCnFromSans()
	ResetExcludedDnsDomains()
	ResetExcludedEmailAddresses()
	ResetExcludedIpRanges()
	ResetExcludedUriDomains()
	ResetFormat()
	ResetId()
	ResetIpSans()
	ResetIssuerRef()
	ResetKeyUsage()
	ResetLocality()
	ResetMaxPathLength()
	ResetNamespace()
	ResetNotAfter()
	ResetNotBeforeDuration()
	ResetOrganization()
	ResetOtherSans()
	ResetOu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermittedDnsDomains()
	ResetPermittedEmailAddresses()
	ResetPermittedIpRanges()
	ResetPermittedUriDomains()
	ResetPostalCode()
	ResetProvince()
	ResetRevoke()
	ResetSignatureBits()
	ResetSkid()
	ResetStreetAddress()
	ResetTtl()
	ResetUriSans()
	ResetUseCsrValues()
	ResetUsePss()
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

// The jsii proxy struct for PkiSecretBackendRootSignIntermediate
type jsiiProxy_PkiSecretBackendRootSignIntermediate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) AltNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) AltNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CaChain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"caChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CertificateBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Csr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) CsrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludeCnFromSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludeCnFromSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedDnsDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedDnsDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedUriDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUriDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ExcludedUriDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUriDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IpSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IpSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IssuerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IssuerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) IssuingCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuingCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) KeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) KeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) MaxPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) MaxPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) NotAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) NotBeforeDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) NotBeforeDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) OtherSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) OtherSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Ou() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ou",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) OuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ouInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedDnsDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedDnsDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedUriDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUriDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PermittedUriDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUriDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Revoke() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) RevokeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) SignatureBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) SignatureBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Skid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) SkidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UseCsrValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UseCsrValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UsePss() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate) UsePssInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePssInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_root_sign_intermediate vault_pki_secret_backend_root_sign_intermediate} Resource.
func NewPkiSecretBackendRootSignIntermediate(scope constructs.Construct, id *string, config *PkiSecretBackendRootSignIntermediateConfig) PkiSecretBackendRootSignIntermediate {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendRootSignIntermediateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendRootSignIntermediate{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_root_sign_intermediate vault_pki_secret_backend_root_sign_intermediate} Resource.
func NewPkiSecretBackendRootSignIntermediate_Override(p PkiSecretBackendRootSignIntermediate, scope constructs.Construct, id *string, config *PkiSecretBackendRootSignIntermediateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetAltNames(val *[]*string) {
	if err := j.validateSetAltNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"altNames",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetCsr(val *string) {
	if err := j.validateSetCsrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csr",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetExcludeCnFromSans(val interface{}) {
	if err := j.validateSetExcludeCnFromSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCnFromSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetExcludedDnsDomains(val *[]*string) {
	if err := j.validateSetExcludedDnsDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedDnsDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetExcludedEmailAddresses(val *[]*string) {
	if err := j.validateSetExcludedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetExcludedIpRanges(val *[]*string) {
	if err := j.validateSetExcludedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetExcludedUriDomains(val *[]*string) {
	if err := j.validateSetExcludedUriDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUriDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetIpSans(val *[]*string) {
	if err := j.validateSetIpSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetIssuerRef(val *string) {
	if err := j.validateSetIssuerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerRef",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetKeyUsage(val *[]*string) {
	if err := j.validateSetKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyUsage",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetMaxPathLength(val *float64) {
	if err := j.validateSetMaxPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPathLength",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetNotAfter(val *string) {
	if err := j.validateSetNotAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAfter",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetNotBeforeDuration(val *string) {
	if err := j.validateSetNotBeforeDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notBeforeDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetOtherSans(val *[]*string) {
	if err := j.validateSetOtherSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otherSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetOu(val *string) {
	if err := j.validateSetOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ou",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetPermittedDnsDomains(val *[]*string) {
	if err := j.validateSetPermittedDnsDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedDnsDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetPermittedEmailAddresses(val *[]*string) {
	if err := j.validateSetPermittedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetPermittedIpRanges(val *[]*string) {
	if err := j.validateSetPermittedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetPermittedUriDomains(val *[]*string) {
	if err := j.validateSetPermittedUriDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedUriDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetProvince(val *string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetRevoke(val interface{}) {
	if err := j.validateSetRevokeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revoke",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetSignatureBits(val *float64) {
	if err := j.validateSetSignatureBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetSkid(val *string) {
	if err := j.validateSetSkidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skid",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetUriSans(val *[]*string) {
	if err := j.validateSetUriSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetUseCsrValues(val interface{}) {
	if err := j.validateSetUseCsrValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsrValues",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootSignIntermediate)SetUsePss(val interface{}) {
	if err := j.validateSetUsePssParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePss",
		val,
	)
}

// Generates CDKTN code for importing a PkiSecretBackendRootSignIntermediate resource upon running "cdktn plan <stack-name>".
func PkiSecretBackendRootSignIntermediate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootSignIntermediate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
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
func PkiSecretBackendRootSignIntermediate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootSignIntermediate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRootSignIntermediate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootSignIntermediate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRootSignIntermediate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootSignIntermediate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendRootSignIntermediate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiSecretBackendRootSignIntermediate.PkiSecretBackendRootSignIntermediate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetAltNames() {
	_jsii_.InvokeVoid(
		p,
		"resetAltNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetCountry() {
	_jsii_.InvokeVoid(
		p,
		"resetCountry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetExcludeCnFromSans() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeCnFromSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetExcludedDnsDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedDnsDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetExcludedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetExcludedUriDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedUriDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetIpSans() {
	_jsii_.InvokeVoid(
		p,
		"resetIpSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetIssuerRef() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetLocality() {
	_jsii_.InvokeVoid(
		p,
		"resetLocality",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetMaxPathLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxPathLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetNotAfter() {
	_jsii_.InvokeVoid(
		p,
		"resetNotAfter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetNotBeforeDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetNotBeforeDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetOrganization() {
	_jsii_.InvokeVoid(
		p,
		"resetOrganization",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetOtherSans() {
	_jsii_.InvokeVoid(
		p,
		"resetOtherSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetOu() {
	_jsii_.InvokeVoid(
		p,
		"resetOu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetPermittedDnsDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedDnsDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetPermittedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetPermittedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetPermittedUriDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedUriDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetPostalCode() {
	_jsii_.InvokeVoid(
		p,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetProvince() {
	_jsii_.InvokeVoid(
		p,
		"resetProvince",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetRevoke() {
	_jsii_.InvokeVoid(
		p,
		"resetRevoke",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetSignatureBits() {
	_jsii_.InvokeVoid(
		p,
		"resetSignatureBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetSkid() {
	_jsii_.InvokeVoid(
		p,
		"resetSkid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		p,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetUriSans() {
	_jsii_.InvokeVoid(
		p,
		"resetUriSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetUseCsrValues() {
	_jsii_.InvokeVoid(
		p,
		"resetUseCsrValues",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ResetUsePss() {
	_jsii_.InvokeVoid(
		p,
		"resetUsePss",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootSignIntermediate) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

