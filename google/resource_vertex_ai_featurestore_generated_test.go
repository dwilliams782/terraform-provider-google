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

func TestAccVertexAIFeaturestore_vertexAiFeaturestoreExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          GetTestOrgFromEnv(t),
		"billing_account": GetTestBillingAccountFromEnv(t),
		"kms_key_name":    BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":   RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeaturestoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestore_vertexAiFeaturestoreExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_featurestore.featurestore",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "etag", "region", "force_destroy"},
			},
		},
	})
}

func testAccVertexAIFeaturestore_vertexAiFeaturestoreExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    fixed_node_count = 2
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}
`, context)
}

func TestAccVertexAIFeaturestore_vertexAiFeaturestoreScalingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          GetTestOrgFromEnv(t),
		"billing_account": GetTestBillingAccountFromEnv(t),
		"kms_key_name":    BootstrapKMSKeyInLocation(t, "us-central1").CryptoKey.Name,
		"random_suffix":   RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeaturestoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeaturestore_vertexAiFeaturestoreScalingExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_featurestore.featurestore",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "etag", "region", "force_destroy"},
			},
		},
	})
}

func testAccVertexAIFeaturestore_vertexAiFeaturestoreScalingExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vertex_ai_featurestore" "featurestore" {
  name     = "terraform3%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region   = "us-central1"
  online_serving_config {
    scaling {
      min_node_count = 2
      max_node_count = 10
    }
  }
  encryption_spec {
    kms_key_name = "%{kms_key_name}"
  }
  force_destroy = true
}
`, context)
}

func testAccCheckVertexAIFeaturestoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_featurestore" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featurestores/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("VertexAIFeaturestore still exists at %s", url)
			}
		}

		return nil
	}
}
