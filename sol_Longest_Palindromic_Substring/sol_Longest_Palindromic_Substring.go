package sol_Longest_Palindromic_Substring

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func longestPalindrome(s string) string {
	mx := 0
	id := 0
	resLen := 0
	resCenter := 0
	t := "$#"
	for _, v := range s {
		t += fmt.Sprintf("%c", v)
		t += "#"
	}

	t += "#"

	p := make([]int, len(t))
	for i := 1; i < len(t); i++ {
		if mx > i {
			p[i] = min(p[2 * id - i], mx - i)
		} else {
			p[i] = 1
		}
		for (i + p[i] < len(t)) && t[i + p[i]] == t[i - p[i]] {
			p[i] += 1
		}
		if mx < i + p[i] {
			mx = i + p[i]
			id = i
		}
		if resLen < p[i] {
			resLen = p[i]
			resCenter = i
		}
	}

	return s[(resCenter - resLen) / 2:(resCenter - resLen) / 2 + resLen - 1]
}

func main() {
	fmt.Println(longestPalindrome("waabwswfd"))
}
