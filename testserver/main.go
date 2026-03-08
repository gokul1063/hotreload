package main

import (
	"fmt"
	"net/http"

	"hotreload/testserver/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HelloHandler)

	fmt.Println("server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
