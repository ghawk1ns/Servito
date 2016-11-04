package servito

import (
	"fmt"
	"github.com/ghawk1ns/servito/routes"
	"log"
	"net/http"
	"time"
    "github.com/gorilla/csrf"
)

func StartServer() {
	addr := fmt.Sprintf("%v:%v", config.Address, config.Port)
    var handler http.Handler
    if (config.CSRFEnable) {
        CSRF := csrf.Protect([]byte(config.CSRFKey))
        handler = CSRF(routes.NewRouter())
    } else {
        handler = routes.NewRouter()
    }
    csrf.Secure(!config.Debug)
    srv := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
	}
    log.Printf("Servito listening at %v debug=%v\n", addr, config.Debug)
	log.Fatal(srv.ListenAndServe())
}