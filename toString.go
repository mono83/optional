package optional

import (
	"fmt"
	"strconv"
)

const optEmptyString = "Optional.Empty"

func (b Bool) String() string {
	if !b.IsPresent() {
		return "Optional.Empty"
	} else if b.value {
		return "Optional[true]"
	}

	return "Optional[false]"
}

func (i Int) String() string {
	if i.IsPresent() {
		return "Optional[" + strconv.Itoa(i.value) + "]"
	}
	return optEmptyString
}

func (s String) String() string {
	if s.IsPresent() {
		return "Optional[" + s.value + "]"
	}
	return optEmptyString
}

func (f Float64) String() string {
	if f.IsPresent() {
		return fmt.Sprintf("Optional[%f]", f.value)
	}

	return optEmptyString
}

func (t Time) String() string {
	if t.IsPresent() {
		return "Optional[" + t.value.String() + "]"
	}

	return optEmptyString
}

func (d Duration) String() string {
	if d.IsPresent() {
		return "Optional[" + d.value.String() + "]"
	}

	return optEmptyString
}

func (e Error) String() string {
	if e.IsPresent() {
		return "Optional[" + e.value.Error() + "]"
	}

	return optEmptyString
}
