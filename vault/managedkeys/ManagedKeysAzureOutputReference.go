// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/managedkeys/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedKeysAzureOutputReference interface {
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
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyBits() *string
	SetKeyBits(val *string)
	KeyBitsInput() *string
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Usages() *[]*string
	SetUsages(val *[]*string)
	UsagesInput() *[]*string
	Uuid() *string
	VaultName() *string
	SetVaultName(val *string)
	VaultNameInput() *string
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
	ResetEnvironment()
	ResetKeyBits()
	ResetResource()
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

// The jsii proxy struct for ManagedKeysAzureOutputReference
type jsiiProxy_ManagedKeysAzureOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowGenerateKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowGenerateKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowGenerateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowReplaceKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowReplaceKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowReplaceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowStoreKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AllowStoreKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoreKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AnyMount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) AnyMountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anyMountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyBits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyBitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Usages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) UsagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) VaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference) VaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultNameInput",
		&returns,
	)
	return returns
}


func NewManagedKeysAzureOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ManagedKeysAzureOutputReference {
	_init_.Initialize()

	if err := validateNewManagedKeysAzureOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedKeysAzureOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewManagedKeysAzureOutputReference_Override(m ManagedKeysAzureOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.managedKeys.ManagedKeysAzureOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetAllowGenerateKey(val interface{}) {
	if err := j.validateSetAllowGenerateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowGenerateKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetAllowReplaceKey(val interface{}) {
	if err := j.validateSetAllowReplaceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowReplaceKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetAllowStoreKey(val interface{}) {
	if err := j.validateSetAllowStoreKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoreKey",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetAnyMount(val interface{}) {
	if err := j.validateSetAnyMountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anyMount",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetKeyBits(val *string) {
	if err := j.validateSetKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetResource(val *string) {
	if err := j.validateSetResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetUsages(val *[]*string) {
	if err := j.validateSetUsagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usages",
		val,
	)
}

func (j *jsiiProxy_ManagedKeysAzureOutputReference)SetVaultName(val *string) {
	if err := j.validateSetVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vaultName",
		val,
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetAllowGenerateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowGenerateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetAllowReplaceKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowReplaceKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetAllowStoreKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowStoreKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetAnyMount() {
	_jsii_.InvokeVoid(
		m,
		"resetAnyMount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		m,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetKeyBits() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyBits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		m,
		"resetResource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ResetUsages() {
	_jsii_.InvokeVoid(
		m,
		"resetUsages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedKeysAzureOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedKeysAzureOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

