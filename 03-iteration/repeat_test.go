package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected: %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; 1 > b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: "aaaaaaaaa"
}
