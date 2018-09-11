package main

import "fmt"

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}

	i := 0
	j := len(s) - 1

	newlist := []byte(s)

	for i < j {
		if isVowel(newlist[i]) && isVowel(newlist[j]) {
			newlist[i], newlist[j] = newlist[j], newlist[i]
			i++
			j--
		} else if !isVowel(newlist[i]) {
			i++
		} else {
			j--
		}
	}

	return string(newlist)
}

func main() {
	ret := reverseVowels("ai")
	fmt.Println(ret)
}
