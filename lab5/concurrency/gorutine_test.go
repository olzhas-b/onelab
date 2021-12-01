package concurrency

import (
	"errors"
	"testing"
)

// here we can change for all benchmark numbers
const (
	OPERATIONS      = 1000
	OperationsCycle = 100000
	Divider = 100
	ErrorsCnt = 100
)
func simpleTask() error {
	sum := 0
	for i := 0; i < OperationsCycle; i++ {
		sum += i
	}
	return nil
}
func taskWithError() error {
	sum := 0
	for i := 0; i < OperationsCycle; i++ {
		sum += i
	}
	return errors.New("greetings from error")
}

func TestExecuteExpectedErrors(t *testing.T) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	if err := Execute(tasks, 4, 5); err == nil {
		t.Error(err)
	}
}

func TestExecuteExpectedNil(t *testing.T) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	if err := Execute(tasks, 4, 1000000); err != nil {
		t.Error(err)
	}
}

func BenchmarkExecuteOneWorker(b *testing.B) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	Execute(tasks, 1, ErrorsCnt)
}

func BenchmarkExecuteTwoWorker(b *testing.B) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	Execute(tasks, 2, ErrorsCnt)
}

func BenchmarkExecuteFourWorker(b *testing.B) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	Execute(tasks, 4, ErrorsCnt)
}

func BenchmarkExecuteEightWorker(b *testing.B) {
	var tasks []func() error
	for i := 0; i < OPERATIONS; i++ {
		if i % Divider == 0 {
			tasks = append(tasks, taskWithError)
		} else {
			tasks = append(tasks, simpleTask)
		}
	}
	Execute(tasks, 8, ErrorsCnt)


}