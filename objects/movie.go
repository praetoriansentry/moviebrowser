package objects

type (
	Movie struct {
		Ssid       string
		Title      string
		RawTitle   string
		Year       string `json:"-"`
		Rated      string `json:"-"`
		Released   string `json:"-"`
		Runtime    string `json:"-"`
		Genre      string `json:"-"`
		Director   string `json:"-"`
		Language   string `json:"-"`
		PosterUrl  string
		Metascore  string `json:"-"`
		ImdbRating string `json:"-"`
		ImdbVotes  string `json:"-"`
		ImdbId     string `json:"-"`
	}
)
