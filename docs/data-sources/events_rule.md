---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_events_rule Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Events::Rule
---

# awscc_events_rule (Data Source)

Data Source schema for AWS::Events::Rule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The ARN of the rule, such as arn:aws:events:us-east-2:123456789012:rule/example.
- `description` (String) The description of the rule.
- `event_bus_name` (String) The name or ARN of the event bus associated with the rule. If you omit this, the default event bus is used.
- `event_pattern` (String) The event pattern of the rule. For more information, see Events and Event Patterns in the Amazon EventBridge User Guide.
- `name` (String) The name of the rule.
- `role_arn` (String) The Amazon Resource Name (ARN) of the role that is used for target invocation.
- `schedule_expression` (String) The scheduling expression. For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see Creating an Amazon EventBridge rule that runs on a schedule.
- `state` (String) The state of the rule.
- `targets` (Attributes Set) Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.
Targets are the resources that are invoked when a rule is triggered. (see [below for nested schema](#nestedatt--targets))

<a id="nestedatt--targets"></a>
### Nested Schema for `targets`

Read-Only:

- `arn` (String)
- `batch_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--batch_parameters))
- `dead_letter_config` (Attributes) (see [below for nested schema](#nestedatt--targets--dead_letter_config))
- `ecs_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--ecs_parameters))
- `http_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--http_parameters))
- `id` (String)
- `input` (String)
- `input_path` (String)
- `input_transformer` (Attributes) (see [below for nested schema](#nestedatt--targets--input_transformer))
- `kinesis_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--kinesis_parameters))
- `redshift_data_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--redshift_data_parameters))
- `retry_policy` (Attributes) (see [below for nested schema](#nestedatt--targets--retry_policy))
- `role_arn` (String)
- `run_command_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--run_command_parameters))
- `sage_maker_pipeline_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--sage_maker_pipeline_parameters))
- `sqs_parameters` (Attributes) (see [below for nested schema](#nestedatt--targets--sqs_parameters))

<a id="nestedatt--targets--batch_parameters"></a>
### Nested Schema for `targets.batch_parameters`

Read-Only:

- `array_properties` (Attributes) (see [below for nested schema](#nestedatt--targets--batch_parameters--array_properties))
- `job_definition` (String)
- `job_name` (String)
- `retry_strategy` (Attributes) (see [below for nested schema](#nestedatt--targets--batch_parameters--retry_strategy))

<a id="nestedatt--targets--batch_parameters--array_properties"></a>
### Nested Schema for `targets.batch_parameters.array_properties`

Read-Only:

- `size` (Number)


<a id="nestedatt--targets--batch_parameters--retry_strategy"></a>
### Nested Schema for `targets.batch_parameters.retry_strategy`

Read-Only:

- `attempts` (Number)



<a id="nestedatt--targets--dead_letter_config"></a>
### Nested Schema for `targets.dead_letter_config`

Read-Only:

- `arn` (String)


<a id="nestedatt--targets--ecs_parameters"></a>
### Nested Schema for `targets.ecs_parameters`

Read-Only:

- `capacity_provider_strategy` (Attributes List) (see [below for nested schema](#nestedatt--targets--ecs_parameters--capacity_provider_strategy))
- `enable_ecs_managed_tags` (Boolean)
- `enable_execute_command` (Boolean)
- `group` (String)
- `launch_type` (String)
- `network_configuration` (Attributes) (see [below for nested schema](#nestedatt--targets--ecs_parameters--network_configuration))
- `placement_constraints` (Attributes List) (see [below for nested schema](#nestedatt--targets--ecs_parameters--placement_constraints))
- `placement_strategies` (Attributes List) (see [below for nested schema](#nestedatt--targets--ecs_parameters--placement_strategies))
- `platform_version` (String)
- `propagate_tags` (String)
- `reference_id` (String)
- `tag_list` (Attributes List) (see [below for nested schema](#nestedatt--targets--ecs_parameters--tag_list))
- `task_count` (Number)
- `task_definition_arn` (String)

<a id="nestedatt--targets--ecs_parameters--capacity_provider_strategy"></a>
### Nested Schema for `targets.ecs_parameters.capacity_provider_strategy`

Read-Only:

- `base` (Number)
- `capacity_provider` (String)
- `weight` (Number)


<a id="nestedatt--targets--ecs_parameters--network_configuration"></a>
### Nested Schema for `targets.ecs_parameters.network_configuration`

Read-Only:

- `aws_vpc_configuration` (Attributes) (see [below for nested schema](#nestedatt--targets--ecs_parameters--network_configuration--aws_vpc_configuration))

<a id="nestedatt--targets--ecs_parameters--network_configuration--aws_vpc_configuration"></a>
### Nested Schema for `targets.ecs_parameters.network_configuration.aws_vpc_configuration`

Read-Only:

- `assign_public_ip` (String)
- `security_groups` (List of String)
- `subnets` (List of String)



<a id="nestedatt--targets--ecs_parameters--placement_constraints"></a>
### Nested Schema for `targets.ecs_parameters.placement_constraints`

Read-Only:

- `expression` (String)
- `type` (String)


<a id="nestedatt--targets--ecs_parameters--placement_strategies"></a>
### Nested Schema for `targets.ecs_parameters.placement_strategies`

Read-Only:

- `field` (String)
- `type` (String)


<a id="nestedatt--targets--ecs_parameters--tag_list"></a>
### Nested Schema for `targets.ecs_parameters.tag_list`

Read-Only:

- `key` (String)
- `value` (String)



<a id="nestedatt--targets--http_parameters"></a>
### Nested Schema for `targets.http_parameters`

Read-Only:

- `header_parameters` (Map of String)
- `path_parameter_values` (List of String)
- `query_string_parameters` (Map of String)


<a id="nestedatt--targets--input_transformer"></a>
### Nested Schema for `targets.input_transformer`

Read-Only:

- `input_paths_map` (Map of String)
- `input_template` (String)


<a id="nestedatt--targets--kinesis_parameters"></a>
### Nested Schema for `targets.kinesis_parameters`

Read-Only:

- `partition_key_path` (String)


<a id="nestedatt--targets--redshift_data_parameters"></a>
### Nested Schema for `targets.redshift_data_parameters`

Read-Only:

- `database` (String)
- `db_user` (String)
- `secret_manager_arn` (String)
- `sql` (String)
- `sqls` (List of String)
- `statement_name` (String)
- `with_event` (Boolean)


<a id="nestedatt--targets--retry_policy"></a>
### Nested Schema for `targets.retry_policy`

Read-Only:

- `maximum_event_age_in_seconds` (Number)
- `maximum_retry_attempts` (Number)


<a id="nestedatt--targets--run_command_parameters"></a>
### Nested Schema for `targets.run_command_parameters`

Read-Only:

- `run_command_targets` (Attributes List) (see [below for nested schema](#nestedatt--targets--run_command_parameters--run_command_targets))

<a id="nestedatt--targets--run_command_parameters--run_command_targets"></a>
### Nested Schema for `targets.run_command_parameters.run_command_targets`

Read-Only:

- `key` (String)
- `values` (List of String)



<a id="nestedatt--targets--sage_maker_pipeline_parameters"></a>
### Nested Schema for `targets.sage_maker_pipeline_parameters`

Read-Only:

- `pipeline_parameter_list` (Attributes List) (see [below for nested schema](#nestedatt--targets--sage_maker_pipeline_parameters--pipeline_parameter_list))

<a id="nestedatt--targets--sage_maker_pipeline_parameters--pipeline_parameter_list"></a>
### Nested Schema for `targets.sage_maker_pipeline_parameters.pipeline_parameter_list`

Read-Only:

- `name` (String)
- `value` (String)



<a id="nestedatt--targets--sqs_parameters"></a>
### Nested Schema for `targets.sqs_parameters`

Read-Only:

- `message_group_id` (String)