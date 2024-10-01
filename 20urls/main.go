package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Hello urls")
	u, err := url.Parse("https://example.com/search?q=golang")
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)

}
