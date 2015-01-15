package handlers

import (
	"../gateways/movie"
	"../gateways/voter"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Movie(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the movie handler")
	// Here I'm manually parsing the url and grabbing the various
	// parts. There are libraries that do a better job of this,
	// but if you are doing something really basic, this is fine
	pathParts := strings.Split(rq.URL.Path, "/")
	pathPartCount := len(pathParts)
	if pathPartCount < 3 {
		// Convenient...
		http.NotFound(rw, rq)
	}
	movieId := pathParts[2]

	if pathPartCount == 4 && pathParts[3] == "image" {
		getMovieImage(rw, rq, movieId)
		return
	}

	currentMovie := movie.GetMovie(movieId)
	if currentMovie == nil {
		http.NotFound(rw, rq)
		return
	}
	// Fetch all the data and prepare it for transformation with
	// the templates.
	movieData := make(map[string]interface{})
	movieData["voters"] = voter.GetVotersForMovie(movieId)
	movieData["similar"] = movie.GetSimilarMovies(movieId)
	movieData["movie"] = currentMovie

	sendResponse(rw, rq, "movie", movieData)
}

// This is handler is for the autocomplete list
func MovieJson(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the movie json handler")
	movieList := movie.GetAllMovieSummary()
	marshalToJsonAndSend(rw, movieList)
}

// If a movie image is request, we look that up and retrieve it
// directly. Proxying is necessary because the IMDb server won't allow
// image requests with unrecognized referrers. I've added cache
// headers (hopefully correct) to help offset the load.
func getMovieImage(rw http.ResponseWriter, rq *http.Request, movieId string) {
	log.Printf("Getting image for movie %s", movieId)
	movie := movie.GetMovie(movieId)
	if movie == nil {
		http.NotFound(rw, rq)
		return
	}

	if movie.PosterUrl == "N/A" {
		http.ServeFile(rw, rq, "./static/img/clapper.png")
		return
	}

	resp, err := http.Get(movie.PosterUrl)
	if err != nil {
		log.Printf("There was an issue fetching the image: %q", err)
		http.NotFound(rw, rq)
		return
	}

	defer resp.Body.Close()
	imageData, imgErr := ioutil.ReadAll(resp.Body)
	if imgErr != nil {
		log.Printf("There was an issue reading the image: %q", imgErr)
		http.NotFound(rw, rq)
		return
	}
	rw.Header().Add("Content-type", "image/jpeg")
	rw.Header().Add("Cache-Control", "max-age=86400")
	rw.Write(imageData)

}
