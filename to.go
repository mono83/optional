package optional

import (
	"strconv"
	"time"
)

// ToString returns optional string, built from value of optional int
func (i Int) ToString() String {
	if i.int == nil {
		return emptyString
	}

	return OfString(strconv.Itoa(*i.int))
}

// ToDurationSeconds return optional time.Duration, built from value of optional int
func (i Int) ToDurationSeconds() DurationSeconds {
	if i.int == nil {
		return emptyDurationSeconds
	}

	return OfSeconds(*i.int)
}

// ToUnixSeconds return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in seconds
func (t Time) ToUnixSeconds() Int {
	if t.time == nil {
		return emptyInt
	}

	return OfInt(int(t.time.Unix()))
}

// ToUnixMillis return optional int, built from value of optional time
// If optional was not empty, resulting optional int will contain
// unix timestamp in milliseconds
func (t Time) ToUnixMillis() Int {
	if t.time == nil {
		return emptyInt
	}

	return OfInt(int(t.time.UnixNano() / 1000000))
}

// ToSeconds returns optional float64, built from value of optional duration
func (d Duration) ToSeconds() Float64 {
	if d.duration == nil {
		return emptyFloat64
	}

	return OfFloat64(d.duration.Seconds())
}

// ToMillis returns optional int, built from value of optional duration
func (d Duration) ToMillis() Int {
	if d.duration == nil {
		return emptyInt
	}

	return OfInt(int(*d.duration / time.Millisecond))
}