package main

/*
class Solution {
public:
    int compress(vector<char>& chars) {
        int n = chars.size(), cur = 0;
        for (int i = 0, j = 0; i < n; i = j) {
            while (j < n && chars[j] == chars[i]) ++j;
            chars[cur++] = chars[i];
            if (j - i == 1) continue;
            for (char c : to_string(j - i)) chars[cur++] = c;
        }
        return cur;
    }
};
*/

func toString(num int) []byte {
	ret := make([]byte, 0)
	for num > 0 {
		ret = append(ret, byte(num%10+48))
		num /= 10
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return ret
}

func compress(chars []byte) int {
	n := len(chars)
	cur := 0
	i := 0
	j := 0
	for i < n {
		for j < n && chars[j] == chars[i] {
			j++
		}
		chars[cur] = chars[i]
		cur++
		if j-i == 1 {
			i = j
			continue
		}
		for _, c := range toString(j - i) {
			chars[cur] = byte(c)
			cur++
		}
		i = j
	}
	return cur
}

func main() {
}
