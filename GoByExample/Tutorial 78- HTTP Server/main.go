package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) { // Just responds on request

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) { // Uses the data in the request

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)     // Handler on endpoint /hello
	http.HandleFunc("/headers", headers) // Handler on endpoint /headers

	http.ListenAndServe(":8090", nil) // Serve on port
}
