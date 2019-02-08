package optional

import "time"

// OrElse returns value from optional or provided value if optional is empty
func (b Bool) OrElse(els bool) bool {
	if b.bool == nil {
		return els
	}

	return *b.bool
}

// OrElse returns value from optional or provided value if optional is empty
func (s String) OrElse(els string) string {
	if s.string == nil {
		return els
	}

	return *s.string
}

// OrElse returns value from optional or provided value if optional is empty
func (i Int) OrElse(els int) int {
	if i.int == nil {
		return els
	}

	return *i.int
}

// OrElse returns value from optional or provided value if optional is empty
func (f Float64) OrElse(els float64) float64 {
	if f.float64 == nil {
		return els
	}

	return *f.float64
}

// OrElse returns value from optional or provided value if optional is empty
func (t Time) OrElse(els time.Time) time.Time {
	if t.time == nil {
		return els
	}

	return *t.time
}
