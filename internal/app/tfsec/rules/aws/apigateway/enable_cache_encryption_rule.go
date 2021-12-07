package apigateway
// 
// // ATTENTION!
// // This rule was autogenerated!
// // Before making changes, consider updating the generator.
// 
// // generator-locked
// import (
// 	"github.com/aquasecurity/defsec/provider"
// 	"github.com/aquasecurity/defsec/result"
// 	"github.com/aquasecurity/defsec/severity"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
// 	"github.com/aquasecurity/tfsec/pkg/rule"
// )
// 
// func init() {
// 	scanner.RegisterCheckRule(rule.Rule{
// 		Provider:  provider.AWSProvider,
// 		Service:   "api-gateway",
// 		ShortCode: "enable-cache-encryption",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "API Gateway must have cache enabled",
// 			Explanation: `Method cache encryption ensures that any sensitive data in the cache is not vulnerable to compromise in the event of interception`,
// 			Impact:      "Data stored in the cache that is unencrypted may be vulnerable to compromise",
// 			Resolution:  "Enable cache encryption",
// 			BadExample: []string{`
// resource "aws_api_gateway_method_settings" "bad_example" {
//   rest_api_id = aws_api_gateway_rest_api.example.id
//   stage_name  = aws_api_gateway_stage.example.stage_name
//   method_path = "path1/GET"
// 
//   settings {
//     metrics_enabled = true
//     logging_level   = "INFO"
//     cache_data_encrypted = false
//   }
// }
// `},
// 			GoodExample: []string{`
// resource "aws_api_gateway_method_settings" "good_example" {
//   rest_api_id = aws_api_gateway_rest_api.example.id
//   stage_name  = aws_api_gateway_stage.example.stage_name
//   method_path = "path1/GET"
// 
//   settings {
//     metrics_enabled = true
//     logging_level   = "INFO"
//     cache_data_encrypted = true
//   }
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method_settings#cache_data_encrypted",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"aws_api_gateway_method_settings",
// 		},
// 		DefaultSeverity: severity.Medium,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
// 			if cacheDataEncryptedAttr := resourceBlock.GetBlock("settings").GetAttribute("cache_data_encrypted"); cacheDataEncryptedAttr.IsNil() { // alert on use of default value
// 				set.AddResult().
// 					WithDescription("Resource '%s' uses default value for settings.cache_data_encrypted", resourceBlock.FullName())
// 			} else if cacheDataEncryptedAttr.IsFalse() {
// 				set.AddResult().
// 					WithDescription("Resource '%s' does not have settings.cache_data_encrypted set to true", resourceBlock.FullName()).
// 					WithAttribute("")
// 			}
// 		},
// 	})
// }
