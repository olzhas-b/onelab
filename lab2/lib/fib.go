package lib

func Fibonacci() func() int {
	var num1, num2 = 0, 1
	return func() int {
		num0 := num1
		num1 = num2
		num2 = num0 + num2
		return num0
	}
}
