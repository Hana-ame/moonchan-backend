package utils

import (
	"net/http"
	"strings"
)

func Get(url string) (*http.Response, error) {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return "", err
	// }
	return http.Get(url)
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }
	// return string(body), nil
}

func ParseAcct(acct string) (username string, domain string) {
	// if strings.HasPrefix(acct, "acct:") {
	// 	acct = acct[5:]
	// }
	acct = strings.TrimPrefix(acct, "acct:")

	arr := strings.Split(acct, "@")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return acct, ""
}
