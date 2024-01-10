// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3express

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3express_directory_bucket", directoryBucketDataSource)
}

// directoryBucketDataSource returns the Terraform awscc_s3express_directory_bucket data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3Express::DirectoryBucket resource.
func directoryBucketDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Returns the Amazon Resource Name (ARN) of the specified bucket.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Returns the Amazon Resource Name (ARN) of the specified bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-z0-9][a-z0-9//.//-]*[a-z0-9]$",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataRedundancy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the number of Availability Zone that's used for redundancy for the bucket.",
		//	  "enum": [
		//	    "SingleAvailabilityZone"
		//	  ],
		//	  "type": "string"
		//	}
		"data_redundancy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the number of Availability Zone that's used for redundancy for the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.",
		//	  "type": "string"
		//	}
		"location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3Express::DirectoryBucket",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Express::DirectoryBucket").WithTerraformTypeName("awscc_s3express_directory_bucket")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"bucket_name":     "BucketName",
		"data_redundancy": "DataRedundancy",
		"location_name":   "LocationName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
