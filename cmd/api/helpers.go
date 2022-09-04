// Filename: cmd/api/helpers

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Utility function for reading ID in Endpoint
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the param
	// Use the ParamsFormContext
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid :id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// convert map to JSON object
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Add a newline to make viewing on terminal easier
	jsonData = append(jsonData, '\n')

	// Add the headers
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Specify that we will serve our response in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// Write the []byte slice containing the JSON response body
	w.Write(jsonData)
	return nil
}
