package webfinger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const WebFingerPath = "/.well-known/webfinger"

var Domain = "moonchan.xyz"
var isAcctExist func(username string) bool = func(username string) bool { return true }

func WebfingerHandler(c *gin.Context) {
	acct, ok := c.GetQuery("resource")
	if !ok {
		c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if !strings.HasPrefix(acct, "acct:") {
		c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	username, domain := parseAcct(acct)
	if domain != Domain {
		c.String(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	if !isAcctExist(username) {
		c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// all right
	data := getResource(username)
	c.AsciiJSON(http.StatusOK, data)
}

// just a helper function for finding self.href from resource object
func GetIdFromWebfinger(res *Resource) (string, error) {
	for _, link := range res.Links {
		if link.Rel == "self" {
			return link.HRef, nil
		}
	}
	return "", fmt.Errorf("not webfinger, find no links with rel=self")
}

func parseAcct(acct string) (username string, domain string) {
	arr := strings.Split(acct[5:], "@")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return acct, ""
}

type Link struct {
	HRef       string             `json:"href,omitempty"`
	Type       string             `json:"type,omitempty"`
	Rel        string             `json:"rel,omitempty"`
	Template   string             `json:"template,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
	Titles     map[string]string  `json:"titles,omitempty"`
}

type Resource struct {
	Subject    string            `json:"subject,omitempty"`
	Aliases    []string          `json:"aliases,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Links      []Link            `json:"links"`
}

func getResource(username string) *Resource {
	// return mockGet(username)

	// TODO: sould connect to database to find Aliases
	self := &Link{
		Rel:  "self",
		Type: "application/activity+json",
		HRef: fmt.Sprintf("https://%s/users/%s", Domain, username),
	}
	// porfile
	profilePage := &Link{
		Rel:  "http://webfinger.net/rel/profile-page",
		Type: "text/html",
		HRef: fmt.Sprintf("https://%s/@%s", Domain, username),
	}
	// what's this used for?
	subscribe := &Link{
		Rel:      "http://ostatus.org/schema/1.0/subscribe",
		Template: fmt.Sprintf("https://%s/authorize-follow?acct={uri}", Domain),
	}

	r := &Resource{
		Subject: fmt.Sprintf("acct:%s@%s", username, Domain),
		Aliases: nil, // TODO
		Links: []Link{
			*self,
			*profilePage,
			*subscribe,
		},
	}
	return r
	// return mockGet(username)
}
