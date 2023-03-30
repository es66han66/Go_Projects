package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var m Movie
	json.NewDecoder(r.Body).Decode(&m)
	m.ID = strconv.Itoa((rand.Int()))
	movies = append(movies, m)
	json.NewEncoder(w).Encode(m)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	var m Movie
	json.NewDecoder(r.Body).Decode(&m)

	for index, item := range movies {

		if item.ID == params["id"] {
			m.ID = item.ID
			movies = append(movies[:index], movies[index+1:]...)
			movies = append(movies, m)
			break
		}
	}

	json.NewEncoder(w).Encode(m)
}

func main() {

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "First_Movie",
		Director: &Director{
			Firstname: "Eshan",
			Lastname:  "Gupta",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "5678",
		Title: "Second_Movie",
		Director: &Director{
			Firstname: "Eshan1",
			Lastname:  "Gupta1",
		},
	})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
