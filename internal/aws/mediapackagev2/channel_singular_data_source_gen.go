// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediapackagev2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediapackagev2_channel", channelDataSource)
}

// channelDataSource returns the Terraform awscc_mediapackagev2_channel data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaPackageV2::Channel resource.
func channelDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Resource Name (ARN) associated with the resource.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Resource Name (ARN) associated with the resource.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ChannelGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"channel_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ChannelName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"channel_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe date and time the channel was created.\u003c/p\u003e",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "<p>The date and time the channel was created.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eEnter any descriptive text that helps you to identify the channel.\u003c/p\u003e",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>Enter any descriptive text that helps you to identify the channel.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IngestEndpointUrls
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"ingest_endpoint_urls": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IngestEndpoints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe list of ingest endpoints.\u003c/p\u003e",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe ingest domain URL where the source stream should be sent.\u003c/p\u003e",
		//	    "properties": {
		//	      "Id": {
		//	        "description": "\u003cp\u003eThe system-generated unique identifier for the IngestEndpoint.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "Url": {
		//	        "description": "\u003cp\u003eThe ingest domain URL where the source stream should be sent.\u003c/p\u003e",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"ingest_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The system-generated unique identifier for the IngestEndpoint.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Url
					"url": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The ingest domain URL where the source stream should be sent.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>The list of ingest endpoints.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InputType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "HLS",
		//	    "CMAF"
		//	  ],
		//	  "type": "string"
		//	}
		"input_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe date and time the channel was modified.\u003c/p\u003e",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "<p>The date and time the channel was modified.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaPackageV2::Channel",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackageV2::Channel").WithTerraformTypeName("awscc_mediapackagev2_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                  "Arn",
		"channel_group_name":   "ChannelGroupName",
		"channel_name":         "ChannelName",
		"created_at":           "CreatedAt",
		"description":          "Description",
		"id":                   "Id",
		"ingest_endpoint_urls": "IngestEndpointUrls",
		"ingest_endpoints":     "IngestEndpoints",
		"input_type":           "InputType",
		"key":                  "Key",
		"modified_at":          "ModifiedAt",
		"tags":                 "Tags",
		"url":                  "Url",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
