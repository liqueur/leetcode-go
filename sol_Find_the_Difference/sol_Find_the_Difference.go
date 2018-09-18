package main

import "fmt"

func findTheDifference(s string, t string) byte {
	smap := make(map[byte]int, 0)
	tmap := make(map[byte]int, 0)

	for _, c := range s {
		smap[byte(c)] += 1
	}

	for _, c := range t {
		tmap[byte(c)] += 1
	}

	for k, _ := range tmap {
		_, in := smap[k]
		if !in || tmap[k] > smap[k] {
			return k
		}
	}

	return 0
}

func main() {
	ret := findTheDifference("a", "aa")
	fmt.Printf("%c\n", ret)
}
