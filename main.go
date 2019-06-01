package main

import (
	"flag"
	"fmt"

	"github.com/Codefresh-Examples/multi-containers-in-pod/pkg/demo"
)

func main() {
	port := flag.String("port", "1111", "server's port")
	runClient := flag.Bool("client", false, "run client")
	runServer := flag.Bool("server", true, "run server")

	flag.Parse()

	fmt.Printf("%v %v %v\n", *port, *runClient, *runServer)

	var s chan string
	var c chan string
	if *runServer {
		fmt.Println("start server ...")
		go func() {
			s = demo.StartServer(*port)
		}()
	}

	if *runClient {
		fmt.Println("start client ...")

		c = demo.StartClient(*port)
	}

	for {
		select {
		case <-c:
			fmt.Printf("[client]%s\n", <-c)

		case <-s:
			fmt.Printf("[server]%s\n", <-s)
		}
	}

}
