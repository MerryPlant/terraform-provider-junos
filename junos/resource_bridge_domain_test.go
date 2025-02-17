package junos_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJunosBridgeDomain_basic(t *testing.T) {
	if os.Getenv("TESTACC_ROUTER") != "" {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: testAccJunosBridgeDomainSwConfigCreate(),
				},
				{
					ResourceName:      "junos_bridge_domain.testacc_default",
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					ResourceName:      "junos_bridge_domain.testacc_bridge_ri",
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: testAccJunosBridgeDomainSwConfigUpdate(),
				},
			},
		})
	}
}

func testAccJunosBridgeDomainSwConfigCreate() string {
	return `
resource "junos_bridge_domain" "testacc_default" {
  name               = "testacc_bd_def"
  description        = "testacc bridge domain default"
  domain_type_bridge = true
  community_vlans    = [4]
  isolated_vlan      = 5
  routing_interface  = "irb.6"
  service_id         = 7
  vlan_id            = 8
}
resource "junos_bridge_domain" "testacc_default2" {
  name         = "testacc_bd_def2"
  vlan_id_list = [9]
}

resource "junos_interface_logical" "testacc_bridge_ri" {
  name = "lo0.1"
  family_inet {
    address {
      cidr_ip = "${junos_routing_options.testacc_bridge_ri.router_id}/32"
    }
  }
}
resource "junos_routing_options" "testacc_bridge_ri" {
  clean_on_destroy = true
  router_id        = "192.0.2.5"
}

resource "junos_routing_instance" "testacc_bridge_ri" {
  name                  = "testacc_bridge_ri"
  type                  = "virtual-switch"
  route_distinguisher   = "10:11"
  vrf_target            = "target:1:200"
  vtep_source_interface = junos_interface_logical.testacc_bridge_ri.name
}
resource "junos_evpn" "testacc_bridge_ri" {
  routing_instance = junos_routing_instance.testacc_bridge_ri.name
  encapsulation    = "vxlan"
  multicast_mode   = "ingress-replication"
}
resource "junos_bridge_domain" "testacc_bridge_ri" {
  depends_on = [
    junos_evpn.testacc_bridge_ri
  ]
  name               = "testacc_bd_ri"
  routing_instance   = junos_routing_instance.testacc_bridge_ri.name
  description        = "testacc bridge domain routing instance"
  domain_id          = 9
  domain_type_bridge = true
  routing_interface  = "irb.13"
  service_id         = 12
  vlan_id            = 13
  vxlan {
    vni                           = 14
    vni_extend_evpn               = true
    decapsulate_accept_inner_vlan = true
    encapsulate_inner_vlan        = true
    ingress_node_replication      = true
    unreachable_vtep_aging_timer  = 600
  }
}

`
}

func testAccJunosBridgeDomainSwConfigUpdate() string {
	return `
resource "junos_bridge_domain" "testacc_default" {
  name        = "testacc_bd_def"
  description = "testacc bridge domain default update"
  vlan_id     = 8
}
resource "junos_bridge_domain" "testacc_default2" {
  name         = "testacc_bd_def2"
  vlan_id_list = [9]
}

resource "junos_interface_logical" "testacc_bridge_ri" {
  name = "lo0.1"
  family_inet {
    address {
      cidr_ip = "${junos_routing_options.testacc_bridge_ri.router_id}/32"
    }
  }
}
resource "junos_routing_options" "testacc_bridge_ri" {
  clean_on_destroy = true
  router_id        = "192.0.2.5"
}

resource "junos_routing_instance" "testacc_bridge_ri" {
  name                  = "testacc_bridge_ri"
  type                  = "virtual-switch"
  route_distinguisher   = "10:11"
  vrf_target            = "target:1:200"
  vtep_source_interface = junos_interface_logical.testacc_bridge_ri.name
}
resource "junos_evpn" "testacc_bridge_ri" {
  routing_instance = junos_routing_instance.testacc_bridge_ri.name
  encapsulation    = "vxlan"
  multicast_mode   = "ingress-replication"
}
resource "junos_bridge_domain" "testacc_bridge_ri" {
  depends_on = [
    junos_evpn.testacc_bridge_ri
  ]
  name              = "testacc_bd_ri"
  routing_instance  = junos_routing_instance.testacc_bridge_ri.name
  description       = "testacc bridge domain routing instance"
  routing_interface = "irb.13"
  service_id        = 12
  vlan_id           = 13
  vxlan {
    vni = 15
  }
}
`
}
