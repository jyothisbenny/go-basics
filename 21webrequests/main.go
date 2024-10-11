package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	performGetRequest()
}

func performGetRequest() {
	const url = "http://192.168.10.54:7000/health/ping"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(content))
}

func performPostRequest() {
	const url = "http://192.168.10.54:7000/health/ping"
	body := strings.NewReader(`
		{ 
			"hy": "hello",
			"how": "are"
		}
 	`)

	response, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Body:", string(content))
}

//skipping form data post, video 30
