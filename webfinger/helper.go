package webfinger

import "fmt"

// just a helper function for finding self.href from resource object
func GetIdFromResource(resource *Resource) (string, error) {
	for _, link := range resource.Links {
		if link.Rel == "self" {
			return link.HRef, nil
		}
	}
	return "", fmt.Errorf("not webfinger, find no links with rel=self")
}
