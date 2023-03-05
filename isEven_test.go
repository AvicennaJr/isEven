package iseven

import "testing"

func TestIsEven(t *testing.T) {
	got := IsEven(6)
	want := true

	if got != want {
		t.Errorf("Got %t, want %t", got, want)
	}
}
