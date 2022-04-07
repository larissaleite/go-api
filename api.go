package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type show struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Network string `json:"network"`
}

var shows = []show{
	{ID: "1", Name: "Homeland", Network: "Showtime"},
	{ID: "2", Name: "Breaking Bad", Network: "AMC"},
	{ID: "3", Name: "Dark", Network: "Netflix"},
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/shows", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, shows)
	})

	r.POST("/shows", func(c *gin.Context) {
		var newShow show

		if err := c.BindJSON(&newShow); err != nil {
			return
		}

		shows = append(shows, newShow)
		c.IndentedJSON(http.StatusCreated, newShow)
	})

	r.GET("/shows/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, show := range shows {
			if show.ID == id {
				c.IndentedJSON(http.StatusOK, show)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TV show not found"})
	})

	r.Run() // runs on 0.0.0.0:8080
}
