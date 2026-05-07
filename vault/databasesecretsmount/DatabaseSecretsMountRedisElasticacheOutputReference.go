// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/databasesecretsmount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretsMountRedisElasticacheOutputReference interface {
	cdktn.ComplexObject
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
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
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	PluginVersion() *string
	SetPluginVersion(val *string)
	PluginVersionInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RootRotationStatements() *[]*string
	SetRootRotationStatements(val *[]*string)
	RootRotationStatementsInput() *[]*string
	RotationPeriod() *float64
	SetRotationPeriod(val *float64)
	RotationPeriodInput() *float64
	RotationSchedule() *string
	SetRotationSchedule(val *string)
	RotationScheduleInput() *string
	RotationWindow() *float64
	SetRotationWindow(val *float64)
	RotationWindowInput() *float64
	SkipStaticRoleImportRotation() interface{}
	SetSkipStaticRoleImportRotation(val interface{})
	SkipStaticRoleImportRotationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VerifyConnection() interface{}
	SetVerifyConnection(val interface{})
	VerifyConnectionInput() interface{}
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
	ResetAllowedRoles()
	ResetData()
	ResetDisableAutomatedRotation()
	ResetPassword()
	ResetPasswordPolicy()
	ResetPluginName()
	ResetPluginVersion()
	ResetRegion()
	ResetRootRotationStatements()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetSkipStaticRoleImportRotation()
	ResetUsername()
	ResetVerifyConnection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseSecretsMountRedisElasticacheOutputReference
type jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PluginVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) PluginVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RootRotationStatements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RootRotationStatementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) SkipStaticRoleImportRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) SkipStaticRoleImportRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretsMountRedisElasticacheOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabaseSecretsMountRedisElasticacheOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretsMountRedisElasticacheOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountRedisElasticacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabaseSecretsMountRedisElasticacheOutputReference_Override(d DatabaseSecretsMountRedisElasticacheOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountRedisElasticacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetPluginVersion(val *string) {
	if err := j.validateSetPluginVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetRootRotationStatements(val *[]*string) {
	if err := j.validateSetRootRotationStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootRotationStatements",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetSkipStaticRoleImportRotation(val interface{}) {
	if err := j.validateSetSkipStaticRoleImportRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipStaticRoleImportRotation",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		d,
		"resetData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetPluginVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetRootRotationStatements() {
	_jsii_.InvokeVoid(
		d,
		"resetRootRotationStatements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetSkipStaticRoleImportRotation() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipStaticRoleImportRotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountRedisElasticacheOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

