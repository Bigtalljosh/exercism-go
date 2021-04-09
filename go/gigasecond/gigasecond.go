/*
	Package gigasecond adds one gigasecond (10^9 (1,000,000,000) seconds) to a given time.
*/
package gigasecond

import "time"

// AddGigasecond takes a time, adds a gigasecond and returns the result
func AddGigasecond(t time.Time) time.Time {
	const gigasecond = time.Second * 1e9
	return t.Add(gigasecond)
}
