---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_codeguruprofiler_profiling_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.
---

# awscc_codeguruprofiler_profiling_group (Resource)

This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **profiling_group_name** (String) The name of the profiling group.

### Optional

- **agent_permissions** (Attributes) The agent permissions attached to this profiling group. (see [below for nested schema](#nestedatt--agent_permissions))
- **anomaly_detection_notification_configuration** (Attributes List) Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency (see [below for nested schema](#nestedatt--anomaly_detection_notification_configuration))
- **compute_platform** (String) The compute platform of the profiling group.
- **tags** (Attributes List) The tags associated with a profiling group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--agent_permissions"></a>
### Nested Schema for `agent_permissions`

Optional:

- **principals** (List of String) The principals for the agent permissions.


<a id="nestedatt--anomaly_detection_notification_configuration"></a>
### Nested Schema for `anomaly_detection_notification_configuration`

Optional:

- **channel_id** (String) Unique identifier for each Channel in the notification configuration of a Profiling Group
- **channel_uri** (String) Unique arn of the resource to be used for notifications. We support a valid SNS topic arn as a channel uri.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
- **value** (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_codeguruprofiler_profiling_group.example <resource ID>
```