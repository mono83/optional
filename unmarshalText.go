package optional

import (
	"errors"
	"strconv"
	"strings"
)

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (b *Bool) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	str := strings.ToLower(string(bts))
	if str == "true" {
		b.set(true)
	} else if str == "false" {
		b.set(false)
	} else {
		return errors.New("unable to produce boolean from " + str)
	}

	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (i *Int) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseInt(string(bts), 10, 64)
	if err != nil {
		return err
	}

	i.set(int(v))
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (s *String) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	s.set(string(bts))
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (f *Float64) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseFloat(string(bts), 64)
	if err != nil {
		return err
	}

	f.set(v)
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (d *DurationMillis) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseInt(string(bts), 10, 64)
	if err != nil {
		return err
	}

	dd := OfMilliseconds(int(v))
	*d = dd
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (d *DurationSeconds) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseInt(string(bts), 10, 64)
	if err != nil {
		return err
	}

	dd := OfSeconds(int(v))
	*d = dd
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (d *DurationMinutes) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseInt(string(bts), 10, 64)
	if err != nil {
		return err
	}

	dd := OfMinutes(int(v))
	*d = dd
	return nil
}

// UnmarshalText is encoding.TextUnmarshaler interface implementation
func (t *TimeUnixSeconds) UnmarshalText(bts []byte) error {
	if len(bts) == 0 {
		// Empty value
		return nil
	}

	v, err := strconv.ParseInt(string(bts), 10, 64)
	if err != nil {
		return err
	}

	tt := OfUnixSeconds(int(v))
	*t = tt
	return nil
}
