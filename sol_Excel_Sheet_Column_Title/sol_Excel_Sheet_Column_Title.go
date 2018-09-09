package main

import "fmt"

func convertToTitle(n int) string {
	title := ""
	for n > 0 {
		title = fmt.Sprintf("%c", (n - 1) % 26 + 65) + title
		n = (n - 1) / 26
	}

	return title
}

func main() {
	fmt.Println(convertToTitle(701))
}
