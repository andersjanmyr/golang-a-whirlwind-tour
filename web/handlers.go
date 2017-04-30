package web

import (
	"io"
	"net/http"
)

// e.g. http.HandleFunc("/health-check", HealthCheckHandler)
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

// Implement SumHandler that sums two parameters a and b
// It should return an error if any of the parameters isn't a number
// Write the test first!
func SumHandler(w http.ResponseWriter, r *http.Request) {
}
