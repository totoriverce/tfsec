package compute
 
 // ATTENTION!
 // This rule was autogenerated!
 // Before making changes, consider updating the generator.
 
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
 		Provider:  provider.GoogleProvider,
 		Service:   "compute",
 		ShortCode: "no-plaintext-disk-keys",
 		Documentation: rule.RuleDocumentation{
 			Summary:     "Disk encryption keys should not be provided in plaintext",
 			Explanation: `Providing your encryption key in plaintext format means anyone with access to the source code also has access to the key.`,
 			Impact:      "Compromise of encryption keys",
 			Resolution:  "Use managed keys or provide the raw key via a secrets manager ",
 			BadExample: []string{`
 resource "google_compute_disk" "bad_example" {
   name  = "test-disk"
   type  = "pd-ssd"
   zone  = "us-central1-a"
   image = "debian-9-stretch-v20200805"
   labels = {
     environment = "dev"
   }
   physical_block_size_bytes = 4096
   disk_encryption_key {
     raw_key = "something"
   }
 }
 `},
 			GoodExample: []string{`
 resource "google_compute_disk" "good_example" {
   name  = "test-disk"
   type  = "pd-ssd"
   zone  = "us-central1-a"
   image = "debian-9-stretch-v20200805"
   labels = {
     environment = "dev"
   }
   physical_block_size_bytes = 4096
 }
 `},
 			Links: []string{
 				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_disk#raw_key",
 			},
 		},
 		RequiredTypes: []string{
 			"resource",
 		},
 		RequiredLabels: []string{
 			"google_compute_disk",
 		},
 		DefaultSeverity: severity.High,
 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
 			if rawKeyAttr := resourceBlock.GetBlock("disk_encryption_key").GetAttribute("raw_key"); rawKeyAttr.IsResolvable() {
 				set.AddResult().
 					WithDescription("Resource '%s' sets disk_encryption_key.raw_key", resourceBlock.FullName()).
 					WithAttribute("")
 			}
 		},
 	})
 }
