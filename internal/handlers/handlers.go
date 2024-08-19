package handlers

import (
	"fmt"
	"groupie-tracker/internal"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ErrorMessage struct {
	Status     int
	StatusText string
}

var artist, err = internal.Unmarshal()

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
		datelist := internal.Suggestions(artist)
		err = tmpl.Execute(w, internal.Search{artist, datelist})
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		tmpl, err := template.ParseFiles("./ui/html/index.html")
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
		datalist := internal.Suggestions(artist)
		search := internal.SearchBar(artist, r.FormValue("search"))

		err = tmpl.Execute(w, internal.Search{search, datalist})
		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		fmt.Println("Method not")
		errorHandler(w, http.StatusMethodNotAllowed)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	if url[2][0] == '0' {
		errorHandler(w, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(url[2])

	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if len(url) > 3 {
		errorHandler(w, http.StatusBadRequest)
		return
	}
	if id <= 0 || id > len(artist) {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artist[id-1])
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	tmpl, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newErr := &ErrorMessage{
		Status:     status,
		StatusText: http.StatusText(status),
	}
	//http.Error(w, http.StatusText(status), status)
	w.WriteHeader(status)
	err = tmpl.Execute(w, newErr)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
