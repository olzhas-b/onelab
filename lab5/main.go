package main

import (
	"errors"
	"fmt"
	"github.com/olzhas-b/lab6/concurrency"
)
func simpleTask() error {
	sum := 0
	for i := 0; i < 100000; i++ {
		sum += i
	}
	return nil
}
func taskWithError() error {
	sum := 0
	for i := 0; i < 100000; i++ {
		sum += i
	}
	return errors.New("greetings from error")
}
func main() {
	var tasks []func() error
	for i := 0; i < 100000; i++ {
		if i % 100 == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	if err := concurrency.Execute(tasks, 8, 100); err != nil {
		fmt.Println(err)
	}
}