// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package oktaauthbackend

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OktaAuthBackendUserList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OktaAuthBackendUserList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OktaAuthBackendUserList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OktaAuthBackendUserList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OktaAuthBackendUserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OktaAuthBackendUserList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OktaAuthBackendUserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOktaAuthBackendUserListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

