package util

import "testing"

func check(t *testing.T, want string, got string) {
	if want != got {
		t.Errorf("wanted %s; got %s", want, got)
	}
}

func TestEllipsis(t *testing.T) {

	got := Ellipsis("", 5)
	check(t, "", got)

	got = Ellipsis("xyz", 4)
	check(t, "xyz", got)

	got = Ellipsis("xyz", 3)
	check(t, "xyz", got)

	got = Ellipsis("xyz", 2)
	check(t, "xy…", got)

	got = Ellipsis("xyz", 1)
	check(t, "x…", got)
}
