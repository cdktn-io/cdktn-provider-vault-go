// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrootcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/pkisecretbackendrootcert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert}.
type PkiSecretBackendRootCert interface {
	cdktn.TerraformResource
	AltNames() *[]*string
	SetAltNames(val *[]*string)
	AltNamesInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Certificate() *string
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
	IssuerId() *string
	IssuerName() *string
	SetIssuerName(val *string)
	IssuerNameInput() *string
	IssuingCa() *string
	JksPassword() *string
	SetJksPassword(val *string)
	JksPasswordInput() *string
	JksPrivateKeyAlias() *string
	SetJksPrivateKeyAlias(val *string)
	JksPrivateKeyAliasInput() *string
	KeyBits() *float64
	SetKeyBits(val *float64)
	KeyBitsInput() *float64
	KeyId() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	KeyRef() *string
	SetKeyRef(val *string)
	KeyRefInput() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
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
	ManagedKeyId() *string
	SetManagedKeyId(val *string)
	ManagedKeyIdInput() *string
	ManagedKeyName() *string
	SetManagedKeyName(val *string)
	ManagedKeyNameInput() *string
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
	Pkcs12Encoder() *string
	SetPkcs12Encoder(val *string)
	Pkcs12EncoderInput() *string
	Pkcs12Password() *string
	SetPkcs12Password(val *string)
	Pkcs12PasswordInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	PrivateKeyFormat() *string
	SetPrivateKeyFormat(val *string)
	PrivateKeyFormatInput() *string
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
	SerialNumber() *string
	SignatureBits() *float64
	SetSignatureBits(val *float64)
	SignatureBitsInput() *float64
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
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UriSans() *[]*string
	SetUriSans(val *[]*string)
	UriSansInput() *[]*string
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
	ResetIssuerName()
	ResetJksPassword()
	ResetJksPrivateKeyAlias()
	ResetKeyBits()
	ResetKeyName()
	ResetKeyRef()
	ResetKeyType()
	ResetKeyUsage()
	ResetLocality()
	ResetManagedKeyId()
	ResetManagedKeyName()
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
	ResetPkcs12Encoder()
	ResetPkcs12Password()
	ResetPostalCode()
	ResetPrivateKeyFormat()
	ResetProvince()
	ResetSignatureBits()
	ResetStreetAddress()
	ResetTtl()
	ResetUriSans()
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

// The jsii proxy struct for PkiSecretBackendRootCert
type jsiiProxy_PkiSecretBackendRootCert struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendRootCert) AltNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) AltNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludeCnFromSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludeCnFromSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedDnsDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedDnsDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedUriDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUriDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludedUriDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUriDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IpSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IpSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IssuerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IssuerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IssuerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IssuingCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuingCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) JksPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jksPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) JksPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jksPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) JksPrivateKeyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jksPrivateKeyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) JksPrivateKeyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jksPrivateKeyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ManagedKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ManagedKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ManagedKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ManagedKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) MaxPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) MaxPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NotAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NotBeforeDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NotBeforeDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OtherSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OtherSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Ou() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ou",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ouInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedDnsDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedDnsDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedUriDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUriDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedUriDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUriDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Pkcs12Encoder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs12Encoder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Pkcs12EncoderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs12EncoderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Pkcs12Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs12Password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Pkcs12PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs12PasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PrivateKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PrivateKeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SignatureBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SignatureBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UsePss() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UsePssInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePssInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert} Resource.
func NewPkiSecretBackendRootCert(scope constructs.Construct, id *string, config *PkiSecretBackendRootCertConfig) PkiSecretBackendRootCert {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendRootCertParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendRootCert{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert} Resource.
func NewPkiSecretBackendRootCert_Override(p PkiSecretBackendRootCert, scope constructs.Construct, id *string, config *PkiSecretBackendRootCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetAltNames(val *[]*string) {
	if err := j.validateSetAltNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"altNames",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetExcludeCnFromSans(val interface{}) {
	if err := j.validateSetExcludeCnFromSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeCnFromSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetExcludedDnsDomains(val *[]*string) {
	if err := j.validateSetExcludedDnsDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedDnsDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetExcludedEmailAddresses(val *[]*string) {
	if err := j.validateSetExcludedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetExcludedIpRanges(val *[]*string) {
	if err := j.validateSetExcludedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetExcludedUriDomains(val *[]*string) {
	if err := j.validateSetExcludedUriDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUriDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetIpSans(val *[]*string) {
	if err := j.validateSetIpSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetIssuerName(val *string) {
	if err := j.validateSetIssuerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetJksPassword(val *string) {
	if err := j.validateSetJksPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jksPassword",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetJksPrivateKeyAlias(val *string) {
	if err := j.validateSetJksPrivateKeyAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jksPrivateKeyAlias",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetKeyBits(val *float64) {
	if err := j.validateSetKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetKeyRef(val *string) {
	if err := j.validateSetKeyRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRef",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetKeyUsage(val *[]*string) {
	if err := j.validateSetKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyUsage",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetManagedKeyId(val *string) {
	if err := j.validateSetManagedKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedKeyId",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetManagedKeyName(val *string) {
	if err := j.validateSetManagedKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedKeyName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetMaxPathLength(val *float64) {
	if err := j.validateSetMaxPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPathLength",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetNotAfter(val *string) {
	if err := j.validateSetNotAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAfter",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetNotBeforeDuration(val *string) {
	if err := j.validateSetNotBeforeDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notBeforeDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetOtherSans(val *[]*string) {
	if err := j.validateSetOtherSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"otherSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetOu(val *string) {
	if err := j.validateSetOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ou",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPermittedDnsDomains(val *[]*string) {
	if err := j.validateSetPermittedDnsDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedDnsDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPermittedEmailAddresses(val *[]*string) {
	if err := j.validateSetPermittedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPermittedIpRanges(val *[]*string) {
	if err := j.validateSetPermittedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPermittedUriDomains(val *[]*string) {
	if err := j.validateSetPermittedUriDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedUriDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPkcs12Encoder(val *string) {
	if err := j.validateSetPkcs12EncoderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkcs12Encoder",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPkcs12Password(val *string) {
	if err := j.validateSetPkcs12PasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkcs12Password",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetPrivateKeyFormat(val *string) {
	if err := j.validateSetPrivateKeyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyFormat",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetProvince(val *string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetSignatureBits(val *float64) {
	if err := j.validateSetSignatureBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetUriSans(val *[]*string) {
	if err := j.validateSetUriSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uriSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert)SetUsePss(val interface{}) {
	if err := j.validateSetUsePssParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePss",
		val,
	)
}

// Generates CDKTN code for importing a PkiSecretBackendRootCert resource upon running "cdktn plan <stack-name>".
func PkiSecretBackendRootCert_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootCert_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
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
func PkiSecretBackendRootCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootCert_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRootCert_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootCert_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRootCert_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRootCert_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendRootCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiSecretBackendRootCert.PkiSecretBackendRootCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRootCert) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetAltNames() {
	_jsii_.InvokeVoid(
		p,
		"resetAltNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetCountry() {
	_jsii_.InvokeVoid(
		p,
		"resetCountry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludeCnFromSans() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeCnFromSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludedDnsDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedDnsDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludedUriDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedUriDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetIpSans() {
	_jsii_.InvokeVoid(
		p,
		"resetIpSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetIssuerName() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetJksPassword() {
	_jsii_.InvokeVoid(
		p,
		"resetJksPassword",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetJksPrivateKeyAlias() {
	_jsii_.InvokeVoid(
		p,
		"resetJksPrivateKeyAlias",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyBits() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyName() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyRef() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyType() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetLocality() {
	_jsii_.InvokeVoid(
		p,
		"resetLocality",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetManagedKeyId() {
	_jsii_.InvokeVoid(
		p,
		"resetManagedKeyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetManagedKeyName() {
	_jsii_.InvokeVoid(
		p,
		"resetManagedKeyName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetMaxPathLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxPathLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetNotAfter() {
	_jsii_.InvokeVoid(
		p,
		"resetNotAfter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetNotBeforeDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetNotBeforeDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOrganization() {
	_jsii_.InvokeVoid(
		p,
		"resetOrganization",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOtherSans() {
	_jsii_.InvokeVoid(
		p,
		"resetOtherSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOu() {
	_jsii_.InvokeVoid(
		p,
		"resetOu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPermittedDnsDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedDnsDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPermittedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPermittedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPermittedUriDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedUriDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPkcs12Encoder() {
	_jsii_.InvokeVoid(
		p,
		"resetPkcs12Encoder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPkcs12Password() {
	_jsii_.InvokeVoid(
		p,
		"resetPkcs12Password",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPostalCode() {
	_jsii_.InvokeVoid(
		p,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPrivateKeyFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKeyFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetProvince() {
	_jsii_.InvokeVoid(
		p,
		"resetProvince",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetSignatureBits() {
	_jsii_.InvokeVoid(
		p,
		"resetSignatureBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		p,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetUriSans() {
	_jsii_.InvokeVoid(
		p,
		"resetUriSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetUsePss() {
	_jsii_.InvokeVoid(
		p,
		"resetUsePss",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

