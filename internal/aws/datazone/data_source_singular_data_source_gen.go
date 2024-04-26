// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datazone_data_source", dataSourceDataSource)
}

// dataSourceDataSource returns the Terraform awscc_datazone_data_source data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataZone::DataSource resource.
func dataSourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssetFormsInput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The metadata forms that are to be attached to the assets that this data source works with.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The details of a metadata form.",
		//	    "properties": {
		//	      "Content": {
		//	        "description": "The content of the metadata form.",
		//	        "maxLength": 75000,
		//	        "type": "string"
		//	      },
		//	      "FormName": {
		//	        "description": "The name of the metadata form.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "TypeIdentifier": {
		//	        "description": "The ID of the metadata form type.",
		//	        "maxLength": 385,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "TypeRevision": {
		//	        "description": "The revision of the metadata form type.",
		//	        "maxLength": 64,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FormName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"asset_forms_input": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Content
					"content": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The content of the metadata form.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: FormName
					"form_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the metadata form.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TypeIdentifier
					"type_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the metadata form type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TypeRevision
					"type_revision": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The revision of the metadata form type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The metadata forms that are to be attached to the assets that this data source works with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.",
		//	  "properties": {
		//	    "GlueRunConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AutoImportDataQualityResult": {
		//	          "description": "Specifies whether to automatically import data quality metrics as part of the data source run.",
		//	          "type": "boolean"
		//	        },
		//	        "DataAccessRole": {
		//	          "description": "The data access role included in the configuration details of the AWS Glue data source.",
		//	          "pattern": "^arn:aws[^:]*:iam::\\d{12}:(role|role/service-role)/[\\w+=,.@-]{1,128}$",
		//	          "type": "string"
		//	        },
		//	        "RelationalFilterConfigurations": {
		//	          "description": "The relational filter configurations included in the configuration details of the AWS Glue data source.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "The relational filter configuration for the data source.",
		//	            "properties": {
		//	              "DatabaseName": {
		//	                "description": "The database name specified in the relational filter configuration for the data source.",
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "FilterExpressions": {
		//	                "description": "The filter expressions specified in the relational filter configuration for the data source.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "additionalProperties": false,
		//	                  "description": "The search filter expression.",
		//	                  "properties": {
		//	                    "Expression": {
		//	                      "maxLength": 2048,
		//	                      "minLength": 1,
		//	                      "type": "string"
		//	                    },
		//	                    "Type": {
		//	                      "description": "The search filter expression type.",
		//	                      "enum": [
		//	                        "INCLUDE",
		//	                        "EXCLUDE"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "Expression",
		//	                    "Type"
		//	                  ],
		//	                  "type": "object"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SchemaName": {
		//	                "description": "The schema name specified in the relational filter configuration for the data source.",
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "DatabaseName"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "RelationalFilterConfigurations"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "RedshiftRunConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "The configuration details of the Amazon Redshift data source.",
		//	      "properties": {
		//	        "DataAccessRole": {
		//	          "description": "The data access role included in the configuration details of the Amazon Redshift data source.",
		//	          "pattern": "^arn:aws[^:]*:iam::\\d{12}:(role|role/service-role)/[\\w+=,.@-]{1,128}$",
		//	          "type": "string"
		//	        },
		//	        "RedshiftCredentialConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "The details of the credentials required to access an Amazon Redshift cluster.",
		//	          "properties": {
		//	            "SecretManagerArn": {
		//	              "description": "The ARN of a secret manager for an Amazon Redshift cluster.",
		//	              "maxLength": 256,
		//	              "pattern": "^arn:aws[^:]*:secretsmanager:[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]:\\d{12}:secret:.*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "SecretManagerArn"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "RedshiftStorage": {
		//	          "description": "The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.",
		//	          "properties": {
		//	            "RedshiftClusterSource": {
		//	              "additionalProperties": false,
		//	              "description": "The name of an Amazon Redshift cluster.",
		//	              "properties": {
		//	                "ClusterName": {
		//	                  "description": "The name of an Amazon Redshift cluster.",
		//	                  "maxLength": 63,
		//	                  "minLength": 1,
		//	                  "pattern": "^[0-9a-z].[a-z0-9\\-]*$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "ClusterName"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "RedshiftServerlessSource": {
		//	              "additionalProperties": false,
		//	              "description": "The details of the Amazon Redshift Serverless workgroup storage.",
		//	              "properties": {
		//	                "WorkgroupName": {
		//	                  "description": "The name of the Amazon Redshift Serverless workgroup.",
		//	                  "maxLength": 64,
		//	                  "minLength": 3,
		//	                  "pattern": "^[a-z0-9-]+$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "WorkgroupName"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "RelationalFilterConfigurations": {
		//	          "description": "The relational filter configurations included in the configuration details of the Amazon Redshift data source.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "The relational filter configuration for the data source.",
		//	            "properties": {
		//	              "DatabaseName": {
		//	                "description": "The database name specified in the relational filter configuration for the data source.",
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "FilterExpressions": {
		//	                "description": "The filter expressions specified in the relational filter configuration for the data source.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "additionalProperties": false,
		//	                  "description": "The search filter expression.",
		//	                  "properties": {
		//	                    "Expression": {
		//	                      "maxLength": 2048,
		//	                      "minLength": 1,
		//	                      "type": "string"
		//	                    },
		//	                    "Type": {
		//	                      "description": "The search filter expression type.",
		//	                      "enum": [
		//	                        "INCLUDE",
		//	                        "EXCLUDE"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "Expression",
		//	                    "Type"
		//	                  ],
		//	                  "type": "object"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SchemaName": {
		//	                "description": "The schema name specified in the relational filter configuration for the data source.",
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "DatabaseName"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "RedshiftCredentialConfiguration",
		//	        "RedshiftStorage",
		//	        "RelationalFilterConfigurations"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GlueRunConfiguration
				"glue_run_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AutoImportDataQualityResult
						"auto_import_data_quality_result": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies whether to automatically import data quality metrics as part of the data source run.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DataAccessRole
						"data_access_role": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The data access role included in the configuration details of the AWS Glue data source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RelationalFilterConfigurations
						"relational_filter_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DatabaseName
									"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The database name specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: FilterExpressions
									"filter_expressions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
										NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Expression
												"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: Type
												"type": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "The search filter expression type.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
										}, /*END NESTED OBJECT*/
										Description: "The filter expressions specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SchemaName
									"schema_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The schema name specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "The relational filter configurations included in the configuration details of the AWS Glue data source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RedshiftRunConfiguration
				"redshift_run_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DataAccessRole
						"data_access_role": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The data access role included in the configuration details of the Amazon Redshift data source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RedshiftCredentialConfiguration
						"redshift_credential_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: SecretManagerArn
								"secret_manager_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of a secret manager for an Amazon Redshift cluster.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The details of the credentials required to access an Amazon Redshift cluster.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RedshiftStorage
						"redshift_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: RedshiftClusterSource
								"redshift_cluster_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ClusterName
										"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name of an Amazon Redshift cluster.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The name of an Amazon Redshift cluster.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: RedshiftServerlessSource
								"redshift_serverless_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: WorkgroupName
										"workgroup_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name of the Amazon Redshift Serverless workgroup.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The details of the Amazon Redshift Serverless workgroup storage.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RelationalFilterConfigurations
						"relational_filter_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DatabaseName
									"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The database name specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: FilterExpressions
									"filter_expressions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
										NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Expression
												"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: Type
												"type": schema.StringAttribute{ /*START ATTRIBUTE*/
													Description: "The search filter expression type.",
													Computed:    true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
										}, /*END NESTED OBJECT*/
										Description: "The filter expressions specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SchemaName
									"schema_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The schema name specified in the relational filter configuration for the data source.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "The relational filter configurations included in the configuration details of the Amazon Redshift data source.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The configuration details of the Amazon Redshift data source.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the data source was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the data source was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the data source.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone domain where the data source is created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone domain where the data source is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone domain where the data source is created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone domain where the data source is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableSetting
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the data source is enabled.",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"enable_setting": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the data source is enabled.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
		//	  "type": "string"
		//	}
		"environment_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the data source.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"data_source_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastRunAssetCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of assets created by the data source during its last run.",
		//	  "type": "number"
		//	}
		"last_run_asset_count": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of assets created by the data source during its last run.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastRunAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp that specifies when the data source was last run.",
		//	  "type": "string"
		//	}
		"last_run_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp that specifies when the data source was last run.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastRunStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the last run of this data source.",
		//	  "type": "string"
		//	}
		"last_run_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the last run of this data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the data source.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone project to which the data source is added.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone project to which the data source is added.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon DataZone project in which you want to add the data source.",
		//	  "type": "string"
		//	}
		"project_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon DataZone project in which you want to add the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublishOnImport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.",
		//	  "type": "boolean"
		//	}
		"publish_on_import": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Recommendation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies whether the business name generation is to be enabled for this data source.",
		//	  "properties": {
		//	    "EnableBusinessNameGeneration": {
		//	      "description": "Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"recommendation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableBusinessNameGeneration
				"enable_business_name_generation": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies whether the business name generation is to be enabled for this data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The schedule of the data source runs.",
		//	  "properties": {
		//	    "Schedule": {
		//	      "description": "The schedule of the data source runs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "cron\\((\\b[0-5]?[0-9]\\b) (\\b2[0-3]\\b|\\b[0-1]?[0-9]\\b) (.*){1,5} (.*){1,5} (.*){1,5} (.*){1,5}\\)",
		//	      "type": "string"
		//	    },
		//	    "Timezone": {
		//	      "description": "The timezone of the data source run.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Schedule
				"schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The schedule of the data source runs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Timezone
				"timezone": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The timezone of the data source run.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The schedule of the data source runs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the data source.",
		//	  "enum": [
		//	    "CREATING",
		//	    "FAILED_CREATION",
		//	    "READY",
		//	    "UPDATING",
		//	    "FAILED_UPDATE",
		//	    "RUNNING",
		//	    "DELETING",
		//	    "FAILED_DELETION"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the data source.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the data source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when this data source was updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when this data source was updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataZone::DataSource",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::DataSource").WithTerraformTypeName("awscc_datazone_data_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"asset_forms_input":                 "AssetFormsInput",
		"auto_import_data_quality_result":   "AutoImportDataQualityResult",
		"cluster_name":                      "ClusterName",
		"configuration":                     "Configuration",
		"content":                           "Content",
		"created_at":                        "CreatedAt",
		"data_access_role":                  "DataAccessRole",
		"data_source_id":                    "Id",
		"database_name":                     "DatabaseName",
		"description":                       "Description",
		"domain_id":                         "DomainId",
		"domain_identifier":                 "DomainIdentifier",
		"enable_business_name_generation":   "EnableBusinessNameGeneration",
		"enable_setting":                    "EnableSetting",
		"environment_id":                    "EnvironmentId",
		"environment_identifier":            "EnvironmentIdentifier",
		"expression":                        "Expression",
		"filter_expressions":                "FilterExpressions",
		"form_name":                         "FormName",
		"glue_run_configuration":            "GlueRunConfiguration",
		"last_run_asset_count":              "LastRunAssetCount",
		"last_run_at":                       "LastRunAt",
		"last_run_status":                   "LastRunStatus",
		"name":                              "Name",
		"project_id":                        "ProjectId",
		"project_identifier":                "ProjectIdentifier",
		"publish_on_import":                 "PublishOnImport",
		"recommendation":                    "Recommendation",
		"redshift_cluster_source":           "RedshiftClusterSource",
		"redshift_credential_configuration": "RedshiftCredentialConfiguration",
		"redshift_run_configuration":        "RedshiftRunConfiguration",
		"redshift_serverless_source":        "RedshiftServerlessSource",
		"redshift_storage":                  "RedshiftStorage",
		"relational_filter_configurations":  "RelationalFilterConfigurations",
		"schedule":                          "Schedule",
		"schema_name":                       "SchemaName",
		"secret_manager_arn":                "SecretManagerArn",
		"status":                            "Status",
		"timezone":                          "Timezone",
		"type":                              "Type",
		"type_identifier":                   "TypeIdentifier",
		"type_revision":                     "TypeRevision",
		"updated_at":                        "UpdatedAt",
		"workgroup_name":                    "WorkgroupName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
