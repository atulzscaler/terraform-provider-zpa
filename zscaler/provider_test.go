package zscaler

import (
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// CharSetAlphaNum is the alphanumeric character set for use with
	// RandStringFromCharSet
	CharSetAlphaNum = "abcdefghijklmnopqrstuvwxyz012346789"

	// CharSetAlpha is the alphabetical character set for use with
	// RandStringFromCharSet
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyz"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

// testAccProviderFactories is a static map containing only the main provider instance
//
// Use other testAccProviderFactories functions, such as testAccProviderFactoriesAlternate,
// for tests requiring special provider configurations.
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zscaler": testAccProvider,
	}

	// Always allocate a new provider instance each invocation, otherwise gRPC
	// ProviderConfigure() can overwrite configuration during concurrent testing.
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"zscaler": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

// RandStringFromCharSet generates a random string by selecting characters from
// the charset provided
func RandStringFromCharSet(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ZPA_CLIENT_ID"); v == "" {
		t.Fatal("ZPA_CLIENT_ID must be set for acceptance tests.")
	}
	if v := os.Getenv("ZPA_CLIENT_SECRET"); v == "" {
		t.Fatal("ZPA_CLIENT_SECRET must be set for acceptance tests.")
	}
	if v := os.Getenv("ZPA_CUSTOMER_ID"); v == "" {
		t.Fatal("ZPA_CUSTOMER_ID must be set for acceptance tests.")
	}
}
