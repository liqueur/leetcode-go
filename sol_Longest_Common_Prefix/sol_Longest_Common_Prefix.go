package sol_Longest_Common_Prefix

import "fmt"

/*
Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
 */


func longestCommonPrefix(strs []string) string {
	com := ""

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return com
			}
		}
		com += string(strs[0][i])
	}

	return com
}


func main() {
	fmt.Println(longestCommonPrefix([]string{"lower","flow","flight"}))
}
