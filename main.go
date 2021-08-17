package main

import (
	"fmt"
	"./handlers"
	"net/http"
	"os"
)

func main() {

	// making a handler which can serve the root
	http.HandleFunc("/", handlers.RootHandler)

	// Checking whether the port 8080 is working or not
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1) // 0 is for success & 1 is for failure
	}
}
