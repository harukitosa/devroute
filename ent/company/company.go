// Code generated by entc, DO NOT EDIT.

package company

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldCompanyID holds the string denoting the company_id field in the database.
	FieldCompanyID = "company_id"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// Table holds the table name of the company in the database.
	Table = "companies"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLocation,
	FieldPostalCode,
	FieldPhoneNumber,
	FieldCompanyID,
	FieldIntroduction,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// CompanyIDValidator is a validator for the "company_id" field. It is called by the builders before save.
	CompanyIDValidator func(string) error
	// IntroductionValidator is a validator for the "introduction" field. It is called by the builders before save.
	IntroductionValidator func(string) error
)
