package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
)

func dataEnvVariable() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAction,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the environment variable. ex. 'HOME' or 'API_KEY'",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value of the environment variable.",
			},
		},
	}
}

func resourceEnvVariable() *schema.Resource {
	resource := dataEnvVariable()

	resource.CreateContext = resourceAction
	resource.UpdateContext = resourceAction
	resource.DeleteContext = resourceAction

	return resource
}

func resourceAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	value := os.Getenv(name)

	d.SetId(name)
	_ = d.Set("value", value)

	return diag.Diagnostics{}
}
