---
layout: "junos"
page_title: "Junos: junos_interface"
sidebar_current: "docs-junos-resource-interface"
description: |-
  Create/configure a interface or interface unit
---

# junos_interface

Provides a interface resource.

!> **NOTE:** Since v1.11.0, this resource is **deprecated**. For more consistency, functionalities of this resource have been splitted in two new resource `junos_interface_physical` and `junos_interface_logical`.
There is a guide for help migrating to the new resources.

## Example Usage

```hcl
# Configure interface of switch
resource junos_interface "interface_switch_demo" {
  name         = "ge-0/0/0"
  description  = "interfaceSwitchDemo"
  trunk        = true
  vlan_members = ["100"]
}
# Configure a L3 interface on Junos Router or firewall
resource junos_interface "interface_fw_demo" {
  name         = "ge-0/0/0"
  description  = "interfaceFwDemo"
  vlan_tagging = true
}
resource junos_interface "interface_fw_demo_100" {
  name        = "${junos_interface.interface_fw_demo.name}.100"
  description = "interfaceFwDemo100"
  inet_address {
    address = "192.0.2.1/25"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource)(`String`) Name of interface or unit interface (with dot).
* `description` - (Optional)(`String`) Description for interface.
* `complete_destroy` - (Optional)(`Bool`) When destroy this resource, delete all configurations => do not add `disable` + `descrition NC` or `apply-groups` with `group_interface_delete` provider argument on **physical** or **st0.x** interfaces.  
(Usually, `st0.x` interfaces are completely deleted with `bind_interface_auto` argument in `junos_security_ipsec_vpn` resource or by `junos_interface_st0_unit` resource because of the dependency, but only if st0.x interface is empty or disable.)
* `vlan_tagging` - (Optional)(`Bool`) Add 802.1q VLAN tagging support.
* `vlan_tagging_id` - (Optional,Computed)(`Int`) 802.1q VLAN ID for unit interface. If not set, computed with `name` of interface (ge-0/0/0.100 = 100)
* `inet` - (Optional,Computed)(`Bool`) Enable family inet.
* `inet6` - (Optional,Computed)(`Bool`) Enable family inet6.
* `inet_address` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each address to declare.
  * `address` - (Required)(`String`) Address IP/Mask v4.
  * `vrrp_group` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each vrrp group to declare. See the [`vrrp_group` arguments for inet_address](#vrrp_group-arguments-for-inet_address) block.

* `inet6_address` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each ipv6 address to declare.
  * `address` - (Required)(`String`) Address IP/Mask v6.
  * `vrrp_group` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each vrrp group to declare. See the [`vrrp_group` arguments for inet6_address](#vrrp_group-arguments-for-inet6_address) block.
* `inet_mtu` - (Optional)(`Int`) Protocol family inet maximum transmission unit.
* `inet6_mtu` - (Optional)(`Int`) Protocol family inet6 maximum transmission unit.
* `inet_filter_input` - (Optional)(`String`) Filter to be applied to received packets for family inet.
* `inet_filter_output` - (Optional)(`String`) Filter to be applied to transmitted packets for family inet.
* `inet6_filter_input` - (Optional)(`String`) Filter to be applied to received packets for family inet6.
* `inet6_filter_output` - (Optional)(`String`) Filter to be applied to transmitted packets for family inet6.
* `inet_rpf_check` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified only once for enable reverse-path-forwarding checks with family inet on this interface with optional arguments.
  * `fail_filter` - (Optional)(`String`) Name of filter applied to packets failing RPF check.
  * `mode_loose` - (Optional)(`Bool`) Use reverse-path-forwarding loose mode instead the strict mode.
* `inet6_rpf_check` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified only once for enable reverse-path-forwarding checks with family inet6 on this interface with optional arguments. Arguments is same as `inet_rpf_check`.
* `ether802_3ad` - (Optional)(`String`) Name of aggregated device for add this interface to link of 802.3ad interface.
* `trunk` - (Optional)(`Bool`) Interface mode is trunk.
* `vlan_members` - (Optional)(`ListOfString`) List of vlan for membership for this interface.
* `vlan_native` - (Optional)(`Int`) Vlan for untagged frames
* `ae_lacp` - (Optional)(`String`) Add lacp option in aggregated-ether-options. Need to be 'active' or 'passive' for initiate transmission or respond.
* `ae_link_speed` - (Optional)(`String`) Link speed of individual interface that joins the AE.
* `ae_minimum_links` - (Optional)(`Int`) Minimum number of aggregated links (1..8).
* `security_zone` - (Optional)(`String`) Add this interface in security_zone. Need to be created before.
* `routing_instance` - (Optional)(`String`) Add this interface in routing_instance. Need to be created before.

---

### vrrp_group arguments for inet_address

* `identifier` - (Required)(`Int`) ID for vrrp
* `virtual_address` - (Required)(`ListOfString`) List of address IP v4.
* `accept_data` - (Optional)(`Bool`) Accept packets destined for virtual IP address. Conflict with `no_accept_data` when apply.
* `advertise_interval` - (Optional)(`Int`) Advertisement interval (seconds)
* `advertisements_threshold` - (Optional)(`Int`)  Number of vrrp advertisements missed before declaring master down.
* `authentication_key` - (Optional)(`String`) Authentication key.  
**WARNING** Clear in tfstate.
* `authentication_type` - (Optional)(`String`) Authentication type. Need to be 'md5' or 'simple'.
* `no_accept_data` - (Optional)(`Bool`) Don't accept packets destined for virtual IP address. Conflict with `accept_data` when apply.
* `no_preempt` - (Optional)(`Bool`) Don't allow preemption. Conflict with `preempt` when apply.
* `preempt` - (Optional)(`Bool`) Allow preemption. Conflict with `no_preempt` when apply.
* `priority` - (Optional)(`Int`) Virtual router election priority.
* `track_interface` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each track_interface to declare.
  * `interface` - (Required)(`String`) Name of interface.
  * `priority_cost` - (Required)(`Int`) Value to subtract from priority when interface is down
* `track_route` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Can be specified multiple times for each track_route to declare.
  * `route` - (Required)(`String`) Route address.
  * `routing_instance` - (Required)(`String`) Routing instance to which route belongs, or 'default'.
  * `priority_cost` - (Required)(`Int`) Value to subtract from priority when route is down.

---

### vrrp_group arguments for inet6_address

Same as [`vrrp_group` arguments for inet_address](#vrrp_group-arguments-for-inet_address) block but without `authentication_key`, `authentication_type` and with

* `virtual_link_local_address` - (Required)(`String`) Address IPv6 for Virtual link-local addresses.

## Import

Junos interface can be imported using an id made up of `<name>`, e.g.

```shell
$ terraform import junos_interface.interface_switch_demo ge-0/0/0

$ terraform import junos_interface.interface_fw_demo_100 ge-0/0/0.0
```
