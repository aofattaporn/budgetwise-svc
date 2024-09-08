package utils

import "encoding/json"

func NilToNULLString(value *string) string {
	if value == nil {
		return "NULL"
	}
	return *value
}

// Return empty string when value is either "nil" or "null"
//
// Otherwise, return string represent json.RawMessage
func SafelyGetString(value json.RawMessage) string {
	if value != nil {
		str := string(value)
		if str != "null" {
			return str
		}
	}
	return ""
}
