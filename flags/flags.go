package flags

import (
	"flag"
	"fmt"
)

// ParseFlags parses command line arguments to fetch Hacker News stories
func ParseFlags() {
	countPtr := flag.Int("count", 5, "Number of stories to fetch")
	flag.Parse()

	fmt.Println("Getting", *countPtr, "stories")
}
