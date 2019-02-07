package optional

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
