---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "Machine Group"
layout: "zpa"
page_title: "ZPA: machine_group"
description: |-
  Gets a ZPA Machine Group details.
---

# zpa_machine_group

The **zpa_machine_group** data source provides details about a specific machine group created in the Zscaler Private Access cloud.
This data source is required when creating:

1. Access policy Rule
2. Access policy timeout rule
3. Access policy forwarding rule

## Example Usage

```hcl
# ZPA Machine Group Data Source
data "zpa_machine_group" "example" {
  name = "MGR01"
}
```

## Argument Reference

The following arguments are supported:

* ` name` - (Required) Name. The name of the machine group to be exported.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **list** (List of Object) (see [below for nested schema](#nestedatt--list))

<a id="nestedatt--list"></a>
### `list`

- **creation_time** (Number)
- **description** (String)
- **enabled** (Boolean)
- **id** (Number)
- **machines** (List of Object) (see [below for nested schema](#nestedobjatt--list--machines))
- **modified_time** (Number)
- **modifiedby** (Number)
- **name** (String)

<a id="nestedobjatt--list--machines"></a>
### `list.machines`

- **creation_time** (Number)
- **description** (String)
- **fingerprint** (String)
- **id** (String)
- **issued_cert_id** (Number)
- **machine_group_id** (String)
- **machine_group_name** (String)
- **machine_token_id** (Number)
- **modified_time** (Number)
- **modifiedby** (Number)
- **name** (String)
- **signing_cert** (Map of String)

