package middlewares

import (
	"net/http"
	"os"
)

// CORS middleware to handle CORS requests
func CorsMiddleware(next http.Handler) http.Handler {

	// ANONYMOUS FUNCTION TO RETURN HTTP HANDLER
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOWED-ORIGIN")) // Allow any origin
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// w.Header().Set("Access-Control-Allow-Headers", os.Getenv("HEADERS")) // Add any other headers you need

		// If this is a preflight request, the method will be OPTIONS,
		// so just send an OK status and return
		// if r.Method == "OPTIONS" {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		// Handle the request
		next.ServeHTTP(w, r)
	})
}
