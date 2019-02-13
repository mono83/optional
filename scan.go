package optional

import (
	"fmt"
	"strconv"
	"time"
)

// Scan is sql.Scanner interface implementation
func (b *Bool) Scan(src interface{}) error {
	if src == nil {
		*b = Bool{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		b.set(sc == "true" || sc == "1")
	case bool:
		b.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (i *Int) Scan(src interface{}) error {
	if src == nil {
		*i = Int{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		i.set(ic)
	case int64:
		i.set(int(x))
	case int:
		i.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (s *String) Scan(src interface{}) error {
	if src == nil {
		*s = String{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		s.set(string(x))
	case string:
		s.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (f *Float64) Scan(src interface{}) error {
	if src == nil {
		*f = Float64{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.ParseFloat(sc, 64)
		if err != nil {
			return err
		}
		f.set(ic)
	case float64:
		f.set(x)
	case int:
		f.set(float64(x))
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (t *TimeUnixSeconds) Scan(src interface{}) error {
	if src == nil {
		*t = TimeUnixSeconds{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		t.set(time.Unix(int64(ic), 0).UTC())
	case int64:
		t.set(time.Unix(x, 0).UTC())
	case int:
		t.set(time.Unix(int64(x), 0).UTC())
	case time.Time:
		t.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (d *DurationSeconds) Scan(src interface{}) error {
	if src == nil {
		*d = DurationSeconds{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		d.set(time.Duration(int64(ic)) * time.Second)
	case int64:
		d.set(time.Duration(x) * time.Second)
	case int:
		d.set(time.Duration(int64(x)) * time.Second)
	case time.Duration:
		d.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (d *DurationMillis) Scan(src interface{}) error {
	if src == nil {
		*d = DurationMillis{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		d.set(time.Duration(int64(ic)) * time.Millisecond)
	case int64:
		d.set(time.Duration(x) * time.Millisecond)
	case int:
		d.set(time.Duration(int64(x)) * time.Millisecond)
	case time.Duration:
		d.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (d *DurationMinutes) Scan(src interface{}) error {
	if src == nil {
		*d = DurationMinutes{}
		return nil
	}

	switch x := src.(type) {
	case []byte:
		sc := string(x)
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		d.set(time.Duration(int64(ic)) * time.Minute)
	case int64:
		d.set(time.Duration(x) * time.Minute)
	case int:
		d.set(time.Duration(int64(x)) * time.Minute)
	case time.Duration:
		d.set(x)
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}
