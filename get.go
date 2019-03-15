package optional

import "time"

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (b Bool) Get() bool {
	if !b.IsPresent() {
		panic("empty optional.Bool value")
	}
	return b.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (i Int) Get() int {
	if !i.IsPresent() {
		panic("empty optional.Int value")
	}
	return i.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (s String) Get() string {
	if !s.IsPresent() {
		panic("empty optional.String value")
	}
	return s.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (f Float64) Get() float64 {
	if !f.IsPresent() {
		panic("empty optional.Float64 value")
	}
	return f.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (t Time) Get() time.Time {
	if !t.IsPresent() {
		panic("empty optional.Time value")
	}
	return t.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (d Duration) Get() time.Duration {
	if !d.IsPresent() {
		panic("empty optional.Duration value")
	}
	return d.value
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: nil in mixed value
// Perform .IsPresent check or use .OrElse to avoid manic
func (m Mixed) Get() interface{} {
	if !m.IsPresent() {
		panic("nil in mixed value")
	}

	return m.mixed
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: nil in optional.Error
// Perform .IsPresent check or use .OrElse to avoid manic
func (e Error) Get() error {
	if !e.IsPresent() {
		panic("nil in optional.Error")
	}

	return e.value
}
