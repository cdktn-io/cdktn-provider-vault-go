// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpolicydocument


type DataVaultPolicyDocumentRuleDeniedParameter struct {
	// Name of denied key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/policy_document#key DataVaultPolicyDocument#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A list of values what are denied by policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/policy_document#value DataVaultPolicyDocument#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

