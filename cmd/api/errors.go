// Filename : cmd/api/errors.go

package main

import (
	"fmt"
	"net/http"
)

// Log errors
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Send JSON-formatted error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// create the json response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// Server error message
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log the error
	app.logError(r, err)
	//prepare a message with error
	message := "the server encountered an problem and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Method not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//prepare a message with error
	message := "the requested resources could not be found."
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Method not Allowed response
func (app *application) MethodNotAllowedReponse(w http.ResponseWriter, r *http.Request) {
	//prepare a message with error
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}
