package main

import (
	"time"
	"volcanofm/vlc"
)

func main() {

	tracks := []string{
		"/Users/jon/Desktop/radio/000001.mp3",
		"/Users/jon/Desktop/radio/000002.mp3",
	}

	vlc := volcanofm.VLC{
		Socket: "/tmp/vlc.sock",
	}

	vlc.Clear()

	for _, path := range tracks {
		vlc.Enqueue(path)
	}

	vlc.Play()
	time.Sleep(10 * time.Second)
	vlc.Stop()
}
