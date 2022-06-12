package main

import (
	"strconv"
	"strings"
)

// 997430
// 997.4.3.0
func ValidIPAddresses(str string) []string {
	result := []string{}
	if len(str) < 4 {
		return result
	}

	i := 1
	j := 2
	k := 3
	min := 1
	max := 3

	if len(str) == 4 {
		max = 1
	} else if len(str) == 5 {
		max = 2
	} else if len(str) == 11 {
		min = 2
	} else if len(str) == 12 {
		min = 3
	}

	for i <= max {
		ip := str[:i] + "." + str[i:j] + "." + str[j:k] + "." + str[k:]
		if isValidIP(ip) {
			result = append(result, ip)
		}

		if len(str)-k > min {
			k++
		} else if j-i < max && k < len(str) {
			j++
			k = j + 1
		} else {
			i++
			j = i + 1
			k = j + 1
		}
	}

	return result
}

func isValidIP(str string) bool {
	ips := strings.Split(str, ".")
	for _, v := range ips {
		if len(v) < 1 || len(v) > 3 {
			return false
		}

		if len(v) > 1 {
			if string(v[0]) == "0" {
				return false
			}
		}

		ip, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if ip < 0 || ip > 255 {
			return false
		}
	}

	return true
}
