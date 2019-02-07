package optional

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (b Bool) Filter(f func(bool) bool) Bool {
	if f == nil || b.bool == nil {
		return emptyBool
	}
	if !f(*b.bool) {
		return emptyBool
	}
	return b
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (s String) Filter(f func(string) bool) String {
	if f == nil || s.string == nil {
		return emptyString
	}
	if !f(*s.string) {
		return emptyString
	}
	return s
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (i Int) Filter(f func(int) bool) Int {
	if f == nil || i.int == nil {
		return emptyInt
	}
	if !f(*i.int) {
		return emptyInt
	}
	return i
}

// Filter method applies predicate on optional content
// Resulting optional will be non empty only if predicate returns true and original value inside
// optional was not empty.
func (f Float64) Filter(c func(float64) bool) Float64 {
	if c == nil || f.float64 == nil {
		return emptyFloat64
	}
	if !c(*f.float64) {
		return emptyFloat64
	}
	return f
}
