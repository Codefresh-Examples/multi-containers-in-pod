package demo

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("start server ...")
	var s chan string
	var c chan string
	go func() {
		s = StartServer(":1111")
	}()
	fmt.Println("start client ...")

	c = StartClient("1111")

	for {
		select {
		case <-c:
			fmt.Printf("[client]%s\n", <-c)

		case <-s:
			fmt.Printf("[server]%s\n", <-s)
		}
	}

}
