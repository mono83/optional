package optional

// Map applies mapping function on optional value if it presents
func (b Bool) Map(f func(bool) bool) Bool {
	if f == nil || b.bool == nil {
		return emptyBool
	}

	return OfBool(f(*b.bool))
}

// MapToInt applies mapping function on optional value if it presents
func (b Bool) MapToInt(f func(bool) int) Int {
	if f == nil || b.bool == nil {
		return emptyInt
	}

	return OfInt(f(*b.bool))
}

// MapToString applies mapping function on optional value if it presents
func (b Bool) MapToString(f func(bool) string) String {
	if f == nil || b.bool == nil {
		return emptyString
	}

	return OfString(f(*b.bool))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (b Bool) MapToFloat64(f func(bool) float64) Float64 {
	if f == nil || b.bool == nil {
		return emptyFloat64
	}

	return OfFloat64(f(*b.bool))
}

// Map applies mapping function on optional value if it presents
func (i Int) Map(f func(int) int) Int {
	if f == nil || i.int == nil {
		return emptyInt
	}

	return OfInt(f(*i.int))
}

// MapToBool applies mapping function on optional value if it presents
func (i Int) MapToBool(f func(int) bool) Bool {
	if f == nil || i.int == nil {
		return emptyBool
	}

	return OfBool(f(*i.int))
}

// MapToString applies mapping function on optional value if it presents
func (i Int) MapToString(f func(int) string) String {
	if f == nil || i.int == nil {
		return emptyString
	}

	return OfString(f(*i.int))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (i Int) MapToFloat64(f func(int) float64) Float64 {
	if f == nil || i.int == nil {
		return emptyFloat64
	}

	return OfFloat64(f(*i.int))
}

// Map applies mapping function on optional value if it presents
func (s String) Map(f func(string) string) String {
	if f == nil || s.string == nil {
		return emptyString
	}

	return OfString(f(*s.string))
}

// MapToBool applies mapping function on optional value if it presents
func (s String) MapToBool(f func(string) bool) Bool {
	if f == nil || s.string == nil {
		return emptyBool
	}

	return OfBool(f(*s.string))
}

// MapToInt applies mapping function on optional value if it presents
func (s String) MapToInt(f func(string) int) Int {
	if f == nil || s.string == nil {
		return emptyInt
	}

	return OfInt(f(*s.string))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (s String) MapToFloat64(f func(string) float64) Float64 {
	if f == nil || s.string == nil {
		return emptyFloat64
	}

	return OfFloat64(f(*s.string))
}

// Map applies mapping function on optional value if it presents
func (f Float64) Map(c func(float64) float64) Float64 {
	if c == nil || f.float64 == nil {
		return emptyFloat64
	}

	return OfFloat64(c(*f.float64))
}

// MapToBool applies mapping function on optional value if it presents
func (f Float64) MapToBool(c func(float64) bool) Bool {
	if c == nil || f.float64 == nil {
		return emptyBool
	}

	return OfBool(c(*f.float64))
}

// MapToInt applies mapping function on optional value if it presents
func (f Float64) MapToInt(c func(float64) int) Int {
	if c == nil || f.float64 == nil {
		return emptyInt
	}

	return OfInt(c(*f.float64))
}

// MapToString applies mapping function on optional value if it presents
func (f Float64) MapToString(c func(float64) string) String {
	if c == nil || f.float64 == nil {
		return emptyString
	}

	return OfString(c(*f.float64))
}
