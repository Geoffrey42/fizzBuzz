package fizzbuzz

import "testing"

func TestDoFizzBuzz(t *testing.T) {
	t.Run("basic case: up to 16 limit with 'fizz' and 'buzz'", func(t *testing.T) {
		got, _ := doFizzBuzz(3, 5, 16, "fizz", "buzz")
		want := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("limit greater than 100", func(t *testing.T) {
		_, err := doFizzBuzz(3, 5, 101, "fizz", "buzz")

		assertError(t, err)
	})

	t.Run("limit lesser than 1", func(t *testing.T) {
		_, err := doFizzBuzz(3, 5, 0, "fizz", "buzz")

		assertError(t, err)
	})
}

func assertError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("wanted an error but didn't get one")
	}
}
