package main

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	for num > 1 {
		if num%4 == 0 {
			num /= 4
		} else {
			return false
		}
	}

	return true
}

func main() {

}
