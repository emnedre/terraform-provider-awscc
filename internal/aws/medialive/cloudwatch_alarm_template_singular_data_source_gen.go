// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package medialive

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
	registry.AddDataSourceFactory("awscc_medialive_cloudwatch_alarm_template", cloudWatchAlarmTemplateDataSource)
}

// cloudWatchAlarmTemplateDataSource returns the Terraform awscc_medialive_cloudwatch_alarm_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaLive::CloudWatchAlarmTemplate resource.
func cloudWatchAlarmTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A cloudwatch alarm template's ARN (Amazon Resource Name)",
		//	  "pattern": "^arn:.+:medialive:.+:cloudwatch-alarm-template:.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A cloudwatch alarm template's ARN (Amazon Resource Name)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ComparisonOperator
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The comparison operator used to compare the specified statistic and the threshold.",
		//	  "enum": [
		//	    "GreaterThanOrEqualToThreshold",
		//	    "GreaterThanThreshold",
		//	    "LessThanThreshold",
		//	    "LessThanOrEqualToThreshold"
		//	  ],
		//	  "type": "string"
		//	}
		"comparison_operator": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The comparison operator used to compare the specified statistic and the threshold.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: DatapointsToAlarm
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "The number of datapoints within the evaluation period that must be breaching to trigger the alarm.",
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"datapoints_to_alarm": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of datapoints within the evaluation period that must be breaching to trigger the alarm.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource's optional description.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource's optional description.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EvaluationPeriods
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "The number of periods over which data is compared to the specified threshold.",
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"evaluation_periods": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of periods over which data is compared to the specified threshold.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`",
		//	  "maxLength": 11,
		//	  "minLength": 7,
		//	  "pattern": "^(aws-)?[0-9]{7}$",
		//	  "type": "string"
		//	}
		"group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A cloudwatch alarm template group's identifier. Can be either be its id or current name.",
		//	  "pattern": "^[^\\s]+$",
		//	  "type": "string"
		//	}
		"group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A cloudwatch alarm template group's identifier. Can be either be its id or current name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A cloudwatch alarm template's id. AWS provided templates have ids that start with `aws-`",
		//	  "maxLength": 11,
		//	  "minLength": 7,
		//	  "pattern": "^(aws-)?[0-9]{7}$",
		//	  "type": "string"
		//	}
		"cloudwatch_alarm_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A cloudwatch alarm template's id. AWS provided templates have ids that start with `aws-`",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Identifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MetricName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the metric associated with the alarm. Must be compatible with targetResourceType.",
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the metric associated with the alarm. Must be compatible with targetResourceType.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[^\\s]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Period
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "The period, in seconds, over which the specified statistic is applied.",
		//	  "maximum": 86400,
		//	  "minimum": 10,
		//	  "type": "number"
		//	}
		"period": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The period, in seconds, over which the specified statistic is applied.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Statistic
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The statistic to apply to the alarm's metric data.",
		//	  "enum": [
		//	    "SampleCount",
		//	    "Average",
		//	    "Sum",
		//	    "Minimum",
		//	    "Maximum"
		//	  ],
		//	  "type": "string"
		//	}
		"statistic": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The statistic to apply to the alarm's metric data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Represents the tags associated with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Represents the tags associated with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource type this template should dynamically generate cloudwatch metric alarms for.",
		//	  "enum": [
		//	    "CLOUDFRONT_DISTRIBUTION",
		//	    "MEDIALIVE_MULTIPLEX",
		//	    "MEDIALIVE_CHANNEL",
		//	    "MEDIALIVE_INPUT_DEVICE",
		//	    "MEDIAPACKAGE_CHANNEL",
		//	    "MEDIAPACKAGE_ORIGIN_ENDPOINT",
		//	    "MEDIACONNECT_FLOW",
		//	    "S3_BUCKET"
		//	  ],
		//	  "type": "string"
		//	}
		"target_resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource type this template should dynamically generate cloudwatch metric alarms for.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Threshold
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "The threshold value to compare with the specified statistic.",
		//	  "type": "number"
		//	}
		"threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The threshold value to compare with the specified statistic.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TreatMissingData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies how missing data points are treated when evaluating the alarm's condition.",
		//	  "enum": [
		//	    "notBreaching",
		//	    "breaching",
		//	    "ignore",
		//	    "missing"
		//	  ],
		//	  "type": "string"
		//	}
		"treat_missing_data": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies how missing data points are treated when evaluating the alarm's condition.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaLive::CloudWatchAlarmTemplate").WithTerraformTypeName("awscc_medialive_cloudwatch_alarm_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"cloudwatch_alarm_template_id": "Id",
		"comparison_operator":          "ComparisonOperator",
		"created_at":                   "CreatedAt",
		"datapoints_to_alarm":          "DatapointsToAlarm",
		"description":                  "Description",
		"evaluation_periods":           "EvaluationPeriods",
		"group_id":                     "GroupId",
		"group_identifier":             "GroupIdentifier",
		"identifier":                   "Identifier",
		"metric_name":                  "MetricName",
		"modified_at":                  "ModifiedAt",
		"name":                         "Name",
		"period":                       "Period",
		"statistic":                    "Statistic",
		"tags":                         "Tags",
		"target_resource_type":         "TargetResourceType",
		"threshold":                    "Threshold",
		"treat_missing_data":           "TreatMissingData",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}