package utilities

func IsPandigital(n int) bool {
	m := make([]bool, 10)
	l := 0

	for n > 0 {
		digit := n % 10
		n /= 10
		if m[digit] == true || digit == 0 {
			return false
		}

		m[digit] = true
		l++
	}

	for i := 1; i <= l; i++ {
		if m[i] == false {
			return false
		}
	}

	return true
}
