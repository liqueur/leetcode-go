package main

import "fmt"

func isNumber(c int) bool {
	return c >= 48 && c <= 57
}

func isLetter(c int) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
}

func isNumOrLetter(c int) bool {
	return isNumber(c) || isLetter(c)
}

func abs(a, b int) int {
	c := a - b

	if c < 0 {
		return -c
	}

	return c
}

func equalChar(a, b int) bool {
	if isLetter(a) && isLetter(b) {
		return a == b || abs(a, b) == 32
	}

	return a == b
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	l := 0
	r := len(s) - 1

	for l <= r {
		if !isNumOrLetter(int(s[l])) {
			l++
			continue
		}

		if !isNumOrLetter(int(s[r])) {
			r--
			continue
		}

		if !equalChar(int(s[l]), int(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
