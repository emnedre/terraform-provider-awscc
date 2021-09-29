---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoveryreadiness_resource_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.
---

# awscc_route53recoveryreadiness_resource_set (Resource)

Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **resource_set_name** (String) The name of the resource set to create.
- **resource_set_type** (String) The resource type of the resources in the resource set. Enter one of the following values for resource type: 

AWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource
- **resources** (Attributes List) A list of resource objects in the resource set. (see [below for nested schema](#nestedatt--resources))

### Optional

- **tags** (Attributes List) A tag to associate with the parameters for a resource set. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **resource_set_arn** (String) The Amazon Resource Name (ARN) of the resource set.

<a id="nestedatt--resources"></a>
### Nested Schema for `resources`

Required:

- **component_id** (String) The component identifier of the resource, generated when DNS target resource is used.
- **dns_target_resource** (Attributes) A component for DNS/routing control readiness checks. (see [below for nested schema](#nestedatt--resources--dns_target_resource))
- **readiness_scopes** (List of String) A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.
- **resource_arn** (String) The Amazon Resource Name (ARN) of the AWS resource.

<a id="nestedatt--resources--dns_target_resource"></a>
### Nested Schema for `resources.dns_target_resource`

Required:

- **domain_name** (String) The domain name that acts as an ingress point to a portion of the customer application.
- **hosted_zone_arn** (String) The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.
- **record_set_id** (String) The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.
- **record_type** (String) The type of DNS record of the target resource.
- **target_resource** (Attributes) The target resource that the Route 53 record points to. (see [below for nested schema](#nestedatt--resources--dns_target_resource--target_resource))

<a id="nestedatt--resources--dns_target_resource--target_resource"></a>
### Nested Schema for `resources.dns_target_resource.target_resource`

Required:

- **nlb_resource** (Attributes) The Network Load Balancer resource that a DNS target resource points to. (see [below for nested schema](#nestedatt--resources--dns_target_resource--target_resource--nlb_resource))
- **r53_resource** (Attributes) The Route 53 resource that a DNS target resource record points to. (see [below for nested schema](#nestedatt--resources--dns_target_resource--target_resource--r53_resource))

<a id="nestedatt--resources--dns_target_resource--target_resource--nlb_resource"></a>
### Nested Schema for `resources.dns_target_resource.target_resource.r53_resource`

Required:

- **arn** (String) A Network Load Balancer resource Amazon Resource Name (ARN).


<a id="nestedatt--resources--dns_target_resource--target_resource--r53_resource"></a>
### Nested Schema for `resources.dns_target_resource.target_resource.r53_resource`

Required:

- **domain_name** (String) The DNS target domain name.
- **record_set_id** (String) The Resource Record set id.





<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53recoveryreadiness_resource_set.example <resource ID>
```