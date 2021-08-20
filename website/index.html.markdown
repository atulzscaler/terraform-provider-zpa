---
layout: 'zscaler'
page_title: 'Provider: ZPA (Zscaler Private Access)'
sidebar_current: 'docs-zpa-index'
description: |-
  The ZPA provider is used to interact with the resources supported by Zscaler. The provider needs to be configured with the proper credentials before it can be used.
---

# ZPA Provider

The ZPA provider is used to interact with the resources supported by Zscaler. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources and data sources.

## Example Usage

Terraform 0.13 and later:

```hcl
terraform {
  required_providers {
    zpa = {
      source = "zpa/zpa"
      version = "~> 1.0"
    }
  }
}

# Configure the ZPA Provider
provider "zpa" {
  client_id  = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  client_secret  = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  customerid= "XXXXXXXXXXXXXXXXX"
}
```

Terraform 0.12 and earlier:

```hcl

# Configure the Okta Provider
provider "zpa" {
  client_id  = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  client_secret  = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  customerid= "XXXXXXXXXXXXXXXXX"
}
```

## Authentication

The ZPA provider offers flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Environment variables
- Provider Config

### Environment variables

You can provide your credentials via the `ZPA_CLIENT_ID`, `ZPA_CLIENT_SECRET`, `ZPA_CUSTOMER_ID`,
environment variables, representing your Zscaler ZPA Client ID, ZPA Client Secret and ZPA Customer ID respectively.

```hcl
provider "zpa" {}
```

Usage:

```sh
$ export ZPA_CLIENT_ID="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
$ export ZPA_CLIENT_SECRET="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
$ export ZPA_CUSTOMER_ID="XXXXXXXXXXXXXXXXX"
$ terraform plan
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the Okta `provider` block:

- `org_name` - (Optional) This is the org name of your Okta account, for example `dev-123456.oktapreview.com` would have an org name of `dev-123456`. It must be provided, but it can also be sourced from the `OKTA_ORG_NAME` environment variable.

- `base_url` - (Optional) This is the domain of your Okta account, for example `dev-123456.oktapreview.com` would have a base url of `oktapreview.com`. It must be provided, but it can also be sourced from the `OKTA_BASE_URL` environment variable.

- `api_token` - (Optional) This is the API token to interact with your Okta org (either `"api_token"` or `"client_id"`, `"scopesv` and `"private_key"` must be provided). It can also be sourced from the `OKTA_API_TOKEN` environment variable.

- `client_id` - (Optional) This is the client ID for obtaining the API token. It can also be sourced from the `OKTA_API_CLIENT_ID` environment variable. 

- `scopes` - (Optional) These are scopes for obtaining the API token in form of a comma separated list. It can also be sourced from the `OKTA_API_SCOPES` environment variable.

- `private_key` - (Optional) This is the private key for obtaining the API token (can be represented by a filepath, or the key itself). It can also be sourced from the `OKTA_API_PRIVATE_KEY` environment variable.

- `backoff` - (Optional) Whether to use exponential back off strategy for rate limits, the default is `true`.

- `min_wait_seconds` - (Optional) Minimum seconds to wait when rate limit is hit, the default is `30`.

- `max_wait_seconds` - (Optional) Maximum seconds to wait when rate limit is hit, the default is `300`.

- `max_retries` - (Optional) Maximum number of retries to attempt before returning an error, the default is `5`.

- `request_timeout` - (Optional) Timeout for single request (in seconds) which is made to Okta, the default is `0` (means no limit is set). The maximum value can be `300`.

- `max_api_capacity` - (Optional, experimental) sets what percentage of capacity the provider can use of the total 
  rate limit capacity while making calls to the Okta management API endpoints. Okta API operates in one minute buckets. 
  See Okta Management API Rate Limits: https://developer.okta.com/docs/reference/rl-global-mgmt. Can be set to a value between 1 and 100.