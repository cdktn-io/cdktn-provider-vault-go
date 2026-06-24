// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/managedkeys/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedKeysPkcsOutputReference interface {
	cdktn.ComplexObject
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
	ForceRwSession() *string
	SetForceRwSession(val *string)
	ForceRwSessionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyBits() *string
	SetKeyBits(val *string)
	KeyBitsInput() *string
	KeyId() *string
	SetKeyId(val *string)
	KeyIdInput() *string
	KeyLabel() *string
	SetKeyLabel(val *string)
	KeyLabelInput() *string
	Library() *string
	SetLibrary(val *string)
	LibraryInput() *string
	MaxParallel() *float64
	SetMaxParallel(val *float64)
	MaxParallelInput() *float64
	Mechanism() *string
	SetMechanism(val *string)
	MechanismInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Pin() *string
	SetPin(val *string)
	PinInput() *string
	Slot() *string
	SetSlot(val *string)
	SlotInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenLabel() *string
	SetTokenLabel(val *string)
	TokenLabelInput() *string
	Usages() *[]*string
	SetUsages(val *[]*string)
	UsagesInput() *[]*string
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
	ResetForceRwSession()
	ResetKeyBits()
	ResetKeyId()
	ResetKeyLabel()
	ResetMaxParallel()
	ResetSlot()
	ResetTokenLabel()
	ResetUsages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedKeysPkcsOutputReference
type jsiiProxy_ManagedKeysPkcsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowGenerateKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowGenerateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowReplaceKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowReplaceKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowStoreKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AllowStoreKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AnyMount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) AnyMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Curve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"curve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) CurveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"curveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) ForceRwSession() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceRwSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) ForceRwSessionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceRwSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyBits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyBitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) KeyLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Library() *string {
	var returns *string
	_jsii_.Get(
		j,
		"library",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) LibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"libraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) MaxParallel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) MaxParallelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Mechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) MechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Pin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) PinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Slot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) SlotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) TokenLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) TokenLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Usages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) UsagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewManagedKeysPkcsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedKeysPkcsOutputReference {
	_init_.Initialize()

	if err := validateNewManagedKeysPkcsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKeysPkcsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysPkcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedKeysPkcsOutputReference_Override(m ManagedKeysPkcsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysPkcsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetAllowGenerateKey(val interface{}) {
	if err := j.validateSetAllowGenerateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGenerateKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetAllowReplaceKey(val interface{}) {
	if err := j.validateSetAllowReplaceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplaceKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetAllowStoreKey(val interface{}) {
	if err := j.validateSetAllowStoreKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoreKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetAnyMount(val interface{}) {
	if err := j.validateSetAnyMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyMount",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetCurve(val *string) {
	if err := j.validateSetCurveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"curve",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetForceRwSession(val *string) {
	if err := j.validateSetForceRwSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRwSession",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetKeyBits(val *string) {
	if err := j.validateSetKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetKeyId(val *string) {
	if err := j.validateSetKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyId",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetKeyLabel(val *string) {
	if err := j.validateSetKeyLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyLabel",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetLibrary(val *string) {
	if err := j.validateSetLibraryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"library",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetMaxParallel(val *float64) {
	if err := j.validateSetMaxParallelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallel",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetMechanism(val *string) {
	if err := j.validateSetMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mechanism",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetPin(val *string) {
	if err := j.validateSetPinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pin",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetSlot(val *string) {
	if err := j.validateSetSlotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slot",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetTokenLabel(val *string) {
	if err := j.validateSetTokenLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenLabel",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysPkcsOutputReference)SetUsages(val *[]*string) {
	if err := j.validateSetUsagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usages",
		val,
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetAllowGenerateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowGenerateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetAllowReplaceKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowReplaceKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetAllowStoreKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowStoreKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetAnyMount() {
	_jsii_.InvokeVoid(
		m,
		"resetAnyMount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetCurve() {
	_jsii_.InvokeVoid(
		m,
		"resetCurve",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetForceRwSession() {
	_jsii_.InvokeVoid(
		m,
		"resetForceRwSession",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetKeyBits() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyBits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetKeyLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetMaxParallel() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxParallel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetSlot() {
	_jsii_.InvokeVoid(
		m,
		"resetSlot",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetTokenLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ResetUsages() {
	_jsii_.InvokeVoid(
		m,
		"resetUsages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedKeysPkcsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

