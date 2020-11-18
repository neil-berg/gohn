package flags

import (
	"flag"

	"github.com/neil-berg/gohn/utils"
)

// ParseFlags parses command line arguments to fetch Hacker News stories
func ParseFlags() (*int, *string) {
	storyTypes := []string{"top", "ask", "show", "job"}

	countPtr := flag.Int("count", 5, "Number of stories to fetch between 1 and 10")
	storyTypePtr := flag.String("type", "top", "Type of stories to fetch (top, ask, show, job)")

	flag.Parse()

	if *countPtr < 1 {
		*countPtr = 1
	}

	if *countPtr > 10 {
		*countPtr = 10
	}

	if !utils.Contains(storyTypes, *storyTypePtr) {
		*storyTypePtr = "top"
	}

	return countPtr, storyTypePtr
}
