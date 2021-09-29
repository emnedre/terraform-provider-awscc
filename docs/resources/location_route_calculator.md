---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_location_route_calculator Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Location::RouteCalculator Resource Type
---

# awscc_location_route_calculator (Resource)

Definition of AWS::Location::RouteCalculator Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **calculator_name** (String)
- **data_source** (String)
- **pricing_plan** (String)

### Optional

- **description** (String)

### Read-Only

- **arn** (String)
- **calculator_arn** (String)
- **create_time** (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)
- **id** (String) Uniquely identifies the resource.
- **update_time** (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_location_route_calculator.example <resource ID>
```