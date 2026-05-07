// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datavaultpluginruntimes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataVaultPluginRuntimesRuntimesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataVaultPluginRuntimesRuntimesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

