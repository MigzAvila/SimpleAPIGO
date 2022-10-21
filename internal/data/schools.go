// Filename : internal/data/schools.go

package data

import (
	"database/sql"
	"time"

	"appletree.miguelavila.net/internal/validator"
	"github.com/lib/pq"
)

type School struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	Version   int32     `json:"version"`
}

func ValidateSchool(v *validator.Validator, school *School) {

	v.Check(school.Name != "", "name", "must be provided")
	v.Check(len(school.Name) <= 200, "name", "must no more 200 characters")

	v.Check(school.Level != "", "level", "must be provided")
	v.Check(len(school.Level) <= 200, "level", "must no more 200 characters")

	v.Check(school.Contact != "", "contact", "must be provided")
	v.Check(len(school.Contact) <= 200, "contact", "must no more 200 characters")

	v.Check(school.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(school.Phone, validator.PhoneRX), "phone", "must be a valid phone number")

	v.Check(school.Email != "", "email", "must be provided")
	v.Check(validator.Matches(school.Email, validator.EmailRX), "email", "must be a valid email")

	v.Check(school.Website != "", "website", "must be provided")
	v.Check(validator.ValidWebsite(school.Website), "website", "must be a valid website")

	v.Check(school.Address != "", "address", "must be provided")
	v.Check(len(school.Address) <= 500, "address", "must no more 500 characters")

	v.Check(school.Mode != nil, "mode", "must be provided")
	v.Check(len(school.Mode) >= 1, "mode", "must contain at least one mode")
	v.Check(len(school.Mode) <= 5, "mode", "must contain at most five mode")
	v.Check(validator.Unique(school.Mode), "mode", "must not contain duplicates")

}

// define a SchoolModel object that wraps a sql.DB connection pool
type SchoolModel struct {
	DB *sql.DB
}

// insert() allows us to create a new School
func (m SchoolModel) Insert(school *School) error {
	query := `
		INSERT INTO schools (name, level, contact, phone, email, website, address, mode)	
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, create_at, version
	`
	// collect data fields into a slice
	args := []interface{}{
		school.Name,
		school.Level,
		school.Contact,
		school.Phone,
		school.Email,
		school.Website,
		school.Address,
		pq.Array(school.Mode),
	}
	// run query ... -> expand the slice
	return m.DB.QueryRow(query, args...).Scan(&school.ID, &school.CreatedAt, &school.Version)
}

// Get() allows us to retrieve a specific School
func (m SchoolModel) Get(id int64) (*School, error) {
	return nil, nil
}

// Update() allows us to update a specific School
func (m SchoolModel) Update(school *School) error {
	return nil
}

// Delete() allows us to delete a specific School
func (m SchoolModel) Delete(id int64) error {
	return nil
}
