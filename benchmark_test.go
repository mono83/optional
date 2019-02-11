package optional

import (
	"testing"
	"time"
)

const sliceSize = 1000

func BenchmarkOfBool(b *testing.B) {
	slice := make([]Bool, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfBool(true)
		}
	}
}

func BenchmarkOfInt(b *testing.B) {
	slice := make([]Int, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfInt(42)
		}
	}
}

func BenchmarkOfString(b *testing.B) {
	slice := make([]String, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfString("foo")
		}
	}
}

func BenchmarkOfFloat64(b *testing.B) {
	slice := make([]Float64, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfFloat64(.12345)
		}
	}
}

func BenchmarkOfTime(b *testing.B) {
	slice := make([]Time, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfTime(time.Time{})
		}
	}
}

func BenchmarkOfDuration(b *testing.B) {
	slice := make([]Duration, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfDuration(time.Second)
		}
	}
}

func BenchmarkOfMixed(b *testing.B) {
	slice := make([]Mixed, sliceSize)
	for x := 0; x < b.N; x++ {
		for i := 0; i < sliceSize; i++ {
			slice[i] = OfMixed("foo")
		}
	}
}

func BenchmarkFilterOrElseBool(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfBool(true).Filter(func(bool) bool { return false }).OrElse(false)
	}
}

func BenchmarkFilterOrElseInt(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfInt(42).Filter(func(int) bool { return false }).OrElse(12)
	}
}

func BenchmarkFilterOrElseString(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfString("foo").Filter(func(string) bool { return false }).OrElse("bar")
	}
}

func BenchmarkFilterOrElseFloat64(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfFloat64(.1).Filter(func(float64) bool { return false }).OrElse(.2)
	}
}

func BenchmarkFilterOrElseTime(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfTime(time.Time{}).Filter(func(time.Time) bool { return false }).OrElse(time.Time{})
	}
}

func BenchmarkFilterOrElseDuration(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfDuration(time.Second).Filter(func(time.Duration) bool { return false }).OrElse(time.Minute)
	}
}

func BenchmarkFilterOrElseMixed(b *testing.B) {
	for x := 0; x < b.N; x++ {
		OfMixed("foo").Filter(func(interface{}) bool { return false }).OrElse(42)
	}
}
