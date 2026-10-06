package handler

import (
	"fmt"
	"net/http"
)

// Handler is the Vercel serverless function for the '/' URL.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow MGT656/660 FTW!@$#")
}
