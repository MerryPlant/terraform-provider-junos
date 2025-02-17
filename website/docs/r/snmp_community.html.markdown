---
layout: "junos"
page_title: "Junos: junos_snmp_community"
sidebar_current: "docs-junos-resource-snmp-community"
description: |-
  Create a snmp community
---

# junos_snmp_community

Provides a snmp community resource.

## Example Usage

```hcl
# Add a snmp community
resource junos_snmp_community "public" {
  name                    = "public"
  authorization_read_only = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource)(`String`) The name of snmp community.
* `authorization_read_only` - (Optional)(`Bool`) Allow read-only access. Conflict with `authorization_read_write`.
* `authorization_read_write` - (Optional)(`Bool`) Allow read and write access. Conflict with `authorization_read_only`.
* `client_list_name` - (Optional)(`String`) The name of client list or prefix list. Conflict with `clients`.
* `clients` - (Optional)(`ListOfString`) List of source address prefix ranges to accept. Conflict with `client_list_name`.
* `routing_instance` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) Use routing-instance name for v1/v2c clients. Can be specified multiple times for each routing-instance.
  * `name` - (Required)(`String`) Name of routing-instance.
  * `client_list_name` - (Optional)(`String`) The name of client list or prefix list.
  * `clients` - (Optional)(`ListOfString`) List of source address prefix ranges to accept.
* `view` - (Optional)(`String`) View name.

## Import

Junos snmp community can be imported using an id made up of `<name>`, e.g.

```shell
$ terraform import junos_snmp_community.public public
```
