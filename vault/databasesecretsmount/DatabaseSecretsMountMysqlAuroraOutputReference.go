// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/databasesecretsmount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretsMountMysqlAuroraOutputReference interface {
	cdktn.ComplexObject
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	ConnectionUrl() *string
	SetConnectionUrl(val *string)
	ConnectionUrlInput() *string
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
	MaxConnectionLifetime() *float64
	SetMaxConnectionLifetime(val *float64)
	MaxConnectionLifetimeInput() *float64
	MaxIdleConnections() *float64
	SetMaxIdleConnections(val *float64)
	MaxIdleConnectionsInput() *float64
	MaxOpenConnections() *float64
	SetMaxOpenConnections(val *float64)
	MaxOpenConnectionsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
	PasswordWo() *string
	SetPasswordWo(val *string)
	PasswordWoInput() *string
	PasswordWoVersion() *float64
	SetPasswordWoVersion(val *float64)
	PasswordWoVersionInput() *float64
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	PluginVersion() *string
	SetPluginVersion(val *string)
	PluginVersionInput() *string
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
	ServiceAccountJson() *string
	SetServiceAccountJson(val *string)
	ServiceAccountJsonInput() *string
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
	TlsCa() *string
	SetTlsCa(val *string)
	TlsCaInput() *string
	TlsCertificateKey() *string
	SetTlsCertificateKey(val *string)
	TlsCertificateKeyInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
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
	ResetAuthType()
	ResetConnectionUrl()
	ResetData()
	ResetDisableAutomatedRotation()
	ResetMaxConnectionLifetime()
	ResetMaxIdleConnections()
	ResetMaxOpenConnections()
	ResetPassword()
	ResetPasswordPolicy()
	ResetPasswordWo()
	ResetPasswordWoVersion()
	ResetPluginName()
	ResetPluginVersion()
	ResetRootRotationStatements()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetServiceAccountJson()
	ResetSkipStaticRoleImportRotation()
	ResetTlsCa()
	ResetTlsCertificateKey()
	ResetUsername()
	ResetUsernameTemplate()
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

// The jsii proxy struct for DatabaseSecretsMountMysqlAuroraOutputReference
type jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxConnectionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxConnectionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxIdleConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxIdleConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxOpenConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) MaxOpenConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PluginVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) PluginVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RootRotationStatements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RootRotationStatementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ServiceAccountJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ServiceAccountJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) SkipStaticRoleImportRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) SkipStaticRoleImportRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipStaticRoleImportRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TlsCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TlsCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TlsCertificateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) TlsCertificateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretsMountMysqlAuroraOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabaseSecretsMountMysqlAuroraOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretsMountMysqlAuroraOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountMysqlAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabaseSecretsMountMysqlAuroraOutputReference_Override(d DatabaseSecretsMountMysqlAuroraOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountMysqlAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetMaxConnectionLifetime(val *float64) {
	if err := j.validateSetMaxConnectionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetMaxIdleConnections(val *float64) {
	if err := j.validateSetMaxIdleConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetMaxOpenConnections(val *float64) {
	if err := j.validateSetMaxOpenConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOpenConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPasswordWoVersion(val *float64) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetPluginVersion(val *string) {
	if err := j.validateSetPluginVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetRootRotationStatements(val *[]*string) {
	if err := j.validateSetRootRotationStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootRotationStatements",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetServiceAccountJson(val *string) {
	if err := j.validateSetServiceAccountJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountJson",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetSkipStaticRoleImportRotation(val interface{}) {
	if err := j.validateSetSkipStaticRoleImportRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipStaticRoleImportRotation",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetTlsCa(val *string) {
	if err := j.validateSetTlsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCa",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetTlsCertificateKey(val *string) {
	if err := j.validateSetTlsCertificateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCertificateKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		d,
		"resetData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetMaxConnectionLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetMaxIdleConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetMaxOpenConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOpenConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetPluginVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetRootRotationStatements() {
	_jsii_.InvokeVoid(
		d,
		"resetRootRotationStatements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetServiceAccountJson() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetSkipStaticRoleImportRotation() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipStaticRoleImportRotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetTlsCa() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetTlsCertificateKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCertificateKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountMysqlAuroraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

