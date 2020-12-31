package utils

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
