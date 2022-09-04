// Filename: cmd/api/healthcheck.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Formatting in json format
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Specify that we will serve our response in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body
	w.Write([]byte(js))
}
