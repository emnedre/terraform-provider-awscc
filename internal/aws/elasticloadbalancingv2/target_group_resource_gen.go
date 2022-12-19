// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_elasticloadbalancingv2_target_group", targetGroupResource)
}

// targetGroupResource returns the Terraform awscc_elasticloadbalancingv2_target_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::TargetGroup resource.
func targetGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"health_check_enabled": {
			// Property: HealthCheckEnabled
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.",
			//	  "type": "boolean"
			//	}
			Description: "Indicates whether health checks are enabled. If the target type is lambda, health checks are disabled by default but can be enabled. If the target type is instance, ip, or alb, health checks are always enabled and cannot be disabled.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"health_check_interval_seconds": {
			// Property: HealthCheckIntervalSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The approximate amount of time, in seconds, between health checks of an individual target.",
			//	  "type": "integer"
			//	}
			Description: "The approximate amount of time, in seconds, between health checks of an individual target.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"health_check_path": {
			// Property: HealthCheckPath
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "[HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.",
			//	  "type": "string"
			//	}
			Description: "[HTTP/HTTPS health checks] The destination for health checks on the targets. [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC protocol version] The path of a custom health check method with the format /package.service/method. The default is /AWS.ALB/healthcheck.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"health_check_port": {
			// Property: HealthCheckPort
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The port the load balancer uses when performing health checks on targets. ",
			//	  "type": "string"
			//	}
			Description: "The port the load balancer uses when performing health checks on targets. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"health_check_protocol": {
			// Property: HealthCheckProtocol
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The protocol the load balancer uses when performing health checks on targets. ",
			//	  "type": "string"
			//	}
			Description: "The protocol the load balancer uses when performing health checks on targets. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"health_check_timeout_seconds": {
			// Property: HealthCheckTimeoutSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The amount of time, in seconds, during which no response from a target means a failed health check.",
			//	  "type": "integer"
			//	}
			Description: "The amount of time, in seconds, during which no response from a target means a failed health check.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"healthy_threshold_count": {
			// Property: HealthyThresholdCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of consecutive health checks successes required before considering an unhealthy target healthy. ",
			//	  "type": "integer"
			//	}
			Description: "The number of consecutive health checks successes required before considering an unhealthy target healthy. ",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"ip_address_type": {
			// Property: IpAddressType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The type of IP address used for this target group. The possible values are ipv4 and ipv6. ",
			//	  "type": "string"
			//	}
			Description: "The type of IP address used for this target group. The possible values are ipv4 and ipv6. ",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"load_balancer_arns": {
			// Property: LoadBalancerArns
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "The Amazon Resource Names (ARNs) of the load balancers that route traffic to this target group.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"matcher": {
			// Property: Matcher
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "[HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.",
			//	  "properties": {
			//	    "GrpcCode": {
			//	      "description": "You can specify values between 0 and 99. You can specify multiple values, or a range of values. The default value is 12.",
			//	      "type": "string"
			//	    },
			//	    "HttpCode": {
			//	      "description": "For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200. You can specify multiple values or a range of values. ",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "[HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a successful response from a target.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"grpc_code": {
						// Property: GrpcCode
						Description: "You can specify values between 0 and 99. You can specify multiple values, or a range of values. The default value is 12.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"http_code": {
						// Property: HttpCode
						Description: "For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200. You can specify multiple values or a range of values. ",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the target group.",
			//	  "type": "string"
			//	}
			Description: "The name of the target group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.",
			//	  "type": "integer"
			//	}
			Description: "The port on which the targets receive traffic. This port is used unless you specify a port override when registering the target. If the target is a Lambda function, this parameter does not apply. If the protocol is GENEVE, the supported port is 6081.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The protocol to use for routing traffic to the targets.",
			//	  "type": "string"
			//	}
			Description: "The protocol to use for routing traffic to the targets.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"protocol_version": {
			// Property: ProtocolVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "[HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2.",
			//	  "type": "string"
			//	}
			Description: "[HTTP/HTTPS protocol] The protocol version. The possible values are GRPC, HTTP1, and HTTP2.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The tags.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The value for the tag. ",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The key name of the tag. ",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "The tags.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The value for the tag. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The key name of the tag. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"target_group_arn": {
			// Property: TargetGroupArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of the Target Group",
			//	  "type": "string"
			//	}
			Description: "The ARN of the Target Group",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"target_group_attributes": {
			// Property: TargetGroupAttributes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "arrayType": "AttributeList",
			//	  "description": "The attributes.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The value of the attribute.",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The name of the attribute.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "The attributes.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The value of the attribute.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"value": {
						// Property: Value
						Description: "The name of the attribute.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"target_group_full_name": {
			// Property: TargetGroupFullName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The full name of the target group.",
			//	  "type": "string"
			//	}
			Description: "The full name of the target group.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"target_group_name": {
			// Property: TargetGroupName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the target group.",
			//	  "type": "string"
			//	}
			Description: "The name of the target group.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The type of target that you must specify when registering targets with this target group. You can't specify targets for a target group using more than one target type.",
			//	  "type": "string"
			//	}
			Description: "The type of target that you must specify when registering targets with this target group. You can't specify targets for a target group using more than one target type.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The targets.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "AvailabilityZone": {
			//	        "description": "An Availability Zone or all. This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.",
			//	        "type": "string"
			//	      },
			//	      "Id": {
			//	        "description": "The ID of the target. If the target type of the target group is instance, specify an instance ID. If the target type is ip, specify an IP address. If the target type is lambda, specify the ARN of the Lambda function. If the target type is alb, specify the ARN of the Application Load Balancer target. ",
			//	        "type": "string"
			//	      },
			//	      "Port": {
			//	        "description": "The port on which the target is listening. If the target group protocol is GENEVE, the supported port is 6081. If the target type is alb, the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.",
			//	        "type": "integer"
			//	      }
			//	    },
			//	    "required": [
			//	      "Id"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "The targets.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"availability_zone": {
						// Property: AvailabilityZone
						Description: "An Availability Zone or all. This determines whether the target receives traffic from the load balancer nodes in the specified Availability Zone or from all enabled Availability Zones for the load balancer.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"id": {
						// Property: Id
						Description: "The ID of the target. If the target type of the target group is instance, specify an instance ID. If the target type is ip, specify an IP address. If the target type is lambda, specify the ARN of the Lambda function. If the target type is alb, specify the ARN of the Application Load Balancer target. ",
						Type:        types.StringType,
						Required:    true,
					},
					"port": {
						// Property: Port
						Description: "The port on which the target is listening. If the target group protocol is GENEVE, the supported port is 6081. If the target type is alb, the targeted Application Load Balancer must have at least one listener whose port matches the target group port. Not used if the target is a Lambda function.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"unhealthy_threshold_count": {
			// Property: UnhealthyThresholdCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of consecutive health check failures required before considering a target unhealthy.",
			//	  "type": "integer"
			//	}
			Description: "The number of consecutive health check failures required before considering a target unhealthy.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply.",
			//	  "type": "string"
			//	}
			Description: "The identifier of the virtual private cloud (VPC). If the target is a Lambda function, this parameter does not apply.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::TargetGroup").WithTerraformTypeName("awscc_elasticloadbalancingv2_target_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"availability_zone":             "AvailabilityZone",
		"grpc_code":                     "GrpcCode",
		"health_check_enabled":          "HealthCheckEnabled",
		"health_check_interval_seconds": "HealthCheckIntervalSeconds",
		"health_check_path":             "HealthCheckPath",
		"health_check_port":             "HealthCheckPort",
		"health_check_protocol":         "HealthCheckProtocol",
		"health_check_timeout_seconds":  "HealthCheckTimeoutSeconds",
		"healthy_threshold_count":       "HealthyThresholdCount",
		"http_code":                     "HttpCode",
		"id":                            "Id",
		"ip_address_type":               "IpAddressType",
		"key":                           "Key",
		"load_balancer_arns":            "LoadBalancerArns",
		"matcher":                       "Matcher",
		"name":                          "Name",
		"port":                          "Port",
		"protocol":                      "Protocol",
		"protocol_version":              "ProtocolVersion",
		"tags":                          "Tags",
		"target_group_arn":              "TargetGroupArn",
		"target_group_attributes":       "TargetGroupAttributes",
		"target_group_full_name":        "TargetGroupFullName",
		"target_group_name":             "TargetGroupName",
		"target_type":                   "TargetType",
		"targets":                       "Targets",
		"unhealthy_threshold_count":     "UnhealthyThresholdCount",
		"value":                         "Value",
		"vpc_id":                        "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
