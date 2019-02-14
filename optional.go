// Package optional package contain primitives to work with optional data
package optional

// Optional is minimal non-generic interface for optional data holders
type Optional interface {
	IsPresent() bool
}

// Scanner represents subset of optional values, that supports mapping from SQL and JSON
type Scanner interface {
	Optional
	Scan(interface{}) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}
