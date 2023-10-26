---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_healthimaging_datastore Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::HealthImaging::Datastore Resource Type
---

# awscc_healthimaging_datastore (Resource)

Definition of AWS::HealthImaging::Datastore Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `datastore_name` (String) User friendly name for Datastore.
- `kms_key_arn` (String) ARN referencing a KMS key or KMS key alias.
- `tags` (Map of String) A Map of key value pairs for Tags.

### Read-Only

- `created_at` (String) The timestamp when the data store was created.
- `datastore_arn` (String) The Datastore's ARN.
- `datastore_id` (String)
- `datastore_status` (String) A string to denote the Datastore's state.
- `id` (String) Uniquely identifies the resource.
- `updated_at` (String) The timestamp when the data store was created.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_healthimaging_datastore.example <resource ID>
```