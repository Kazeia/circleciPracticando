package main

import (
    "fmt"
    "net/http"
)

// Circleci Test1
func HelloWorld(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World")
}

func main() {
    http.HandleFunc("/", HelloWorld)
    http.ListenAndServe(":3000", nil)
}
