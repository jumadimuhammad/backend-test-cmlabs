package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{
		"SAIPPUAKIVIKAUPPIAS",
		"Aibohphobia",
		"Anna",
		"Civic",
		"My gym",
		"No lemon, no melon",
		"kasur RUsaK",
	}

	for _, v := range arr {
		fmt.Println(isPalindrome(v))
	}
}

func isPalindrome(word string) (flag bool) {

	str := strings.ReplaceAll(word, " ", "")

	for i := 0; i < len(str)/2; i++ {
		if !strings.EqualFold(string(str[i]), strings.ToLower(string(str[len(str)-i-1]))) {
			return
		}
	}

	flag = true
	return
}
