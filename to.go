package optional

import (
	"strconv"
	"time"
)

// ToString returns optional string, built from value of optional int
func (i Int) ToString() String {
	if i.IsPresent() {
		return OfString(strconv.Itoa(i.value))
	}

	return String{}
}

// ToDurationSeconds return optional time.Duration, built from value of optional int
func (i Int) ToDurationSeconds() DurationSeconds {
	if i.IsPresent() {
		return OfSeconds(i.value)
	}

	return DurationSeconds{}
}

// ToUnixSeconds return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in seconds
func (t Time) ToUnixSeconds() Int {
	if t.IsPresent() {
		return OfInt(int(t.value.Unix()))
	}

	return Int{}
}

// ToUnixMillis return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in milliseconds
func (t Time) ToUnixMillis() Int {
	if t.IsPresent() {
		return OfInt(int(t.value.UnixNano() / 1000000))
	}
	return Int{}

}

// ToSeconds returns optional float64, built from value of optional duration
func (d Duration) ToSeconds() Float64 {
	if d.IsPresent() {
		return OfFloat64(d.value.Seconds())
	}

	return Float64{}
}

// ToMillis returns optional int, built from value of optional duration
func (d Duration) ToMillis() Int {
	if d.IsPresent() {
		return OfInt(int(d.value / time.Millisecond))
	}

	return Int{}
}

// ToString returns optional string, built from value of optional error
func (e Error) ToString() String {
	if e.IsPresent() {
		return OfString(e.value.Error())
	}

	return String{}
}
