package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	//gin.SetMode(gin.ReleaseMode) //PROD
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAblbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("0.0.0.0:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	res := c.BindJSON(&newAlbum)
	if res != nil {
		return
	}
	for _, a := range albums {
		if a.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "id ja cadastrado"})
			return
		}
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAblbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum album
			res := c.BindJSON(&updatedAlbum)
			if res != nil {
				return
			}
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusCreated, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
