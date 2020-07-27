package urlutil

import (
	"fmt"
	"net/url"
	"strings"
)

// MustParse attempts to parse an string to url, but ignores any errors
// if the parsing errors, the returned value is nil
func MustParse(rawUrl string) (parsedUrl *url.URL) {
	parsedUrl, _ = url.Parse(rawUrl)
	return parsedUrl
}

// WithPath copies the url object and appends a path to it
func WithPath(path string, u *url.URL) *url.URL  {
	if !strings.HasPrefix(path, "/") {
		path = "/"+path
	}

	return MustParse(fmt.Sprintf("%s%s", u.String(), path))
}