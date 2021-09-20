package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Recipe struct {
	Name string `json:"name"`
	Tag []string `json:"tag"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

func main()  {
	router := gin.Default()

	router.Run(":8080")
}
