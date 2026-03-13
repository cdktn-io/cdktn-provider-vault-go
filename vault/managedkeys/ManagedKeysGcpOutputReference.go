// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/managedkeys/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedKeysGcpOutputReference interface {
	cdktn.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	CryptoKey() *string
	SetCryptoKey(val *string)
	CryptoKeyInput() *string
	CryptoKeyVersion() *string
	SetCryptoKeyVersion(val *string)
	CryptoKeyVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyRing() *string
	SetKeyRing(val *string)
	KeyRingInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetCryptoKeyVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedKeysGcpOutputReference
type jsiiProxy_ManagedKeysGcpOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowGenerateKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowGenerateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowReplaceKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowReplaceKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowStoreKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AllowStoreKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AnyMount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) AnyMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CryptoKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CryptoKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CryptoKeyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) CryptoKeyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) KeyRing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) KeyRingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewManagedKeysGcpOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedKeysGcpOutputReference {
	_init_.Initialize()

	if err := validateNewManagedKeysGcpOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKeysGcpOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysGcpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedKeysGcpOutputReference_Override(m ManagedKeysGcpOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysGcpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetAllowGenerateKey(val interface{}) {
	if err := j.validateSetAllowGenerateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGenerateKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetAllowReplaceKey(val interface{}) {
	if err := j.validateSetAllowReplaceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplaceKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetAllowStoreKey(val interface{}) {
	if err := j.validateSetAllowStoreKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoreKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetAnyMount(val interface{}) {
	if err := j.validateSetAnyMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyMount",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetCredentials(val *string) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetCryptoKey(val *string) {
	if err := j.validateSetCryptoKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cryptoKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetCryptoKeyVersion(val *string) {
	if err := j.validateSetCryptoKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cryptoKeyVersion",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetKeyRing(val *string) {
	if err := j.validateSetKeyRingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRing",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysGcpOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ResetAllowGenerateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowGenerateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ResetAllowReplaceKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowReplaceKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ResetAllowStoreKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowStoreKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ResetAnyMount() {
	_jsii_.InvokeVoid(
		m,
		"resetAnyMount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ResetCryptoKeyVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetCryptoKeyVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysGcpOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedKeysGcpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

