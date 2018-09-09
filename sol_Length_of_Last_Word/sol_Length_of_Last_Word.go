package sol_Length_of_Last_Word

import "fmt"

/*
Input: "Hello World"
Output: 5
 */


func lengthOfLastWord(s string) int {
	var start, end int

	for end = len(s) - 1; end >= 0; end-- {
		if s[end] != ' ' {
			break
		}
	}

	for start = end; start >= 0; start-- {
		if s[start] == ' ' {
			break
		}
	}

	return end - start
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World "))
}
