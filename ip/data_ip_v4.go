package ip

import (
	"context"

	"github.com/Miro-Ecosystem/go-miro/miro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	c := meta.(*miro.Client)
	var diags diag.Diagnostics

	board, err := c.Boards.Get(ctx, data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if board == nil {
		data.SetId("")
		return diags
	}

	if err := data.Set("boards", board); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)
	return diags
}
