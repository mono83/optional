package optional

import "time"

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (b Bool) Get() bool {
	return *b.bool
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (i Int) Get() int {
	return *i.int
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (s String) Get() string {
	return *s.string
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (f Float64) Get() float64 {
	return *f.float64
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (t Time) Get() time.Time {
	return *t.time
}

// Get returns value from optional.
// Invocation on empty optional will cause panic: invalid memory address or nil pointer dereference
// Perform .IsPresent check or use .OrElse to avoid manic
func (d Duration) Get() time.Duration {
	return *d.duration
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
