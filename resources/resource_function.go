package resources

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceFunction() *schema.Resource {
        return &schema.Resource{
                Create: resourceFunctionCreate,
                Read:   resourceFunctionRead,
                Update: resourceFunctionUpdate,
                Delete: resourceFunctionDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}


func resourceFunctionCreate(d *schema.ResourceData, m interface{}) error {
	return resourceFunctionRead(d, m)
}

func resourceFunctionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFunctionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFunctionRead(d, m)
}

func resourceFunctionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}