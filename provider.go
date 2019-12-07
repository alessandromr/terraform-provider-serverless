package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/alessandromr/terraform-provider-aws-serverless/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"serverless_aws_function": resources.ResourceFunction(),
		},
	}
}
