package handlers

import (
    "net/http"
    "io"
    "github.com/ghawk1ns/servito/stats"
    "encoding/json"
)

// e.g. http.HandleFunc("/health", Health)
func Health(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    healthMap := make(map[string]string)
    healthMap["alive"] = "true"
    healthMap["uptime"] = stats.Uptime()
    healthMap["requests"] = stats.GetRequests()
    v,_ := json.Marshal(healthMap)
    io.WriteString(w, string(v))
}
