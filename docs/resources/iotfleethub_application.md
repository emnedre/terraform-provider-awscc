---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotfleethub_application Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::IoTFleetHub::Application
---

# awscc_iotfleethub_application (Resource)

Resource schema for AWS::IoTFleetHub::Application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **application_name** (String) Application Name, should be between 1 and 256 characters.
- **role_arn** (String) The ARN of the role that the web application assumes when it interacts with AWS IoT Core. For more info on configuring this attribute, see https://docs.aws.amazon.com/iot/latest/apireference/API_iotfleethub_CreateApplication.html#API_iotfleethub_CreateApplication_RequestSyntax

### Optional

- **application_description** (String) Application Description, should be between 1 and 2048 characters.
- **tags** (Attributes Set) A list of key-value pairs that contain metadata for the application. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **application_arn** (String) The ARN of the application.
- **application_creation_date** (Number) When the Application was created
- **application_id** (String) The ID of the application.
- **application_last_update_date** (Number) When the Application was last updated
- **application_state** (String) The current state of the application.
- **application_url** (String) The URL of the application.
- **error_message** (String) A message indicating why Create or Delete Application failed.
- **id** (String) Uniquely identifies the resource.
- **sso_client_id** (String) The AWS SSO application generated client ID (used with AWS SSO APIs).

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iotfleethub_application.example <resource ID>
```