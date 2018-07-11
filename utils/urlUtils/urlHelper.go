package urlUtils

import (
  "net/url"
  "path"
)

func Join(domain string, paths ...string) string {
  u, _ := url.Parse(domain)
  for path := paths {
      u.Path = path.Join(u.Path, path)
  }
  return u.String()
}
