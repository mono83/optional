package optional

// IsPresent returns true if optional contains value and false for null
func (b Bool) IsPresent() bool {
	return b.bool != nil
}

// IsPresent returns true if optional contains value and false for null
func (s String) IsPresent() bool {
	return s.string != nil
}

// IsPresent returns true if optional contains value and false for null
func (i Int) IsPresent() bool {
	return i.int != nil
}

// IsPresent returns true if optional contains value and false for null
func (f Float64) IsPresent() bool {
	return f.float64 != nil
}
