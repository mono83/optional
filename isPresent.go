package optional

import "reflect"

// IsPresent analyzes provided value for emptiness
// It will return false only if provided nil or empty Optional
func IsPresent(mix interface{}) bool {
	if mix == nil {
		return false
	}

	if v, ok := mix.(Optional); ok {
		return v.IsPresent()
	}

	vo := reflect.ValueOf(mix)
	return !(vo.Kind() == reflect.Ptr && vo.IsNil())
}

// IsPresent returns true if optional contains value and false for null
func (b Bool) IsPresent() bool {
	return b.presents
}

// IsPresent returns true if optional contains value and false for null
func (s String) IsPresent() bool {
	return s.presents
}

// IsPresent returns true if optional contains value and false for null
func (i Int) IsPresent() bool {
	return i.presents
}

// IsPresent returns true if optional contains value and false for null
func (f Float64) IsPresent() bool {
	return f.presents
}

// IsPresent returns true if optional contains value and false for null
func (t Time) IsPresent() bool {
	return t.presents
}

// IsPresent returns true if optional contains value and false for null
func (d Duration) IsPresent() bool {
	return d.presents
}

// IsPresent returns true if optional contains value and false for null
func (m Mixed) IsPresent() bool {
	if m.mixed == nil {
		return false
	}
	vo := reflect.ValueOf(m.mixed)
	return !(vo.Kind() == reflect.Ptr && vo.IsNil())
}
