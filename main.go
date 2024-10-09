package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
}

func main() {
	router := gin.Default()

	log.Println("Starting server...")
	router.Run(":8080")
}
