package goawesomehelper

import (
	"net/url"
	"path"
)

//URL a url
type URL url.URL

//JoinPath join path to url
func (u *URL) JoinPath(spath string) {
	u.Path = path.Join(u.Path, spath)
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
