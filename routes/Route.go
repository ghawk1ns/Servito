package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"github.com/ghawk1ns/servito/stats"
)

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := handlerInterceptor(&route)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}

func handlerInterceptor(route *Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		start := time.Now()
		route.HandlerFunc.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			route.Name,
			time.Since(start),
			r.UserAgent(),
		)
	})
}