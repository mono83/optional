package optional

import "time"

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (b Bool) IfPresent(f func(bool)) Bool {
	if f != nil && b.IsPresent() {
		f(b.value)
	}

	return b
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (s String) IfPresent(f func(string)) String {
	if f != nil && s.IsPresent() {
		f(s.value)
	}

	return s
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (i Int) IfPresent(f func(int)) Int {
	if f != nil && i.IsPresent() {
		f(i.value)
	}

	return i
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (f Float64) IfPresent(c func(float64)) Float64 {
	if c != nil && f.IsPresent() {
		c(f.value)
	}

	return f
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (t Time) IfPresent(f func(time.Time)) Time {
	if f != nil && t.IsPresent() {
		f(t.value)
	}

	return t
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (d Duration) IfPresent(f func(time.Duration)) Duration {
	if f != nil && d.IsPresent() {
		f(d.value)
	}

	return d
}

// IfPresent invokes provided callback if optional contains some value
// If optional is empty, callback will not be invoked
func (m Mixed) IfPresent(f func(interface{})) Mixed {
	if f != nil && m.IsPresent() {
		f(m.mixed)
	}

	return m
}
