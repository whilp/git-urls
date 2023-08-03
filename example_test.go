package giturls_test

import (
	"fmt"

	giturls "github.com/whilp/git-urls"
)

func ExampleParse() {
	result, _ := giturls.Parse("https://user:password@host.xz/organization/repo.git?ref=feature/test")

	fmt.Printf("Protcol: %s\n", result.Scheme)
	fmt.Printf("User/Password: %s\n", result.User)
	fmt.Printf("Host: %s\n", result.Host)
	fmt.Printf("Path: %s\n", result.Path)
	fmt.Printf("Query: %s\n", result.RawQuery)
	// Output:
	// Protcol: https
	// User/Password: user:password
	// Host: host.xz
	// Path: /organization/repo.git
	// Query: ref=feature/test
}
