package movie_lookup

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vishaltelangre/cowboy/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/app/cowboy/utils"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

const (
	// slackRespTmpl is a customised template to display rich formatted movie details on Slack
	slackRespTmpl = `
:movie_camera: <http://imdb.com/title/{{.IMDbID}}|{{.Title}}> ({{.Year}})

*Rating:* IMDb - *{{.IMDbRating}}* | Rotten Tomatoes - *{{.TomatoRating}}*
*Genre:* {{.Genre}}
*Language:* {{.Language}}
*Duration:* {{.Runtime}}
*Director:* {{.Director}}
*Star Cast:* {{.Actors}}
*Awards:* {{.Awards}}

*Plot:*
> _{{.Plot}}_

*Tomato Consensus:*
> _{{.TomatoConsensus}}_
`
)

// Movie defines structure of a movie with a lot details
type Movie struct {
	Title           string `json:"Title, omitempty"`
	Year            string `json:"Year, omitempty"`
	Rated           string `json:"Rated, omitempty"`
	ReleasedOn      string `json:"Released, omitempty"`
	Runtime         string `json:"Runtime, omitempty"`
	Genre           string `json:"Genre, omitempty"`
	Director        string `json:"Director, omitempty"`
	Writer          string `json:"Writer, omitempty"`
	Actors          string `json:"Actors, omitempty"`
	Plot            string `json:"Plot, omitempty"`
	Language        string `json:"Language, omitempty"`
	Country         string `json:"Country, omitempty"`
	Awards          string `json:"Awards, omitempty"`
	PosterURL       string `json:"Poster, omitempty"`
	IMDbRating      string `json:"imdbRating, omitempty"`
	IMDbVotes       string `json:"imdbVotes, omitempty"`
	IMDbID          string `json:"imdbID, omitempty"`
	Type            string `json:"Type, omitempty"`
	TomatoRating    string `json:"TomatoRating, omitempty"`
	TomatoConsensus string `json:"tomatoConsensus, omitempty"`

	IsValidResponse string `json:"Response"`
}

// lookupMovie fetches movie details from a third-party API
func lookupMovie(term string) (*Movie, error) {
	var apiURL *url.URL
	apiURL, err := url.Parse("http://www.omdbapi.com")
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("t", term)
	parameters.Add("plot", "full")
	parameters.Add("r", "json")
	parameters.Add("tomatoes", "true")
	parameters.Add("v", "1")
	apiURL.RawQuery = parameters.Encode()

	content, err := utils.GetContent(apiURL.String())

	if err != nil {
		return nil, err
	}

	var movie Movie
	err = json.Unmarshal(content, &movie)
	if err != nil {
		return nil, err
	}

	if movie.IsValidResponse == "False" {
		return nil, errors.New(movie.IsValidResponse)
	}

	return &movie, err
}

// formatSlackResp creates Slack-compatible movie details string
func formatSlackResp(movie *Movie) (string, error) {
	buf := new(bytes.Buffer)
	t := template.New("SlackMovie")
	t, err := t.Parse(slackRespTmpl)
	if err != nil {
		return "", err
	}

	err = t.Execute(buf, movie)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// MovieHandler is a route handler for '/movie.:format' route
func MovieHandler(c *gin.Context) {
	requestType := c.Param("format")
	term := c.Request.PostFormValue("text")

	log.Printf("Movie Query: %s", term)

	movie, err := lookupMovie(term)

	switch requestType {
	case "json":
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Response": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, movie)
		}
	case "slack":
		text, err := formatSlackResp(movie)
		if err != nil {
			c.String(http.StatusNotFound, "Not Found")
		}

		c.String(http.StatusOK, text)
	default:
		c.JSON(http.StatusUnsupportedMediaType, nil)
	}
}
