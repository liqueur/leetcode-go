package sol_Valid_Parentheses

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := list.List{}
	for _, v := range s {
		if stack.Len() == 0 {
			stack.PushBack(v)
		} else {
			switch {
			case stack.Back().Value == '(' && v == ')',
				stack.Back().Value == '{' && v == '}',
				stack.Back().Value == '[' && v == ']':
				stack.Remove(stack.Back())
			default:
				stack.PushBack(v)
			}
		}
	}

	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("("))
}
