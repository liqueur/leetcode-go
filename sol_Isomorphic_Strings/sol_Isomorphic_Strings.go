package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := make([]int, len(s))
	tt := make([]int, len(t))

	ssCheck := make(map[int]int, 0)
	ttCheck := make(map[int]int, 0)

	for i, v := range s {
		_, in := ssCheck[int(v)]
		if !in {
			ssCheck[int(v)] = i
		}

		ss[i] = ssCheck[int(v)]
	}

	for i, v := range t {
		_, in := ttCheck[int(v)]
		if !in {
			ttCheck[int(v)] = i
		}

		tt[i] = ttCheck[int(v)]
	}

	for i := 0; i < len(s); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}

	return true
}

func main() {

}
