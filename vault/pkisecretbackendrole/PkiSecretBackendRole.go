// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/pkisecretbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_role vault_pki_secret_backend_role}.
type PkiSecretBackendRole interface {
	cdktn.TerraformResource
	AllowAnyName() interface{}
	SetAllowAnyName(val interface{})
	AllowAnyNameInput() interface{}
	AllowBareDomains() interface{}
	SetAllowBareDomains(val interface{})
	AllowBareDomainsInput() interface{}
	AllowedDomains() *[]*string
	SetAllowedDomains(val *[]*string)
	AllowedDomainsInput() *[]*string
	AllowedDomainsTemplate() interface{}
	SetAllowedDomainsTemplate(val interface{})
	AllowedDomainsTemplateInput() interface{}
	AllowedOtherSans() *[]*string
	SetAllowedOtherSans(val *[]*string)
	AllowedOtherSansInput() *[]*string
	AllowedSerialNumbers() *[]*string
	SetAllowedSerialNumbers(val *[]*string)
	AllowedSerialNumbersInput() *[]*string
	AllowedUriSans() *[]*string
	SetAllowedUriSans(val *[]*string)
	AllowedUriSansInput() *[]*string
	AllowedUriSansTemplate() interface{}
	SetAllowedUriSansTemplate(val interface{})
	AllowedUriSansTemplateInput() interface{}
	AllowedUserIds() *[]*string
	SetAllowedUserIds(val *[]*string)
	AllowedUserIdsInput() *[]*string
	AllowGlobDomains() interface{}
	SetAllowGlobDomains(val interface{})
	AllowGlobDomainsInput() interface{}
	AllowIpSans() interface{}
	SetAllowIpSans(val interface{})
	AllowIpSansInput() interface{}
	AllowLocalhost() interface{}
	SetAllowLocalhost(val interface{})
	AllowLocalhostInput() interface{}
	AllowSubdomains() interface{}
	SetAllowSubdomains(val interface{})
	AllowSubdomainsInput() interface{}
	AllowWildcardCertificates() interface{}
	SetAllowWildcardCertificates(val interface{})
	AllowWildcardCertificatesInput() interface{}
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BasicConstraintsValidForNonCa() interface{}
	SetBasicConstraintsValidForNonCa(val interface{})
	BasicConstraintsValidForNonCaInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientFlag() interface{}
	SetClientFlag(val interface{})
	ClientFlagInput() interface{}
	CnValidations() *[]*string
	SetCnValidations(val *[]*string)
	CnValidationsInput() *[]*string
	CodeSigningFlag() interface{}
	SetCodeSigningFlag(val interface{})
	CodeSigningFlagInput() interface{}
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
	Country() *[]*string
	SetCountry(val *[]*string)
	CountryInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmailProtectionFlag() interface{}
	SetEmailProtectionFlag(val interface{})
	EmailProtectionFlagInput() interface{}
	EnforceHostnames() interface{}
	SetEnforceHostnames(val interface{})
	EnforceHostnamesInput() interface{}
	ExtKeyUsage() *[]*string
	SetExtKeyUsage(val *[]*string)
	ExtKeyUsageInput() *[]*string
	ExtKeyUsageOids() *[]*string
	SetExtKeyUsageOids(val *[]*string)
	ExtKeyUsageOidsInput() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GenerateLease() interface{}
	SetGenerateLease(val interface{})
	GenerateLeaseInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuerRef() *string
	SetIssuerRef(val *string)
	IssuerRefInput() *string
	KeyBits() *float64
	SetKeyBits(val *float64)
	KeyBitsInput() *float64
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
	Locality() *[]*string
	SetLocality(val *[]*string)
	LocalityInput() *[]*string
	MaxTtl() *string
	SetMaxTtl(val *string)
	MaxTtlInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NoStore() interface{}
	SetNoStore(val interface{})
	NoStoreInput() interface{}
	NoStoreMetadata() interface{}
	SetNoStoreMetadata(val interface{})
	NoStoreMetadataInput() interface{}
	NotAfter() *string
	SetNotAfter(val *string)
	NotAfterInput() *string
	NotBeforeDuration() *string
	SetNotBeforeDuration(val *string)
	NotBeforeDurationInput() *string
	Organization() *[]*string
	SetOrganization(val *[]*string)
	OrganizationInput() *[]*string
	Ou() *[]*string
	SetOu(val *[]*string)
	OuInput() *[]*string
	PolicyIdentifier() PkiSecretBackendRolePolicyIdentifierList
	PolicyIdentifierInput() interface{}
	PolicyIdentifiers() *[]*string
	SetPolicyIdentifiers(val *[]*string)
	PolicyIdentifiersInput() *[]*string
	PostalCode() *[]*string
	SetPostalCode(val *[]*string)
	PostalCodeInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	Province() *[]*string
	SetProvince(val *[]*string)
	ProvinceInput() *[]*string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RequireCn() interface{}
	SetRequireCn(val interface{})
	RequireCnInput() interface{}
	SerialNumberSource() *string
	SetSerialNumberSource(val *string)
	SerialNumberSourceInput() *string
	ServerFlag() interface{}
	SetServerFlag(val interface{})
	ServerFlagInput() interface{}
	SignatureBits() *float64
	SetSignatureBits(val *float64)
	SignatureBitsInput() *float64
	StreetAddress() *[]*string
	SetStreetAddress(val *[]*string)
	StreetAddressInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	UseCsrCommonName() interface{}
	SetUseCsrCommonName(val interface{})
	UseCsrCommonNameInput() interface{}
	UseCsrSans() interface{}
	SetUseCsrSans(val interface{})
	UseCsrSansInput() interface{}
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
	PutPolicyIdentifier(value interface{})
	ResetAllowAnyName()
	ResetAllowBareDomains()
	ResetAllowedDomains()
	ResetAllowedDomainsTemplate()
	ResetAllowedOtherSans()
	ResetAllowedSerialNumbers()
	ResetAllowedUriSans()
	ResetAllowedUriSansTemplate()
	ResetAllowedUserIds()
	ResetAllowGlobDomains()
	ResetAllowIpSans()
	ResetAllowLocalhost()
	ResetAllowSubdomains()
	ResetAllowWildcardCertificates()
	ResetBasicConstraintsValidForNonCa()
	ResetClientFlag()
	ResetCnValidations()
	ResetCodeSigningFlag()
	ResetCountry()
	ResetEmailProtectionFlag()
	ResetEnforceHostnames()
	ResetExtKeyUsage()
	ResetExtKeyUsageOids()
	ResetGenerateLease()
	ResetId()
	ResetIssuerRef()
	ResetKeyBits()
	ResetKeyType()
	ResetKeyUsage()
	ResetLocality()
	ResetMaxTtl()
	ResetNamespace()
	ResetNoStore()
	ResetNoStoreMetadata()
	ResetNotAfter()
	ResetNotBeforeDuration()
	ResetOrganization()
	ResetOu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyIdentifier()
	ResetPolicyIdentifiers()
	ResetPostalCode()
	ResetProvince()
	ResetRequireCn()
	ResetSerialNumberSource()
	ResetServerFlag()
	ResetSignatureBits()
	ResetStreetAddress()
	ResetTtl()
	ResetUseCsrCommonName()
	ResetUseCsrSans()
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

// The jsii proxy struct for PkiSecretBackendRole
type jsiiProxy_PkiSecretBackendRole struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowAnyName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowAnyNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAnyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowBareDomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowBareDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowBareDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowBareDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedDomainsTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedDomainsTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedDomainsTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedDomainsTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedOtherSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOtherSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedOtherSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOtherSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedSerialNumbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedSerialNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedSerialNumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedSerialNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUriSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUriSansTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedUriSansTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUriSansTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedUriSansTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUserIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUserIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowedUserIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUserIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowGlobDomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGlobDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowGlobDomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGlobDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowIpSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIpSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowIpSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowIpSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowLocalhost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalhost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowLocalhostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLocalhostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowWildcardCertificates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWildcardCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) AllowWildcardCertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWildcardCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) BasicConstraintsValidForNonCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicConstraintsValidForNonCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) BasicConstraintsValidForNonCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicConstraintsValidForNonCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ClientFlag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ClientFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CnValidations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cnValidations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CnValidationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cnValidationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CodeSigningFlag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigningFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CodeSigningFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigningFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Country() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) CountryInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) EmailProtectionFlag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtectionFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) EmailProtectionFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailProtectionFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) EnforceHostnames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) EnforceHostnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ExtKeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ExtKeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ExtKeyUsageOids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extKeyUsageOids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ExtKeyUsageOidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extKeyUsageOidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) GenerateLease() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateLease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) GenerateLeaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateLeaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) IssuerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) IssuerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyUsage() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) KeyUsageInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Locality() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) LocalityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) MaxTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) MaxTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NoStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NoStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NoStoreMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStoreMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NoStoreMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noStoreMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NotAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NotBeforeDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) NotBeforeDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBeforeDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Organization() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) OrganizationInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Ou() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ou",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) OuInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ouInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PolicyIdentifier() PkiSecretBackendRolePolicyIdentifierList {
	var returns PkiSecretBackendRolePolicyIdentifierList
	_jsii_.Get(
		j,
		"policyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PolicyIdentifierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PolicyIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PolicyIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PostalCode() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) PostalCodeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Province() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ProvinceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) RequireCn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) RequireCnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) SerialNumberSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumberSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) SerialNumberSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumberSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ServerFlag() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) ServerFlagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverFlagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) SignatureBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) SignatureBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) StreetAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) StreetAddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UseCsrCommonName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrCommonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UseCsrCommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrCommonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UseCsrSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UseCsrSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsrSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UsePss() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePss",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRole) UsePssInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePssInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_role vault_pki_secret_backend_role} Resource.
func NewPkiSecretBackendRole(scope constructs.Construct, id *string, config *PkiSecretBackendRoleConfig) PkiSecretBackendRole {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendRole{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_role vault_pki_secret_backend_role} Resource.
func NewPkiSecretBackendRole_Override(p PkiSecretBackendRole, scope constructs.Construct, id *string, config *PkiSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowAnyName(val interface{}) {
	if err := j.validateSetAllowAnyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAnyName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowBareDomains(val interface{}) {
	if err := j.validateSetAllowBareDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowBareDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedDomains(val *[]*string) {
	if err := j.validateSetAllowedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedDomainsTemplate(val interface{}) {
	if err := j.validateSetAllowedDomainsTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDomainsTemplate",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedOtherSans(val *[]*string) {
	if err := j.validateSetAllowedOtherSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOtherSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedSerialNumbers(val *[]*string) {
	if err := j.validateSetAllowedSerialNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedSerialNumbers",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedUriSans(val *[]*string) {
	if err := j.validateSetAllowedUriSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUriSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedUriSansTemplate(val interface{}) {
	if err := j.validateSetAllowedUriSansTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUriSansTemplate",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowedUserIds(val *[]*string) {
	if err := j.validateSetAllowedUserIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUserIds",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowGlobDomains(val interface{}) {
	if err := j.validateSetAllowGlobDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGlobDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowIpSans(val interface{}) {
	if err := j.validateSetAllowIpSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowIpSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowLocalhost(val interface{}) {
	if err := j.validateSetAllowLocalhostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLocalhost",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowSubdomains(val interface{}) {
	if err := j.validateSetAllowSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowSubdomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetAllowWildcardCertificates(val interface{}) {
	if err := j.validateSetAllowWildcardCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowWildcardCertificates",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetBasicConstraintsValidForNonCa(val interface{}) {
	if err := j.validateSetBasicConstraintsValidForNonCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicConstraintsValidForNonCa",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetClientFlag(val interface{}) {
	if err := j.validateSetClientFlagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientFlag",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetCnValidations(val *[]*string) {
	if err := j.validateSetCnValidationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cnValidations",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetCodeSigningFlag(val interface{}) {
	if err := j.validateSetCodeSigningFlagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningFlag",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetCountry(val *[]*string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetEmailProtectionFlag(val interface{}) {
	if err := j.validateSetEmailProtectionFlagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailProtectionFlag",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetEnforceHostnames(val interface{}) {
	if err := j.validateSetEnforceHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceHostnames",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetExtKeyUsage(val *[]*string) {
	if err := j.validateSetExtKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extKeyUsage",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetExtKeyUsageOids(val *[]*string) {
	if err := j.validateSetExtKeyUsageOidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extKeyUsageOids",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetGenerateLease(val interface{}) {
	if err := j.validateSetGenerateLeaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generateLease",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetIssuerRef(val *string) {
	if err := j.validateSetIssuerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerRef",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetKeyBits(val *float64) {
	if err := j.validateSetKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetKeyUsage(val *[]*string) {
	if err := j.validateSetKeyUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyUsage",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetLocality(val *[]*string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetMaxTtl(val *string) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetNoStore(val interface{}) {
	if err := j.validateSetNoStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStore",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetNoStoreMetadata(val interface{}) {
	if err := j.validateSetNoStoreMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noStoreMetadata",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetNotAfter(val *string) {
	if err := j.validateSetNotAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notAfter",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetNotBeforeDuration(val *string) {
	if err := j.validateSetNotBeforeDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notBeforeDuration",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetOrganization(val *[]*string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetOu(val *[]*string) {
	if err := j.validateSetOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ou",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetPolicyIdentifiers(val *[]*string) {
	if err := j.validateSetPolicyIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyIdentifiers",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetPostalCode(val *[]*string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetProvince(val *[]*string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetRequireCn(val interface{}) {
	if err := j.validateSetRequireCnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetSerialNumberSource(val *string) {
	if err := j.validateSetSerialNumberSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serialNumberSource",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetServerFlag(val interface{}) {
	if err := j.validateSetServerFlagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverFlag",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetSignatureBits(val *float64) {
	if err := j.validateSetSignatureBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetStreetAddress(val *[]*string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetUseCsrCommonName(val interface{}) {
	if err := j.validateSetUseCsrCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsrCommonName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetUseCsrSans(val interface{}) {
	if err := j.validateSetUseCsrSansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsrSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRole)SetUsePss(val interface{}) {
	if err := j.validateSetUsePssParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePss",
		val,
	)
}

// Generates CDKTN code for importing a PkiSecretBackendRole resource upon running "cdktn plan <stack-name>".
func PkiSecretBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
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
func PkiSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRole) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) PutPolicyIdentifier(value interface{}) {
	if err := p.validatePutPolicyIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPolicyIdentifier",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowAnyName() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowAnyName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowBareDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowBareDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedDomainsTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDomainsTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedOtherSans() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedOtherSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedSerialNumbers() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedSerialNumbers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedUriSans() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedUriSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedUriSansTemplate() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedUriSansTemplate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowedUserIds() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedUserIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowGlobDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowGlobDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowIpSans() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowIpSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowLocalhost() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowLocalhost",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowSubdomains() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowSubdomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetAllowWildcardCertificates() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowWildcardCertificates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetBasicConstraintsValidForNonCa() {
	_jsii_.InvokeVoid(
		p,
		"resetBasicConstraintsValidForNonCa",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetClientFlag() {
	_jsii_.InvokeVoid(
		p,
		"resetClientFlag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetCnValidations() {
	_jsii_.InvokeVoid(
		p,
		"resetCnValidations",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetCodeSigningFlag() {
	_jsii_.InvokeVoid(
		p,
		"resetCodeSigningFlag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetCountry() {
	_jsii_.InvokeVoid(
		p,
		"resetCountry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetEmailProtectionFlag() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailProtectionFlag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetEnforceHostnames() {
	_jsii_.InvokeVoid(
		p,
		"resetEnforceHostnames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetExtKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetExtKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetExtKeyUsageOids() {
	_jsii_.InvokeVoid(
		p,
		"resetExtKeyUsageOids",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetGenerateLease() {
	_jsii_.InvokeVoid(
		p,
		"resetGenerateLease",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetIssuerRef() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerRef",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetKeyBits() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetKeyType() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetLocality() {
	_jsii_.InvokeVoid(
		p,
		"resetLocality",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetNoStore() {
	_jsii_.InvokeVoid(
		p,
		"resetNoStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetNoStoreMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetNoStoreMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetNotAfter() {
	_jsii_.InvokeVoid(
		p,
		"resetNotAfter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetNotBeforeDuration() {
	_jsii_.InvokeVoid(
		p,
		"resetNotBeforeDuration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetOrganization() {
	_jsii_.InvokeVoid(
		p,
		"resetOrganization",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetOu() {
	_jsii_.InvokeVoid(
		p,
		"resetOu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetPolicyIdentifier() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyIdentifier",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetPolicyIdentifiers() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyIdentifiers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetPostalCode() {
	_jsii_.InvokeVoid(
		p,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetProvince() {
	_jsii_.InvokeVoid(
		p,
		"resetProvince",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetRequireCn() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireCn",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetSerialNumberSource() {
	_jsii_.InvokeVoid(
		p,
		"resetSerialNumberSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetServerFlag() {
	_jsii_.InvokeVoid(
		p,
		"resetServerFlag",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetSignatureBits() {
	_jsii_.InvokeVoid(
		p,
		"resetSignatureBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		p,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetUseCsrCommonName() {
	_jsii_.InvokeVoid(
		p,
		"resetUseCsrCommonName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetUseCsrSans() {
	_jsii_.InvokeVoid(
		p,
		"resetUseCsrSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) ResetUsePss() {
	_jsii_.InvokeVoid(
		p,
		"resetUsePss",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

