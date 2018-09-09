package main

/*
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 */

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}

	ret := make([]int, 0)

	for i := 0; i < len(numbers) - 1; i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target && i != j {
				return append(ret, i + 1, j + 1)
			}
		}
	}

	return ret
}

func main() {
	
}
