package flags

import (
	"flag"
)

// ParseFlags parses command line arguments to fetch Hacker News stories
func ParseFlags() *int {
	countPtr := flag.Int("count", 5, "Number of stories to fetch between 1 and 10")

	flag.Parse()

	if *countPtr < 1 {
		*countPtr = 1
	}

	if *countPtr > 10 {
		*countPtr = 10
	}

	return countPtr
}
