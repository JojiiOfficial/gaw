package goawesomehelper

import (
	"net/url"
	"path"
)

//URLJoinPath join path to url
func URLJoinPath(surl, spath string) (string, error) {
	u, err := url.Parse(surl)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, spath)
	return u.String(), nil
}
