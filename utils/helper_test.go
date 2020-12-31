package utils

import (
	"testing"
)

func TestBuildMemberFromParams(t *testing.T) {
	params := make(map[string][]string)

	params["int1"] = []string{"3"}
	params["int2"] = []string{"5"}
	params["limit"] = []string{"16"}
	params["str1"] = []string{"fizz"}
	params["str2"] = []string{"buzz"}

	got := BuildMemberFromParams(params)
	want := "3-5-16-fizz-buzz"

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
