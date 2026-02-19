// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kvsecret

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KvSecret) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateImportFromParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateMoveToIdParameters(id *string) error {
	return nil
}

func (k *jsiiProxy_KvSecret) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateKvSecret_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateKvSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKvSecret_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateKvSecret_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetDataJsonParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KvSecret) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewKvSecretParameters(scope constructs.Construct, id *string, config *KvSecretConfig) error {
	return nil
}

