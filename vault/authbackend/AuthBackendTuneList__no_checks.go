// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package authbackend

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthBackendTuneList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AuthBackendTuneList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AuthBackendTuneList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AuthBackendTuneList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthBackendTuneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthBackendTuneList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AuthBackendTuneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAuthBackendTuneListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

