// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/databasesecretsmount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseSecretsMountCassandraOutputReference interface {
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
	Data() *map[string]*string
	SetData(val *map[string]*string)
	DataInput() *map[string]*string
	DisableAutomatedRotation() interface{}
	SetDisableAutomatedRotation(val interface{})
	DisableAutomatedRotationInput() interface{}
	// Experimental.
	Fqn() *string
	Hosts() *[]*string
	SetHosts(val *[]*string)
	HostsInput() *[]*string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LocalDatacenter() *string
	SetLocalDatacenter(val *string)
	LocalDatacenterInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PemBundle() *string
	SetPemBundle(val *string)
	PemBundleInput() *string
	PemJson() *string
	SetPemJson(val *string)
	PemJsonInput() *string
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProtocolVersion() *float64
	SetProtocolVersion(val *float64)
	ProtocolVersionInput() *float64
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
	ResetConnectTimeout()
	ResetConsistency()
	ResetData()
	ResetDisableAutomatedRotation()
	ResetHosts()
	ResetInsecureTls()
	ResetLocalDatacenter()
	ResetPassword()
	ResetPemBundle()
	ResetPemJson()
	ResetPluginName()
	ResetPort()
	ResetProtocolVersion()
	ResetRootRotationStatements()
	ResetRotationPeriod()
	ResetRotationSchedule()
	ResetRotationWindow()
	ResetSkipVerification()
	ResetSocketKeepAlive()
	ResetTls()
	ResetTlsServerName()
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

// The jsii proxy struct for DatabaseSecretsMountCassandraOutputReference
type jsiiProxy_DatabaseSecretsMountCassandraOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Consistency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ConsistencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consistencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) DisableAutomatedRotation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) DisableAutomatedRotationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutomatedRotationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) LocalDatacenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDatacenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) LocalDatacenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDatacenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PemBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PemBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PemJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PemJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pemJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ProtocolVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ProtocolVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RootRotationStatements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RootRotationStatementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) RotationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rotationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) SkipVerification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) SkipVerificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipVerificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) SocketKeepAlive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) SocketKeepAliveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Tls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) TlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) TlsServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) TlsServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretsMountCassandraOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabaseSecretsMountCassandraOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretsMountCassandraOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretsMountCassandraOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountCassandraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabaseSecretsMountCassandraOutputReference_Override(d DatabaseSecretsMountCassandraOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.databaseSecretsMount.DatabaseSecretsMountCassandraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetConsistency(val *string) {
	if err := j.validateSetConsistencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consistency",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetDisableAutomatedRotation(val interface{}) {
	if err := j.validateSetDisableAutomatedRotationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutomatedRotation",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetHosts(val *[]*string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetLocalDatacenter(val *string) {
	if err := j.validateSetLocalDatacenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localDatacenter",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetPemBundle(val *string) {
	if err := j.validateSetPemBundleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemBundle",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetPemJson(val *string) {
	if err := j.validateSetPemJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pemJson",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetProtocolVersion(val *float64) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetRootRotationStatements(val *[]*string) {
	if err := j.validateSetRootRotationStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootRotationStatements",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetRotationPeriod(val *float64) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetRotationSchedule(val *string) {
	if err := j.validateSetRotationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationSchedule",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetRotationWindow(val *float64) {
	if err := j.validateSetRotationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationWindow",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetSkipVerification(val interface{}) {
	if err := j.validateSetSkipVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipVerification",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetSocketKeepAlive(val *string) {
	if err := j.validateSetSocketKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"socketKeepAlive",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetTls(val interface{}) {
	if err := j.validateSetTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tls",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetTlsServerName(val *string) {
	if err := j.validateSetTlsServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsServerName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountCassandraOutputReference)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetConsistency() {
	_jsii_.InvokeVoid(
		d,
		"resetConsistency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		d,
		"resetData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetDisableAutomatedRotation() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableAutomatedRotation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		d,
		"resetHosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		d,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetLocalDatacenter() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalDatacenter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetPemBundle() {
	_jsii_.InvokeVoid(
		d,
		"resetPemBundle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetPemJson() {
	_jsii_.InvokeVoid(
		d,
		"resetPemJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetRootRotationStatements() {
	_jsii_.InvokeVoid(
		d,
		"resetRootRotationStatements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetRotationSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetRotationWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetRotationWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetSkipVerification() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipVerification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetSocketKeepAlive() {
	_jsii_.InvokeVoid(
		d,
		"resetSocketKeepAlive",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		d,
		"resetTls",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetTlsServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountCassandraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

