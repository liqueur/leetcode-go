package main

func canConstruct(ransomNote string, magazine string) bool {
	smap1 := make(map[interface{}]int, 0)
	smap2 := make(map[interface{}]int, 0)

	for _, c := range ransomNote {
		smap1[c] += 1
	}

	for _, c := range magazine {
		smap2[c] += 1
	}

	if len(smap1) > len(smap2) {
		return false
	}

	for k, _ := range smap1 {
		if smap1[k] > smap2[k] {
			return false
		}
	}

	return true
}

func main() {

}
