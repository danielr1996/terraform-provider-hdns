package terraform

import (
	"context"
	"errors"
	"github.com/danielr1996/hdns-go/src/hdns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
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
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"hdns_zones": dataSourceZones(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if token == ""{
		return nil, diag.FromErr(errors.New("token must not be empty"))
	}

	client := hdns.NewClient().WithToken(token)
	return client, diags
}
