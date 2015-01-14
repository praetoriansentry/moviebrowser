package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/movie/", handlers.Movie)
	http.HandleFunc("/movie/list.json", handlers.MovieJson)
	http.HandleFunc("/voter/", handlers.Voter)

	http.Handle("/js/", http.FileServer(http.Dir("./static")))
	http.Handle("/css/", http.FileServer(http.Dir("./static")))
	http.Handle("/img/", http.FileServer(http.Dir("./static")))
	http.Handle("/favicon.ico", http.FileServer(http.Dir("./static/img")))

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
