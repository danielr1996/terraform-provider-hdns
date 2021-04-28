package record

import (
	"context"
	"fmt"
	"github.com/danielr1996/hdns-go/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource defines the terraform resource for a record
func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	recordType := d.Get("type").(string)
	value := d.Get("value").(string)
	name := d.Get("name").(string)
	zoneId := d.Get("zone_id").(string)

	record, err := c.Record.Create(name, recordType, value, zoneId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error creating record %s IN %s %s", name, recordType, value),
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(record.Id)
	read(ctx, d, m)
	return diags

}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	record, err := c.Record.GetById(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading record %s", d.Id()),
			Detail:   err.Error(),
		})
		return diags
	}
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
	return diags
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)

	if d.HasChanges("type", "name", "value", "zone_id") {
		recordType := d.Get("type").(string)
		value := d.Get("value").(string)
		name := d.Get("name").(string)
		zoneId := d.Get("zone_id").(string)
		_, err := c.Record.Update(name, recordType, value, zoneId, d.Id())
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Error updating record %s IN %s %s with id %s", name, recordType, value, d.Id()),
				Detail:   err.Error(),
			})
			return diags
		}

	}
	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	err := c.Record.Delete(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error deleting record %s", d.Id()),
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId("")
	return diags
}
