// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_codedeploy_deployment_config", deploymentConfigResource)
}

// deploymentConfigResource returns the Terraform awscc_codedeploy_deployment_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::CodeDeploy::DeploymentConfig resource.
func deploymentConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ComputePlatform
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The destination platform type for the deployment (Lambda, Server, or ECS).",
		//	  "type": "string"
		//	}
		"compute_platform": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The destination platform type for the deployment (Lambda, Server, or ECS).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentConfigName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
		//	  "type": "string"
		//	}
		"deployment_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MinimumHealthyHosts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
		//	  "properties": {
		//	    "Type": {
		//	      "type": "string"
		//	    },
		//	    "Value": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "Value"
		//	  ],
		//	  "type": "object"
		//	}
		"minimum_healthy_hosts": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Value
				"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrafficRoutingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration that specifies how the deployment traffic is routed.",
		//	  "properties": {
		//	    "TimeBasedCanary": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CanaryInterval": {
		//	          "type": "integer"
		//	        },
		//	        "CanaryPercentage": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "CanaryPercentage",
		//	        "CanaryInterval"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeBasedLinear": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "LinearInterval": {
		//	          "type": "integer"
		//	        },
		//	        "LinearPercentage": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "LinearInterval",
		//	        "LinearPercentage"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"traffic_routing_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: TimeBasedCanary
				"time_based_canary": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CanaryInterval
						"canary_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: CanaryPercentage
						"canary_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TimeBasedLinear
				"time_based_linear": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LinearInterval
						"linear_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: LinearPercentage
						"linear_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration that specifies how the deployment traffic is routed.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ZonalConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The zonal deployment config that specifies how the zonal deployment behaves",
		//	  "properties": {
		//	    "FirstZoneMonitorDurationInSeconds": {
		//	      "format": "int64",
		//	      "type": "integer"
		//	    },
		//	    "MinimumHealthyHostsPerZone": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Type": {
		//	          "type": "string"
		//	        },
		//	        "Value": {
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "Type",
		//	        "Value"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "MonitorDurationInSeconds": {
		//	      "format": "int64",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"zonal_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FirstZoneMonitorDurationInSeconds
				"first_zone_monitor_duration_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MinimumHealthyHostsPerZone
				"minimum_healthy_hosts_per_zone": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Type
						"type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: Value
						"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MonitorDurationInSeconds
				"monitor_duration_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The zonal deployment config that specifies how the zonal deployment behaves",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::CodeDeploy::DeploymentConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeDeploy::DeploymentConfig").WithTerraformTypeName("awscc_codedeploy_deployment_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"canary_interval":                        "CanaryInterval",
		"canary_percentage":                      "CanaryPercentage",
		"compute_platform":                       "ComputePlatform",
		"deployment_config_name":                 "DeploymentConfigName",
		"first_zone_monitor_duration_in_seconds": "FirstZoneMonitorDurationInSeconds",
		"linear_interval":                        "LinearInterval",
		"linear_percentage":                      "LinearPercentage",
		"minimum_healthy_hosts":                  "MinimumHealthyHosts",
		"minimum_healthy_hosts_per_zone":         "MinimumHealthyHostsPerZone",
		"monitor_duration_in_seconds":            "MonitorDurationInSeconds",
		"time_based_canary":                      "TimeBasedCanary",
		"time_based_linear":                      "TimeBasedLinear",
		"traffic_routing_config":                 "TrafficRoutingConfig",
		"type":                                   "Type",
		"value":                                  "Value",
		"zonal_config":                           "ZonalConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
