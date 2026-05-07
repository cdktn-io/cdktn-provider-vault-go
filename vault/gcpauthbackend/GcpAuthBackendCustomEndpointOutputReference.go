// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-vault-go/vault/v17/gcpauthbackend/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpAuthBackendCustomEndpointOutputReference interface {
	cdktn.ComplexObject
	Api() *string
	SetApi(val *string)
	ApiInput() *string
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
	Compute() *string
	SetCompute(val *string)
	ComputeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Crm() *string
	SetCrm(val *string)
	CrmInput() *string
	// Experimental.
	Fqn() *string
	Iam() *string
	SetIam(val *string)
	IamInput() *string
	InternalValue() *GcpAuthBackendCustomEndpoint
	SetInternalValue(val *GcpAuthBackendCustomEndpoint)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetApi()
	ResetCompute()
	ResetCrm()
	ResetIam()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GcpAuthBackendCustomEndpointOutputReference
type jsiiProxy_GcpAuthBackendCustomEndpointOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Api() *string {
	var returns *string
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Compute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ComputeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Crm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) CrmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Iam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) IamInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) InternalValue() *GcpAuthBackendCustomEndpoint {
	var returns *GcpAuthBackendCustomEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGcpAuthBackendCustomEndpointOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GcpAuthBackendCustomEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewGcpAuthBackendCustomEndpointOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GcpAuthBackendCustomEndpointOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-vault.gcpAuthBackend.GcpAuthBackendCustomEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGcpAuthBackendCustomEndpointOutputReference_Override(g GcpAuthBackendCustomEndpointOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-vault.gcpAuthBackend.GcpAuthBackendCustomEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetApi(val *string) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetCompute(val *string) {
	if err := j.validateSetComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compute",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetCrm(val *string) {
	if err := j.validateSetCrmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crm",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetIam(val *string) {
	if err := j.validateSetIamParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iam",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetInternalValue(val *GcpAuthBackendCustomEndpoint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ResetApi() {
	_jsii_.InvokeVoid(
		g,
		"resetApi",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ResetCompute() {
	_jsii_.InvokeVoid(
		g,
		"resetCompute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ResetCrm() {
	_jsii_.InvokeVoid(
		g,
		"resetCrm",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ResetIam() {
	_jsii_.InvokeVoid(
		g,
		"resetIam",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GcpAuthBackendCustomEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

