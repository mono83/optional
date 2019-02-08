package optional

import "time"

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (b Bool) IfPresent(f func(bool)) {
	if f != nil && b.bool != nil {
		f(*b.bool)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (s String) IfPresent(f func(string)) {
	if f != nil && s.string != nil {
		f(*s.string)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (i Int) IfPresent(f func(int)) {
	if f != nil && i.int != nil {
		f(*i.int)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (f Float64) IfPresent(c func(float64)) {
	if c != nil && f.float64 != nil {
		c(*f.float64)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (t Time) IfPresent(f func(time.Time)) {
	if f != nil && t.time != nil {
		f(*t.time)
	}
}
