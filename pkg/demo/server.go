package demo

import (
	"fmt"
	"net/http"
)

func StartServer(port string) chan string {

	fmt.Printf("starting server on port %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	c := make(chan string)
	go func(port string) {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		fmt.Println("Server error]", err)
		c <- fmt.Sprintf("server is shutted down %s", err.Error())
		close(c)
	}(port)

	fmt.Printf("%v listening on port", port)
	return c
}
