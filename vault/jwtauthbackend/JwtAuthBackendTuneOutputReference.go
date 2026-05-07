// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package jwtauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jwtauthbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JwtAuthBackendTuneOutputReference interface {
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

// The jsii proxy struct for JwtAuthBackendTuneOutputReference
type jsiiProxy_JwtAuthBackendTuneOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AllowedResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AllowedResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) DefaultLeaseTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLeaseTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) DefaultLeaseTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLeaseTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ListingVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ListingVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) MaxLeaseTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLeaseTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) MaxLeaseTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxLeaseTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) PassthroughRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) PassthroughRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"passthroughRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


func NewJwtAuthBackendTuneOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) JwtAuthBackendTuneOutputReference {
	_init_.Initialize()

	if err := validateNewJwtAuthBackendTuneOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_JwtAuthBackendTuneOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.jwtAuthBackend.JwtAuthBackendTuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewJwtAuthBackendTuneOutputReference_Override(j JwtAuthBackendTuneOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.jwtAuthBackend.JwtAuthBackendTuneOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		j,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetAllowedResponseHeaders(val *[]*string) {
	if err := j.validateSetAllowedResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetAuditNonHmacRequestKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacRequestKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetAuditNonHmacResponseKeys(val *[]*string) {
	if err := j.validateSetAuditNonHmacResponseKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetDefaultLeaseTtl(val *string) {
	if err := j.validateSetDefaultLeaseTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetListingVisibility(val *string) {
	if err := j.validateSetListingVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listingVisibility",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetMaxLeaseTtl(val *string) {
	if err := j.validateSetMaxLeaseTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetPassthroughRequestHeaders(val *[]*string) {
	if err := j.validateSetPassthroughRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetAllowedResponseHeaders() {
	_jsii_.InvokeVoid(
		j,
		"resetAllowedResponseHeaders",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		j,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetDefaultLeaseTtl() {
	_jsii_.InvokeVoid(
		j,
		"resetDefaultLeaseTtl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetListingVisibility() {
	_jsii_.InvokeVoid(
		j,
		"resetListingVisibility",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetMaxLeaseTtl() {
	_jsii_.InvokeVoid(
		j,
		"resetMaxLeaseTtl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetPassthroughRequestHeaders() {
	_jsii_.InvokeVoid(
		j,
		"resetPassthroughRequestHeaders",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ResetTokenType() {
	_jsii_.InvokeVoid(
		j,
		"resetTokenType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := j.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackendTuneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

