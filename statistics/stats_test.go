package statistics

import (
	"reflect"
	"testing"

	"github.com/Geoffrey42/fizzbuzz/models"
	"github.com/go-redis/redis"
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

func TestGetTopRequestFromList(t *testing.T) {
	t.Run("The is one top request", func(t *testing.T) {
		var topRequests = []redis.Z{
			{
				Score:  42,
				Member: "3-5-16-fizz-buzz",
			},
			{
				Score:  21,
				Member: "2-8-75-Perceval-Karadoc",
			},
		}

		got, _ := GetTopRequestFromList(topRequests)
		want := models.Stat{
			Hit:   42,
			Int1:  3,
			Int2:  5,
			Limit: 16,
			Str1:  "fizz",
			Str2:  "buzz",
		}

		if !reflect.DeepEqual(*got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("There are several top requests", func(t *testing.T) {
		var topRequests = []redis.Z{
			{
				Score:  42,
				Member: "3-5-16-fizz-buzz",
			},
			{
				Score:  42,
				Member: "2-8-75-Perceval-Karadoc",
			},
		}

		_, err := GetTopRequestFromList(topRequests)

		if err == nil {
			t.Errorf("got no error but expect one")
		}
	})
}
