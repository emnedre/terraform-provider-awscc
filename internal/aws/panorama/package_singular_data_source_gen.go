// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package panorama

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_panorama_package", packageDataSource)
}

// packageDataSource returns the Terraform awscc_panorama_package data source.
// This Terraform data source corresponds to the CloudFormation AWS::Panorama::Package resource.
func packageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"created_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_\\/]+$",
		//	  "type": "string"
		//	}
		"package_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the package.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\-\\_]+$",
		//	  "type": "string"
		//	}
		"package_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the package.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A storage location.",
		//	  "properties": {
		//	    "BinaryPrefixLocation": {
		//	      "description": "The location's binary prefix.",
		//	      "type": "string"
		//	    },
		//	    "Bucket": {
		//	      "description": "The location's bucket.",
		//	      "type": "string"
		//	    },
		//	    "GeneratedPrefixLocation": {
		//	      "description": "The location's generated prefix.",
		//	      "type": "string"
		//	    },
		//	    "ManifestPrefixLocation": {
		//	      "description": "The location's manifest prefix.",
		//	      "type": "string"
		//	    },
		//	    "RepoPrefixLocation": {
		//	      "description": "The location's repo prefix.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"storage_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BinaryPrefixLocation
				"binary_prefix_location": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The location's binary prefix.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The location's bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GeneratedPrefixLocation
				"generated_prefix_location": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The location's generated prefix.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ManifestPrefixLocation
				"manifest_prefix_location": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The location's manifest prefix.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RepoPrefixLocation
				"repo_prefix_location": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The location's repo prefix.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A storage location.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for the package.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^.+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^.+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for the package.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Panorama::Package",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Panorama::Package").WithTerraformTypeName("awscc_panorama_package")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"binary_prefix_location":    "BinaryPrefixLocation",
		"bucket":                    "Bucket",
		"created_time":              "CreatedTime",
		"generated_prefix_location": "GeneratedPrefixLocation",
		"key":                       "Key",
		"manifest_prefix_location":  "ManifestPrefixLocation",
		"package_id":                "PackageId",
		"package_name":              "PackageName",
		"repo_prefix_location":      "RepoPrefixLocation",
		"storage_location":          "StorageLocation",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
