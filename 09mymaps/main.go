package main

import (
	"fmt"
)

func main() {
	fmt.Println("my maps")
	names := make(map[string]string)
	fmt.Println(names)

	names["JY"] = "jyothis"
	names["AN"] = "anash"
	names["KR"] = "krishna"

	fmt.Println(names)

	delete(names, "KR")
	fmt.Println(names)

	for sn, fullName := range names {
		fmt.Println(sn, fullName)
	}
}
