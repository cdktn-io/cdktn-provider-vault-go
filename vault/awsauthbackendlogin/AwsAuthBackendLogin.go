// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awsauthbackendlogin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v16/awsauthbackendlogin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_login vault_aws_auth_backend_login}.
type AwsAuthBackendLogin interface {
	cdktn.TerraformResource
	Accessor() *string
	AuthType() *string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientToken() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamHttpRequestMethod() *string
	SetIamHttpRequestMethod(val *string)
	IamHttpRequestMethodInput() *string
	IamRequestBody() *string
	SetIamRequestBody(val *string)
	IamRequestBodyInput() *string
	IamRequestHeaders() *string
	SetIamRequestHeaders(val *string)
	IamRequestHeadersInput() *string
	IamRequestUrl() *string
	SetIamRequestUrl(val *string)
	IamRequestUrlInput() *string
	Id() *string
	SetId(val *string)
	Identity() *string
	SetIdentity(val *string)
	IdentityInput() *string
	IdInput() *string
	LeaseDuration() *float64
	LeaseStartTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Metadata() cdktn.StringMap
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Nonce() *string
	SetNonce(val *string)
	NonceInput() *string
	Pkcs7() *string
	SetPkcs7(val *string)
	Pkcs7Input() *string
	Policies() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Renewable() cdktn.IResolvable
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Signature() *string
	SetSignature(val *string)
	SignatureInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetBackend()
	ResetIamHttpRequestMethod()
	ResetIamRequestBody()
	ResetIamRequestHeaders()
	ResetIamRequestUrl()
	ResetId()
	ResetIdentity()
	ResetNamespace()
	ResetNonce()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPkcs7()
	ResetRole()
	ResetSignature()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AwsAuthBackendLogin
type jsiiProxy_AwsAuthBackendLogin struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAuthBackendLogin) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamHttpRequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamHttpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamHttpRequestMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamHttpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestHeadersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IamRequestUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRequestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) LeaseStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Metadata() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Nonce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) NonceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Pkcs7() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs7",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Pkcs7Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pkcs7Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Policies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Renewable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"renewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) Signature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) SignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendLogin) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_login vault_aws_auth_backend_login} Resource.
func NewAwsAuthBackendLogin(scope constructs.Construct, id *string, config *AwsAuthBackendLoginConfig) AwsAuthBackendLogin {
	_init_.Initialize()

	if err := validateNewAwsAuthBackendLoginParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAuthBackendLogin{}

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/aws_auth_backend_login vault_aws_auth_backend_login} Resource.
func NewAwsAuthBackendLogin_Override(a AwsAuthBackendLogin, scope constructs.Construct, id *string, config *AwsAuthBackendLoginConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetIamHttpRequestMethod(val *string) {
	if err := j.validateSetIamHttpRequestMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamHttpRequestMethod",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetIamRequestBody(val *string) {
	if err := j.validateSetIamRequestBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRequestBody",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetIamRequestHeaders(val *string) {
	if err := j.validateSetIamRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetIamRequestUrl(val *string) {
	if err := j.validateSetIamRequestUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetIdentity(val *string) {
	if err := j.validateSetIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetNonce(val *string) {
	if err := j.validateSetNonceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonce",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetPkcs7(val *string) {
	if err := j.validateSetPkcs7Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pkcs7",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendLogin)SetSignature(val *string) {
	if err := j.validateSetSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signature",
		val,
	)
}

// Generates CDKTN code for importing a AwsAuthBackendLogin resource upon running "cdktn plan <stack-name>".
func AwsAuthBackendLogin_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAuthBackendLogin_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AwsAuthBackendLogin_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendLogin_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendLogin_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendLogin_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAuthBackendLogin_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAuthBackendLogin_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAuthBackendLogin_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-vault.awsAuthBackendLogin.AwsAuthBackendLogin",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetIamHttpRequestMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIamHttpRequestMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetIamRequestBody() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRequestBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetIamRequestHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRequestHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetIamRequestUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRequestUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetNonce() {
	_jsii_.InvokeVoid(
		a,
		"resetNonce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetPkcs7() {
	_jsii_.InvokeVoid(
		a,
		"resetPkcs7",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetRole() {
	_jsii_.InvokeVoid(
		a,
		"resetRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) ResetSignature() {
	_jsii_.InvokeVoid(
		a,
		"resetSignature",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendLogin) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendLogin) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

