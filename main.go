package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	cookieName = "andrewmthomas87-coffee-auth"
)

var addr = flag.String("addr", ":7777", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		r.Cookie("")
		serveWs(w, r)
	})

	log.Printf("Starting HTTP Server at '%s'", *addr)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
