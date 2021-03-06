---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "Application Segment"
layout: "zpa"
page_title: "ZPA: application_server"
description: |-
  Gets a ZPA Application Server details.

---

# zpa_application_server (Data Source)

The **zpa_application_server** data source provides details about a specific application server created in the Zscaler Private Access cloud.
This data source must be used in the following circumstances:

1. Server Group (When Dynamic Discovery is set to false)

## Example Usage

```hcl
# ZPA Application Server Data Source
data "zpa_application_server" "example" {
 name = "server.example.com"
}
```

## Argument Reference

The following arguments are supported:

* ` name` - (Required) Name. The name of the application server to be exported.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

### Read-Only

- **address** (String)
- **app_server_group_ids** (Set of String)
- **config_space** (String)
- **creation_time** (String)
- **description** (String)
- **enabled** (Boolean)
- **modified_time** (String)
- **modifiedby** (String)
- **id** (String)


