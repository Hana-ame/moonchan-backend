package webfinger

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/moonchan-backend/utils"
)

func TestParseAcct(t *testing.T) {
	u, d := utils.ParseAcct("http://example.com")
	fmt.Println(u, d)

	cd := CheckDomain("CheckDomain")
	fmt.Println(domain)
	fmt.Println(cd)
}
