package main

import (
	"fmt"
	"github.com/olzhas-b/lab4/tetris"
	"github.com/olzhas-b/lab4/top"
)

func outputFormat(p tetris.Pair ) string {
	return fmt.Sprintf(" \nstruct with memory %d {\n %s }", p.Sz, p.Order)
}

func main() {
	// task1
	res := top.TopWords("aaabb    d", 5)
	fmt.Printf("%+v\n", res)

	// task2
	ans := tetris.GetTopMinimizedStruct("tetris/generator/example_struct.go")
	for i := 0; i < len(ans); i++ {
		print(outputFormat(ans[i]))
	}



}
