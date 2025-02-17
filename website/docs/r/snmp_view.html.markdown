---
layout: "junos"
page_title: "Junos: junos_snmp_view"
sidebar_current: "docs-junos-resource-snmp-view"
description: |-
  Create a snmp view
---

# junos_snmp_view

Provides a snmp view resource.

## Example Usage

```hcl
# Add a snmp view
resource junos_snmp_view "view1" {
  name        = "view1"
  oid_include = [".1"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource)(`String`) The name of snmp view.
* `oid_include` - (Optional)(`ListOfString`) OID include list.
* `oid_exclude` - (Optional)(`ListOfString`) OID exclude list.

## Import

Junos snmp view can be imported using an id made up of `<name>`, e.g.

```shell
$ terraform import junos_snmp_view.view1 view1
```
