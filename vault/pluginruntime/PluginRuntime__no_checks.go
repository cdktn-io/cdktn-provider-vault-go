// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pluginruntime

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PluginRuntime) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateImportFromParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateMoveToIdParameters(id *string) error {
	return nil
}

func (p *jsiiProxy_PluginRuntime) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validatePluginRuntime_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validatePluginRuntime_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePluginRuntime_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validatePluginRuntime_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetCgroupParentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetCpuNanosParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetMemoryBytesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetOciRuntimeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetRootlessParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PluginRuntime) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewPluginRuntimeParameters(scope constructs.Construct, id *string, config *PluginRuntimeConfig) error {
	return nil
}

