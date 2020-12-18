package ip

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rdegges/go-ipify"
)

func datasourceV4() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": {
				Description: "v4 IP address",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
		ReadContext: resourceBoardRead,
	}
}

func resourceBoardRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	ip, err := ipify.GetIp()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("ip", ip); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(ip)
	return diags
}
