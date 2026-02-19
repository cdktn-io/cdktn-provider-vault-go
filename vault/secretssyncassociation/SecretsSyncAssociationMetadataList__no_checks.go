// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretssyncassociation

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsSyncAssociationMetadataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretsSyncAssociationMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretsSyncAssociationMetadataList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretsSyncAssociationMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretsSyncAssociationMetadataList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretsSyncAssociationMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretsSyncAssociationMetadataListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

