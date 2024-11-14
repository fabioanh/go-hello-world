package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []album{
	{ID: "1", Title: "Dookie", Artist: "Green Day", Price: 53.76},
	{ID: "2", Title: "Lo mejor que hay en mi vida", Artist: "Andres Cepeda", Price: 30.0},
	{ID: "3", Title: "Whatever people say I am, that's what I'm not", Artist: "Arctic Monkeys", Price: 45.13},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbum)

	router.Run("localhost:8080")
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(context *gin.Context) {
	// id, _ := context.Params.Get("id")
	id := context.Param("id")

	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	context.Status(http.StatusNotFound)
}

func postAlbum(context *gin.Context) {
	fmt.Println("Posting album")
	var newAlbum album
	if err := context.BindJSON(&newAlbum); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
