// Filename: cmd/api/schools

package main

import (
	"fmt"
	"net/http"
	"time"

	"appletree.miguelavila.net/internal/data"
	"appletree.miguelavila.net/internal/validator"
)

//createSchoolHandler for POST /v1/schools endpoints

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Target decode destination
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badResquestReponse(w, r, err)
		return
	}
	// Copy the values from the input struct to a new School struct
	school := &data.School{
		Name:    input.Name,
		Level:   input.Level,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode:    input.Mode,
	}

	// Initialize a new instance of validator
	v := validator.New()

	// Check the errors maps if there were any errors validation
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// create a school
	err = app.models.Schools.Insert(school)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// create a Location header for the newly created resource/school
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/schools/%d", school.ID))
	// write the json response with 201 - created status code with the body
	// being the school data and the headers being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"school": school}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

// createSchoolHandler for GET /v1/schools endpoints
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	//Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the School struct containing the ID we extracted from
	// From URL and sample data
	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Apple Tree",
		Level:     "High School",
		Contact:   "Anna Smith",
		Phone:     "601-4411",
		Address:   "14 Apple Street",
		Mode:      []string{"Blended", "Online"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
