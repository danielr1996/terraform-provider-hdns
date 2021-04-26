package zone

import (
	"context"
	"github.com/danielr1996/hdns-go/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSource defines the terraform datasource for a zone
func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: read,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	name, ok := d.GetOk("name")
	if !ok {
		return diag.Errorf("Please provide a name argument")
	}
	zone, err := c.Zone.GetByName(name.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(zone.Id)
	if err := d.Set("name", zone.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ttl", zone.Ttl); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
