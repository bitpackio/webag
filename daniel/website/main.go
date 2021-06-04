
// Setup
// apt install golang
//
// Run
// go run main.go
//
// Usage
// Serve static html content
// browser "http://localhost:8080"

package main

import (
    "fmt"
    "net/http"
)

const (
    verbose bool = true
)

var (
    serverPort string = "0.0.0.0:8080"
)

func log(s string) {
    if(verbose) { 
        fmt.Printf("%s", s + "\n") 
    }
}

func logMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log("Got a " + r.Method + " request for: " + r.URL.Path)
        handler.ServeHTTP(w, r)
        log("Handler finished processing request.")
    })
}

func main() {
    fs := http.FileServer(http.Dir("html"))
    http.Handle("/", fs)
    logHandler := logMiddleware(http.DefaultServeMux)
    log("WebServer ready and waiting for processing requests.")

    if err := http.ListenAndServe(serverPort, logHandler); err != nil {
        fmt.Println(err)
    }
}
