package example

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"example": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("EXAMPLE_URL"); v == "" {
		t.Fatal("EXAMPLE_URL must be set for acceptance tests")
	}
	if v := os.Getenv("EXAMPLE_AUTH_TOKEN"); v == "" {
		t.Fatal("EXAMPLE_AUTH_TOKEN must be set for acceptance tests")
	}
}
