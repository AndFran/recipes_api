package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)

}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = uuid.New().String()
	recipe.PublishedAt = time.Now().UTC()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusCreated, recipe)
}

func ListRecipesHandler(c *gin.Context) {

}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = id
	// todo find in slice and replace

	c.JSON(http.StatusOK, recipe)
}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	// TODO find in list

	index := 0
	recipes = append(recipes[:index], recipes[index+1:]...) // everything up to (not including) and everything 1 after
	c.JSON(http.StatusNoContent, gin.H{
		"id": id,
	})
}

func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	results := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				results = append(results, recipes[i])
				break
			}
		}
	}
	c.JSON(http.StatusOK, results)
}

func main() {
	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/:id", UpdateRecipeHandler)
	router.GET("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search", SearchRecipesHandler)

	log.Println("Starting server...")
	router.Run(":8080")
}
