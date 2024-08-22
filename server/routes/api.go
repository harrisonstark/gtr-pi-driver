package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PushCarEvent(w http.ResponseWriter, r *http.Request) {
	return
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	db := client.Database("myDatabase")
	collection := db.Collection("myCollection")

	cursor, err := collection.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var results []map[string]interface{}
	if err := cursor.All(context.Background(), &results); err != nil {
		http.Error(w, "Error decoding the database results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
