---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_domain_name_v2 Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::ApiGateway::DomainNameV2.
---

# awscc_apigateway_domain_name_v2 (Resource)

Resource Type definition for AWS::ApiGateway::DomainNameV2.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `certificate_arn` (String)
- `domain_name` (String)
- `endpoint_configuration` (Attributes) (see [below for nested schema](#nestedatt--endpoint_configuration))
- `policy` (String)
- `security_policy` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `domain_name_arn` (String) The amazon resource name (ARN) of the domain name resource.
- `domain_name_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--endpoint_configuration"></a>
### Nested Schema for `endpoint_configuration`

Optional:

- `types` (List of String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_apigateway_domain_name_v2.example "domain_name_arn"
```