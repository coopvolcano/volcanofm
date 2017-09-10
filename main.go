package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	volcanofm "volcanofm/vlc"

	"github.com/dhowden/tag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

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
	musicFiles := []string{
		"01.flac",
		"02.flac",
	}
	tracks := []string{}

	for _, sub := range musicFiles {
		path := filepath.Join("C:\\", "Users", "binarycleric", "Desktop", sub)
		file, err := os.Open(path)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}

		metadata, err := tag.ReadFrom(file)

		fmt.Println(metadata.Album())
		fmt.Println(metadata.Title())
		fmt.Println(metadata.Artist())

		s3Key := "jondev/" + sub
		sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
		uploader := s3manager.NewUploader(sess)
		response, err2 := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("volcano-fm"),
			Key:    aws.String(s3Key),
			Body:   file,
		})

		if err2 != nil {
			fmt.Println(err2)
		}

		fmt.Println(response)

		tracks = append(tracks, "https://s3.amazonaws.com/volcano-fm/jondev/"+sub)
	}

	fmt.Println(tracks)

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
