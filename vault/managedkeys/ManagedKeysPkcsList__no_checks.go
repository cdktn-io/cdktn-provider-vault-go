// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedkeys

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedKeysPkcsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysPkcsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedKeysPkcsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedKeysPkcsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedKeysPkcsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

