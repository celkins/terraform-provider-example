package example

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccExampleProject_basic(t *testing.T) {
	var proj Project

	testAccExampleProjectConfig := `
		resource "example_project" "test_proj" {
			name = "test-proj-for-project-test"
		}`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExampleProjectConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleProjectExists("example_project.test_proj", &proj),
				),
			},
		},
	})
}

func testAccCheckExampleProjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ExampleClient)
	rs, ok := s.RootModule().Resources["example_project.test_proj"]
	if !ok {
		return fmt.Errorf("Not found %s", "example_project.test_proj")
	}

	response, _ := client.Get(fmt.Sprintf("projects/%s", rs.Primary.Attributes["name"]))

	if response.StatusCode != 404 {
		return fmt.Errorf("Project still exists")
	}

	return nil
}

func testAccCheckExampleProjectExists(n string, proj *Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No project ID is set")
		}
		return nil
	}
}
