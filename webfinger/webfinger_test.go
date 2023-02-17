package webfinger

import (
	"fmt"
	"testing"
)

func TestParseAcct(t *testing.T) {
	u, d := ParseAcct("http://example.com")
	fmt.Println(u, d)

	cd := CheckDomain("CheckDomain")
	fmt.Println(domain)
	fmt.Println(cd)
}
