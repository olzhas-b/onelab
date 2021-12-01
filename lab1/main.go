package main

import (
	"fmt"
	"github.com/olzhas-b/one_lab_test_module/lib"
)

func main() {
	println(lib.LowerToUpper("WhatIsThat123IsItNumberOrWHat"))
	println(lib.UpperToLower("WhatIsThat123IsItNumberOrWHat"))

	x1, x2, err := lib.FindRootQuEq(8, 2, -3)
	if err != nil {
		println(err)
	} else {
		fmt.Printf("first root is: %f \nsecond root is:%f\n", x1, x2)
	}
	fmt.Printf("generated id is:%s\n", lib.GetId())

}
