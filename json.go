package optional

import (
	"bytes"
	"encoding/json"
	"time"
)

var jsonNull = []byte("null")

// MarshalJSON is json.Marshaler interface implementation
func (b Bool) MarshalJSON() (text []byte, err error) {
	return json.Marshal(b.bool)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (b *Bool) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*b = emptyBool
		return nil
	}
	t := OfBool(false)
	if err := json.Unmarshal(text, t.bool); err != nil {
		return err
	}
	*b = t
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (i Int) MarshalJSON() (text []byte, err error) {
	return json.Marshal(i.int)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (i *Int) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*i = emptyInt
		return nil
	}
	t := OfInt(0)
	if err := json.Unmarshal(text, t.int); err != nil {
		return err
	}
	*i = t
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (s String) MarshalJSON() (text []byte, err error) {
	return json.Marshal(s.string)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (s *String) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*s = emptyString
		return nil
	}
	t := OfString("")
	if err := json.Unmarshal(text, t.string); err != nil {
		return err
	}
	*s = t
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (f Float64) MarshalJSON() (text []byte, err error) {
	return json.Marshal(f.float64)
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (f *Float64) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		*f = emptyFloat64
		return nil
	}
	t := OfFloat64(0)
	if err := json.Unmarshal(text, t.float64); err != nil {
		return err
	}
	*f = t
	return nil
}

// MarshalJSON is json.Marshaler interface implementation
func (t TimeUnixSeconds) MarshalJSON() (text []byte, err error) {
	if t.time == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(t.time.Unix())
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (t *TimeUnixSeconds) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		t.time = nil
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
	if d.duration == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(int(*d.duration / time.Second))
}

// UnmarshalJSON is json.Unmarshaler interface implementation
func (d *DurationSeconds) UnmarshalJSON(text []byte) error {
	if bytes.Equal(jsonNull, text) {
		d.duration = nil
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
