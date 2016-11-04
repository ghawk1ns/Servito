package stats

import (
    "time"
    "sync/atomic"
    "strconv"
)

var startTime time.Time

func init() {
    startTime = time.Now()
}

func Uptime() string {
    return time.Since(startTime).String()
}

var requests int64

func IncRequests() {
    atomic.AddInt64(&requests, 1)
}

func GetRequests() string {
    return strconv.FormatInt(requests, 10)
}


