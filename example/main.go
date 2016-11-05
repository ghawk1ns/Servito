package main

import (
    "github.com/ghawk1ns/servito"
    "net/http"
    "fmt"
    "log"
)

func main() {
    // Create a new route and provide a HandlerFunc
    servito.AddRoute("Index", "GET", "/", Index)

    // You can set optional config values, otherwise the defaults will be used
    // servito.SetAddress("localhost")
    // servito.SetPort("8000")
    // servito.SetDebug(true)

    // optionally: you can pass an entire config block
    // servito.LoadConfigFromFile(...)
    // servito.LoadConfigFromJSON(...)
    // servito.LoadConfigFromPath(...)

    // Add your own interceptor, this is a great place to add logs, debug, validate params and auth
    servito.AddInterceptor(myInterceptor)

    // Finally, simply start the server!
    servito.StartServer()
}

// a HandlerFunc takes request r and writes the output to writer w
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}

// return true if you "handle" the request. For example, specific requests should include a specific parameter.
// If you intercept that request and it doesn't the value you expect, you can handle the request here without passing it on
var myInterceptor = func(w http.ResponseWriter, r *http.Request) bool {
        log.Printf("Intercepting request to %v", r.RequestURI)
        // returning true would result in stopping the request execution
        return false
    }