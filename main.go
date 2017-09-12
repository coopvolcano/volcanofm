package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	volcanofm "volcanofm/vlc"

	"github.com/dhowden/tag"

	_ "github.com/lib/pq"
)

// Endpoint to upload file
// Read file and parse metadata
// Push metadata to database
// Push file to S3 naming the file the primary id from the database
// Enqueue file to be played

// http://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-2
// https://n0tablog.wordpress.com/2009/02/09/controlling-vlc-via-rc-remote-control-interface-using-a-unix-domain-socket-and-no-programming/

/*
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"AddPerm",
      "Effect":"Allow",
      "Principal": "*",
      "Action":["s3:GetObject"],
      "Resource":["arn:aws:s3:::examplebucket/*"]
    }
  ]
}
*/

func main() {
	musicFiles, _ := filepath.Glob("/uploads/**/*.flac")
	tracks := []string{}

	for _, path := range musicFiles {
		file, err := os.Open(path)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}

		metadata, err := tag.ReadFrom(file)

		fmt.Println(metadata.Album() + "***" + metadata.Title() + "***" + metadata.Artist())

		// TODO: This needs to be a database ID or something.
		sub := "00001"
		tracks = append(tracks, "https://s3.amazonaws.com/volcano-fm/jondev/"+sub)
	}

	vlc := volcanofm.VLC{}
	defer vlc.Stop()

	vlc.Clear()

	for _, path := range tracks {
		vlc.Enqueue(path)
	}

	vlc.Play()
	time.Sleep(10 * time.Second)

	vlc.Next()
	time.Sleep(10 * time.Second)
}
