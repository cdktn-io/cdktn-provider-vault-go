// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package audit

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Audit) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_Audit) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_Audit) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAudit_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAudit_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAudit_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAudit_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetLocalParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetOptionsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Audit) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAuditParameters(scope constructs.Construct, id *string, config *AuditConfig) error {
	return nil
}

