package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albuns = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbuns(c *gin.Context) {
	log.Println("Start returning all albuns")
	c.IndentedJSON(http.StatusOK, albuns)
	log.Println("Object seacherd ", albuns)
}

func postAlbuns(c *gin.Context) {
	var newAlbum album

	log.Println("Save the albuns")
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	log.Println("Object seacherd ", newAlbum)
	albuns = append(albuns, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbunsByID(c *gin.Context) {
	id := c.Param("id")
	log.Println("Returning albuns by id")
	for _, a := range albuns {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			log.Println("Object seacherd ", a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func putByID(c *gin.Context) {
	id := c.Param("id")
	var album album
	for _, a := range albuns {
		if a.ID == id {
			log.Println("object searched :", a)
			if err := c.ShouldBindJSON(&album); err == nil {

				a = album
				log.Println("Updated Object", a)
				c.JSON(http.StatusOK, a)

			}
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albuns", getAlbuns)
	router.GET("/album/:id", getAlbunsByID)
	router.POST("/albuns", postAlbuns)
	router.PUT("/albuns/:id", putByID)

	router.Run("localhost:8080")

}
