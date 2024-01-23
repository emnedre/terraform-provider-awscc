---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_b2bi_capability Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::B2BI::Capability Resource Type
---

# awscc_b2bi_capability (Resource)

Definition of AWS::B2BI::Capability Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `type` (String)

### Optional

- `instructions_documents` (Attributes List) (see [below for nested schema](#nestedatt--instructions_documents))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `capability_arn` (String)
- `capability_id` (String)
- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `modified_at` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Optional:

- `edi` (Attributes) (see [below for nested schema](#nestedatt--configuration--edi))

<a id="nestedatt--configuration--edi"></a>
### Nested Schema for `configuration.edi`

Required:

- `input_location` (Attributes) (see [below for nested schema](#nestedatt--configuration--edi--input_location))
- `output_location` (Attributes) (see [below for nested schema](#nestedatt--configuration--edi--output_location))
- `transformer_id` (String)
- `type` (Attributes) (see [below for nested schema](#nestedatt--configuration--edi--type))

<a id="nestedatt--configuration--edi--input_location"></a>
### Nested Schema for `configuration.edi.input_location`

Optional:

- `bucket_name` (String)
- `key` (String)


<a id="nestedatt--configuration--edi--output_location"></a>
### Nested Schema for `configuration.edi.output_location`

Optional:

- `bucket_name` (String)
- `key` (String)


<a id="nestedatt--configuration--edi--type"></a>
### Nested Schema for `configuration.edi.type`

Optional:

- `x12_details` (Attributes) (see [below for nested schema](#nestedatt--configuration--edi--type--x12_details))

<a id="nestedatt--configuration--edi--type--x12_details"></a>
### Nested Schema for `configuration.edi.type.x12_details`

Optional:

- `transaction_set` (String)
- `version` (String)





<a id="nestedatt--instructions_documents"></a>
### Nested Schema for `instructions_documents`

Optional:

- `bucket_name` (String)
- `key` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_b2bi_capability.example <resource ID>
```