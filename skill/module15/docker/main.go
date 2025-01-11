package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!!\n")
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println(("Go backend"))
	http.ListenAndServe(":8787", nil)
}
