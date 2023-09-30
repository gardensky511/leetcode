package defanging_an_ip_address

import "strings"

// topic 👉 https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
