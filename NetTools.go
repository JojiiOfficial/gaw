package goawesomehelper

import (
	"net/http"
	"strings"
)

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
