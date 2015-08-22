package producthunt

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/app/cowboy/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"text/template"
)

const (
	slackPostsRespTmpl = `
{{ range .Posts }}
>{{ if .Featured }} :collision: {{ end }}*<{{.GetURL}}|{{.Name}}>* - {{.Tagline}}
> :thumbsup: {{.VotesCount}}, :speech_balloon: {{.CommentsCount}} [<{{.DiscussionURL}}|Discuss>]
{{ if .HasMakers }}> *Makers*: {{ range .Makers }}{{ if .HasWebsite }}<{{.WebsiteURL}}|{{.Name}}>{{ else }}{{.Name}}{{ end }} {{ if .HasHeadline }}({{.Headline}}){{ end }}, {{ end }}{{ end }}
> *Submitted On*: {{.SubmittedOn}}
{{ end }}
`
)

type Posts struct {
	Posts []Post `json:"posts, omitempty"`
}

type Maker struct {
	Name       string `json:"name, omitempty"`
	Headline   string `json:"headline, omitempty"`
	WebsiteURL string `json:"website_url, omitempty"`
	ProfileURL string `json:"profile_url, omitempty"`
}

type Post struct {
	Name          string            `json:"name, omitempty"`
	Tagline       string            `json:"tagline, omitempty"`
	Featured      bool              `json:"featured, omitempty"`
	CommentsCount int               `json:"comments_count, omitempty"`
	VotesCount    int               `json:"votes_count, omitempty"`
	DiscussionURL string            `json:"discussion_url, omitempty"`
	GetURL        string            `json:"redirect_url, omitempty"`
	ScreenshotURL map[string]string `json:"screenshot_url, omitempty"`
	Makers        []Maker           `json:makers, omitempty`
	SubmittedOn   string            `json:"day, omitempty"`
}

func (m *Maker) HasWebsite() bool {
	return !(m.WebsiteURL == "")
}

func (m *Maker) HasHeadline() bool {
	return !(m.Headline == "")
}

func (p *Post) HasMakers() bool {
	return len(p.Makers) > 0
}

func getPosts(daysAgo int) (*Posts, error) {
	var apiURL *url.URL
	apiURL, err := url.Parse(APIBaseURL + "/v1/posts")
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("days_ago", strconv.Itoa(daysAgo))
	apiURL.RawQuery = parameters.Encode()

	content, err := utils.GetContent(apiURL.String(), requestHeaders)

	if err != nil {
		return nil, err
	}

	var posts Posts
	err = json.Unmarshal(content, &posts)
	if err != nil {
		return nil, err
	}

	return &posts, err
}

// formatPostsSlackResp creates Slack-compatible post listing string
func formatPostsSlackResp(posts *Posts) (string, error) {
	buf := new(bytes.Buffer)
	t := template.New("PHPosts")
	t, err := t.Parse(slackPostsRespTmpl)
	if err != nil {
		return "", err
	}

	err = t.Execute(buf, posts)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// PostsHandler is a route handler for '/producthunt/posts.:format' route
func PostsHandler(c *gin.Context) {
	requestType := c.Param("format")
	daysAgo, err := strconv.Atoi(c.Request.PostFormValue("text"))
	if err != nil {
		daysAgo = 0
	}

	log.Printf("Days Ago: %d", daysAgo)

	posts, err := getPosts(daysAgo)

	switch requestType {
	case "json":
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Response": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, posts)
		}
	case "slack":
		text, err := formatPostsSlackResp(posts)
		if err != nil {
			c.String(http.StatusNotFound, "Not Found")
		}

		c.String(http.StatusOK, text)
	default:
		c.JSON(http.StatusUnsupportedMediaType, nil)
	}
}
