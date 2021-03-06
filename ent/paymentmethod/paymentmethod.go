// Code generated by entc, DO NOT EDIT.

package paymentmethod

const (
	// Label holds the string label denoting the paymentmethod type in the database.
	Label = "payment_method"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the paymentmethod in the database.
	Table = "payment_methods"
)

// Columns holds all SQL columns for paymentmethod fields.
var Columns = []string{
	FieldID,
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
