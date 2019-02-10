package optional

import (
	"strconv"
	"time"
)

// ToString returns optional string, built from value of optional int
func (i Int) ToString() String {
	if i.IsPresent() {
		return OfString(strconv.Itoa(*i.int))
	}

	return emptyString
}

// ToDurationSeconds return optional time.Duration, built from value of optional int
func (i Int) ToDurationSeconds() DurationSeconds {
	if i.IsPresent() {
		return OfSeconds(*i.int)
	}

	return emptyDurationSeconds
}

// ToUnixSeconds return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in seconds
func (t Time) ToUnixSeconds() Int {
	if t.IsPresent() {
		return OfInt(int(t.time.Unix()))
	}

	return emptyInt
}

// ToUnixMillis return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in milliseconds
func (t Time) ToUnixMillis() Int {
	if t.IsPresent() {
		return OfInt(int(t.time.UnixNano() / 1000000))
	}
	return emptyInt

}

// ToSeconds returns optional float64, built from value of optional duration
func (d Duration) ToSeconds() Float64 {
	if d.IsPresent() {
		return OfFloat64(d.duration.Seconds())
	}

	return emptyFloat64
}

// ToMillis returns optional int, built from value of optional duration
func (d Duration) ToMillis() Int {
	if d.IsPresent() {
		return OfInt(int(*d.duration / time.Millisecond))
	}

	return emptyInt
}
