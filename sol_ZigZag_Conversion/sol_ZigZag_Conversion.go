package sol_ZigZag_Conversion

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ns := ""
	interval := numRows * 2 - 2

	for i := 0; i < numRows; i++ {
		inter := interval - 2 * i
		for j := i; j < len(s); {
			ns += fmt.Sprintf("%c", s[j])
			if i == 0 || i == numRows - 1 {
				j += interval
			} else {
				j += inter
				inter = interval - inter
			}
		}
	}

	return ns
}


func main() {
	fmt.Println(convert("A", 1))
}
