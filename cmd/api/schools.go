// Filename: cmd/api/schools

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//createSchoolHandler for POST /v1/schools endpoints

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new School...")

}

// createSchoolHandler for GET /v1/schools endpoints
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Use the ParamsFormContext
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details for the School %d\n", id)

}
