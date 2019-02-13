package optional

import (
	"bytes"
	"encoding/json"
	"time"
)

var jsonNull = []byte("null")

// MarshalJSON is json.Marshaler interface implementation
func (b Bool) MarshalJSON() (text []byte, err error) {
	if !b.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(b.value)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (b *Bool) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*b = Bool{}
		return nil
	}
	var value bool
	if err := json.Unmarshal(text, &value); err != nil {
		return err
	}
	b.set(value)
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (i Int) MarshalJSON() (text []byte, err error) {
	if !i.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (i *Int) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*i = Int{}
		return nil
	}
	var value int
	if err := json.Unmarshal(text, &value); err != nil {
		return err
	}
	i.set(value)
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (s String) MarshalJSON() (text []byte, err error) {
	if !s.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(s.value)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (s *String) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*s = String{}
		return nil
	}
	var value string
	if err := json.Unmarshal(text, &value); err != nil {
		return err
	}
	s.set(value)
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (f Float64) MarshalJSON() (text []byte, err error) {
	if !f.presents {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (f *Float64) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*f = Float64{}
		return nil
	}

	var value float64
	if err := json.Unmarshal(text, &value); err != nil {
		return err
	}
	f.set(value)
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (t TimeUnixSeconds) MarshalJSON() (text []byte, err error) {
	if !t.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(t.value.Unix())
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (t *TimeUnixSeconds) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*t = TimeUnixSeconds{}
		return nil
	}
	var sec int
	if err := json.Unmarshal(text, &sec); err != nil {
		return err
	}
	tc := OfUnixSeconds(sec)
	*t = tc
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (d DurationSeconds) MarshalJSON() (text []byte, err error) {
	if !d.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(int(d.value / time.Second))
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (d *DurationSeconds) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*d = DurationSeconds{}
		return nil
	}
	var sec int
	if err := json.Unmarshal(text, &sec); err != nil {
		return err
	}
	tc := OfSeconds(sec)
	*d = tc
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (d DurationMillis) MarshalJSON() (text []byte, err error) {
	if !d.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(int(d.value / time.Millisecond))
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (d *DurationMillis) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*d = DurationMillis{}
		return nil
	}
	var millis int
	if err := json.Unmarshal(text, &millis); err != nil {
		return err
	}
	tc := OfMilliseconds(millis)
	*d = tc
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (d DurationMinutes) MarshalJSON() (text []byte, err error) {
	if !d.IsPresent() {
		return json.Marshal(nil)
	}
	return json.Marshal(int(d.value / time.Minute))
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (d *DurationMinutes) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*d = DurationMinutes{}
		return nil
	}
	var minutes int
	if err := json.Unmarshal(text, &minutes); err != nil {
		return err
	}
	tc := OfMinutes(minutes)
	*d = tc
	return nil
}
