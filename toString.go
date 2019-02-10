package optional

import (
	"fmt"
	"strconv"
)

const optEmptyString = "Optional.Empty"

func (b Bool) String() string {
	if !b.IsPresent() {
		return "Optional.Empty"
	} else if *b.bool {
		return "Optional[true]"
	}

	return "Optional[false]"
}

func (i Int) String() string {
	if i.IsPresent() {
		return "Optional[" + strconv.Itoa(*i.int) + "]"
	}
	return optEmptyString
}

func (s String) String() string {
	if s.IsPresent() {
		return "Optional[" + *s.string + "]"
	}
	return optEmptyString
}

func (f Float64) String() string {
	if f.IsPresent() {
		return fmt.Sprintf("Optional[%f]", *f.float64)
	}

	return optEmptyString
}

func (t Time) String() string {
	if t.IsPresent() {
		return "Optional[" + t.time.String() + "]"
	}

	return optEmptyString
}

func (d Duration) String() string {
	if d.IsPresent() {
		return "Optional[" + d.duration.String() + "]"
	}

	return optEmptyString
}
