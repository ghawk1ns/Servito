package handlers

import (
    "net/http"
    "io"
    "encoding/json"
    "time"
    "sync/atomic"
    "strconv"
)

func Health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    healthMap := make(map[string]string)
    healthMap["alive"] = "true"
    healthMap["uptime"] = uptime()
    healthMap["requests"] = getRequests()
    v,_ := json.Marshal(healthMap)
    io.WriteString(w, string(v))
}

var startTime time.Time

func init() {
    startTime = time.Now()
}

func uptime() string {
    return time.Since(startTime).String()
}

var requests int64

func IncRequests() {
    atomic.AddInt64(&requests, 1)
}

func getRequests() string {
    return strconv.FormatInt(requests, 10)
}