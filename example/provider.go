package example

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"net/http"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("EXAMPLE_URL", nil),
				Description: "URL of the root of the target example server; MUST include trailing slash.",
			},
			"auth_token": {
				Required:    true,
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("EXAMPLE_AUTH_TOKEN", nil),
				Description: "Auth token to use with the example API.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"example_project": resourceProject(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := &ExampleClient{
		URL:        d.Get("url").(string),
		AuthToken:  d.Get("auth_token").(string),
		HTTPClient: &http.Client{},
	}

	return client, nil
}
