package servito

import (
	"net/http"
    "github.com/gorilla/mux"
    "time"
    "log"
    "github.com/ghawk1ns/servito/debugHandlers"
)

// Internal list of routes
var routes []route

// Defines a Route for a request to take
type route struct {
	Name        string // Name of the route
	Method      string // HTTP Method
	Pattern     string // Path Pattern
	HandlerFunc http.HandlerFunc // Function to handle this request
}

func AddRoute(name string, method string, pattern string, handlerFunction http.HandlerFunc) {
    addRoute(&route{
        name,
        method,
        pattern,
        handlerFunction,
    })
}

// Adds a new route
func addRoute(r *route) {
    if routes == nil  {
        routes = make([]route, 0)
    }
    if (config.Debug) {
        log.Printf("Added route: %v\n", *r)
    }
    routes = append(routes, *r)
}

// Creates a new Router
func newRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        // TODO: Let users add their own interceptor
        handler := internalInterceptor(route)
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}

// Adds some additional logs and stats if debug=true
func internalInterceptor(route route) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if (config.Debug) {
            handlers.IncRequests()
            start := time.Now()
            defer log.Printf(
                "%s\t%s\t%s\t%s\t%s",
                r.Method,
                r.RequestURI,
                route.Name,
                time.Since(start),
                r.UserAgent(),
            )
        }
        // Give the user the option to intercept the request and handle it manually
        handled := userInterceptor !=  nil && userInterceptor(w, r)
        if !handled {
            route.HandlerFunc.ServeHTTP(w, r)
        }
    }
}

// Intercepts all incoming requests and passes control over the user. If the request was handled,
// RequestInterceptor should return true. This means the request will not be routed to the original target
type RequestInterceptor func(http.ResponseWriter, *http.Request) bool

var userInterceptor RequestInterceptor

func AddInterceptor(fn RequestInterceptor) {
    userInterceptor = fn
}