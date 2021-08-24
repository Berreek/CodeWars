package ipValidation

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	ipSections := strings.Split(ip, ".")
	if len(ipSections) != 4 {
		return false
	}

	for _, section := range ipSections {
		if strings.HasPrefix(section, "0") && len(section) > 1 {
			return false
		}

		num, err := strconv.Atoi(section)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}
