---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sns_subscription Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SNS::Subscription
---

# awscc_sns_subscription (Data Source)

Data Source schema for AWS::SNS::Subscription



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) Arn of the subscription
- `delivery_policy` (String) The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.
- `endpoint` (String) The subscription's endpoint. The endpoint value depends on the protocol that you specify.
- `filter_policy` (String) The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.
- `filter_policy_scope` (String) This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.
- `protocol` (String) The subscription's protocol.
- `raw_message_delivery` (Boolean) When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.
- `redrive_policy` (String) When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.
- `region` (String) For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.
- `replay_policy` (String) Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.
- `subscription_role_arn` (String) This property applies only to Amazon Data Firehose delivery stream subscriptions.
- `topic_arn` (String) The ARN of the topic to subscribe to.