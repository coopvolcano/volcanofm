package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"

	_ "github.com/lib/pq"
)

// Endpoint to upload file
// Read file and parse metadata
// Push metadata to database
// Push file to S3 naming the file the primary id from the database
// Enqueue file to be played

func main() {
	path := filepath.Join("C:\\", "Users", "binarycleric", "Desktop", "01.flac")
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	metadata, err := tag.ReadFrom(file)

	fmt.Println(metadata.Album())
	fmt.Println(metadata.Title())
	fmt.Println(metadata.Artist())

	/*
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
		})
	*/
	/*
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
	*/
}
