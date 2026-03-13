// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedkeys

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKeysGcpList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysGcpList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysGcpList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysGcpList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysGcpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysGcpList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysGcpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKeysGcpListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

