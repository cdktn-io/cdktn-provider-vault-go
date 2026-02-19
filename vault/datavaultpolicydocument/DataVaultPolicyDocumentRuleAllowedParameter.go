// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpolicydocument


type DataVaultPolicyDocumentRuleAllowedParameter struct {
	// Name of permitted key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/data-sources/policy_document#key DataVaultPolicyDocument#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A list of values what are permitted by policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/data-sources/policy_document#value DataVaultPolicyDocument#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

