package main

import (
	"fmt"
	"goindextemplate/handlers"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8090")

	handlers.InitData()

	http.HandleFunc("/status", handlers.Status)
	http.HandleFunc("/rerun", handlers.Rerun)
	http.ListenAndServe(":8090", nil)
}
