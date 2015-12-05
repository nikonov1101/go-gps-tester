package main

import (
	"net/http"
	"log"
	"github.com/macaron-contrib/pongo2"
	"github.com/Unknwon/macaron"
)

const (
	LISTEN_ADDR = "0.0.0.0:8095"
)

func main() {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer())

	// routing
	m.Get("/", HttpIndex)
	m.Get("/put", HttpPutGps)

	log.Printf("Serve on %s", LISTEN_ADDR)
	http.ListenAndServe(LISTEN_ADDR, m)
}