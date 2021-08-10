package compute

// ATTENTION!
// This rule was autogenerated!
// Before making changes, consider updating the generator.

import (
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/provider"
	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/rule"
	"github.com/aquasecurity/tfsec/pkg/severity"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		Provider:       provider.GoogleProvider,
		Service:   "compute",
		ShortCode: "vm-disk-encryption-customer-key",
		Documentation: rule.RuleDocumentation{
			Summary:     "VM disks should be encrypted with Customer Supplied Encryption Keys",
			Explanation: `Using unmanaged keys makes rotation and general management difficult.`,
			Impact:      "Using unmanaged keys does not allow for proper management",
			Resolution:  "Use managed keys ",
			BadExample: []string{  `
resource "google_service_account" "default" {
  account_id   = "service_account_id"
  display_name = "Service Account"
}

resource "google_compute_instance" "bad_example" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }

  metadata = {
    foo = "bar"
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    email  = google_service_account.default.email
    scopes = ["cloud-platform"]
  }
}
`},
			GoodExample: []string{ `
resource "google_service_account" "default" {
  account_id   = "service_account_id"
  display_name = "Service Account"
}

resource "google_compute_instance" "good_example" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
    kms_key_self_link = "something"
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }

  metadata = {
    foo = "bar"
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    email  = google_service_account.default.email
    scopes = ["cloud-platform"]
  }
}
`},
			Links: []string{
				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#kms_key_self_link",
			},
		},
		RequiredTypes:  []string{ 
			"resource",
		},
		RequiredLabels: []string{ 
			"google_compute_instance",
		},
		DefaultSeverity: severity.Low, 
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context){
			if kmsKeySelfLinkAttr := resourceBlock.GetBlock("boot_disk").GetAttribute("kms_key_self_link"); kmsKeySelfLinkAttr.IsNil() { // alert on use of default value
				set.AddResult().
					WithDescription("Resource '%s' uses default value for boot_disk.kms_key_self_link", resourceBlock.FullName())
			} else if kmsKeySelfLinkAttr.IsNotResolvable() {
				set.AddResult().
					WithDescription("Resource '%s' does not set boot_disk.kms_key_self_link", resourceBlock.FullName()).
					WithAttribute(kmsKeySelfLinkAttr)
			}
		},
	})
}
