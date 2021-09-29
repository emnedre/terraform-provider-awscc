---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_domain_name Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::ApiGateway::DomainName.
---

# awscc_apigateway_domain_name (Resource)

Resource Type definition for AWS::ApiGateway::DomainName.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **certificate_arn** (String)
- **domain_name** (String)
- **endpoint_configuration** (Attributes) (see [below for nested schema](#nestedatt--endpoint_configuration))
- **mutual_tls_authentication** (Attributes) (see [below for nested schema](#nestedatt--mutual_tls_authentication))
- **ownership_verification_certificate_arn** (String)
- **regional_certificate_arn** (String)
- **security_policy** (String)
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **distribution_domain_name** (String)
- **distribution_hosted_zone_id** (String)
- **id** (String) Uniquely identifies the resource.
- **regional_domain_name** (String)
- **regional_hosted_zone_id** (String)

<a id="nestedatt--endpoint_configuration"></a>
### Nested Schema for `endpoint_configuration`

Optional:

- **types** (List of String)


<a id="nestedatt--mutual_tls_authentication"></a>
### Nested Schema for `mutual_tls_authentication`

Optional:

- **truststore_uri** (String)
- **truststore_version** (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_apigateway_domain_name.example <resource ID>
```