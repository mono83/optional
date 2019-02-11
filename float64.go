package optional

var emptyFloat64 = Float64{}

// Float64 contains optional float64 value
type Float64 struct {
	*float64
}

// OfFloat64 creates new optional float64 containing provided value
func OfFloat64(f float64) Float64 {
	return Float64{float64: &f}
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (f Float64) FilterZero() Float64 {
	if !f.IsPresent() || *f.float64 == 0 {
		return emptyFloat64
	}

	return f
}
