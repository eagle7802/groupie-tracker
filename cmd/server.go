package server

import (
	"groupie-tracker/internal/handlers"
	"log"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist/", handlers.ArtistHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	img := http.FileServer(http.Dir("./ui/html/img/"))
	mux.Handle("/img/", http.StripPrefix("/img", img))

	log.Println("Running server on http://localhost:7000")
	if err := http.ListenAndServe(":7000", mux); err != nil {
		log.Println(err)
		return
	}
}
