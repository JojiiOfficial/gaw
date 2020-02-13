package gaw

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//ReservedIPs that are reserved (RFC 1918)
var ReservedIPs = []string{
	"0.0.0.0/8",
	"10.0.0.0/8",
	"127.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.0.2.0/24",
	"192.88.99.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
	"224.0.0.0/4",
	"240.0.0.0/4",
}

//GetIPFromHTTPrequest gets the real IP from the request
func GetIPFromHTTPrequest(r *http.Request) string {
	ipheader := []string{"X-Forwarded-For", "X-Real-Ip", "HTTP_CLIENT_IP", "HTTP_X_FORWARDED_FOR", "HTTP_X_FORWARDED", "HTTP_X_CLUSTER_CLIENT_IP", "HTTP_FORWARDED_FOR", "HTTP_FORWARDED", "REMOTE_ADDR"}
	var repIP string
	for _, header := range ipheader {
		cip := r.Header.Get(header)
		cip = strings.Trim(cip, " ")
		if len(cip) > 0 {
			repIP = cip
			break
		}
	}
	if len(strings.Trim(repIP, " ")) == 0 {
		repIP = r.RemoteAddr
	}
	if strings.Contains(repIP, ":") {
		repIP = repIP[:(strings.LastIndex(repIP, ":"))]
	}
	return repIP
}

func inPortValid(port int) bool {
	return port > 0 && port < 65535
}

// GetHTMLFromURL returns the whole Body of the given website by doing a simple GET request
func GetHTMLFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		return "", err
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}
