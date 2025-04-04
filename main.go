package main

import (
	"fmt"
	"net/http"
	"stockchart/handlers"
	"time"

	"github.com/go-chi/chi/v5"
)

// main function initializes the router and starts the server
// It sets up the routes for the home page and WebSocket connection
// and starts listening on port 3000
func main() {

	// Setting up refresh rate once here
	handlers.SetTickerInterval(3 * time.Second)

	r := chi.NewRouter()
	r.Get("/", handlers.HomeHandler())
	r.Get("/ws", handlers.HandleWebSocket)

	fmt.Println("Server running on :3000")
	http.ListenAndServe(":3000", r)
}
