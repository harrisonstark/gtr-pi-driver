package main

import (
	"log"
	"net/http"

	"gtr/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/push-car-event", routes.PushCarEvent).Methods("POST")

	// Serve static files
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../client"))))

	handler := cors.Default().Handler(r)

	// Start server
	port := "7070"
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
