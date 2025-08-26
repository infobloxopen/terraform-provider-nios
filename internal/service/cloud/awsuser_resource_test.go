package cloud_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForAwsuser = "access_key_id,account_id,govcloud_enabled,last_used,name,nios_user_name,status"

func TestAccAwsuserResource_basic(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")
	secret_access_key := "S1JGWfwcZWEY+hSkfpyhxigL9A/uaJ96mY"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserBasicConfig(accessKeyId, accountId, name, secret_access_key),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "access_key_id", accessKeyId),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "govcloud_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsuserResource_disappears(t *testing.T) {
	resourceName := "nios_cloud_awsuser.test"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")
	secret_access_key := "S1JGWfwcZWEY+jhSkfpyhxigL9A/ua6mY"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAwsuserDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAwsuserBasicConfig(accessKeyId, accountId, name, secret_access_key),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					testAccCheckAwsuserDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAwsuserResource_AccessKeyId(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_access_key_id"
	var v cloud.Awsuser
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")
	accessKeyId1 := "AKIA" + acctest.RandomAlphaNumeric(16)
	accessKeyId2 := "AKIA" + acctest.RandomAlphaNumeric(16)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserAccessKeyId(accessKeyId1, accountId, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "access_key_id"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAwsuserAccessKeyId(accessKeyId2, accountId, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "access_key_id"),
					resource.TestCheckResourceAttr(resourceName, "account_id", accountId),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsuserResource_AccountId(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_account_id"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	name := acctest.RandomNameWithPrefix("aws-user")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserAccountId("33773173961", accessKeyId, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "account_id", "33773173961"),
				),
			},
			//Update and Read
			{
				Config: testAccAwsuserAccountId("12345689012", accessKeyId, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "account_id", "12345689012"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsuserResource_GovcloudEnabled(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_govcloud_enabled"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserGovcloudEnabled(accountId, accessKeyId, name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "govcloud_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsuserGovcloudEnabled(accountId, accessKeyId, name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "govcloud_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsuserResource_Name(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_name"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserName(accountId, accessKeyId, name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAwsuserName(accountId, accessKeyId, name+"-updated", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO : create a aww user called aws1  and aws2
func TestAccAwsuserResource_NiosUserName(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_nios_user_name"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserNiosUserName(accountId, accessKeyId, name, "aws1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nios_user_name", "aws1"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsuserNiosUserName(accountId, accessKeyId, name, "aws2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nios_user_name", "aws2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAwsuserResource_SecretAccessKey(t *testing.T) {
	var resourceName = "nios_cloud_awsuser.test_secret_access_key"
	var v cloud.Awsuser
	accessKeyId := "AKIA" + acctest.RandomAlphaNumeric(16)
	accountId := "337773173961"
	name := acctest.RandomNameWithPrefix("aws-user")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAwsuserSecretAccessKey(accountId, accessKeyId, name, "S1JGWfwcZWEYhSkfpyhxigL9A/J96mY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secret_access_key", "S1JGWfwcZWEYhSkfpyhxigL9A/J96mY"),
				),
			},
			// Update and Read
			{
				Config: testAccAwsuserSecretAccessKey(accountId, accessKeyId, name, "K1JGWfwcZWEYYhSkfpyhxigL9A/J96mY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsuserExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secret_access_key", "K1JGWfwcZWEYYhSkfpyhxigL9A/J96mY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAwsuserExists(ctx context.Context, resourceName string, v *cloud.Awsuser) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.CloudAPI.
			AwsuserAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			//Read(ctx, rs.Primary.Attributes["ref"]).
			ReturnFieldsPlus(readableAttributesForAwsuser).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetAwsuserResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetAwsuserResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckAwsuserDestroy(ctx context.Context, v *cloud.Awsuser) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.CloudAPI.
			AwsuserAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			//Read(ctx, *v.Ref).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForAwsuser).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckAwsuserDisappears(ctx context.Context, v *cloud.Awsuser) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.CloudAPI.
			AwsuserAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAwsuserBasicConfig(accessKeyId, accountId, name, secretAccessKey string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test" {
    access_key_id = %q
    account_id    = %q
    name          = %q
    secret_access_key = %q
	govcloud_enabled = false
}
`, accessKeyId, accountId, name, secretAccessKey)
}

func testAccAwsuserAccessKeyId(accessKeyId, accountId, name string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_access_key_id" {
    access_key_id = %q
    account_id    = %q
    name          = %q
    govcloud_enabled = false
	secret_access_key = "S1JGWfwcZWEY+hSkfpyhxigL9A/ua96mY"
}
`, accessKeyId, accountId, name)
}

func testAccAwsuserAccountId(accountId, accessKeyId, name string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_account_id" {
    account_id = %q
    access_key_id = %q
    name = %q
    secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    govcloud_enabled = false
}
`, accountId, accessKeyId, name)
}

func testAccAwsuserGovcloudEnabled(accountId, accessKeyId, name, govcloudEnabled string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_govcloud_enabled" {
    account_id = %q
    access_key_id = %q
    name = %q
    secret_access_key = "S1JGWfwcZWEY+hSkfpyhxigL9A/ua96mY"
    govcloud_enabled = %q
}
`, accountId, accessKeyId, name, govcloudEnabled)
}

func testAccAwsuserName(accountId, accessKeyId, name, govcloudEnabled string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_name" {
    account_id = %q
    access_key_id = %q
    name = %q
    secret_access_key = "S1JGWfwcZWEYhSkfpyhxigL9A/uaJ6mY"
    govcloud_enabled = %q
}
`, accountId, accessKeyId, name, govcloudEnabled)
}

func testAccAwsuserNiosUserName(accountId, accessKeyId, name, niosUserName string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_nios_user_name" {
    account_id = %q
    access_key_id = %q
    name = %q
    secret_access_key = "S1JGWfwcZWEY+jhSkfpyhxigL9A/uaJ6mY"
    govcloud_enabled = false
    nios_user_name = %q
}
`, accountId, accessKeyId, name, niosUserName)
}

func testAccAwsuserSecretAccessKey(accountId, accessKeyId, name, secretAccessKey string) string {
	return fmt.Sprintf(`
resource "nios_cloud_awsuser" "test_secret_access_key" {
    account_id = %q
    access_key_id = %q
    name = %q
    govcloud_enabled = false
	secret_access_key = %q
}
`, accountId, accessKeyId, name, secretAccessKey)
}
