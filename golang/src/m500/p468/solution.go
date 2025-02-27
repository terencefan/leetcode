package main

import (
	"strings"
)

func ValidIPv4(segments []string) bool {
	for _, segment := range segments {
		var n = 0
		if len(segment) > 3 || len(segment) == 0 {
			return false
		} else if len(segment) > 1 && segment[0] == '0' {
			return false
		}
		for _, b := range []byte(segment) {
			n = n * 10
			if b < '0' || b > '9' {
				return false
			}
			n += int(b - 48)
			if n > 255 {
				return false
			}
		}
	}
	return true
}

func ValidIPv6(segments []string) bool {
	for _, segment := range segments {
		if len(segment) > 4 || len(segment) == 0 {
			return false
		}
		for _, b := range []byte(segment) {
			if '0' <= b && b <= '9' {
			} else if 'A' <= b && b <= 'F' {
			} else if 'a' <= b && b <= 'f' {
			} else {
				return false
			}
		}
	}
	return true
}

func ValidIPAddress(IP string) string {
	var s []string

	s = strings.SplitN(IP, ".", -1)
	if len(s) == 4 && ValidIPv4(s) {
		return "IPv4"
	}

	s = strings.SplitN(IP, ":", -1)
	if len(s) == 8 && ValidIPv6(s) {
		return "IPv6"
	}

	return "Neither"
}
