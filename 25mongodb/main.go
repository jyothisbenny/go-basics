package main

import (
	"fmt"
	"github.com/jyothisbenny/mongodb/model"
	"github.com/jyothisbenny/mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("Server is getting started!")
	model.Init()
	http.ListenAndServe(":8080", router.Router())
	fmt.Println("Server started!")

}
