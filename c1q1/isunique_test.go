package c1q1

import "testing"

func TestUniqueMap(t *testing.T) {
	want := true
	if got := UniqueMap("abcde"); got != want {
		t.Errorf("UniqueMap() = %v, want %v", got, want)
	}
}
