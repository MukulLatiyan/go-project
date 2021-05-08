package main

import (
	"fmt"
	"net/http"

	"github.com/MukulLatiyan/go-project/pkg/handlers"
)

const portNumber = ":8090"

func main() {
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server has started running on:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
