package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Received %q, expected %q", got, want)
		}
	}

	t.Run("Return 5 A's", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEqual(t, got, want)
	})

	t.Run("Return 4 B's", func(t *testing.T) {
		got := Repeat("b", 4)
		want := "bbbb"
		assertEqual(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	got := Repeat("!", 10)
	fmt.Println(got)
	// Output: !!!!!!!!!!
}
