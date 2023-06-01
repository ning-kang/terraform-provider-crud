package crud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrderResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "crud_unicorn" "test" {
  name = "hello"
  age = 1
  colour = "red"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify first order item
					resource.TestCheckResourceAttr("crud_unicorn.test", "name", "hello"),
					resource.TestCheckResourceAttr("crud_unicorn.test", "age", "1"),
					resource.TestCheckResourceAttr("crud_unicorn.test", "colour", "red"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("crud_unicorn.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "crud_unicorn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "crud_unicorn" "test" {
  name = "world"
  age = 2
  colour = "green"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("crud_unicorn.test", "name", "world"),
					resource.TestCheckResourceAttr("crud_unicorn.test", "age", "2"),
					resource.TestCheckResourceAttr("crud_unicorn.test", "colour", "green"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
