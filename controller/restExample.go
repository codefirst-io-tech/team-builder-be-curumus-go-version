package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type errorDto struct {
	Status           bool   `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

func main() {
	router := gin.Default()
	router.GET("/album/all", GetAlbums)
	router.GET("/album/:id", GetAlbumById)
	router.POST("/album", AddAlbums)
	router.DELETE("/album/:id", DeleteAlbum)
	router.Run("localhost:8080")
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(context *gin.Context) {
	id := context.Params.ByName("id")
	var albumById album
	for _, a := range albums {
		if a.ID == id {
			albumById = a
			break
		}
	}
	fmt.Print(albumById)
	if albumById.ID == "" {
		e := errorDto{false, "id bulunamadı", "421"}
		context.IndentedJSON(http.StatusBadRequest, e)
	} else {
		context.IndentedJSON(http.StatusOK, albumById)
	}
}

func AddAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		e := errorDto{false, err.Error(), err.Error()}
		c.IndentedJSON(http.StatusBadRequest, e)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbum(context *gin.Context) {
	id := context.Params.ByName("id")
	var deletedAlbum album
	for i := range albums {
		if albums[i].ID == id {
			deletedAlbum = albums[i]
		}
	}
	deletedAlbumIndex := getIndexByAlbum(deletedAlbum, albums)
	albums = append(albums[:deletedAlbumIndex], albums[deletedAlbumIndex+1:]...)
}

func getIndexByAlbum(album album, albums []album) int {
	for i := range albums {
		if albums[i] == album {
			return i
		}
	}
	return -1
}
