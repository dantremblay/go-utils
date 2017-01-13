package validation

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	NotValidFQDN      = fmt.Errorf("fqdn is not valid")
	NotValidGroupname = fmt.Errorf("groupname is not valid")
	NotValidHostname  = fmt.Errorf("hostname is not valid")
	NotValidIPAddr    = fmt.Errorf("ip address is not valid")
	NotValidName      = fmt.Errorf("name is not valid")
	NotValidUsername  = fmt.Errorf("username is not valid")
)

func IsValidName(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)

	if !reName.MatchString(name) {
		return NotValidName
	}

	return nil
}

func IsValidUsername(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_\.]+$`)

	if !reName.MatchString(name) {
		return NotValidUsername
	}

	return nil
}

func IsValidGroupname(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)

	if !reName.MatchString(name) {
		return NotValidGroupname
	}

	return nil
}

func IsValidHostname(name string) error {
	reName := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]{1,63}$`)

	if !reName.MatchString(name) {
		return NotValidHostname
	}

	return nil
}

func IsValidIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return NotValidIPAddr
	}

	return nil
}

func IsValidFQDN(s string) error {
	if len(s) == 0 || len(s) > 254 {
		return NotValidFQDN
	}

	parts := strings.Split(s, ".")

	for i, p := range parts {
		rePart := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]{1,63}$`)

		if i == len(parts)-1 {
			rePart = regexp.MustCompile(`^[a-zA-Z]{2,63}$`)
		}

		if !rePart.MatchString(p) {
			return NotValidFQDN
		}
	}

	return nil
}
