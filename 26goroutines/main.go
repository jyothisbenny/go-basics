package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//first method using timer
	go greeter("Hello")
	greeter("World")

	//second method using sync wait groups
	websites := []string{
		"https://google.com/",
		"https://fb.com/",
		"https://instagram.com/",
		"https://gmail.com/",
		"https://github.com/",
	}
	for _, endponint := range websites {
		go webRequest(endponint)
		wg.Add(1)
	}

	wg.Wait()

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func webRequest(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("%d status code on endpoint %s\n", res.StatusCode, endpoint)
	}
}
