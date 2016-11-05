package servito

import (
	"fmt"
	"log"
	"net/http"
	"time"
    "github.com/gorilla/csrf"
    "github.com/ghawk1ns/servito/debugHandlers"
)

// Starts a servito with the default config
func StartServer() {
    if (config.Debug) {
        log.Println("Adding Debug Handlers")
        AddRoute("Index", "GET", "/", handlers.Index)
        AddRoute("Health", "GET", "/health", handlers.Health)
    }
    var handler http.Handler
    if (config.CSRFEnable) {
        CSRF := csrf.Protect([]byte(config.CSRFKey))
        handler = CSRF(newRouter())
        if (config.Debug) {
            log.Println("CSRF is enabled")
        }
    } else {
        if (config.Debug) {
            log.Println("CSRF is disabled")
        }
        handler = newRouter()
    }
    addr := fmt.Sprintf("%v:%v", config.Address, config.Port)
    srv := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
	}
    if (config.Debug) {
        log.Printf("listening at %v debug=%v\n", addr, config.Debug)
    }
	log.Fatal(srv.ListenAndServe())
}