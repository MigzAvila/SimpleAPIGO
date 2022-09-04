// Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// create a map to hold the healthcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	// convert map to JSON object
	jsonData, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered an error and could not process the request", http.StatusInternalServerError)
		return
	}

	// Add a newline to make viewing on terminal easier
	jsonData = append(jsonData, '\n')
	// Specify that we will serve our response in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Write the []byte slice containing the JSON response body
	w.Write([]byte(jsonData))
}
