// Code generated by entc, DO NOT EDIT.

package paymentchannel

const (
	// Label holds the string label denoting the paymentchannel type in the database.
	Label = "payment_channel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the paymentchannel in the database.
	Table = "payment_channels"
)

// Columns holds all SQL columns for paymentchannel fields.
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
