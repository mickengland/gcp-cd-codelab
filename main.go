package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello Kubernetes. Let's test a rollback!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":80", nil)
}
