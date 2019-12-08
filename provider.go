package main

import (
	"github.com/alessandromr/terraform-provider-aws-serverless/aws"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"serverless_aws_function_s3":   aws.ResourceFunctionS3(),
			"serverless_aws_function_http": aws.ResourceFunctionHTTP(),
		},
	}
}
