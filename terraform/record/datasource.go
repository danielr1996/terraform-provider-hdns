package record

import (
	"context"
	"github.com/danielr1996/hdns-go/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSource defines the terraform datasource for a record
func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRecordRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceRecordRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)

	id, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("Please provide a id argument")
	}
	record, err := c.Record.GetById(id.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(record.Id)
	if err := d.Set("type", record.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", record.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value", record.Value); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("zone_id", record.ZoneId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created", record.Created); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("modified", record.Modified); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
