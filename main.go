package main

import (
	"time"
	"volcanofm/vlc"

	_ "github.com/lib/pq"
)

func main() {
	tracks := []string{
		"https://s3.amazonaws.com/volcano-fm/jondev/06+Only+One.flac",
	}

	vlc := volcanofm.VLC{}

	vlc.Clear()

	for _, path := range tracks {
		vlc.Enqueue(path)
	}

	vlc.Play()
	time.Sleep(10 * time.Second)
	vlc.Stop()
}
