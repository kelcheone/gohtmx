package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {

	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	flag.Parse()

	films := map[string][]Film{
		"comedy": {
			{Title: "Ghostbusters", Director: "Ivan Reitman"},
			{Title: "Bridesmaids", Director: "Paul Feig"},
			{Title: "Knives Out", Director: "Rian Johnson"},
		},
		"drama": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Schindler's List", Director: "Steven Spielberg"},
			{Title: "Casablanca", Director: "Michael Curtiz"},
		},
		"thriller": {
			{Title: "The Prestige", Director: "Christopher Nolan"},
			{Title: "The Usual Suspects", Director: "Bryan Singer"},
			{Title: "North by Northwest", Director: "Alfred Hitchcock"},
		},
	}

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		tmp1 := template.Must(template.ParseFiles("template/index.html"))
		tmp1.Execute(w, films)
	}
	AddFilm := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)

		filmTitle := r.PostFormValue("title")
		filmDirector := r.PostFormValue("director")
		filmGenre := r.PostFormValue("genre")

		fmt.Printf("Adding %s by %s to %s films\n", filmTitle, filmDirector, filmGenre)

		// add the film to the database
		films[filmGenre] = append(films[filmGenre], Film{Title: filmTitle, Director: filmDirector})
		tmp1 := template.Must(template.ParseFiles("template/index.html"))
		tmp1.ExecuteTemplate(w, "comedy-list-element", films[filmGenre][len(films[filmGenre])-1])

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", AddFilm)

	fmt.Printf("Starting server at %s\n", *listenAddr)

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
