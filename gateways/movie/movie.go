package movie

import (
	"../../data"
	"../../objects"
	"database/sql"
	"log"
)

func GetMovie(movieId string) *objects.Movie {
	e := data.GetQueryEngine()
	rows, err := e.Query("select * from movie where ssid = ? limit 1;", movieId)

	if err != nil {
		log.Printf("There was an issue fetching the movies: %q", err)
		return nil
	}
	defer rows.Close()
	return rowToMovie(rows)
}

func UpdateImdbId(movieId string, imdbId string) bool {
	e := data.GetQueryEngine()
	_, err := e.Exec("update movie set imdbid = ? where ssid = ?;", imdbId, movieId)

	if err != nil {
		log.Printf("There was an issue updating the imdb id: %q", err)
		return false
	}

	return true

}
func GetAllMovieSummary() []*objects.Movie {

	e := data.GetQueryEngine()
	rows, err := e.Query("select ssid, title, raw_title, poster_url from movie;")
	movies := make([]*objects.Movie, 0)

	if err != nil {
		log.Printf("There was an issue fetching the movies: %q", err)
		return movies
	}
	defer rows.Close()

	for rows.Next() {
		movie := new(objects.Movie)
		err := rows.Scan(
			&movie.Ssid,
			&movie.Title,
			&movie.RawTitle,
			&movie.PosterUrl,
		)
		if err != nil {
			log.Printf("There was an issue scanning the movie results: %q", err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func GetSimilarMovies(movieId string) []*objects.Movie {
	e := data.GetQueryEngine()
	rows, err := e.Query(`SELECT movie.*
	    FROM MOVIE
	      JOIN (

		     SELECT
		       count(*) vote_count,
		       vote.movie_id
		     FROM VOTE
		       JOIN (
			      SELECT DISTINCT voter_id
			      FROM vote
			      WHERE movie_id = ?
			    ) movie_voters
			 ON movie_voters.voter_id = vote.voter_id
		     WHERE vote.movie_id != ?
		     GROUP BY vote.movie_id
		   ) similar_movies
		ON similar_movies.movie_id = movie.ssid
	    ORDER BY vote_count
	      DESC
	    LIMIT 25;`, movieId, movieId)
	movies := make([]*objects.Movie, 0)

	if err != nil {
		log.Printf("There was an issue fetching the movies: %q", err)
		return movies
	}
	defer rows.Close()

	for rows.Next() {
		movies = append(movies, scanMovie(rows))
	}
	return movies
}

func GetMoviesByVoter(voterId string) []*objects.Movie {
	e := data.GetQueryEngine()
	rows, err := e.Query(`SELECT *
	    FROM movie
	    WHERE ssid IN (
	      SELECT vote.movie_id
	      FROM VOTE
	      WHERE vote.voter_id = ?)
	    ORDER BY movie.raw_title;`, voterId)
	movies := make([]*objects.Movie, 0)

	if err != nil {
		log.Printf("There was an issue fetching the movies: %q", err)
		return movies
	}
	defer rows.Close()

	for rows.Next() {
		movies = append(movies, scanMovie(rows))
	}
	return movies
}

func rowToMovie(row *sql.Rows) *objects.Movie {
	if row.Next() {
		return scanMovie(row)
	}
	return nil
}

func scanMovie(row *sql.Rows) *objects.Movie {
	movie := new(objects.Movie)
	err := row.Scan(
		&movie.Ssid,
		&movie.Title,
		&movie.RawTitle,
		&movie.Year,
		&movie.Rated,
		&movie.Released,
		&movie.Runtime,
		&movie.Genre,
		&movie.Director,
		&movie.Language,
		&movie.PosterUrl,
		&movie.Metascore,
		&movie.ImdbRating,
		&movie.ImdbVotes,
		&movie.ImdbId,
	)
	if err != nil {
		log.Printf("There was an issue scanning the movie results: %q", err)
		return nil
	}
	return movie
}
