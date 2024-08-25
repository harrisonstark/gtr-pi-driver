package routes

import (
	"fmt"
	"io"
	"net/http"
)

func PushCarEvent(w http.ResponseWriter, r *http.Request) {
	// Extract buttonId from query parameters
	buttonId := r.URL.Query().Get("buttonId")

	if buttonId == "" {
		http.Error(w, "Missing buttonId", http.StatusBadRequest)
		return
	}

	// Prepare the POST request to localhost:7171/event
	targetURL := fmt.Sprintf("http://localhost:7171/push_event?event=%s", buttonId)

	req, err := http.NewRequest("POST", targetURL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the POST request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response from the target server back to the original client
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
