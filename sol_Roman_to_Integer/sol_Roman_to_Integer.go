package sol_Roman_to_Integer

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return stack[len(stack) - 1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if theStack.Len() == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func symbol2Value(symbol int) int {
	switch symbol {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func romanToInt(s string) int {
	sum := 0
	for i, v := range s {
		if i == len(s) - 1 || symbol2Value(int(s[i + 1])) <= symbol2Value(int(v)) {
			sum += symbol2Value(int(v))
		} else {
			sum -= symbol2Value(int(v))
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("MDCCCLXXXIV"))
}
