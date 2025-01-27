---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sagemaker_endpoint Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SageMaker::Endpoint
---

# awscc_sagemaker_endpoint (Data Source)

Data Source schema for AWS::SageMaker::Endpoint



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `deployment_config` (Attributes) Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies. (see [below for nested schema](#nestedatt--deployment_config))
- `endpoint_arn` (String) The Amazon Resource Name (ARN) of the endpoint.
- `endpoint_config_name` (String) The name of the endpoint configuration for the SageMaker endpoint. This is a required property.
- `endpoint_name` (String) The name of the SageMaker endpoint. This name must be unique within an AWS Region.
- `exclude_retained_variant_properties` (Attributes List) Specifies a list of variant properties that you want to exclude when updating an endpoint. (see [below for nested schema](#nestedatt--exclude_retained_variant_properties))
- `retain_all_variant_properties` (Boolean) When set to true, retains all variant properties for an endpoint when it is updated.
- `retain_deployment_config` (Boolean) When set to true, retains the deployment configuration during endpoint updates.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--deployment_config"></a>
### Nested Schema for `deployment_config`

Read-Only:

- `auto_rollback_configuration` (Attributes) Configuration for automatic rollback if an error occurs during deployment. (see [below for nested schema](#nestedatt--deployment_config--auto_rollback_configuration))
- `blue_green_update_policy` (Attributes) Configuration for blue-green update deployment policies. (see [below for nested schema](#nestedatt--deployment_config--blue_green_update_policy))
- `rolling_update_policy` (Attributes) Configuration for rolling update deployment policies. (see [below for nested schema](#nestedatt--deployment_config--rolling_update_policy))

<a id="nestedatt--deployment_config--auto_rollback_configuration"></a>
### Nested Schema for `deployment_config.auto_rollback_configuration`

Read-Only:

- `alarms` (Attributes List) List of CloudWatch alarms to monitor during the deployment. If any alarm goes off, the deployment is rolled back. (see [below for nested schema](#nestedatt--deployment_config--auto_rollback_configuration--alarms))

<a id="nestedatt--deployment_config--auto_rollback_configuration--alarms"></a>
### Nested Schema for `deployment_config.auto_rollback_configuration.alarms`

Read-Only:

- `alarm_name` (String) The name of the CloudWatch alarm.



<a id="nestedatt--deployment_config--blue_green_update_policy"></a>
### Nested Schema for `deployment_config.blue_green_update_policy`

Read-Only:

- `maximum_execution_timeout_in_seconds` (Number) The maximum time allowed for the blue/green update, in seconds.
- `termination_wait_in_seconds` (Number) The wait time before terminating the old endpoint during a blue/green deployment.
- `traffic_routing_configuration` (Attributes) The traffic routing configuration for the blue/green deployment. (see [below for nested schema](#nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration))

<a id="nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration"></a>
### Nested Schema for `deployment_config.blue_green_update_policy.traffic_routing_configuration`

Read-Only:

- `canary_size` (Attributes) Specifies the size of the canary traffic in a canary deployment. (see [below for nested schema](#nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration--canary_size))
- `linear_step_size` (Attributes) Specifies the step size for linear traffic routing. (see [below for nested schema](#nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration--linear_step_size))
- `type` (String) Specifies the type of traffic routing (e.g., 'AllAtOnce', 'Canary', 'Linear').
- `wait_interval_in_seconds` (Number) Specifies the wait interval between traffic shifts, in seconds.

<a id="nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration--canary_size"></a>
### Nested Schema for `deployment_config.blue_green_update_policy.traffic_routing_configuration.canary_size`

Read-Only:

- `type` (String) Specifies whether the `Value` is an instance count or a capacity unit.
- `value` (Number) The value representing either the number of instances or the number of capacity units.


<a id="nestedatt--deployment_config--blue_green_update_policy--traffic_routing_configuration--linear_step_size"></a>
### Nested Schema for `deployment_config.blue_green_update_policy.traffic_routing_configuration.linear_step_size`

Read-Only:

- `type` (String) Specifies whether the `Value` is an instance count or a capacity unit.
- `value` (Number) The value representing either the number of instances or the number of capacity units.




<a id="nestedatt--deployment_config--rolling_update_policy"></a>
### Nested Schema for `deployment_config.rolling_update_policy`

Read-Only:

- `maximum_batch_size` (Attributes) Specifies the maximum batch size for each rolling update. (see [below for nested schema](#nestedatt--deployment_config--rolling_update_policy--maximum_batch_size))
- `maximum_execution_timeout_in_seconds` (Number) The maximum time allowed for the rolling update, in seconds.
- `rollback_maximum_batch_size` (Attributes) The maximum batch size for rollback during an update failure. (see [below for nested schema](#nestedatt--deployment_config--rolling_update_policy--rollback_maximum_batch_size))
- `wait_interval_in_seconds` (Number) The time to wait between steps during the rolling update, in seconds.

<a id="nestedatt--deployment_config--rolling_update_policy--maximum_batch_size"></a>
### Nested Schema for `deployment_config.rolling_update_policy.maximum_batch_size`

Read-Only:

- `type` (String) Specifies whether the `Value` is an instance count or a capacity unit.
- `value` (Number) The value representing either the number of instances or the number of capacity units.


<a id="nestedatt--deployment_config--rolling_update_policy--rollback_maximum_batch_size"></a>
### Nested Schema for `deployment_config.rolling_update_policy.rollback_maximum_batch_size`

Read-Only:

- `type` (String) Specifies whether the `Value` is an instance count or a capacity unit.
- `value` (Number) The value representing either the number of instances or the number of capacity units.




<a id="nestedatt--exclude_retained_variant_properties"></a>
### Nested Schema for `exclude_retained_variant_properties`

Read-Only:

- `variant_property_type` (String) The type of variant property (e.g., 'DesiredInstanceCount', 'DesiredWeight', 'DataCaptureConfig').


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key of the tag.
- `value` (String) The value of the tag.
