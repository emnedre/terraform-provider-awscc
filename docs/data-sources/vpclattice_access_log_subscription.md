---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_access_log_subscription Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::AccessLogSubscription
---

# awscc_vpclattice_access_log_subscription (Data Source)

Data Source schema for AWS::VpcLattice::AccessLogSubscription



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `access_log_subscription_id` (String)
- `arn` (String)
- `destination_arn` (String)
- `resource_arn` (String)
- `resource_id` (String)
- `resource_identifier` (String)
- `service_network_log_type` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
