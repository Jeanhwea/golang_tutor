package nc0085

import (
	"strconv"
	"strings"
)

func solve(s string) (ans string) {
	if isIpv4(s) {
		return "IPv4"
	}
	if isIPv6(s) {
		return "IPv6"
	}
	return "Neither"
}

func isIpv4(s string) (ans bool) {
	nums := strings.Split(s, ".")
	if len(nums) != 4 {
		return false
	}

	for _, num := range nums {
		if num != "0" && strings.HasPrefix(num, "0") {
			return false
		}

		n, err := strconv.Atoi(num)
		if err != nil {
			return false
		}

		if n < 0 || n > 255 {
			return false
		}
	}

	return true
}

func isIPv6(s string) (ans bool) {
	nums := strings.Split(s, ":")
	if len(nums) != 8 {
		return false
	}

	for _, num := range nums {
		if len(num) <= 0 || len(num) > 4 {
			return false
		}
		for _, d := range []byte(num) {
			if (d >= 'a' && d <= 'f') || (d >= 'A' && d <= 'F') || (d >= '0' && d <= '9') {
				continue
			}
			return false
		}
	}
	return true
}
