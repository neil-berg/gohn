package flags

import (
	"flag"
)

// ParseFlags parses command line arguments to fetch Hacker News stories
func ParseFlags() *int {
	defaultCount := 5
	countPtr := flag.Int("count", defaultCount, "Number of stories to fetch")

	flag.Parse()

	if *countPtr < 0 {
		*countPtr = defaultCount
	}

	return countPtr
}
