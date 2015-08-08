package main

import (
	"github.com/vishaltelangre/cowboy/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/vishaltelangre/cowboy/app/cowboy/powers"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := gin.Default()

	r.POST("/movie.:format", movie_lookup.MovieHandler)

	// TODO:
	// r.POST("/fortune.:format", fortune.FortuneHandler)
	// r.POST("/forecast.:format", forecast.ForecastHandler)
	// r.POST("/define.:format", encyclopedia.DefinitionHandler)

	r.Run(":" + port)
}
