// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wisdom_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSWisdomAIAgentVersionDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Wisdom::AIAgentVersion", "awscc_wisdom_ai_agent_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSWisdomAIAgentVersionDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Wisdom::AIAgentVersion", "awscc_wisdom_ai_agent_version", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
