package main

import (
	"fmt"
	"net/http"

	ascii "ascii/funcs"
)

const port = ":8080"

func main() {
	// Register handler function with URL pattern "/"
	http.HandleFunc("/", ascii.Handler)

	fmt.Println("(http://localhost:8080/)server started on port :", port)
	http.ListenAndServe(port, nil)
}
