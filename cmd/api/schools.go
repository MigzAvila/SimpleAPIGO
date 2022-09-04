// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
)

//createSchoolHandler for POST /v1/schools endpoints

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new School...")

}

// createSchoolHandler for GET /v1/schools endpoints
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//Utilize Utility Methods/Functions from helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Displaying out the school ID
	fmt.Fprintf(w, "show the details for the School %d\n", id)

}
