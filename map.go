package optional

import "time"

/*
   Bool map functions
*/

// Map applies mapping function on optional value if it presents
func (b Bool) Map(f func(bool) bool) Bool {
	if f == nil || !b.IsPresent() {
		return emptyBool
	}

	return OfBool(f(*b.bool))
}

// MapToInt applies mapping function on optional value if it presents
func (b Bool) MapToInt(f func(bool) int) Int {
	if f == nil || !b.IsPresent() {
		return emptyInt
	}

	return OfInt(f(*b.bool))
}

// MapToString applies mapping function on optional value if it presents
func (b Bool) MapToString(f func(bool) string) String {
	if f == nil || !b.IsPresent() {
		return emptyString
	}

	return OfString(f(*b.bool))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (b Bool) MapToFloat64(f func(bool) float64) Float64 {
	if f == nil || !b.IsPresent() {
		return emptyFloat64
	}

	return OfFloat64(f(*b.bool))
}

/*
   Int map functions
*/

// Map applies mapping function on optional value if it presents
func (i Int) Map(f func(int) int) Int {
	if f == nil || !i.IsPresent() {
		return emptyInt
	}

	return OfInt(f(*i.int))
}

// MapToBool applies mapping function on optional value if it presents
func (i Int) MapToBool(f func(int) bool) Bool {
	if f == nil || !i.IsPresent() {
		return emptyBool
	}

	return OfBool(f(*i.int))
}

// MapToString applies mapping function on optional value if it presents
func (i Int) MapToString(f func(int) string) String {
	if f == nil || !i.IsPresent() {
		return emptyString
	}

	return OfString(f(*i.int))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (i Int) MapToFloat64(f func(int) float64) Float64 {
	if f == nil || !i.IsPresent() {
		return emptyFloat64
	}

	return OfFloat64(f(*i.int))
}

/*
   String map functions
*/

// Map applies mapping function on optional value if it presents
func (s String) Map(f func(string) string) String {
	if f == nil || !s.IsPresent() {
		return emptyString
	}

	return OfString(f(*s.string))
}

// MapToBool applies mapping function on optional value if it presents
func (s String) MapToBool(f func(string) bool) Bool {
	if f == nil || !s.IsPresent() {
		return emptyBool
	}

	return OfBool(f(*s.string))
}

// MapToInt applies mapping function on optional value if it presents
func (s String) MapToInt(f func(string) int) Int {
	if f == nil || !s.IsPresent() {
		return emptyInt
	}

	return OfInt(f(*s.string))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (s String) MapToFloat64(f func(string) float64) Float64 {
	if f == nil || !s.IsPresent() {
		return emptyFloat64
	}

	return OfFloat64(f(*s.string))
}

/*
   Float64 map functions
*/

// Map applies mapping function on optional value if it presents
func (f Float64) Map(c func(float64) float64) Float64 {
	if c == nil || !f.IsPresent() {
		return emptyFloat64
	}

	return OfFloat64(c(*f.float64))
}

// MapToBool applies mapping function on optional value if it presents
func (f Float64) MapToBool(c func(float64) bool) Bool {
	if c == nil || !f.IsPresent() {
		return emptyBool
	}

	return OfBool(c(*f.float64))
}

// MapToInt applies mapping function on optional value if it presents
func (f Float64) MapToInt(c func(float64) int) Int {
	if c == nil || !f.IsPresent() {
		return emptyInt
	}

	return OfInt(c(*f.float64))
}

// MapToString applies mapping function on optional value if it presents
func (f Float64) MapToString(c func(float64) string) String {
	if c == nil || !f.IsPresent() {
		return emptyString
	}

	return OfString(c(*f.float64))
}

/*
   Time map functions
*/

// Map applies mapping function on optional value if it presents
func (t Time) Map(f func(time.Time) time.Time) Time {
	if f == nil || !t.IsPresent() {
		return emptyTime
	}

	return OfTime(f(*t.time))
}

/*
   Int map functions
*/

// Map applies mapping function on optional value if it presents
func (d Duration) Map(f func(time.Duration) time.Duration) Duration {
	if f == nil || !d.IsPresent() {
		return emptyDuration
	}

	return OfDuration(f(*d.duration))
}

/*
   Mixed map functions
*/

// Map applies mapping function on optional value if it presents
func (m Mixed) Map(f func(interface{}) interface{}) Mixed {
	if f != nil && m.IsPresent() {
		return OfMixed(f(m.mixed))
	}

	return emptyMixed
}

// MapToBool applies mapping function on optional value if it presents
func (m Mixed) MapToBool(f func(interface{}) bool) Bool {
	if f != nil && m.IsPresent() {
		return OfBool(f(m.mixed))
	}

	return emptyBool
}

// MapToInt applies mapping function on optional value if it presents
func (m Mixed) MapToInt(f func(interface{}) int) Int {
	if f != nil && m.IsPresent() {
		return OfInt(f(m.mixed))
	}

	return emptyInt
}

// MapToString applies mapping function on optional value if it presents
func (m Mixed) MapToString(f func(interface{}) string) String {
	if f != nil && m.IsPresent() {
		return OfString(f(m.mixed))
	}

	return emptyString
}

// MapToFloat64 applies mapping function on optional value if it presents
func (m Mixed) MapToFloat64(f func(interface{}) float64) Float64 {
	if f != nil && m.IsPresent() {
		return OfFloat64(f(m.mixed))
	}

	return emptyFloat64
}
