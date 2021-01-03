// Package statistics provides helpuf functions to handle redis Sorted Set interactions
package statistics

import (
	"strconv"
	"strings"

	"github.com/Geoffrey42/fizzbuzz/models"
	"github.com/go-redis/redis"
)

// Key is the Redis Sorted Set storing the request stats
const Key string = "top-requests"

// BuildMemberFromParams converts request's query parameters into a
// Sorted Set member e.g '/api/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz'
// turns into "3-5-16-fizz-buzz".
func BuildMemberFromParams(p map[string][]string) string {
	res := p["int1"][0] + "-" + p["int2"][0] + "-" +
		p["limit"][0] + "-" + p["str1"][0] + "-" + p["str2"][0]

	return res
}

// GetTopRequestFromList return the most hit request from a redis Sorted Set
// or an error if several has the same score.
func GetTopRequestFromList(topRequests []redis.Z) (*models.Stat, *models.Error) {
	if len(topRequests) > 1 && topRequests[0].Score == topRequests[1].Score {
		return nil, &models.Error{Code: 404, Message: "Several requests have the top hits."}
	}

	if str, ok := topRequests[0].Member.(string); ok {
		params := strings.Split(str, "-")

		int1, _ := strconv.ParseInt(params[0], 10, 64)
		int2, _ := strconv.ParseInt(params[1], 10, 64)
		limit, _ := strconv.ParseInt(params[2], 10, 64)

		topRequest := models.Stat{
			Hit:   int64(topRequests[0].Score),
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  params[3],
			Str2:  params[4],
		}

		return &topRequest, nil
	}

	return nil, &models.Error{Code: 404, Message: "No request can be found."}
}
