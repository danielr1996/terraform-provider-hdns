package terraform

import (
	"context"
	"github.com/danielr1996/hdns-go/client"
	"github.com/danielr1996/terraform-provider-hdns/terraform/record"
	"github.com/danielr1996/terraform-provider-hdns/terraform/zone"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider defines the terraform provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HDNS_TOKEN", nil),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"hdns_record": record.Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hdns_zone":   zone.DataSource(),
			"hdns_record": record.DataSource(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var diags diag.Diagnostics
	if token == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HDNS client",
			Detail:   "token must not be null",
		})
		return nil, diags
	}
	return client.New().WithToken(token), diags
}
