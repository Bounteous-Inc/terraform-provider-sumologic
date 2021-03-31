// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSumologicLookupTable_basic(t *testing.T) {
	var lookupTable LookupTable
	testName := "SampleLookupTable"
	testFieldName := "FieldName1"
	testFieldType := "boolean"
	testTtl := 100
	testPrimaryKeys := "FieldName1"
	testSizeLimitAction := "StopIncomingMessages"
	testDescription := "This is a sample lookup table description."

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLookupTableDestroy(lookupTable),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicLookupTableConfigImported(testName, testFieldName, testFieldType, testTtl, testPrimaryKeys, testSizeLimitAction, testDescription),
			},
			{
				ResourceName:      "sumologic_lookup_table.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSumologicLookupTable_create(t *testing.T) {
	var lookupTable LookupTable
	testName := "SampleLookupTable"
	testFieldName := "FieldName1"
	testFieldType := "boolean"
	testTtl := 100
	testPrimaryKeys := "FieldName1"
	testSizeLimitAction := "StopIncomingMessages"
	testDescription := "This is a sample lookup table description."
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLookupTableDestroy(lookupTable),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicLookupTable(testName, testFieldName, testFieldType, testTtl, testPrimaryKeys, testSizeLimitAction, testDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLookupTableExists("sumologic_lookup_table.test", &lookupTable, t),
					testAccCheckLookupTableAttributes("sumologic_lookup_table.test"),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "name", testName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.#", "1"),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_name", testFieldName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_type", testFieldType),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "ttl", strconv.Itoa(testTtl)),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "primary_keys.0", testPrimaryKeys),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "size_limit_action", testSizeLimitAction),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "description", testDescription),
				),
			},
		},
	})
}

func TestAccSumologicLookupTable_update(t *testing.T) {
	var lookupTable LookupTable
	testName := "SampleLookupTable"
	testFieldName := "FieldName1"
	testFieldType := "boolean"
	testTtl := 100
	testPrimaryKeys := "FieldName1"
	testSizeLimitAction := "StopIncomingMessages"
	testDescription := "This is a sample lookup table description."

	testUpdatedTtl := 101
	testUpdatedSizeLimitAction := "DeleteOldData"
	testUpdatedDescription := "This is a sample lookup table description Updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLookupTableDestroy(lookupTable),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicLookupTable(testName, testFieldName, testFieldType, testTtl, testPrimaryKeys, testSizeLimitAction, testDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLookupTableExists("sumologic_lookup_table.test", &lookupTable, t),
					testAccCheckLookupTableAttributes("sumologic_lookup_table.test"),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "name", testName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.#", "1"),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_name", testFieldName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_type", testFieldType),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "ttl", strconv.Itoa(testTtl)),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "primary_keys.0", testPrimaryKeys),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "size_limit_action", testSizeLimitAction),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "description", testDescription),
				),
			},
			{
				Config: testAccSumologicLookupTableUpdate(testName, testFieldName, testFieldType, testUpdatedTtl, testPrimaryKeys, testUpdatedSizeLimitAction, testUpdatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "name", testName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.#", "1"),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_name", testFieldName),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "fields.0.field_type", testFieldType),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "ttl", strconv.Itoa(testUpdatedTtl)),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "primary_keys.0", testPrimaryKeys),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "size_limit_action", testUpdatedSizeLimitAction),
					resource.TestCheckResourceAttr("sumologic_lookup_table.test", "description", testUpdatedDescription),
				),
			},
		},
	})
}

func testAccCheckLookupTableDestroy(lookupTable LookupTable) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		_, err := client.GetLookupTable(lookupTable.ID)
		if err == nil {
			return fmt.Errorf("Lookup Table still exists")
		}
		return nil
	}
}

func testAccCheckLookupTableExists(name string, lookupTable *LookupTable, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			//need this so that we don't get an unused import error for strconv in some cases
			return fmt.Errorf("Error = %s. LookupTable not found: %s", strconv.FormatBool(ok), name)
		}

		//need this so that we don't get an unused import error for strings in some cases
		if strings.EqualFold(rs.Primary.ID, "") {
			return fmt.Errorf("LookupTable ID is not set")
		}

		id := rs.Primary.ID
		client := testAccProvider.Meta().(*Client)
		newLookupTable, err := client.GetLookupTable(id)
		if err != nil {
			return fmt.Errorf("LookupTable %s not found", id)
		}
		lookupTable = newLookupTable
		return nil
	}
}

func testAccCheckSumologicLookupTableConfigImported(name string, testFieldName string, testFieldType string, ttl int, primaryKeys string, sizeLimitAction string, description string) string {
	return fmt.Sprintf(`
data "sumologic_personal_folder" "personalFolder" {}
resource "sumologic_lookup_table" "foo" {
      name = "%s"
      fields {
        field_name = "%s"
        field_type = "%s"
      }
      ttl = %d
      primary_keys = ["%s"]
      parent_folder_id = "${data.sumologic_personal_folder.personalFolder.id}"
      size_limit_action = "%s"
      description = "%s"
}
`, name, testFieldName, testFieldType, ttl, primaryKeys, sizeLimitAction, description)
}

func testAccSumologicLookupTable(name string, testFieldName string, testFieldType string, ttl int, primaryKeys string, sizeLimitAction string, description string) string {
	return fmt.Sprintf(`
data "sumologic_personal_folder" "personalFolder" {}
resource "sumologic_lookup_table" "test" {
    name = "%s"
    fields {
      field_name = "%s"
      field_type = "%s"
    }
    ttl = %d
    primary_keys = ["%s"]
    parent_folder_id = "${data.sumologic_personal_folder.personalFolder.id}"
    size_limit_action = "%s"
    description = "%s"
}
`, name, testFieldName, testFieldType, ttl, primaryKeys, sizeLimitAction, description)
}

func testAccSumologicLookupTableUpdate(name string, testFieldName string, testFieldType string, ttl int, primaryKeys string, sizeLimitAction string, description string) string {
	return fmt.Sprintf(`
data "sumologic_personal_folder" "personalFolder" {}
resource "sumologic_lookup_table" "test" {
      name = "%s"
      fields {
        field_name = "%s"
        field_type = "%s"
      }
      ttl = %d
      primary_keys = ["%s"]
      parent_folder_id = "${data.sumologic_personal_folder.personalFolder.id}"
      size_limit_action = "%s"
      description = "%s"
}
`, name, testFieldName, testFieldType, ttl, primaryKeys, sizeLimitAction, description)
}

func testAccCheckLookupTableAttributes(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		f := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(name, "name"),
			resource.TestCheckResourceAttrSet(name, "ttl"),
			resource.TestCheckResourceAttrSet(name, "parent_folder_id"),
			resource.TestCheckResourceAttrSet(name, "size_limit_action"),
			resource.TestCheckResourceAttrSet(name, "description"),
		)
		return f(s)
	}
}
