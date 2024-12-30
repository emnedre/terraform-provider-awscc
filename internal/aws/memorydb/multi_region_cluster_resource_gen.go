// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package memorydb

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_memorydb_multi_region_cluster", multiRegionClusterResource)
}

// multiRegionClusterResource returns the Terraform awscc_memorydb_multi_region_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::MemoryDB::MultiRegionCluster resource.
func multiRegionClusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the multi region cluster.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the multi region cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the multi region cluster.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the multi region cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The engine type used by the multi region cluster.",
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The engine type used by the multi region cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Redis engine version used by the multi region cluster.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Redis engine version used by the multi region cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MultiRegionClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Global Datastore, it is generated by MemoryDB adding a prefix to MultiRegionClusterNameSuffix.",
		//	  "type": "string"
		//	}
		"multi_region_cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Global Datastore, it is generated by MemoryDB adding a prefix to MultiRegionClusterNameSuffix.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MultiRegionClusterNameSuffix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.",
		//	  "pattern": "[a-z][a-z0-9\\-]*",
		//	  "type": "string"
		//	}
		"multi_region_cluster_name_suffix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Multi Region cluster. This value must be unique as it also serves as the multi region cluster identifier.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[a-z][a-z0-9\\-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// MultiRegionClusterNameSuffix is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: MultiRegionParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the parameter group associated with the multi region cluster.",
		//	  "type": "string"
		//	}
		"multi_region_parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the parameter group associated with the multi region cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NodeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The compute and memory capacity of the nodes in the multi region cluster.",
		//	  "type": "string"
		//	}
		"node_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The compute and memory capacity of the nodes in the multi region cluster.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: NumShards
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of shards the multi region cluster will contain.",
		//	  "type": "integer"
		//	}
		"num_shards": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of shards the multi region cluster will contain.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the multi region cluster. For example, Available, Updating, Creating.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the multi region cluster. For example, Available, Updating, Creating.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TLSEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
		//	  "type": "boolean"
		//	}
		"tls_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag that enables in-transit encryption when set to true.\n\nYou cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this multi region cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for the tag. May not be null.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value. May be null.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for the tag. May not be null.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value. May be null.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this multi region cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdateStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.",
		//	  "enum": [
		//	    "COORDINATED",
		//	    "UNCOORDINATED"
		//	  ],
		//	  "type": "string"
		//	}
		"update_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An enum string value that determines the update strategy for scaling. Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"COORDINATED",
					"UNCOORDINATED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// UpdateStrategy is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::MemoryDB::Multi Region Cluster resource creates an Amazon MemoryDB Multi Region Cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::MultiRegionCluster").WithTerraformTypeName("awscc_memorydb_multi_region_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                               "ARN",
		"description":                       "Description",
		"engine":                            "Engine",
		"engine_version":                    "EngineVersion",
		"key":                               "Key",
		"multi_region_cluster_name":         "MultiRegionClusterName",
		"multi_region_cluster_name_suffix":  "MultiRegionClusterNameSuffix",
		"multi_region_parameter_group_name": "MultiRegionParameterGroupName",
		"node_type":                         "NodeType",
		"num_shards":                        "NumShards",
		"status":                            "Status",
		"tags":                              "Tags",
		"tls_enabled":                       "TLSEnabled",
		"update_strategy":                   "UpdateStrategy",
		"value":                             "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/MultiRegionClusterNameSuffix",
		"/properties/UpdateStrategy",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}