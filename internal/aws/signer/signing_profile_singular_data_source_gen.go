// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package signer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_signer_signing_profile", signingProfileDataSource)
}

// signingProfileDataSource returns the Terraform awscc_signer_signing_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::Signer::SigningProfile resource.
func signingProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified signing profile.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified signing profile.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PlatformId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the target signing platform.",
		//	  "enum": [
		//	    "AWSLambda-SHA384-ECDSA",
		//	    "Notation-OCI-SHA384-ECDSA"
		//	  ],
		//	  "type": "string"
		//	}
		"platform_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the target signing platform.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the signing profile. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
		//	  "maxLength": 64,
		//	  "minLength": 2,
		//	  "pattern": "^[0-9a-zA-Z_]$",
		//	  "type": "string"
		//	}
		"profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the signing profile. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"profile_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified signing profile version.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"profile_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified signing profile version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SignatureValidityPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Signature validity period of the profile.",
		//	  "properties": {
		//	    "Type": {
		//	      "enum": [
		//	        "DAYS",
		//	        "MONTHS",
		//	        "YEARS"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Value": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"signature_validity_period": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Value
				"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Signature validity period of the profile.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags associated with the signing profile.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
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
			Description: "A list of tags associated with the signing profile.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Signer::SigningProfile",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Signer::SigningProfile").WithTerraformTypeName("awscc_signer_signing_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"key":                       "Key",
		"platform_id":               "PlatformId",
		"profile_name":              "ProfileName",
		"profile_version":           "ProfileVersion",
		"profile_version_arn":       "ProfileVersionArn",
		"signature_validity_period": "SignatureValidityPeriod",
		"tags":                      "Tags",
		"type":                      "Type",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
