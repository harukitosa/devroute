// Code generated by entc, DO NOT EDIT.

package companyuser

import (
	"fmt"
)

const (
	// Label holds the string label denoting the companyuser type in the database.
	Label = "company_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastNameFurigana holds the string denoting the last_name_furigana field in the database.
	FieldLastNameFurigana = "last_name_furigana"
	// FieldFirstNameFurigana holds the string denoting the first_name_furigana field in the database.
	FieldFirstNameFurigana = "first_name_furigana"
	// FieldProfileName holds the string denoting the profile_name field in the database.
	FieldProfileName = "profile_name"
	// FieldIconURL holds the string denoting the icon_url field in the database.
	FieldIconURL = "icon_url"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// Table holds the table name of the companyuser in the database.
	Table = "company_users"
)

// Columns holds all SQL columns for companyuser fields.
var Columns = []string{
	FieldID,
	FieldLastName,
	FieldFirstName,
	FieldLastNameFurigana,
	FieldFirstNameFurigana,
	FieldProfileName,
	FieldIconURL,
	FieldGender,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameFuriganaValidator is a validator for the "last_name_furigana" field. It is called by the builders before save.
	LastNameFuriganaValidator func(string) error
	// FirstNameFuriganaValidator is a validator for the "first_name_furigana" field. It is called by the builders before save.
	FirstNameFuriganaValidator func(string) error
	// ProfileNameValidator is a validator for the "profile_name" field. It is called by the builders before save.
	ProfileNameValidator func(string) error
	// IconURLValidator is a validator for the "icon_url" field. It is called by the builders before save.
	IconURLValidator func(string) error
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMale, GenderFemale, GenderOther:
		return nil
	default:
		return fmt.Errorf("companyuser: invalid enum value for gender field: %q", ge)
	}
}