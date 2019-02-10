package optional

import "time"

// OrElse returns value from optional or provided value if optional is empty
func (b Bool) OrElse(els bool) bool {
	if b.IsPresent() {
		return *b.bool
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (s String) OrElse(els string) string {
	if s.IsPresent() {
		return *s.string
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (i Int) OrElse(els int) int {
	if i.IsPresent() {
		return *i.int
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (f Float64) OrElse(els float64) float64 {
	if f.IsPresent() {
		return *f.float64
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (t Time) OrElse(els time.Time) time.Time {
	if t.IsPresent() {
		return *t.time
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (d Duration) OrElse(els time.Duration) time.Duration {
	if d.IsPresent() {
		return *d.duration
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (m Mixed) OrElse(els interface{}) interface{} {
	if m.IsPresent() {
		return m.mixed
	}
	return els
}
