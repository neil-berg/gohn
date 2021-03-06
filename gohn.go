package main

import (
	"github.com/neil-berg/gohn/flags"
	"github.com/neil-berg/gohn/requests"
)

func main() {
	count, storyType := flags.ParseFlags()
	requests.GetTopStories(count, storyType)
}
