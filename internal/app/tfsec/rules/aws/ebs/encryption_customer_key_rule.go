package ebs
 
 // ATTENTION!
 // This rule was autogenerated!
 // Before making changes, consider updating the generator.
 
 // generator-locked
 
 import (
 	"github.com/aquasecurity/defsec/provider"
 	"github.com/aquasecurity/defsec/result"
 	"github.com/aquasecurity/defsec/severity"
 	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
 	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
 	"github.com/aquasecurity/tfsec/pkg/rule"
 )
 
 func init() {
 	scanner.RegisterCheckRule(rule.Rule{
 		Provider:  provider.AWSProvider,
 		Service:   "ebs",
 		ShortCode: "encryption-customer-key",
 		Documentation: rule.RuleDocumentation{
 			Summary:     "EBS volume encryption should use Customer Managed Keys",
 			Explanation: `Encryption using AWS keys provides protection for your EBS volume. To increase control of the encryption and manage factors like rotation use customer managed keys.`,
 			Impact:      "Using AWS managed keys does not allow for fine grained control",
 			Resolution:  "Enable encryption using customer managed keys",
 			BadExample: []string{`
 resource "aws_ebs_volume" "example" {
   availability_zone = "us-west-2a"
   size              = 40
 
   tags = {
     Name = "HelloWorld"
   }
 }
 `},
 			GoodExample: []string{`
 resource "aws_kms_key" "ebs_encryption" {
 	enable_key_rotation = true
 }
 
 resource "aws_ebs_volume" "example" {
   availability_zone = "us-west-2a"
   size              = 40
 
   kms_key_id = aws_kms_key.ebs_encryption.arn
 
   tags = {
     Name = "HelloWorld"
   }
 }
 `},
 			Links: []string{
 				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ebs_volume#kms_key_id",
 			},
 		},
 		RequiredTypes: []string{
 			"resource",
 		},
 		RequiredLabels: []string{
 			"aws_ebs_volume",
 		},
 		DefaultSeverity: severity.Low,
 		CheckTerraform: func(set result.Set, resourceBlock block.Block, module block.Module) {
 
 			if resourceBlock.MissingChild("kms_key_id") {
 				set.AddResult().
 					WithDescription("Resource '%s' does not use CMK", resourceBlock.FullName())
 				return
 			}
 
 			kmsKeyAttr := resourceBlock.GetAttribute("kms_key_id")
 			if kmsKeyAttr.IsDataBlockReference() {
 				kmsData, err := module.GetReferencedBlock(kmsKeyAttr, resourceBlock)
 				if err != nil {
 					return
 				}
 				keyIdAttr := kmsData.GetAttribute("key_id")
 				if keyIdAttr.IsNotNil() && keyIdAttr.StartsWith("alias/aws/") {
 					set.AddResult().
 						WithDescription("Resource '%s' explicitly uses the default CMK", resourceBlock.FullName()).
 						WithAttribute("")
 				}
 			}
 		},
 	})
 }
