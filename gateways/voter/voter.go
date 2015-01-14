package voter

import (
	"../../data"
	"../../objects"
	"database/sql"
	"log"
)

func GetVotersForMovie(movieId string) []*objects.Voter {
	e := data.GetQueryEngine()
	voters := make([]*objects.Voter, 0)
	rows, err := e.Query("select voter.ssid, voter.name from vote, voter  where vote.voter_id = voter.ssid and  vote.movie_id = ?;", movieId)

	if err != nil {
		log.Printf("There was an issue fetching the movies: %q", err)
		return voters
	}
	defer rows.Close()

	for rows.Next() {
		voters = append(voters, scanVoter(rows))
	}
	return voters
}

func GetVoterById(voterId string) *objects.Voter {
	e := data.GetQueryEngine()
	rows, err := e.Query("select voter.ssid, voter.name from voter where voter.ssid = ?;", voterId)

	if err != nil {
		log.Printf("There was an issue fetching the voter: %q", err)
		return nil
	}
	defer rows.Close()

	if rows.Next() {
		return scanVoter(rows)
	}
	return nil
}

func scanVoter(rows *sql.Rows) *objects.Voter {
	voter := new(objects.Voter)
	err := rows.Scan(
		&voter.Ssid,
		&voter.Name,
	)
	if err != nil {
		log.Printf("There was an issue scanning the voter information: %q", err)
		return nil
	}
	return voter
}
