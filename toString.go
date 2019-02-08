package optional

import (
	"fmt"
	"strconv"
)

func (b Bool) String() string {
	if b.bool == nil {
		return "Optional.Empty"
	} else if *b.bool {
		return "Optional[true]"
	}

	return "Optional[false]"
}

func (i Int) String() string {
	if i.int == nil {
		return "Optional.Empty"
	}

	return "Optional[" + strconv.Itoa(*i.int) + "]"
}

func (s String) String() string {
	if s.string == nil {
		return "Optional.Empty"
	}

	return "Optional[" + *s.string + "]"
}

func (f Float64) String() string {
	if f.float64 == nil {
		return "Optional.Empty"
	}

	return fmt.Sprintf("Optional[%f]", *f.float64)
}

func (t Time) String() string {
	if t.time == nil {
		return "Optional.Empty"
	}

	return "Optional[" + t.time.String() + "]"
}
