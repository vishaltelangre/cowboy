package excuse

import (
	"github.com/vishaltelangre/cowboy/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/vishaltelangre/cowboy/app/cowboy/utils"
	"net/http"
	"net/url"
	"strings"
)

// Excuse defines structure of a programmer's excuse in respond to a bad happening
type Excuse struct {
	Text     string `json:"Text"`
	Response string `json:"Response"`
}

func getAnExcuse() (*Excuse, error) {
	var endpointURL *url.URL
	endpointURL, err := url.Parse("http://www.programmerexcuses.com/")
	if err != nil {
		return nil, err
	}

	content, err := utils.GetContent(endpointURL.String())
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return nil, err
	}

	var excuseText string

	var f func(*html.Node, bool)
	f = func(n *html.Node, canExtractText bool) {
		// Extract text content if node is a TextNode
		if canExtractText && n.Type == html.TextNode {
			excuseText = n.Data
		}

		// Tell whether text content can be extracted or not.
		// Here, we will recursively traverse for 'a' ElementNode.
		canExtractText = canExtractText || (n.Type == html.ElementNode && n.Data == "a")

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, canExtractText)
		}
	}
	f(doc, false)

	return &Excuse{Text: excuseText, Response: "true"}, err
}

// Handler is a route handler for '/excuse.:format' route
func Handler(c *gin.Context) {
	requestType := c.Param("format")
	excuse, err := getAnExcuse()

	switch requestType {
	case "json":
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Response": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, excuse)
		}
	case "slack":
		if err != nil {
			c.String(http.StatusNotFound, "> I am not feeling well\n")
		}

		c.String(http.StatusOK, "\n> "+excuse.Text+"\n")
	default:
		c.String(http.StatusUnsupportedMediaType, "")
	}
}
