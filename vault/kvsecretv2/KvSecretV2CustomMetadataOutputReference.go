// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kvsecretv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/kvsecretv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KvSecretV2CustomMetadataOutputReference interface {
	cdktn.ComplexObject
	CasRequired() interface{}
	SetCasRequired(val interface{})
	CasRequiredInput() interface{}
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
	Data() *map[string]*string
	SetData(val *map[string]*string)
	DataInput() *map[string]*string
	DeleteVersionAfter() *float64
	SetDeleteVersionAfter(val *float64)
	DeleteVersionAfterInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *KvSecretV2CustomMetadata
	SetInternalValue(val *KvSecretV2CustomMetadata)
	MaxVersions() *float64
	SetMaxVersions(val *float64)
	MaxVersionsInput() *float64
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
	ResetCasRequired()
	ResetData()
	ResetDeleteVersionAfter()
	ResetMaxVersions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KvSecretV2CustomMetadataOutputReference
type jsiiProxy_KvSecretV2CustomMetadataOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) CasRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"casRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) CasRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"casRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) DeleteVersionAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteVersionAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) DeleteVersionAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteVersionAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) InternalValue() *KvSecretV2CustomMetadata {
	var returns *KvSecretV2CustomMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) MaxVersions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) MaxVersionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKvSecretV2CustomMetadataOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KvSecretV2CustomMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewKvSecretV2CustomMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KvSecretV2CustomMetadataOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2CustomMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKvSecretV2CustomMetadataOutputReference_Override(k KvSecretV2CustomMetadataOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.kvSecretV2.KvSecretV2CustomMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetCasRequired(val interface{}) {
	if err := j.validateSetCasRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"casRequired",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetDeleteVersionAfter(val *float64) {
	if err := j.validateSetDeleteVersionAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVersionAfter",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetInternalValue(val *KvSecretV2CustomMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetMaxVersions(val *float64) {
	if err := j.validateSetMaxVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxVersions",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KvSecretV2CustomMetadataOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ResetCasRequired() {
	_jsii_.InvokeVoid(
		k,
		"resetCasRequired",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		k,
		"resetData",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ResetDeleteVersionAfter() {
	_jsii_.InvokeVoid(
		k,
		"resetDeleteVersionAfter",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ResetMaxVersions() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxVersions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KvSecretV2CustomMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

