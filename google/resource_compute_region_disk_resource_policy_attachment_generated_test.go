// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeRegionDiskResourcePolicyAttachment_regionDiskResourcePolicyAttachmentBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionDiskResourcePolicyAttachmentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDiskResourcePolicyAttachment_regionDiskResourcePolicyAttachmentBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_disk_resource_policy_attachment.attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"disk", "region"},
			},
		},
	})
}

func testAccComputeRegionDiskResourcePolicyAttachment_regionDiskResourcePolicyAttachmentBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk_resource_policy_attachment" "attachment" {
  name = google_compute_resource_policy.policy.name
  disk = google_compute_region_disk.ssd.name
  region = "us-central1"
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-base-disk%{random_suffix}"
  image = "debian-cloud/debian-11"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name  = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk" "ssd" {
  name  = "tf-test-my-disk%{random_suffix}"
  replica_zones = ["us-central1-a", "us-central1-f"]
  snapshot = google_compute_snapshot.snapdisk.id
  size  = 50
  type  = "pd-ssd"
  region  = "us-central1"
}

resource "google_compute_resource_policy" "policy" {
  name = "tf-test-my-resource-policy%{random_suffix}"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time = "04:00"
      }
    }
  }
}

data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}
`, context)
}

func testAccCheckComputeRegionDiskResourcePolicyAttachmentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_disk_resource_policy_attachment" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/disks/{{disk}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionDiskResourcePolicyAttachment still exists at %s", url)
			}
		}

		return nil
	}
}
