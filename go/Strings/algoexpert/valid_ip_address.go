package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a string formulate all Valid IP addresses
// onto an array and display
func ValidIpAddresses(str string) []string {
	validIps := []string{}

	fmt.Println(min(len(str), 4))
	for i := 0; i < min(len(str), 4); i++ {
		IpAddressParts := []string{"", "", "", ""}

		IpAddressParts[0] = str[:i]
		if !isValidPart(IpAddressParts[0]) {
			continue
		}

		for j := i + 1; j < min(len(str), i+4); j++ {
			IpAddressParts[1] = str[i:j]

			if !isValidPart(IpAddressParts[1]) {
				continue
			}

			for k := j + 1; k < min(len(str), j+4); k++ {
				IpAddressParts[2] = str[j:k]
				IpAddressParts[3] = str[k:]

				if isValidPart(IpAddressParts[2]) && isValidPart(IpAddressParts[3]) {
					ip := strings.Join(IpAddressParts, ".")

					validIps = append(validIps, ip)
				}
			}
		}
	}

	return validIps
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func isValidPart(s string) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if i > 255 {
		return false
	}

	return len(s) == len(strconv.Itoa(i)) // Covers leading zeroes as well
}

func main() {
	str := "1921680"
	fmt.Println(ValidIpAddresses(str))
}
