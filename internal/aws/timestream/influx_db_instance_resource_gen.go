// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package timestream

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_timestream_influx_db_instance", influxDBInstanceResource)
}

// influxDBInstanceResource returns the Terraform awscc_timestream_influx_db_instance resource.
// This Terraform resource corresponds to the CloudFormation AWS::Timestream::InfluxDBInstance resource.
func influxDBInstanceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocatedStorage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The allocated storage for the InfluxDB instance.",
		//	  "maximum": 16384,
		//	  "minimum": 20,
		//	  "type": "integer"
		//	}
		"allocated_storage": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The allocated storage for the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(20, 16384),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that is associated with the InfluxDB instance.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[a-z\\-]*:timestream\\-influxdb:[a-z0-9\\-]+:[0-9]{12}:(db\\-instance)/[a-zA-Z0-9]{3,64}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that is associated with the InfluxDB instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone (AZ) where the InfluxDB instance is created.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone (AZ) where the InfluxDB instance is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The bucket for the InfluxDB instance.",
		//	  "maxLength": 64,
		//	  "minLength": 2,
		//	  "pattern": "^[^_][^\"]*$",
		//	  "type": "string"
		//	}
		"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The bucket for the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(2, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[^_][^\"]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Bucket is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DbInstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The compute instance of the InfluxDB instance.",
		//	  "enum": [
		//	    "db.influx.medium",
		//	    "db.influx.large",
		//	    "db.influx.xlarge",
		//	    "db.influx.2xlarge",
		//	    "db.influx.4xlarge",
		//	    "db.influx.8xlarge",
		//	    "db.influx.12xlarge",
		//	    "db.influx.16xlarge"
		//	  ],
		//	  "type": "string"
		//	}
		"db_instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The compute instance of the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"db.influx.medium",
					"db.influx.large",
					"db.influx.xlarge",
					"db.influx.2xlarge",
					"db.influx.4xlarge",
					"db.influx.8xlarge",
					"db.influx.12xlarge",
					"db.influx.16xlarge",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DbParameterGroupIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an existing InfluxDB parameter group.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "^[a-zA-Z0-9]+$",
		//	  "type": "string"
		//	}
		"db_parameter_group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an existing InfluxDB parameter group.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DbStorageType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The storage type of the InfluxDB instance.",
		//	  "enum": [
		//	    "InfluxIOIncludedT1",
		//	    "InfluxIOIncludedT2",
		//	    "InfluxIOIncludedT3"
		//	  ],
		//	  "type": "string"
		//	}
		"db_storage_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The storage type of the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"InfluxIOIncludedT1",
					"InfluxIOIncludedT2",
					"InfluxIOIncludedT3",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeploymentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Deployment type of the InfluxDB Instance.",
		//	  "enum": [
		//	    "SINGLE_AZ",
		//	    "WITH_MULTIAZ_STANDBY"
		//	  ],
		//	  "type": "string"
		//	}
		"deployment_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Deployment type of the InfluxDB Instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SINGLE_AZ",
					"WITH_MULTIAZ_STANDBY",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The connection endpoint for the InfluxDB instance.",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The connection endpoint for the InfluxDB instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The service generated unique identifier for InfluxDB instance.",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "^[a-zA-Z0-9]+$",
		//	  "type": "string"
		//	}
		"influx_db_instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The service generated unique identifier for InfluxDB instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InfluxAuthParametersSecretArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB instance.",
		//	  "pattern": "^arn:[a-z]*:secretsmanager:[a-z\\-0-9]*:[0-9]*:secret:[a-zA-Z0-9\\-]*",
		//	  "type": "string"
		//	}
		"influx_auth_parameters_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogDeliveryConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for sending logs to customer account from the InfluxDB instance.",
		//	  "properties": {
		//	    "S3Configuration": {
		//	      "additionalProperties": false,
		//	      "description": "S3 configuration for sending logs to customer account from the InfluxDB instance.",
		//	      "properties": {
		//	        "BucketName": {
		//	          "description": "The bucket name for logs to be sent from the InfluxDB instance",
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "pattern": "^[0-9a-z]+[0-9a-z\\.\\-]*[0-9a-z]+$",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "Specifies whether logging to customer specified bucket is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enabled",
		//	        "BucketName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Configuration"
		//	  ],
		//	  "type": "object"
		//	}
		"log_delivery_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Configuration
				"s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The bucket name for logs to be sent from the InfluxDB instance",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(3, 63),
								stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-z]+[0-9a-z\\.\\-]*[0-9a-z]+$"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies whether logging to customer specified bucket is enabled.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "S3 configuration for sending logs to customer account from the InfluxDB instance.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for sending logs to customer account from the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name that is associated with the InfluxDB instance.",
		//	  "maxLength": 40,
		//	  "minLength": 3,
		//	  "pattern": "^[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name that is associated with the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 40),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Organization
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The organization for the InfluxDB instance.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"organization": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The organization for the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Organization is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Password
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password for the InfluxDB instance.",
		//	  "maxLength": 64,
		//	  "minLength": 8,
		//	  "pattern": "^[a-zA-Z0-9]+$",
		//	  "type": "string"
		//	}
		"password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password for the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(8, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Password is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Attach a public IP to the customer ENI.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Attach a public IP to the customer ENI.",
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecondaryAvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Secondary Availability Zone (AZ) where the InfluxDB instance is created, if DeploymentType is set as WITH_MULTIAZ_STANDBY.",
		//	  "type": "string"
		//	}
		"secondary_availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Secondary Availability Zone (AZ) where the InfluxDB instance is created, if DeploymentType is set as WITH_MULTIAZ_STANDBY.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Status of the InfluxDB Instance.",
		//	  "enum": [
		//	    "CREATING",
		//	    "AVAILABLE",
		//	    "DELETING",
		//	    "MODIFYING",
		//	    "UPDATING",
		//	    "DELETED",
		//	    "FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Status of the InfluxDB Instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this DB instance.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An arbitrary set of tags (key-value pairs) for this DB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(1, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Username
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The username for the InfluxDB instance.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"username": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The username for the InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Username is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 5),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of EC2 subnet IDs for this InfluxDB instance.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 3,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of EC2 subnet IDs for this InfluxDB instance.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 3),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "The AWS::Timestream::InfluxDBInstance resource creates an InfluxDB instance.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::InfluxDBInstance").WithTerraformTypeName("awscc_timestream_influx_db_instance")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocated_storage":                 "AllocatedStorage",
		"arn":                               "Arn",
		"availability_zone":                 "AvailabilityZone",
		"bucket":                            "Bucket",
		"bucket_name":                       "BucketName",
		"db_instance_type":                  "DbInstanceType",
		"db_parameter_group_identifier":     "DbParameterGroupIdentifier",
		"db_storage_type":                   "DbStorageType",
		"deployment_type":                   "DeploymentType",
		"enabled":                           "Enabled",
		"endpoint":                          "Endpoint",
		"influx_auth_parameters_secret_arn": "InfluxAuthParametersSecretArn",
		"influx_db_instance_id":             "Id",
		"key":                               "Key",
		"log_delivery_configuration":        "LogDeliveryConfiguration",
		"name":                              "Name",
		"organization":                      "Organization",
		"password":                          "Password",
		"publicly_accessible":               "PubliclyAccessible",
		"s3_configuration":                  "S3Configuration",
		"secondary_availability_zone":       "SecondaryAvailabilityZone",
		"status":                            "Status",
		"tags":                              "Tags",
		"username":                          "Username",
		"value":                             "Value",
		"vpc_security_group_ids":            "VpcSecurityGroupIds",
		"vpc_subnet_ids":                    "VpcSubnetIds",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Username",
		"/properties/Password",
		"/properties/Organization",
		"/properties/Bucket",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(2160)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
