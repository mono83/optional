package optional

import "time"

/*
   Bool map functions
*/

// Map applies mapping function on optional value if it presents
func (b Bool) Map(f func(bool) bool) Bool {
	if f == nil || !b.IsPresent() {
		return Bool{}
	}

	return OfBool(f(b.value))
}

// MapToInt applies mapping function on optional value if it presents
func (b Bool) MapToInt(f func(bool) int) Int {
	if f == nil || !b.IsPresent() {
		return Int{}
	}

	return OfInt(f(b.value))
}

// MapToString applies mapping function on optional value if it presents
func (b Bool) MapToString(f func(bool) string) String {
	if f == nil || !b.IsPresent() {
		return String{}
	}

	return OfString(f(b.value))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (b Bool) MapToFloat64(f func(bool) float64) Float64 {
	if f == nil || !b.IsPresent() {
		return Float64{}
	}

	return OfFloat64(f(b.value))
}

/*
   Int map functions
*/

// Map applies mapping function on optional value if it presents
func (i Int) Map(f func(int) int) Int {
	if f == nil || !i.IsPresent() {
		return Int{}
	}

	return OfInt(f(i.value))
}

// MapToBool applies mapping function on optional value if it presents
func (i Int) MapToBool(f func(int) bool) Bool {
	if f == nil || !i.IsPresent() {
		return Bool{}
	}

	return OfBool(f(i.value))
}

// MapToString applies mapping function on optional value if it presents
func (i Int) MapToString(f func(int) string) String {
	if f == nil || !i.IsPresent() {
		return String{}
	}

	return OfString(f(i.value))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (i Int) MapToFloat64(f func(int) float64) Float64 {
	if f == nil || !i.IsPresent() {
		return Float64{}
	}

	return OfFloat64(f(i.value))
}

/*
   String map functions
*/

// Map applies mapping function on optional value if it presents
func (s String) Map(f func(string) string) String {
	if f == nil || !s.IsPresent() {
		return String{}
	}

	return OfString(f(s.value))
}

// MapToBool applies mapping function on optional value if it presents
func (s String) MapToBool(f func(string) bool) Bool {
	if f == nil || !s.IsPresent() {
		return Bool{}
	}

	return OfBool(f(s.value))
}

// MapToInt applies mapping function on optional value if it presents
func (s String) MapToInt(f func(string) int) Int {
	if f == nil || !s.IsPresent() {
		return Int{}
	}

	return OfInt(f(s.value))
}

// MapToFloat64 applies mapping function on optional value if it presents
func (s String) MapToFloat64(f func(string) float64) Float64 {
	if f == nil || !s.IsPresent() {
		return Float64{}
	}

	return OfFloat64(f(s.value))
}

/*
   Float64 map functions
*/

// Map applies mapping function on optional value if it presents
func (f Float64) Map(c func(float64) float64) Float64 {
	if c == nil || !f.IsPresent() {
		return Float64{}
	}

	return OfFloat64(c(f.value))
}

// MapToBool applies mapping function on optional value if it presents
func (f Float64) MapToBool(c func(float64) bool) Bool {
	if c == nil || !f.IsPresent() {
		return Bool{}
	}

	return OfBool(c(f.value))
}

// MapToInt applies mapping function on optional value if it presents
func (f Float64) MapToInt(c func(float64) int) Int {
	if c == nil || !f.IsPresent() {
		return Int{}
	}

	return OfInt(c(f.value))
}

// MapToString applies mapping function on optional value if it presents
func (f Float64) MapToString(c func(float64) string) String {
	if c == nil || !f.IsPresent() {
		return String{}
	}

	return OfString(c(f.value))
}

/*
   Time map functions
*/

// Map applies mapping function on optional value if it presents
func (t Time) Map(f func(time.Time) time.Time) Time {
	if f == nil || !t.IsPresent() {
		return Time{}
	}

	return OfTime(f(t.value))
}

/*
   Int map functions
*/

// Map applies mapping function on optional value if it presents
func (d Duration) Map(f func(time.Duration) time.Duration) Duration {
	if f == nil || !d.IsPresent() {
		return Duration{}
	}

	return OfDuration(f(d.value))
}

/*
   Mixed map functions
*/

// Map applies mapping function on optional value if it presents
func (m Mixed) Map(f func(interface{}) interface{}) Mixed {
	if f != nil && m.IsPresent() {
		return OfMixed(f(m.mixed))
	}

	return Mixed{}
}

// MapToBool applies mapping function on optional value if it presents
func (m Mixed) MapToBool(f func(interface{}) bool) Bool {
	if f != nil && m.IsPresent() {
		return OfBool(f(m.mixed))
	}

	return Bool{}
}

// MapToInt applies mapping function on optional value if it presents
func (m Mixed) MapToInt(f func(interface{}) int) Int {
	if f != nil && m.IsPresent() {
		return OfInt(f(m.mixed))
	}

	return Int{}
}

// MapToString applies mapping function on optional value if it presents
func (m Mixed) MapToString(f func(interface{}) string) String {
	if f != nil && m.IsPresent() {
		return OfString(f(m.mixed))
	}

	return String{}
}

// MapToFloat64 applies mapping function on optional value if it presents
func (m Mixed) MapToFloat64(f func(interface{}) float64) Float64 {
	if f != nil && m.IsPresent() {
		return OfFloat64(f(m.mixed))
	}

	return Float64{}
}
