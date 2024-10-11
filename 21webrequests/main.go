package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
