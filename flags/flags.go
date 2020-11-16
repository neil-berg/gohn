package flags

import (
	"flag"
)

// ParseFlags parses command line arguments to fetch Hacker News stories
func ParseFlags() *int {
	countPtr := flag.Int("count", 5, "Number of stories to fetch")
	// TODO: Validate that count is positive
	flag.Parse()

	return countPtr
}
