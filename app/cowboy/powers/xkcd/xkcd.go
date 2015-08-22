package xkcd

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/app/cowboy/utils"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

const (
	APIBaseURL = "http://xkcd.com"
)

var (
	slackRespTmpl = `<{{.ComicLink}}|Issue #{{.ComicID}}> [<{{.ExplainLink}}|Explain>]
`
)

type Comic struct {
	ComicID    int    `json:"num, omitempty"`
	Title      string `json:"title, omitempty"`
	SafeTitle  string `json:"safe_title, omitempty"`
	HoverText  string `json:"alt, omitempty"`
	Transcript string `json:"transcript, omitempty"`
	Link       string `json:"link, omitempty"`
	News       string `json:"news, omitempty"`
	ImageURL   string `json:"img, omitempty"`
	Year       string `json:"year, omitempty"`
	Month      string `json:"month, omitempty"`
	Day        string `json:"day, omitempty"`
}

func (c *Comic) ComicLink() string {
	return APIBaseURL + "/" + strconv.Itoa(c.ComicID)
}

func (c *Comic) ExplainLink() string {
	return "http://explainxkcd.com/" + strconv.Itoa(c.ComicID)
}

func fetch(apiURL *url.URL) (*Comic, error) {
	content, err := utils.GetContent(apiURL.String(), nil)

	if err != nil {
		return nil, err
	}

	var comic Comic
	err = json.Unmarshal(content, &comic)
	if err != nil {
		return nil, err
	}

	return &comic, err
}

func fetchRecent() (*Comic, error) {
	var apiURL *url.URL
	apiURL, err := url.Parse(APIBaseURL + "/info.0.json")
	if err != nil {
		return nil, err
	}

	return fetch(apiURL)
}

func fetchByID(id string) (*Comic, error) {
	var apiURL *url.URL
	apiURL, err := url.Parse(APIBaseURL + "/" + id + "/info.0.json")
	if err != nil {
		return nil, err
	}

	return fetch(apiURL)
}

func formatSlackResp(comic *Comic) (string, error) {
	buf := new(bytes.Buffer)
	t := template.New("XKCDPost")
	t, err := t.Parse(slackRespTmpl)
	if err != nil {
		return "", err
	}

	err = t.Execute(buf, comic)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Handler is a route handler for '/xkcd.:format' route
func Handler(c *gin.Context) {
	requestType := c.Param("format")
	term := c.Request.PostFormValue("text")

	log.Printf("Comic Query: %s", term)

	var comic *Comic
	var err error

	var lowerQuery = strings.ToLower(term)
	if regexp.MustCompile(`^\d+$`).MatchString(lowerQuery) {
		comic, err = fetchByID(term)
	} else if regexp.MustCompile(`^(?:recent|latest|new|last)?$`).MatchString(lowerQuery) {
		comic, err = fetchRecent()
	}

	switch requestType {
	case "json":
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Response": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, comic)
		}
	case "slack":
		text, err := formatSlackResp(comic)
		if err != nil {
			c.String(http.StatusNotFound, "Not Found")
		}

		c.String(http.StatusOK, text)
	default:
		c.JSON(http.StatusUnsupportedMediaType, nil)
	}
}
