---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "Global Policy Forwarding"
layout: "zpa"
page_title: "ZPA: policy_forwarding"
description: |-
  Gets a ZPA Global Policy Forwarding details.

---

# zpa_policy_forwarding

The **zpa_policy_forwarding** data source provides details about a the ZPA Global Policy Forwarding.
This data source must be used in the following circumstances:

1. Access policy forwarding (To retrieve the Global Policy Forwarding ```PolicySetID```) where forwarding rules will be created.

## Example Usage

```hcl
# ZPA Global Policy Forwarding Data Source
data "zpa_policy_forwarding" "example" {
}
```

## Argument Reference

This data source does not require an argument to be provided.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **creation_time** (String)
- **description** (String)
- **enabled** (Boolean)
- **modified_time** (String)
- **modifiedby** (String)
- **name** (String)
- **policy_type** (String)
- **rules** (List of Object) (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### `rules`

Read-Only:

- **action** (String)
- **action_id** (String)
- **bypass_default_rule** (Boolean)
- **conditions** (List of Object) (see [below for nested schema](#nestedobjatt--rules--conditions))
- **creation_time** (String)
- **custom_msg** (String)
- **description** (String)
- **id** (String)
- **isolation_default_rule** (Boolean)
- **modified_time** (String)
- **modifiedby** (String)
- **name** (String)
- **operator** (String)
- **policy_set_id** (String)
- **policy_type** (String)
- **priority** (String)
- **reauth_default_rule** (Boolean)
- **reauth_idle_timeout** (String)
- **reauth_timeout** (String)
- **rule_order** (String)
- **zpn_cbi_profile_id** (String)
- **zpn_inspection_profile_id** (String)
- **zpn_inspection_profile_name** (String)

<a id="nestedobjatt--rules--conditions"></a>
### `rules.conditions`

Read-Only:

- **creation_time** (String)
- **id** (String)
- **modified_time** (String)
- **modifiedby** (String)
- **negated** (Boolean)
- **operands** (List of Object) (see [below for nested schema](#nestedobjatt--rules--conditions--operands))

<a id="nestedobjatt--rules--conditions--operands"></a>
### `rules.conditions.operands`

Read-Only:

- **creation_time** (String)
- **id** (String)
- **idp_id** (String)
- **lhs** (String)
- **modified_time** (String)
- **modifiedby** (String)
- **name** (String)
- **object_type** (String)
- **operator** (String)
- **rhs** (String)


