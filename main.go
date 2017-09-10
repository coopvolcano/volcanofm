package main

import (
	"log"
	"net"
	"time"
)

func toVLC(command string) {
	conn, err := net.Dial("unix", "/tmp/vlc.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	_, err2 := conn.Write([]byte(command))
	if err != nil {
		log.Fatal("Write error:", err2)
	}

	log.Print("Client sent:", command)
	time.Sleep(2 * time.Second)
}

func main() {
	toVLC("clear")
	toVLC("enqueue /Users/jon/Desktop/radio/000001.mp3")
	toVLC("enqueue /Users/jon/Desktop/radio/000002.mp3")
	toVLC("play")
	time.Sleep(10 * time.Second)
	toVLC("pause")
}
