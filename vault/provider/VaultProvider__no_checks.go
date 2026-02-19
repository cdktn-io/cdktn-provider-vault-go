// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VaultProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (v *jsiiProxy_VaultProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateVaultProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateVaultProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVaultProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateVaultProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginAwsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginAzureParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginCertParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginGcpParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginJwtParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginKerberosParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginOciParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginOidcParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginRadiusParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginTokenFileParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetAuthLoginUserpassParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetClientAuthParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetHeadersParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSetNamespaceFromTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipChildTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipGetVaultVersionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VaultProvider) validateSetSkipTlsVerifyParameters(val interface{}) error {
	return nil
}

func validateNewVaultProviderParameters(scope constructs.Construct, id *string, config *VaultProviderConfig) error {
	return nil
}

