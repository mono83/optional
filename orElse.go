package optional

import "time"

// OrElse returns value from optional or provided value if optional is empty
func (b Bool) OrElse(els bool) bool {
	if b.IsPresent() {
		return b.value
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (s String) OrElse(els string) string {
	if s.IsPresent() {
		return s.value
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (i Int) OrElse(els int) int {
	if i.IsPresent() {
		return i.value
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (f Float64) OrElse(els float64) float64 {
	if f.IsPresent() {
		return f.value
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (t Time) OrElse(els time.Time) time.Time {
	if t.IsPresent() {
		return t.value
	}
	return els
}

// OrElse returns value from optional or provided value if optional is empty
func (d Duration) OrElse(els time.Duration) time.Duration {
	if d.IsPresent() {
		return d.value
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
