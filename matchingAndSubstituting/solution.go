package matchingAndSubstituting

import (
	"fmt"
	"strings"
	"unicode"
)

func areAllDigits(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isPhoneValid(phone string) bool {
	splitPhone := strings.Split(strings.Split(phone, " ")[1], "-")
	if len(splitPhone) != 4 {
		return false
	}

	if len(splitPhone[1]) != 3 || len(splitPhone[2]) != 3 || len(splitPhone[3]) != 4 {
		return false
	}

	return splitPhone[0] == "+1" && areAllDigits(splitPhone[1]) && areAllDigits(splitPhone[2]) && areAllDigits(splitPhone[3])
}

func isVersionValid(v string) bool {
	splitVersion := strings.Split(v, ".")
	return len(splitVersion) == 2 && areAllDigits(splitVersion[0]) && areAllDigits(splitVersion[1])
}

func getVersion(v, versionFromInput string) string {
	if v != "2.0" {
		return versionFromInput
	}

	return v
}

func Change(s, prog, version string) string {
	fmt.Sprintf(s)
	data := strings.Split(s, "\n")
	fileVersion := strings.Split(data[5], " ")[1]

	if !isPhoneValid(data[3]) || !isVersionValid(fileVersion) {
		return "ERROR: VERSION or PHONE"
	}

	return fmt.Sprintf("Program: %s Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: %s", prog, getVersion(fileVersion, version))
}
