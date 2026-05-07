// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/ldapauthbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LdapAuthBackendTuneOutputReference interface {
	cdktn.ComplexObject
	AllowedResponseHeaders() *[]*string
	SetAllowedResponseHeaders(val *[]*string)
	AllowedResponseHeadersInput() *[]*string
	AuditNonHmacRequestKeys() *[]*string
	SetAuditNonHmacRequestKeys(val *[]*string)
	AuditNonHmacRequestKeysInput() *[]*string
	AuditNonHmacResponseKeys() *[]*string
	SetAuditNonHmacResponseKeys(val *[]*string)
	AuditNonHmacResponseKeysInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultLeaseTtl() *string
	SetDefaultLeaseTtl(val *string)
	DefaultLeaseTtlInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListingVisibility() *string
	SetListingVisibility(val *string)
	ListingVisibilityInput() *string
	MaxLeaseTtl() *string
	SetMaxLeaseTtl(val *string)
	MaxLeaseTtlInput() *string
	PassthroughRequestHeaders() *[]*string
	SetPassthroughRequestHeaders(val *[]*string)
	PassthroughRequestHeadersInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAllowedResponseHeaders()
	ResetAuditNonHmacRequestKeys()
	ResetAuditNonHmacResponseKeys()
	ResetDefaultLeaseTtl()
	ResetListingVisibility()
	ResetMaxLeaseTtl()
	ResetPassthroughRequestHeaders()
	ResetTokenType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LdapAuthBackendTuneOutputReference
type jsiiProxy_LdapAuthBackendTuneOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AllowedResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AllowedResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) DefaultLeaseTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLeaseTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) DefaultLeaseTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLeaseTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) ListingVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) ListingVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) MaxLeaseTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLeaseTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) MaxLeaseTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLeaseTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) PassthroughRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) PassthroughRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


func NewLdapAuthBackendTuneOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LdapAuthBackendTuneOutputReference {
	_init_.Initialize()

	if err := validateNewLdapAuthBackendTuneOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LdapAuthBackendTuneOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackendTuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLdapAuthBackendTuneOutputReference_Override(l LdapAuthBackendTuneOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.ldapAuthBackend.LdapAuthBackendTuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetAllowedResponseHeaders(val *[]*string) {
	if err := j.validateSetAllowedResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetAuditNonHmacRequestKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacRequestKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetAuditNonHmacResponseKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacResponseKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetDefaultLeaseTtl(val *string) {
	if err := j.validateSetDefaultLeaseTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtl",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetListingVisibility(val *string) {
	if err := j.validateSetListingVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingVisibility",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetMaxLeaseTtl(val *string) {
	if err := j.validateSetMaxLeaseTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtl",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetPassthroughRequestHeaders(val *[]*string) {
	if err := j.validateSetPassthroughRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LdapAuthBackendTuneOutputReference)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetAllowedResponseHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowedResponseHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		l,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		l,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetDefaultLeaseTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetDefaultLeaseTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetListingVisibility() {
	_jsii_.InvokeVoid(
		l,
		"resetListingVisibility",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetMaxLeaseTtl() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxLeaseTtl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetPassthroughRequestHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetPassthroughRequestHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ResetTokenType() {
	_jsii_.InvokeVoid(
		l,
		"resetTokenType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapAuthBackendTuneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

