package junos_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJunosSecurityUtmProfileWebFL_basic(t *testing.T) {
	if os.Getenv("TESTACC_SWITCH") == "" && os.Getenv("TESTACC_ROUTER") == "" {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: testAccJunosSecurityUtmProfileWebFLConfigCreate(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"custom_block_message", "Blocked by Juniper"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"default_action", "log-and-permit"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.#", "1"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.0.default", "log-and-permit"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.0.server_connectivity", "log-and-permit"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.0.timeout", "log-and-permit"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.0.timeout", "log-and-permit"),
					),
				},
				{
					Config: testAccJunosSecurityUtmProfileWebFLConfigUpdate(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"custom_block_message", "Blocked by Juniper"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"fallback_settings.#", "0"),
						resource.TestCheckResourceAttr("junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
							"timeout", "3"),
					),
				},
				{
					ResourceName:      "junos_security_utm_profile_web_filtering_juniper_local.testacc_ProfileWebFL",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	}
}

func testAccJunosSecurityUtmProfileWebFLConfigCreate() string {
	return `
resource junos_security_utm_profile_web_filtering_juniper_local "testacc_ProfileWebFL" {
  name                 = "testacc ProfileWebFL"
  custom_block_message = "Blocked by Juniper"
  default_action       = "log-and-permit"
  fallback_settings {
    default             = "log-and-permit"
    server_connectivity = "log-and-permit"
    timeout             = "log-and-permit"
  }
}
`
}

func testAccJunosSecurityUtmProfileWebFLConfigUpdate() string {
	return `
resource junos_security_utm_profile_web_filtering_juniper_local "testacc_ProfileWebFL" {
  name                 = "testacc ProfileWebFL"
  custom_block_message = "Blocked by Juniper"
  default_action       = "log-and-permit"
  timeout              = 3
}
`
}
