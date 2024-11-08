// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_rest_api", restApiDataSource)
}

// restApiDataSource returns the Terraform awscc_apigateway_rest_api data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::RestApi resource.
func restApiDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiKeySourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"api_key_source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BinaryMediaTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"binary_media_types": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Body
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.",
		//	  "type": "string"
		//	}
		"body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BodyS3Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.",
		//	  "properties": {
		//	    "Bucket": {
		//	      "description": "The name of the S3 bucket where the OpenAPI file is stored.",
		//	      "relationshipRef": {
		//	        "propertyPath": "/properties/BucketName",
		//	        "typeName": "AWS::S3::Bucket"
		//	      },
		//	      "type": "string"
		//	    },
		//	    "ETag": {
		//	      "description": "The Amazon S3 ETag (a file checksum) of the OpenAPI file. If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.",
		//	      "type": "string"
		//	    },
		//	    "Key": {
		//	      "description": "The file name of the OpenAPI file (Amazon S3 object name).",
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "For versioning-enabled buckets, a specific version of the OpenAPI file.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"body_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bucket
				"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the S3 bucket where the OpenAPI file is stored.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ETag
				"e_tag": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 ETag (a file checksum) of the OpenAPI file. If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Key
				"key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The file name of the OpenAPI file (Amazon S3 object name).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For versioning-enabled buckets, a specific version of the OpenAPI file.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CloneFrom
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"clone_from": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisableExecuteApiEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "boolean"
		//	}
		"disable_execute_api_endpoint": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A list of the endpoint types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.",
		//	  "properties": {
		//	    "Types": {
		//	      "description": "",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "VpcEndpointIds": {
		//	      "description": "",
		//	      "items": {
		//	        "relationshipRef": {
		//	          "propertyPath": "/properties/Id",
		//	          "typeName": "AWS::EC2::VPCEndpoint"
		//	        },
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"endpoint_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Types
				"types": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcEndpointIds
				"vpc_endpoint_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A list of the endpoint types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ``Parameters`` property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailOnWarnings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "boolean"
		//	}
		"fail_on_warnings": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinimumCompressionSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "integer"
		//	}
		"minimum_compression_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Mode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This property applies only when you use OpenAPI to define your REST API. The ``Mode`` determines how API Gateway handles resource updates.\n Valid values are ``overwrite`` or ``merge``. \n For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.\n  For ``merge``, the new API definition is merged with the existing API.\n If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. \n Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.",
		//	  "type": "string"
		//	}
		"mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This property applies only when you use OpenAPI to define your REST API. The ``Mode`` determines how API Gateway handles resource updates.\n Valid values are ``overwrite`` or ``merge``. \n For ``overwrite``, the new API definition replaces the existing one. The existing API identifier remains unchanged.\n  For ``merge``, the new API definition is merged with the existing API.\n If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ``overwrite``. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. \n Use the default mode to define top-level ``RestApi`` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "string"
		//	}
		"parameters": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document that contains the permissions for the ``RestApi`` resource. To set the ARN for the policy, use the ``!Join`` intrinsic function with ``\"\"`` as delimiter and values of ``\"execute-api:/\"`` and ``\"*\"``.",
		//	  "type": "string"
		//	}
		"policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A policy document that contains the permissions for the ``RestApi`` resource. To set the ARN for the policy, use the ``!Join`` intrinsic function with ``\"\"`` as delimiter and values of ``\"execute-api:/\"`` and ``\"*\"``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RootResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"root_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "uniqueItems": false
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
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::RestApi",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::RestApi").WithTerraformTypeName("awscc_apigateway_rest_api")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_source_type":          "ApiKeySourceType",
		"binary_media_types":           "BinaryMediaTypes",
		"body":                         "Body",
		"body_s3_location":             "BodyS3Location",
		"bucket":                       "Bucket",
		"clone_from":                   "CloneFrom",
		"description":                  "Description",
		"disable_execute_api_endpoint": "DisableExecuteApiEndpoint",
		"e_tag":                        "ETag",
		"endpoint_configuration":       "EndpointConfiguration",
		"fail_on_warnings":             "FailOnWarnings",
		"key":                          "Key",
		"minimum_compression_size":     "MinimumCompressionSize",
		"mode":                         "Mode",
		"name":                         "Name",
		"parameters":                   "Parameters",
		"policy":                       "Policy",
		"rest_api_id":                  "RestApiId",
		"root_resource_id":             "RootResourceId",
		"tags":                         "Tags",
		"types":                        "Types",
		"value":                        "Value",
		"version":                      "Version",
		"vpc_endpoint_ids":             "VpcEndpointIds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
