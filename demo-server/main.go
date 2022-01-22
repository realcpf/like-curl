package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)

	w.Header().Add("me", "realcpf")
	fmt.Fprintf(w, "hello word")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("localhost:8080", nil)
}
