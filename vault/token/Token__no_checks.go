// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package token

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Token) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateImportFromParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (t *jsiiProxy_Token) validateMoveToIdParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_Token) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateToken_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateToken_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateToken_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetDisplayNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetExplicitMaxTtlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetMetadataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetNoDefaultPolicyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetNoParentParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetNumUsesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetPeriodParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetPoliciesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetRenewableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetRenewIncrementParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetRenewMinLeaseParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetRoleNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetTtlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Token) validateSetWrappingTtlParameters(val *string) error {
	return nil
}

func validateNewTokenParameters(scope constructs.Construct, id *string, config *TokenConfig) error {
	return nil
}

