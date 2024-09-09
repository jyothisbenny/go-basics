package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	formatedTime := presentTime.Format("2006-01-02 15:04:05 Monday")
	fmt.Println(formatedTime)

	createdDate := time.Date(2000, time.March, 15, 12, 12, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
