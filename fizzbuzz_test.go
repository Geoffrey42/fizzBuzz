package fizzbuzz

import "testing"

func TestDoFizzBuzz(t *testing.T) {
	got := doFizzBuzz(3, 5, 16, "fizz", "buzz")
	want := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
