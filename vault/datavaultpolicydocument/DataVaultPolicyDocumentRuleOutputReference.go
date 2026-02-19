// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpolicydocument

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/datavaultpolicydocument/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultPolicyDocumentRuleOutputReference interface {
	cdktn.ComplexObject
	AllowedParameter() DataVaultPolicyDocumentRuleAllowedParameterList
	AllowedParameterInput() interface{}
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	CapabilitiesInput() *[]*string
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
	DeniedParameter() DataVaultPolicyDocumentRuleDeniedParameterList
	DeniedParameterInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxWrappingTtl() *string
	SetMaxWrappingTtl(val *string)
	MaxWrappingTtlInput() *string
	MinWrappingTtl() *string
	SetMinWrappingTtl(val *string)
	MinWrappingTtlInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RequiredParameters() *[]*string
	SetRequiredParameters(val *[]*string)
	RequiredParametersInput() *[]*string
	SubscribeEventTypes() *[]*string
	SetSubscribeEventTypes(val *[]*string)
	SubscribeEventTypesInput() *[]*string
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
	PutAllowedParameter(value interface{})
	PutDeniedParameter(value interface{})
	ResetAllowedParameter()
	ResetDeniedParameter()
	ResetDescription()
	ResetMaxWrappingTtl()
	ResetMinWrappingTtl()
	ResetRequiredParameters()
	ResetSubscribeEventTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataVaultPolicyDocumentRuleOutputReference
type jsiiProxy_DataVaultPolicyDocumentRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) AllowedParameter() DataVaultPolicyDocumentRuleAllowedParameterList {
	var returns DataVaultPolicyDocumentRuleAllowedParameterList
	_jsii_.Get(
		j,
		"allowedParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) AllowedParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) CapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) DeniedParameter() DataVaultPolicyDocumentRuleDeniedParameterList {
	var returns DataVaultPolicyDocumentRuleDeniedParameterList
	_jsii_.Get(
		j,
		"deniedParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) DeniedParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deniedParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) MaxWrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) MaxWrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWrappingTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) MinWrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minWrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) MinWrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minWrappingTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) RequiredParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) RequiredParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) SubscribeEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subscribeEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) SubscribeEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subscribeEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataVaultPolicyDocumentRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataVaultPolicyDocumentRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataVaultPolicyDocumentRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataVaultPolicyDocumentRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultPolicyDocument.DataVaultPolicyDocumentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataVaultPolicyDocumentRuleOutputReference_Override(d DataVaultPolicyDocumentRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.dataVaultPolicyDocument.DataVaultPolicyDocumentRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetCapabilities(val *[]*string) {
	if err := j.validateSetCapabilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetMaxWrappingTtl(val *string) {
	if err := j.validateSetMaxWrappingTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWrappingTtl",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetMinWrappingTtl(val *string) {
	if err := j.validateSetMinWrappingTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWrappingTtl",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetRequiredParameters(val *[]*string) {
	if err := j.validateSetRequiredParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredParameters",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetSubscribeEventTypes(val *[]*string) {
	if err := j.validateSetSubscribeEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscribeEventTypes",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) PutAllowedParameter(value interface{}) {
	if err := d.validatePutAllowedParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllowedParameter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) PutDeniedParameter(value interface{}) {
	if err := d.validatePutDeniedParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeniedParameter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetAllowedParameter() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedParameter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetDeniedParameter() {
	_jsii_.InvokeVoid(
		d,
		"resetDeniedParameter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetMaxWrappingTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxWrappingTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetMinWrappingTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetMinWrappingTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetRequiredParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetRequiredParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ResetSubscribeEventTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetSubscribeEventTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultPolicyDocumentRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

