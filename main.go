package main

import (
	"fmt"
	"goindextemplate/handlers"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8090")

	http.HandleFunc("/status", handlers.Status)
	http.ListenAndServe(":8090", nil)
}
