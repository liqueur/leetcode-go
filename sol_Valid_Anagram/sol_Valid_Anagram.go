package main

import "fmt"

func isAnagram(s string, t string) bool {
	smap := make(map[int]int, 0)
	tmap := make(map[int]int, 0)

	for _, v := range s {
		_, in := smap[int(v)]
		if in {
			smap[int(v)] += 1
		} else {
			smap[int(v)] = 1
		}
	}

	for _, v := range t {
		_, in := tmap[int(v)]
		if in {
			tmap[int(v)] += 1
		} else {
			tmap[int(v)] = 1
		}
	}

	if len(smap) != len(tmap) {
		return false
	}

	for k, v := range smap {
		if v != tmap[k] {
			return false
		}
	}

	return true
}

func main() {
	ret := isAnagram("rat", "car")
	fmt.Println(ret)
}
