package optional

import (
	"fmt"
	"strconv"
	"time"
)

// Scan is sql.Scanner interface implementation
func (b *Bool) Scan(src interface{}) error {
	if src == nil {
		*b = emptyBool
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		v := sc == "true" || sc == "1"
		b.bool = &v
	case bool:
		v := src.(bool)
		b.bool = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (i *Int) Scan(src interface{}) error {
	if src == nil {
		*i = emptyInt
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		i.int = &ic
	case int64:
		v := int(src.(int64))
		i.int = &v
	case int:
		v := src.(int)
		i.int = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (s *String) Scan(src interface{}) error {
	if src == nil {
		*s = emptyString
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		s.string = &sc
	case string:
		v := src.(string)
		s.string = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (f *Float64) Scan(src interface{}) error {
	if src == nil {
		*f = emptyFloat64
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		ic, err := strconv.ParseFloat(sc, 64)
		if err != nil {
			return err
		}
		f.float64 = &ic
	case float64:
		v := src.(float64)
		f.float64 = &v
	case int:
		v := float64(src.(int))
		f.float64 = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (t *TimeUnixSeconds) Scan(src interface{}) error {
	if src == nil {
		t.time = nil
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		tc := time.Unix(int64(ic), 0).UTC()
		t.time = &tc
	case int64:
		v := time.Unix(src.(int64), 0).UTC()
		t.time = &v
	case int:
		v := time.Unix(int64(src.(int)), 0).UTC()
		t.time = &v
	case time.Time:
		v := src.(time.Time)
		t.time = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}

// Scan is sql.Scanner interface implementation
func (d *DurationSeconds) Scan(src interface{}) error {
	if src == nil {
		d.duration = nil
		return nil
	}

	switch src.(type) {
	case []byte:
		sc := string(src.([]byte))
		ic, err := strconv.Atoi(sc)
		if err != nil {
			return err
		}
		tc := time.Duration(int64(ic)) * time.Second
		d.duration = &tc
	case int64:
		v := time.Duration(src.(int64)) * time.Second
		d.duration = &v
	case int:
		v := time.Duration(int64(src.(int))) * time.Second
		d.duration = &v
	case time.Duration:
		v := src.(time.Duration)
		d.duration = &v
	default:
		return fmt.Errorf("unsupported type %T for scanning", src)
	}

	return nil
}
