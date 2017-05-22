package jsonip

import (
	"encoding/json"
	"net/http"
)

var (
	ipv4Endpoint = "https://ipv4.jsonip.com/"
	ipv6Endpoint = "https://ipv6.jsonip.com/"
)

// IP information
type JSONIP struct {
	IP string `json:"ip"`
}

// Fetch IPv4
func IPv4() (*JSONIP, error) {
	return fetch(ipv4Endpoint)
}

// Fetch IPv6
func IPv6() (*JSONIP, error) {
	return fetch(ipv6Endpoint)
}

func fetch(target string) (*JSONIP, error) {
	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ip *JSONIP
	err = json.NewDecoder(resp.Body).Decode(&ip)
	return ip, err
}
