package optional

import "time"

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (b Bool) Filter(f func(bool) bool) Bool {
	if f == nil || !b.IsPresent() || !f(b.value) {
		return Bool{}
	}
	return b
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (s String) Filter(f func(string) bool) String {
	if f == nil || !s.IsPresent() || !f(s.value) {
		return String{}
	}
	return s
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (i Int) Filter(f func(int) bool) Int {
	if f == nil || !i.IsPresent() || !f(i.value) {
		return Int{}
	}
	return i
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (f Float64) Filter(c func(float64) bool) Float64 {
	if c == nil || !f.IsPresent() || !c(f.value) {
		return Float64{}
	}
	return f
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (t Time) Filter(f func(time.Time) bool) Time {
	if f == nil || !t.IsPresent() || !f(t.value) {
		return Time{}
	}

	return t
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (d Duration) Filter(f func(time.Duration) bool) Duration {
	if f == nil || !d.IsPresent() || !f(d.value) {
		return Duration{}
	}

	return d
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (m Mixed) Filter(f func(interface{}) bool) Mixed {
	if f == nil || !m.IsPresent() || !f(m.mixed) {
		return Mixed{}
	}

	return m
}
