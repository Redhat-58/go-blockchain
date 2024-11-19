package controllers

import (
	"fmt"
	"net/http"
)

// HelloHandler handles requests to the /hello route.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go HTTP server!")
}
