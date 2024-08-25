package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PushCarEvent(w http.ResponseWriter, r *http.Request) {
	// Connect to the MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Get the collection
	db := client.Database("gtr-pi")
	collection := db.Collection("gtr-pi-events")

	// Extract the query parameter
	event := r.URL.Query().Get("buttonId")
	if event == "" {
		http.Error(w, "Missing 'buttonId' query parameter", http.StatusBadRequest)
		return
	}

	// Create a document with the event and current timestamp
	doc := map[string]interface{}{
		"event":     event,
		"timestamp": time.Now(),
	}

	// Insert the document into the collection
	_, err = collection.InsertOne(context.Background(), doc)
	if err != nil {
		http.Error(w, "Error inserting document into the database", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
