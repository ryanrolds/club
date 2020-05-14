package main

import (
	"log"
	"net/http"

	"github.com/ryanrolds/club/signalling"
)

func main() {
	http.Handle("/room", &signalling.SignallingServer{})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
