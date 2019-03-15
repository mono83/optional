package optional

import (
	"time"
)

// MarshalYAML is yaml.Marshaler interface implementation
func (b Bool) MarshalYAML() (interface{}, error) {
	if !b.IsPresent() {
		return nil, nil
	}
	return b.value, nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (b *Bool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate bool
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	b.set(candidate)
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (i Int) MarshalYAML() (interface{}, error) {
	if !i.IsPresent() {
		return nil, nil
	}
	return i.value, nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (i *Int) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate int
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	i.set(candidate)
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (s String) MarshalYAML() (interface{}, error) {
	if !s.IsPresent() {
		return nil, nil
	}
	return s.value, nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (s *String) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate string
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	s.set(candidate)
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (f Float64) MarshalYAML() (interface{}, error) {
	if !f.IsPresent() {
		return nil, nil
	}
	return f.value, nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (f *Float64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate float64
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	f.set(candidate)
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (t TimeUnixSeconds) MarshalYAML() (interface{}, error) {
	if !t.IsPresent() {
		return nil, nil
	}
	return t.value.Unix(), nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (t *TimeUnixSeconds) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate int
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	tc := OfUnixSeconds(candidate)
	*t = tc
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (d DurationSeconds) MarshalYAML() (interface{}, error) {
	if !d.IsPresent() {
		return nil, nil
	}

	return int(d.value / time.Second), nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (d *DurationSeconds) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate int
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	tc := OfSeconds(candidate)
	*d = tc
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (d DurationMillis) MarshalYAML() (interface{}, error) {
	if !d.IsPresent() {
		return nil, nil
	}
	return int(d.value / time.Millisecond), nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (d *DurationMillis) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate int
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	tc := OfMilliseconds(candidate)
	*d = tc
	return nil
}

// MarshalYAML is yaml.Marshaler interface implementation
func (d DurationMinutes) MarshalYAML() (interface{}, error) {
	if !d.IsPresent() {
		return nil, nil
	}
	return int(d.value / time.Minute), nil
}

// UnmarshalYAML is yaml.Unmarshaler interface implementation
func (d *DurationMinutes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var candidate int
	if err := unmarshal(&candidate); err != nil {
		return err
	}

	tc := OfMinutes(candidate)
	*d = tc
	return nil
}
