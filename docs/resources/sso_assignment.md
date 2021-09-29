---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sso_assignment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for SSO assignmet
---

# awscc_sso_assignment (Resource)

Resource Type definition for SSO assignmet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **instance_arn** (String) The sso instance that the permission set is owned.
- **permission_set_arn** (String) The permission set that the assignemt will be assigned
- **principal_id** (String) The assignee's identifier, user id/group id
- **principal_type** (String) The assignee's type, user/group
- **target_id** (String) The account id to be provisioned.
- **target_type** (String) The type of resource to be provsioned to, only aws account now

### Read-Only

- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_sso_assignment.example <resource ID>
```