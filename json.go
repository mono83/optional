package optional

import (
	"bytes"
	"encoding/json"
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
