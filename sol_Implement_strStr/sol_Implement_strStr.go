package sol_Implement_strStr

import "fmt"

/*
Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
 */

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	idx := -1

	for i := 0; i < len(haystack); i++ {
		idx = i
		k := i
		for j:= 0; j < len(needle); j++ {
			if k == len(haystack) || haystack[k] != needle[j] {
				idx = -1
				break
			}

			k += 1
		}

		if idx != -1 {
			break
		}
	}

	return idx
}

func main() {
	//fmt.Println(strStr("mississippi", "issip"))
	//fmt.Println(strStr("", ""))
	//fmt.Println(strStr("", "issip"))
	fmt.Println(strStr("aaa", "aaa"))
}
