package gke
// 
// // ATTENTION!
// // This rule was autogenerated!
// // Before making changes, consider updating the generator.
// 
// import (
// 	"github.com/aquasecurity/defsec/provider"
// 	"github.com/aquasecurity/defsec/result"
// 	"github.com/aquasecurity/defsec/severity"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/cidr"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
// 	"github.com/aquasecurity/tfsec/pkg/rule"
// 	"github.com/zclconf/go-cty/cty"
// )
// 
// func init() {
// 	scanner.RegisterCheckRule(rule.Rule{
// 		Provider:  provider.GoogleProvider,
// 		Service:   "gke",
// 		ShortCode: "no-public-control-plane",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "GKE Control Plane should not be publicly accessible",
// 			Explanation: `The GKE control plane is exposed to the public internet by default. `,
// 			Impact:      "GKE control plane exposed to public internet",
// 			Resolution:  "Use private nodes and master authorised networks to prevent exposure",
// 			BadExample: []string{`
// resource "google_service_account" "default" {
//   account_id   = "service-account-id"
//   display_name = "Service Account"
// }
// 
// resource "google_container_cluster" "primary" {
//   name     = "my-gke-cluster"
//   location = "us-central1"
// 
//   # We can't create a cluster with no node pool defined, but we want to only use
//   # separately managed node pools. So we create the smallest possible default
//   # node pool and immediately delete it.
//   remove_default_node_pool = true
//   initial_node_count       = 1
//   master_authorized_networks_config = [{
//     cidr_blocks = [{
//       cidr_block = "0.0.0.0/0"
//       display_name = "external"
//     }]
//   }]
// }
// 
// resource "google_container_node_pool" "primary_preemptible_nodes" {
//   name       = "my-node-pool"
//   location   = "us-central1"
//   cluster    = google_container_cluster.primary.name
//   node_count = 1
// 
//   node_config {
//     preemptible  = true
//     machine_type = "e2-medium"
// 
//     # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
//     service_account = google_service_account.default.email
//     oauth_scopes    = [
//       "https://www.googleapis.com/auth/cloud-platform"
//     ]
//   }
// }
// `},
// 			GoodExample: []string{`
// resource "google_service_account" "default" {
//   account_id   = "service-account-id"
//   display_name = "Service Account"
// }
// 
// resource "google_container_cluster" "primary" {
//   name     = "my-gke-cluster"
//   location = "us-central1"
// 
//   # We can't create a cluster with no node pool defined, but we want to only use
//   # separately managed node pools. So we create the smallest possible default
//   # node pool and immediately delete it.
//   remove_default_node_pool = true
//   initial_node_count       = 1
//   master_authorized_networks_config = [{
//     cidr_blocks = [{
//       cidr_block = "10.10.128.0/24"
//       display_name = "internal"
//     }]
//   }]
// }
// 
// resource "google_container_node_pool" "primary_preemptible_nodes" {
//   name       = "my-node-pool"
//   location   = "us-central1"
//   cluster    = google_container_cluster.primary.name
//   node_count = 1
// 
//   node_config {
//     preemptible  = true
//     machine_type = "e2-medium"
// 
//     # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
//     service_account = google_service_account.default.email
//     oauth_scopes    = [
//       "https://www.googleapis.com/auth/cloud-platform"
//     ]
//   }
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"google_container_cluster",
// 		},
// 		DefaultSeverity: severity.High,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
// 			config := resourceBlock.GetAttribute("master_authorized_networks_config")
// 			if config.IsNil() {
// 				return
// 			}
// 			config.Each(func(key cty.Value, val cty.Value) {
// 				if !val.Type().IsObjectType() {
// 					return
// 				}
// 				m := val.AsValueMap()
// 				blocks, ok := m["cidr_blocks"]
// 				if !ok {
// 					return
// 				}
// 				for _, block := range blocks.AsValueSlice() {
// 					if !block.Type().IsObjectType() {
// 						continue
// 					}
// 					blockObj := block.AsValueMap()
// 					cidrBlock, ok := blockObj["cidr_block"]
// 					if !ok {
// 						continue
// 					}
// 					if cidrBlock.Type() != cty.String {
// 						continue
// 					}
// 					if cidr.IsOpen(cidrBlock.AsString()) {
// 						set.AddResult().
// 							WithDescription("Resource '%s' defines a cluster with an internet exposed control plane.", resourceBlock).
// 							WithAttribute("")
// 					}
// 				}
// 			})
// 		},
// 	})
// }
