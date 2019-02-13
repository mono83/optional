package optional

import "strings"

// String contains optional string value
type String struct {
	value    string
	presents bool
}

func (s *String) set(value string) {
	s.value = value
	s.presents = true
}

// OfString creates new optional string containing provided value
func OfString(s string) String {
	return String{value: s, presents: true}
}

// Trim applies whitespace trim mapping and returns optional with new value
func (s String) Trim() String {
	if !s.IsPresent() {
		return s
	}

	return OfString(strings.TrimSpace(s.value))
}

// FilterEmpty applies empty (but not whitespace) filtering and returns optional
func (s String) FilterEmpty() String {
	if !s.IsPresent() || len(s.value) == 0 {
		return String{}
	}

	return s
}
