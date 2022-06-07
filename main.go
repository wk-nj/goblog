package main

import (
	"fmt"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi8888")
}

func main() {
	http.HandleFunc("/test", f)
	http.ListenAndServe("127.0.0.1:3000", nil)
}