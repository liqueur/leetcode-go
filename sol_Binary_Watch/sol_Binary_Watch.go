package main

import "fmt"

func count(num int) int {
	cnt := 0

	for num > 0 {
		if num%2 > 0 {
			cnt++
		}
		num /= 2
	}

	return cnt
}

func readBinaryWatch(num int) []string {
	wmap := make(map[string]int, 0)
	ret := make([]string, 0)

	for h := 0; h <= 11; h++ {
		for m := 0; m <= 59; m++ {
			hs := fmt.Sprintf("%v", h)
			ms := fmt.Sprintf("%v", m)
			if m < 10 {
				ms = "0" + ms
			}
			wmap[hs+":"+ms] = count(h) + count(m)
		}
	}

	for k, v := range wmap {
		if v == num {
			ret = append(ret, k)
		}
	}

	return ret
}

func main() {
	ret := readBinaryWatch(1)
	fmt.Println(ret)
}
