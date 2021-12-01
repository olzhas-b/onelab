package main

import (
	"bufio"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"
)

const (
	port = ":3333"
)
func TestWithOneConnection(t *testing.T) {
	var wg sync.WaitGroup
	go Server(port, 4, &wg)
	time.Sleep(time.Second)
	testTable := struct {
		input []string
		expected []string
		output []string
	}{
		input: []string{"1\n","2\n", "3\n", "quitAll\n" },
		expected: []string{"1\n", "4\n", "9\n"},
		output: []string{},
	}
	conn, _ := net.Dial("tcp", "127.0.0.1" + port)
	for _, num := range testTable.input {
		_, err := conn.Write([]byte(num))
		if err != nil {
			panic(err)
		}
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		print(response)
		testTable.output = append(testTable.output, response)
	}
	conn.Close()
	if !reflect.DeepEqual(testTable.output, testTable.expected) {
		t.Errorf("not equal")
	}

}