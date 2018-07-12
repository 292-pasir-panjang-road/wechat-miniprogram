package urlUtils

import (
	"net/url"
	"path"
)

func Join(domain string, paths ...string) string {
	u, _ := url.Parse(domain)
	for _, pathItem := range paths {
		u.Path = path.Join(u.Path, pathItem)
	}
	return u.String()
}
