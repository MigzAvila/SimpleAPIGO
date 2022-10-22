// Filename: cmd/api/schools

package main

import (
	"errors"
	"fmt"
	"net/http"

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

	// Fetch the specific school
	school, err := app.models.Schools.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// write the data return by the Get method
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// This method does a partial replacement
	// get the id of the school and update the school
	// Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// fetch the original record from database
	school, err := app.models.Schools.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// create an input struct to hold the data read in from the client
	// Update input struct to use pointers because pointers have a default value of nil
	// if field remains nil then we know that the client is not interested in updating the field
	var input struct {
		Name    *string  `json:"name"`
		Level   *string  `json:"level"`
		Contact *string  `json:"contact"`
		Phone   *string  `json:"phone"`
		Email   *string  `json:"email"`
		Website *string  `json:"website"`
		Address *string  `json:"address"`
		Mode    []string `json:"mode"`
	}
	// Decode the data from the client
	err = app.readJSON(w, r, &input)

	// copy / update the fields / values in the school variable using the fields in the input struct
	if err != nil {
		app.badResquestReponse(w, r, err)
		return
	}

	if input.Name != nil {
		school.Name = *input.Name
	}

	if input.Level != nil {
		school.Level = *input.Level
	}

	if input.Contact != nil {
		school.Contact = *input.Contact
	}

	if input.Phone != nil {
		school.Phone = *input.Phone
	}

	if input.Email != nil {
		school.Email = *input.Email
	}

	if input.Website != nil {
		school.Website = *input.Website
	}

	if input.Address != nil {
		school.Address = *input.Address
	}

	if input.Mode != nil {
		school.Mode = input.Mode
	}

	// validate the data provided by the client, if the validation fails,
	// then we send a 422 - Unprocessable responses to the client
	// Initialize a new validation error instance
	v := validator.New()

	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Pass the updated school record to the update method
	err = app.models.Schools.Update(school)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// write the json response by Update
	err = app.writeJSON(w, http.StatusCreated, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) deleteSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// This method does a delete of a specific school
	// get the id of the school and update the school
	//Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// delete the school from the database. send a 404 notFoundResponse status code to the client if there is no matching record

	// fetch the original record from database
	err = app.models.Schools.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	//  return 200 status ok the client with a successful message
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "school successfully deleted"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}
