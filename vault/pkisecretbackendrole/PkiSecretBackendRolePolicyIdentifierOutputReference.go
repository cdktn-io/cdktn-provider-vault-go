// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/pkisecretbackendrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendRolePolicyIdentifierOutputReference interface {
	cdktn.ComplexObject
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
	Cps() *string
	SetCps(val *string)
	CpsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Notice() *string
	SetNotice(val *string)
	NoticeInput() *string
	Oid() *string
	SetOid(val *string)
	OidInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetCps()
	ResetNotice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PkiSecretBackendRolePolicyIdentifierOutputReference
type jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) Cps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) CpsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) Notice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) NoticeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noticeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) Oid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) OidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPkiSecretBackendRolePolicyIdentifierOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PkiSecretBackendRolePolicyIdentifierOutputReference {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendRolePolicyIdentifierOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRolePolicyIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPkiSecretBackendRolePolicyIdentifierOutputReference_Override(p PkiSecretBackendRolePolicyIdentifierOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.pkiSecretBackendRole.PkiSecretBackendRolePolicyIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetCps(val *string) {
	if err := j.validateSetCpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cps",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetNotice(val *string) {
	if err := j.validateSetNoticeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notice",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetOid(val *string) {
	if err := j.validateSetOidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oid",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ResetCps() {
	_jsii_.InvokeVoid(
		p,
		"resetCps",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ResetNotice() {
	_jsii_.InvokeVoid(
		p,
		"resetNotice",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRolePolicyIdentifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

