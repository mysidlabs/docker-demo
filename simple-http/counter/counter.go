package main

import (
	"fmt"
	"net/http"
)

var globalCount int64

func inc(w http.ResponseWriter, r *http.Request) {
	globalCount++
	get(w, r)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, globalCount)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/inc", inc)
	mux.HandleFunc("/get", get)
	http.ListenAndServe(":8181", mux)
}
