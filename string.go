package optional

import "strings"

var emptyString = String{}

// String contains optional string value
type String struct {
	*string
}

// OfString creates new optional string containing provided value
func OfString(s string) String {
	return String{string: &s}
}

// Trim applies whitespace trim mapping and returns optional with new value
func (s String) Trim() String {
	if s.string == nil {
		return s
	}

	return OfString(strings.TrimSpace(*s.string))
}

// FilterEmpty applies empty (but not whitespace) filtering and returns optional
func (s String) FilterEmpty() String {
	if s.string == nil || len(*s.string) == 0 {
		return emptyString
	}

	return s
}
