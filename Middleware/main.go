package main

import (
	"fmt"
	"net/http"
	"time"
)

// Section 1 - Handles the uptime request
type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Current uptime: %s \n", time.Since(h.Started)))
}

// Section 2 - Handles (middleware) the secret token request
// if token is correct, it will go to the uptime section
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("secret") == h.secret {
		h.next.ServeHTTP(w, req)
	} else {
		// Returns 404 not found
		http.NotFound(w, req)
	}
}

func main() {
	// If it passes the middleware - it will access the 'next' part
	// which is to request the uptime
	http.Handle("/", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "SuperSecretToken",
	})
	// curl -i 127.0.0.1:3000/ - This will return 404
	// curl -i 127.0.0.1:3000/?secret=SuperSecretToken - This will return the uptime!
	http.ListenAndServe(":3000", nil)
}
