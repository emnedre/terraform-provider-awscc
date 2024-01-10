// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package emr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_emr_studio", studioDataSource)
}

// studioDataSource returns the Terraform awscc_emr_studio data source.
// This Terraform data source corresponds to the CloudFormation AWS::EMR::Studio resource.
func studioDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the EMR Studio.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the EMR Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AuthMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
		//	  "enum": [
		//	    "SSO",
		//	    "IAM"
		//	  ],
		//	  "type": "string"
		//	}
		"auth_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultS3Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
		//	  "maxLength": 10280,
		//	  "minLength": 6,
		//	  "pattern": "^s3://.*",
		//	  "type": "string"
		//	}
		"default_s3_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A detailed description of the Studio.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A detailed description of the Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS KMS key identifier (ARN) used to encrypt AWS EMR Studio workspace and notebook files when backed up to AWS S3.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"encryption_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS KMS key identifier (ARN) used to encrypt AWS EMR Studio workspace and notebook files when backed up to AWS S3.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineSecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
		//	  "maxLength": 256,
		//	  "minLength": 4,
		//	  "pattern": "^sg-[a-zA-Z0-9\\-._]+$",
		//	  "type": "string"
		//	}
		"engine_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdcInstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM Identity Center instance to create the Studio application.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"idc_instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM Identity Center instance to create the Studio application.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdcUserAssignment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL. If the value is set to REQUIRED, users must be explicitly assigned to the Studio application to access the Studio.",
		//	  "enum": [
		//	    "REQUIRED",
		//	    "OPTIONAL"
		//	  ],
		//	  "type": "string"
		//	}
		"idc_user_assignment": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL. If the value is set to REQUIRED, users must be explicitly assigned to the Studio application to access the Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdpAuthUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.",
		//	  "maxLength": 4096,
		//	  "pattern": "^https://[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])(:[0-9]*)*([?/#].*)?$",
		//	  "type": "string"
		//	}
		"idp_auth_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdpRelayStateParameterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of relay state parameter for external Identity Provider.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"idp_relay_state_parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of relay state parameter for external Identity Provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A descriptive name for the Amazon EMR Studio.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_-]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A descriptive name for the Amazon EMR Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StudioId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the EMR Studio.",
		//	  "maxLength": 256,
		//	  "minLength": 4,
		//	  "pattern": "^es-[0-9A-Z]+",
		//	  "type": "string"
		//	}
		"studio_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the EMR Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
		//	  "items": {
		//	    "description": "Identifier of a subnet",
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An arbitrary set of tags (key-value pairs) for this EMR Studio.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "[a-zA-Z+-=._:/]+$",
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
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TrustedIdentityPropagationEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean indicating whether to enable Trusted identity propagation for the Studio. The default value is false.",
		//	  "type": "boolean"
		//	}
		"trusted_identity_propagation_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean indicating whether to enable Trusted identity propagation for the Studio. The default value is false.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Url
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique Studio access URL.",
		//	  "maxLength": 4096,
		//	  "pattern": "^https://[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])(:[0-9]*)*([?/#].*)?$",
		//	  "type": "string"
		//	}
		"url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique Studio access URL.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies.",
		//	  "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
		//	  "type": "string"
		//	}
		"user_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
		//	  "pattern": "^(vpc-[0-9a-f]{8}|vpc-[0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkspaceSecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
		//	  "pattern": "^sg-[a-zA-Z0-9\\-._]+$",
		//	  "type": "string"
		//	}
		"workspace_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EMR::Studio",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::Studio").WithTerraformTypeName("awscc_emr_studio")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"auth_mode":                            "AuthMode",
		"default_s3_location":                  "DefaultS3Location",
		"description":                          "Description",
		"encryption_key_arn":                   "EncryptionKeyArn",
		"engine_security_group_id":             "EngineSecurityGroupId",
		"idc_instance_arn":                     "IdcInstanceArn",
		"idc_user_assignment":                  "IdcUserAssignment",
		"idp_auth_url":                         "IdpAuthUrl",
		"idp_relay_state_parameter_name":       "IdpRelayStateParameterName",
		"key":                                  "Key",
		"name":                                 "Name",
		"service_role":                         "ServiceRole",
		"studio_id":                            "StudioId",
		"subnet_ids":                           "SubnetIds",
		"tags":                                 "Tags",
		"trusted_identity_propagation_enabled": "TrustedIdentityPropagationEnabled",
		"url":                                  "Url",
		"user_role":                            "UserRole",
		"value":                                "Value",
		"vpc_id":                               "VpcId",
		"workspace_security_group_id":          "WorkspaceSecurityGroupId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
