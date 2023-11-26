package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/rickardl/terraform-provider-basexform/internal"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"basexform_encode": internal.ResourceEncode(),
					"basexform_decode": internal.ResourceDecode(),
				},
			}
		},
	})
}
