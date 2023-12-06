package integration_tests

import "testing"

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}
