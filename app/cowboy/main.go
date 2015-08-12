package main

import (
	"github.com/vishaltelangre/cowboy/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/app/cowboy/powers/excuse"
	"github.com/vishaltelangre/cowboy/app/cowboy/powers/movie_lookup"
	"github.com/vishaltelangre/cowboy/app/cowboy/powers/producthunt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(
			http.StatusMovedPermanently,
			"https://github.com/vishaltelangre/cowboy",
		)
	})
	r.POST("/movie.:format", movie_lookup.Handler)
	r.POST("/excuse.:format", excuse.Handler)
	r.POST("/producthunt/posts.:format", producthunt.PostsHandler)

	// TODO:
	// r.POST("/producthunt/collections.:format", producthunt.CollHandler)
	// r.POST("/hn/best.:format", hn.Handler)
	// r.POST("/fortune.:format", fortune.Handler)
	// r.POST("/forecast.:format", forecast.Handler)
	// r.POST("/define.:format", dict.Handler)
	// r.POST("/wiki.:format", wiki.Handler)

	r.Run(":" + port)
}
