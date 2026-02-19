// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/databasesecretbackendconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretBackendConnectionCassandraOutputReference interface {
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
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	Consistency() *string
	SetConsistency(val *string)
	ConsistencyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	InternalValue() *DatabaseSecretBackendConnectionCassandra
	SetInternalValue(val *DatabaseSecretBackendConnectionCassandra)
	LocalDatacenter() *string
	SetLocalDatacenter(val *string)
	LocalDatacenterInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PemBundle() *string
	SetPemBundle(val *string)
	PemBundleInput() *string
	PemJson() *string
	SetPemJson(val *string)
	PemJsonInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProtocolVersion() *float64
	SetProtocolVersion(val *float64)
	ProtocolVersionInput() *float64
	SkipVerification() interface{}
	SetSkipVerification(val interface{})
	SkipVerificationInput() interface{}
	SocketKeepAlive() *string
	SetSocketKeepAlive(val *string)
	SocketKeepAliveInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Tls() interface{}
	SetTls(val interface{})
	TlsInput() interface{}
	TlsServerName() *string
	SetTlsServerName(val *string)
	TlsServerNameInput() *string
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
	ResetConnectTimeout()
	ResetConsistency()
	ResetHosts()
	ResetInsecureTls()
	ResetLocalDatacenter()
	ResetPassword()
	ResetPemBundle()
	ResetPemJson()
	ResetPort()
	ResetProtocolVersion()
	ResetSkipVerification()
	ResetSocketKeepAlive()
	ResetTls()
	ResetTlsServerName()
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

// The jsii proxy struct for DatabaseSecretBackendConnectionCassandraOutputReference
type jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Consistency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ConsistencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) InternalValue() *DatabaseSecretBackendConnectionCassandra {
	var returns *DatabaseSecretBackendConnectionCassandra
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) LocalDatacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDatacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) LocalDatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDatacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PemBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PemBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PemJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PemJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ProtocolVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ProtocolVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) SkipVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) SkipVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) SocketKeepAlive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) SocketKeepAliveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Tls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) TlsServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) TlsServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretBackendConnectionCassandraOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatabaseSecretBackendConnectionCassandraOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretBackendConnectionCassandraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionCassandraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSecretBackendConnectionCassandraOutputReference_Override(d DatabaseSecretBackendConnectionCassandraOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionCassandraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetConsistency(val *string) {
	if err := j.validateSetConsistencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consistency",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetHosts(val *[]*string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetInternalValue(val *DatabaseSecretBackendConnectionCassandra) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetLocalDatacenter(val *string) {
	if err := j.validateSetLocalDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDatacenter",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetPemBundle(val *string) {
	if err := j.validateSetPemBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemBundle",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetPemJson(val *string) {
	if err := j.validateSetPemJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemJson",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetProtocolVersion(val *float64) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetSkipVerification(val interface{}) {
	if err := j.validateSetSkipVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipVerification",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetSocketKeepAlive(val *string) {
	if err := j.validateSetSocketKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"socketKeepAlive",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetTls(val interface{}) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetTlsServerName(val *string) {
	if err := j.validateSetTlsServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsServerName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetConsistency() {
	_jsii_.InvokeVoid(
		d,
		"resetConsistency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		d,
		"resetHosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		d,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetLocalDatacenter() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalDatacenter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetPemBundle() {
	_jsii_.InvokeVoid(
		d,
		"resetPemBundle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetPemJson() {
	_jsii_.InvokeVoid(
		d,
		"resetPemJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetSkipVerification() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipVerification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetSocketKeepAlive() {
	_jsii_.InvokeVoid(
		d,
		"resetSocketKeepAlive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		d,
		"resetTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetTlsServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionCassandraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

