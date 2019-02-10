package optional

import "time"

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (b Bool) IfPresent(f func(bool)) {
	if f != nil && b.IsPresent() {
		f(*b.bool)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (s String) IfPresent(f func(string)) {
	if f != nil && s.IsPresent() {
		f(*s.string)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (i Int) IfPresent(f func(int)) {
	if f != nil && i.IsPresent() {
		f(*i.int)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (f Float64) IfPresent(c func(float64)) {
	if c != nil && f.IsPresent() {
		c(*f.float64)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (t Time) IfPresent(f func(time.Time)) {
	if f != nil && t.IsPresent() {
		f(*t.time)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (d Duration) IfPresent(f func(time.Duration)) {
	if f != nil && d.IsPresent() {
		f(*d.duration)
	}
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (m Mixed) IfPresent(f func(interface{})) {
	if f != nil && m.IsPresent() {
		f(m.mixed)
	}
}
