---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iam_instance_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates a new instance profile. For information about instance profiles, see Using instance profiles https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html.
  For information about the number of instance profiles you can create, see object quotas https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html in the User Guide.
---

# awscc_iam_instance_profile (Resource)

Creates a new instance profile. For information about instance profiles, see [Using instance profiles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html).
  For information about the number of instance profiles you can create, see [object quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *User Guide*.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `roles` (Set of String) The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.

### Optional

- `instance_profile_name` (String) The name of the instance profile to create.
 This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
- `path` (String) The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.
 This parameter is optional. If it is not included, it defaults to a slash (/).
 This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\u0021``) through the DEL character (``\u007F``), including most punctuation characters, digits, and upper and lowercased letters.

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iam_instance_profile.example <resource ID>
```