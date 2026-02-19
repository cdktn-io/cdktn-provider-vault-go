// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/databasesecretbackendconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretBackendConnectionMysqlOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseSecretBackendConnectionMysql
	SetInternalValue(val *DatabaseSecretBackendConnectionMysql)
	MaxConnectionLifetime() *float64
	SetMaxConnectionLifetime(val *float64)
	MaxConnectionLifetimeInput() *float64
	MaxIdleConnections() *float64
	SetMaxIdleConnections(val *float64)
	MaxIdleConnectionsInput() *float64
	MaxOpenConnections() *float64
	SetMaxOpenConnections(val *float64)
	MaxOpenConnectionsInput() *float64
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordWo() *string
	SetPasswordWo(val *string)
	PasswordWoInput() *string
	PasswordWoVersion() *float64
	SetPasswordWoVersion(val *float64)
	PasswordWoVersionInput() *float64
	ServiceAccountJson() *string
	SetServiceAccountJson(val *string)
	ServiceAccountJsonInput() *string
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
	ResetAuthType()
	ResetConnectionUrl()
	ResetMaxConnectionLifetime()
	ResetMaxIdleConnections()
	ResetMaxOpenConnections()
	ResetPassword()
	ResetPasswordWo()
	ResetPasswordWoVersion()
	ResetServiceAccountJson()
	ResetTlsCa()
	ResetTlsCertificateKey()
	ResetUsername()
	ResetUsernameTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseSecretBackendConnectionMysqlOutputReference
type jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) InternalValue() *DatabaseSecretBackendConnectionMysql {
	var returns *DatabaseSecretBackendConnectionMysql
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxConnectionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxConnectionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxIdleConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxIdleConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxOpenConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) MaxOpenConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) PasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) PasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ServiceAccountJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ServiceAccountJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TlsCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TlsCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TlsCertificateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) TlsCertificateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretBackendConnectionMysqlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatabaseSecretBackendConnectionMysqlOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretBackendConnectionMysqlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionMysqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSecretBackendConnectionMysqlOutputReference_Override(d DatabaseSecretBackendConnectionMysqlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionMysqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetInternalValue(val *DatabaseSecretBackendConnectionMysql) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetMaxConnectionLifetime(val *float64) {
	if err := j.validateSetMaxConnectionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetMaxIdleConnections(val *float64) {
	if err := j.validateSetMaxIdleConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetMaxOpenConnections(val *float64) {
	if err := j.validateSetMaxOpenConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOpenConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetPasswordWoVersion(val *float64) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetServiceAccountJson(val *string) {
	if err := j.validateSetServiceAccountJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountJson",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetTlsCa(val *string) {
	if err := j.validateSetTlsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCa",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetTlsCertificateKey(val *string) {
	if err := j.validateSetTlsCertificateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCertificateKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetMaxConnectionLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetMaxIdleConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetMaxOpenConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOpenConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetServiceAccountJson() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetTlsCa() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetTlsCertificateKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCertificateKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

