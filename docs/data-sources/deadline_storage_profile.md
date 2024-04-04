---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_deadline_storage_profile Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Deadline::StorageProfile
---

# awscc_deadline_storage_profile (Data Source)

Data Source schema for AWS::Deadline::StorageProfile



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `display_name` (String)
- `farm_id` (String)
- `file_system_locations` (Attributes List) (see [below for nested schema](#nestedatt--file_system_locations))
- `os_family` (String)
- `storage_profile_id` (String)

<a id="nestedatt--file_system_locations"></a>
### Nested Schema for `file_system_locations`

Read-Only:

- `name` (String)
- `path` (String)
- `type` (String)