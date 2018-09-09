package sol_Count_and_Say

import "fmt"

/*
Input: 1
Output: "1"

Input: 4
Output: "1211"

1.     1
2.     11
3.     21
4.     1211
5.     111221
 */


func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	bunch := "1"

	for k := 1; k < n; k++ {
		newBunch := ""
		for i := 0; i < len(bunch); i++ {
			cnt := 0
			for j := i; j < len(bunch); j++ {
				if bunch[j] == bunch[i] {
					cnt += 1
				} else {
					break
				}
				i = j
			}
			newBunch += fmt.Sprintf("%v%c", cnt, bunch[i])
		}
		bunch = newBunch
	}

	return bunch
}


func main() {
	fmt.Println(countAndSay(5))
}