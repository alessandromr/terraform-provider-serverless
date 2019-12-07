package main

import (
	"github.com/alessandromr/terraform-provider-aws-serverless/aws"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"serverless_aws_function": aws.ResourceFunction(),
		},
	}
}
