package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello webrequests")
	url := "https://en.wikipedia.org/wiki/Pacific_Ocean"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("response status:", response.Status)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response body:", string(content))
}
