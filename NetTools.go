package gaw

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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

//IsIPReserved returns true if the given IP is a reserved ip
func IsIPReserved(ip string) (bool, error) {
	pip := net.ParseIP(ip)
	if pip.To4() == nil {
		return false, errors.New("No IPv4")
	}

	for _, reservedIP := range ReservedIPs {
		_, subnet, err := net.ParseCIDR(reservedIP)
		if err != nil {
			return false, err
		}
		if subnet.Contains(pip) {
			return true, nil
		}
	}

	return false, nil
}

//IsReserved returns true if inp is reserved (if its a url a dns lookup will be made)
func IsReserved(inp string) (bool, error) {
	trial := net.ParseIP(inp)
	if trial.To4() != nil {
		isRes, err := IsIPReserved(inp)
		if err != nil {
			fmt.Println(err.Error())
			return false, err
		}
		return isRes, nil
	}
	host := inp

	//Get hostname if it is an url
	u, err := url.Parse(host)
	if err == nil {
		if len(u.Hostname()) > 0 {
			host = u.Hostname()
		}
	}

	//Lookup host
	ips, err := net.LookupHost(host)
	if err != nil {
		return false, err
	}

	for _, ip := range ips {
		is, err := IsReserved(ip)
		if err != nil {
			return false, err
		}

		if is {
			return true, nil
		}
	}

	return false, nil
}

//GetIPFromHTTPrequest gets the real IP from the request
func GetIPFromHTTPrequest(r *http.Request) string {
	ipHeader := []string{"X-Forwarded-For", "X-Real-Ip", "HTTP_CLIENT_IP", "HTTP_X_FORWARDED_FOR", "HTTP_X_FORWARDED", "HTTP_X_CLUSTER_CLIENT_IP", "HTTP_FORWARDED_FOR", "HTTP_FORWARDED", "REMOTE_ADDR"}
	var repIP string
	for _, header := range ipHeader {
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

//IsPortValid returns true if the given TCP port is valid
func IsPortValid(port int) bool {
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

//IPtoInt converts an IP to an integer
func IPtoInt(sIP string) uint {
	ip := net.ParseIP(sIP)
	return uint(ip[12])*16777216 + uint(ip[13])*65536 + uint(ip[14])*256 + uint(ip[15])
}

//IntToIP converts an integer to an IP
func IntToIP(ip uint) net.IP {
	return net.IPv4(byte((ip>>24)&0xFF), byte((ip>>16)&0xFF), byte((ip>>8)&0xFF), byte(ip&0xFF))
}
