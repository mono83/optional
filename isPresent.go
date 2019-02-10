package optional

import "reflect"

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

// IsPresent returns true if optional contains value and false for null
func (t Time) IsPresent() bool {
	return t.time != nil
}

// IsPresent returns true if optional contains value and false for null
func (d Duration) IsPresent() bool {
	return d.duration != nil
}

// IsPresent returns true if optional contains value and false for null
func (m Mixed) IsPresent() bool {
	if m.mixed == nil {
		return false
	}
	vo := reflect.ValueOf(m.mixed)
	return !(vo.Kind() == reflect.Ptr && vo.IsNil())
}
