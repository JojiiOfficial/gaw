package gaw

import (
	"net/url"
	"path"
	"strings"
)

//URL a url
type URL url.URL

//JoinPath join path to url
func (u *URL) JoinPath(sPath string) {
	u.Path = path.Join(u.Path, sPath)
}

//RemoveSubdomain removes the subdomain
func (u *URL) RemoveSubdomain(sPath string) string {
	return RemoveSubdomain((url.URL)(*u))
}

//ParseURL parses an URL to a goawesomehelper.URL
func ParseURL(u string) (*URL, error) {
	purl, err := url.Parse(u)
	ue := (URL)(*purl)
	return &ue, err
}

//URLJoinPath join path to url
func URLJoinPath(surl, spath string) (string, error) {
	u, err := url.Parse(surl)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, spath)
	return u.String(), nil
}

//RemoveSubdomain removes subdomains from url
func RemoveSubdomain(u url.URL) string {
	parts := strings.Split(u.Hostname(), ".")
	return parts[len(parts)-2] + "." + parts[len(parts)-1]
}
