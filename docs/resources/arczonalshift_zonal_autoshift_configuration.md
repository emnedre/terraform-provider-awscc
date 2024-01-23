---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_arczonalshift_zonal_autoshift_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type
---

# awscc_arczonalshift_zonal_autoshift_configuration (Resource)

Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `practice_run_configuration` (Attributes) (see [below for nested schema](#nestedatt--practice_run_configuration))
- `resource_identifier` (String)
- `zonal_autoshift_status` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--practice_run_configuration"></a>
### Nested Schema for `practice_run_configuration`

Required:

- `outcome_alarms` (Attributes List) (see [below for nested schema](#nestedatt--practice_run_configuration--outcome_alarms))

Optional:

- `blocked_dates` (List of String)
- `blocked_windows` (List of String)
- `blocking_alarms` (Attributes List) (see [below for nested schema](#nestedatt--practice_run_configuration--blocking_alarms))

<a id="nestedatt--practice_run_configuration--outcome_alarms"></a>
### Nested Schema for `practice_run_configuration.outcome_alarms`

Required:

- `alarm_identifier` (String)
- `type` (String)


<a id="nestedatt--practice_run_configuration--blocking_alarms"></a>
### Nested Schema for `practice_run_configuration.blocking_alarms`

Required:

- `alarm_identifier` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_arczonalshift_zonal_autoshift_configuration.example <resource ID>
```