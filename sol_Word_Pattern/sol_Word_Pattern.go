package main

import "fmt"

func wordPattern(pattern string, str string) bool {
	p1 := makeShortPattern(pattern)
	p2 := makeLongPattern(str)

	if len(p1) != len(p2) {
		return false
	}

	for i := 0; i < len(p1); i++ {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}

func makeLongPattern(word string) []int {
	p := make([]int, 0)
	pmap := make(map[interface{}]int, 0)
	tokens := makeTokens(word)

	for i, c := range tokens {
		_, in := pmap[c]
		if !in {
			pmap[c] = i
		}
	}

	for _, c := range tokens {
		p = append(p, pmap[c])
	}

	return p
}

func makeShortPattern(word string) []int {
	p := make([]int, 0)
	pmap := make(map[interface{}]int, 0)

	for i, c := range word {
		_, in := pmap[c]
		if !in {
			pmap[c] = i
		}
	}

	for _, c := range word {
		p = append(p, pmap[c])
	}

	return p
}

func makeTokens(words string) []string {
	w := ""
	words += " "
	tokens := make([]string, 0)
	for _, c := range words {
		if c != ' ' {
			w += string(c)
		} else if w != "" {
			tokens = append(tokens, w)
			w = ""
		}
	}

	return tokens
}

func main() {
	ret := wordPattern("abba", "dog cat cat fish")
	fmt.Println(ret)
}
