package main

import (
	"fmt"
	"net/http"
)

const Port string = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting server on port %s...", Port))
	_ = http.ListenAndServe(Port, nil)
}
