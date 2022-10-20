package main

import (
	"fmt"
	"strings"
)

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	result := ""
	counter := 0
	for _, c := range s {
		var temp string
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') && (c >= '0' || c <= '9') {
			counter++
			if counter%3 == 0 {
				temp = strings.ToUpper(string(c))
			} else {
				temp = strings.ToLower(string(c))
			}
			result += temp
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}
