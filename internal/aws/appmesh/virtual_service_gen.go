// Code generated by generators/resource/main.go; DO NOT EDIT.

package appmesh

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_appmesh_virtual_service", virtualService)
}

// virtualService returns the Terraform aws_appmesh_virtual_service resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppMesh::VirtualService resource type.
func virtualService(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"mesh_name": {
			// Property: MeshName
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// MeshName is a force-new attribute.
		},
		"mesh_owner": {
			// Property: MeshOwner
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// MeshOwner is a force-new attribute.
		},
		"resource_owner": {
			// Property: ResourceOwner
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"spec": {
			// Property: Spec
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "Provider": {
			         "additionalProperties": false,
			         "properties": {
			           "VirtualNode": {
			             "additionalProperties": false,
			             "properties": {
			               "VirtualNodeName": {
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/VirtualNodeServiceProvider",
			             "required": [
			               "VirtualNodeName"
			             ],
			             "type": "object"
			           },
			           "VirtualRouter": {
			             "additionalProperties": false,
			             "properties": {
			               "VirtualRouterName": {
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/VirtualRouterServiceProvider",
			             "required": [
			               "VirtualRouterName"
			             ],
			             "type": "object"
			           }
			         },
			         "$ref": "#/definitions/VirtualServiceProvider",
			         "type": "object"
			       }
			     },
			     "$ref": "#/definitions/VirtualServiceSpec",
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"provider": {
						// Property: Provider
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "VirtualNode": {
						         "additionalProperties": false,
						         "properties": {
						           "VirtualNodeName": {
						             "type": "string"
						           }
						         },
						         "$ref": "#/definitions/VirtualNodeServiceProvider",
						         "required": [
						           "VirtualNodeName"
						         ],
						         "type": "object"
						       },
						       "VirtualRouter": {
						         "additionalProperties": false,
						         "properties": {
						           "VirtualRouterName": {
						             "type": "string"
						           }
						         },
						         "$ref": "#/definitions/VirtualRouterServiceProvider",
						         "required": [
						           "VirtualRouterName"
						         ],
						         "type": "object"
						       }
						     },
						     "$ref": "#/definitions/VirtualServiceProvider",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"virtual_node": {
									// Property: VirtualNode
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "properties": {
									       "VirtualNodeName": {
									         "type": "string"
									       }
									     },
									     "$ref": "#/definitions/VirtualNodeServiceProvider",
									     "required": [
									       "VirtualNodeName"
									     ],
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"virtual_node_name": {
												// Property: VirtualNodeName
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Required: true,
											},
										},
									),
									Optional: true,
								},
								"virtual_router": {
									// Property: VirtualRouter
									// CloudFormation resource type schema:
									/*
									   {
									     "additionalProperties": false,
									     "properties": {
									       "VirtualRouterName": {
									         "type": "string"
									       }
									     },
									     "$ref": "#/definitions/VirtualRouterServiceProvider",
									     "required": [
									       "VirtualRouterName"
									     ],
									     "type": "object"
									   }
									*/
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"virtual_router_name": {
												// Property: VirtualRouterName
												// CloudFormation resource type schema:
												/*
												   {
												     "type": "string"
												   }
												*/
												Type:     types.StringType,
												Required: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"uid": {
			// Property: Uid
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"virtual_service_name": {
			// Property: VirtualServiceName
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// VirtualServiceName is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: `Resource Type definition for AWS::AppMesh::VirtualService`,
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppMesh::VirtualService").WithTerraformTypeName("aws_appmesh_virtual_service").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithUpdateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_appmesh_virtual_service", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}