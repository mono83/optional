package optional

// Float64 contains optional float64 value
type Float64 struct {
	value    float64
	presents bool
}

func (f *Float64) set(value float64) {
	f.value = value
	f.presents = true
}

// OfFloat64 creates new optional float64 containing provided value
func OfFloat64(f float64) Float64 {
	return Float64{value: f, presents: true}
}

// OfFloat64Ref creates new optional float64 containing provided value
func OfFloat64Ref(f *float64) Float64 {
	if f == nil {
		return Float64{}
	}

	return OfFloat64(*f)
}

// FilterZero applies zero value filtering
// This method will return empty optional if value inside optional is zero or missing
func (f Float64) FilterZero() Float64 {
	if !f.IsPresent() || f.value == 0 {
		return Float64{}
	}

	return f
}
