// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/managedkeys/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedKeysAwsOutputReference interface {
	cdktn.ComplexObject
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	AllowGenerateKey() interface{}
	SetAllowGenerateKey(val interface{})
	AllowGenerateKeyInput() interface{}
	AllowReplaceKey() interface{}
	SetAllowReplaceKey(val interface{})
	AllowReplaceKeyInput() interface{}
	AllowStoreKey() interface{}
	SetAllowStoreKey(val interface{})
	AllowStoreKeyInput() interface{}
	AnyMount() interface{}
	SetAnyMount(val interface{})
	AnyMountInput() interface{}
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
	Curve() *string
	SetCurve(val *string)
	CurveInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyBits() *string
	SetKeyBits(val *string)
	KeyBitsInput() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Uuid() *string
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
	ResetAllowGenerateKey()
	ResetAllowReplaceKey()
	ResetAllowStoreKey()
	ResetAnyMount()
	ResetCurve()
	ResetEndpoint()
	ResetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedKeysAwsOutputReference
type jsiiProxy_ManagedKeysAwsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowGenerateKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowGenerateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowReplaceKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowReplaceKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowStoreKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AllowStoreKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AnyMount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) AnyMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Curve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"curve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) CurveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"curveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KeyBits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KeyBitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewManagedKeysAwsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedKeysAwsOutputReference {
	_init_.Initialize()

	if err := validateNewManagedKeysAwsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKeysAwsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysAwsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedKeysAwsOutputReference_Override(m ManagedKeysAwsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysAwsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetAllowGenerateKey(val interface{}) {
	if err := j.validateSetAllowGenerateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGenerateKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetAllowReplaceKey(val interface{}) {
	if err := j.validateSetAllowReplaceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplaceKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetAllowStoreKey(val interface{}) {
	if err := j.validateSetAllowStoreKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoreKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetAnyMount(val interface{}) {
	if err := j.validateSetAnyMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyMount",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetCurve(val *string) {
	if err := j.validateSetCurveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curve",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetKeyBits(val *string) {
	if err := j.validateSetKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAwsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetAllowGenerateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowGenerateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetAllowReplaceKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowReplaceKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetAllowStoreKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowStoreKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetAnyMount() {
	_jsii_.InvokeVoid(
		m,
		"resetAnyMount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetCurve() {
	_jsii_.InvokeVoid(
		m,
		"resetCurve",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAwsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

