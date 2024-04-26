---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cognito_user_pool_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Cognito::UserPoolGroup
---

# awscc_cognito_user_pool_group (Resource)

Resource Type definition for AWS::Cognito::UserPoolGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `user_pool_id` (String)

### Optional

- `description` (String)
- `group_name` (String)
- `precedence` (Number)
- `role_arn` (String)

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_cognito_user_pool_group.example <resource ID>
```