package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Thanks for using Servito! Have an idea or spotted a bug? Submit a PR! https://github.com/GHawk1ns/Servito")
}
