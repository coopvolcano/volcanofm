package volcanofm

import (
	"log"
	"net"
	"time"
)

type VLC struct {
}

func (vlc VLC) Play() {
	vlc.command("play")
}

func (vlc VLC) Stop() {
	vlc.command("stop")
}

func (vlc VLC) Clear() {
	vlc.command("clear")
}

func (vlc VLC) Next() {
	vlc.command("next")
}

func (vlc VLC) RandomOn() {
	vlc.command("random on")
}

func (vlc VLC) Enqueue(path string) {
	vlc.command("enqueue " + path)
}

func (vlc VLC) command(command string) {
	conn, err := net.Dial("tcp", "vlc:4242")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()

	_, err2 := conn.Write([]byte(command))
	if err != nil {
		log.Fatal("Write error:", err2)
	}

	log.Print("Client sent:", command)
	time.Sleep(250 * time.Millisecond)
}
