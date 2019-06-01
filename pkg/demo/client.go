package demo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Ping(url string) string {
	fmt.Printf("trying to connect %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server answered %s", string(body))
	return string(body)
}
func IsReady(url string) {
	for {
		time.Sleep(time.Second)

		log.Println("Checking if started...")
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Failed:", err)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Println("Not OK:", resp.StatusCode)
			continue
		}

		// Reached this point: server is up and running!
		break
	}
	log.Println("SERVER UP AND RUNNING!")

}
func StartClient(port string) chan string {
	url := fmt.Sprintf("http://localhost:%s",
		port)
	IsReady(url)
	fmt.Println("server is up and running...")
	timeout := time.Tick(1 * time.Second)
	c := make(chan string)
	go func() {
		for {
			select {
			case <-timeout:
				{
					body := Ping(url)
					c <- string(body)
				}
			}
		}
	}()

	return c
}
