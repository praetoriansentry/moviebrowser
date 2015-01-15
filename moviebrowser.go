package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {

	// Handle all of the dynamic pages
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/movie/", handlers.Movie)
	http.HandleFunc("/movie/list.json", handlers.MovieJson)
	http.HandleFunc("/voter/", handlers.Voter)

	// Delegate statick requests to http.FileServer. All of those
	// requests will look inside the /static folder
	http.Handle("/js/", http.FileServer(http.Dir("./static")))
	http.Handle("/css/", http.FileServer(http.Dir("./static")))
	http.Handle("/img/", http.FileServer(http.Dir("./static")))
	http.Handle("/favicon.ico", http.FileServer(http.Dir("./static/img")))

	log.Println("Starting Server")
	// Start the server on port 8888.
	log.Fatal(http.ListenAndServe(":8888", nil))
}
