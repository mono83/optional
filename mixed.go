package optional

// Mixed represents arbitrary interface{} optional
type Mixed struct {
	mixed interface{}
}

// OfMixed creates new optional interface{} holder containing provided value
func OfMixed(src interface{}) Mixed {
	return Mixed{mixed: src}
}
