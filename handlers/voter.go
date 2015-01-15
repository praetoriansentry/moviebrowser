package handlers

import (
	"../gateways/movie"
	"../gateways/voter"
	"log"
	"net/http"
	"strings"
)

func Voter(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the voter handler")
	pathParts := strings.Split(rq.URL.Path, "/")

	// Manually parsing. If this were a bigger application, it might
	// make sense to do something less hacky
	if len(pathParts) < 3 {
		http.NotFound(rw, rq)
		return
	}
	voterId := pathParts[2]
	currentVoter := voter.GetVoterById(voterId)
	if currentVoter == nil {
		http.NotFound(rw, rq)
		return
	}
	otherMovies := movie.GetMoviesByVoter(voterId)
	voterData := make(map[string]interface{})
	voterData["voter"] = currentVoter
	voterData["movies"] = otherMovies
	sendResponse(rw, rq, "voter", voterData)
}
