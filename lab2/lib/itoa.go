package lib

func reverse(s *string) {
	b := []byte(*s)
	for i := 0; i < len(b) / 2; i++ {
		b[i], b[len(b) - i - 1] = b[len(b) - i - 1], b[i]
	}
	*s = string(b)
}

func Itoa(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	var lessZero bool
	if n < 0 {
		lessZero = true
		n = -n
	}
	for n != 0 {
		result += string(byte(n % 10) + 48)
		n /= 10
	}
	if lessZero {
		result += "-"
	}
	reverse(&result)
	return result
}