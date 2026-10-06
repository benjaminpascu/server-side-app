package handler

import (
	"fmt"
	"net/http"
)

// Handler is the Vercel serverless function for the '/nickname' URL.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alert-moose")
}
