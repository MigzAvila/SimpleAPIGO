// Filename: cmd/api/helpers

package main

import (
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
