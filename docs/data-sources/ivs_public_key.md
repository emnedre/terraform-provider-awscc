---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ivs_public_key Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IVS::PublicKey
---

# awscc_ivs_public_key (Data Source)

Data Source schema for AWS::IVS::PublicKey



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) Key-pair identifier.
- `fingerprint` (String) Key-pair identifier.
- `name` (String) Name of the public key to be imported. The value does not need to be unique.
- `public_key_material` (String) The public portion of a customer-generated key pair.
- `tags` (Attributes Set) A list of key-value pairs that contain metadata for the asset model. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)