package main

import (
	"bufio"
	"context"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"net"
	"strconv"
	"sync"
)

func Server(addr string, numClient int64, wg *sync.WaitGroup) {
	defer wg.Done()
	listener, _ := net.Listen("tcp", addr)
	println("tcp server started listen port" + addr)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	sem := semaphore.NewWeighted(numClient)
	for {
		select {
		case <-ctx.Done():
			listener.Close()
			return
		default:
			conn, err := listener.Accept()
			if err == nil {
				conn := conn
				if err := sem.Acquire(ctx, 1); err != nil {

				}
				group.Go( func() error {
					defer sem.Release(1)
					for {
						message, errRead := bufio.NewReader(conn).ReadString('\n')
						if errRead != nil {
							return errRead
						}
						if len(message) >= 1 {
							message = message[0:len(message) - 1]
						}
						if message == "quit" {
							errClose := conn.Close()
							println("conn closed with :" + conn.RemoteAddr().String())
							return errClose
						}
						//fmt.Printf("\n out ans %v  === %v \n", len(message), message)
						if message == "quitAll" {
							println("conn closed with :" + conn.RemoteAddr().String())
							conn.Close()
							cancel()
							return nil
						}
						num, err := strconv.Atoi(message)
						if err != nil {
							return err
						}
						num *= num
						msgNum := strconv.Itoa(num)
						if len(msgNum) != 0 {
							conn.Write([]byte(msgNum+"\n"))
						}
					}
				})
			}
		}
	}
}
