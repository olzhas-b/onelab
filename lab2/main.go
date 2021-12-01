package main

import (
	"github.com/olzhas-b/lab3/lib"
)

func main() {
	// task 1 Itoa
	println("convert int to string", lib.Itoa(123), lib.Itoa(-123), lib.Itoa(0))

	// task 2 Atoi
	num0, _ := lib.Atoi("0")
	num1, _ := lib.Atoi("1")
	num2, _ := lib.Atoi("-2")
	//  -- => +
	num3, _ := lib.Atoi("--3")
	// --- => -
	num4, _ := lib.Atoi("---4")
	// офигенная реализация atoi, сам офигел)
	println("convert string to int", num0, num1, num2, num3, num4)

	// task 3 reverse
	println(lib.Reverse("Олжас"), lib.Reverse("Golang"))

	// task 4 sort imports in golang file
	err := lib.SortImports("/home/olzhas/go/src/github.com/olzhas-b/lab3/lib/text.txt")
	if err != nil {
		println(err)
	} else {
		println("file was reWrite successfully")
	}

	// task 5 fibonacci
	generator := lib.Fibonacci()
	for i := 0; i < 15; i++ {
		print(generator(), " ")
	}
	println()

	// task 6 recover
	var str = "123"
	var index = 3
	lib.RuneByIndex(&str, &index)
}