package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func postAlbums(context *gin.Context) {
	var record Album

	// bind received JSON to `record`
	if err := context.BindJSON(&record); err != nil {
		log.Println(err)
		return
	}

	_, exists := albums[record.ID]

	if exists {
		context.IndentedJSON(http.StatusConflict, gin.H{"message": "album already exists"})
		return
	}

	// add the new album to the slice
	albums[record.ID] = record

	context.IndentedJSON(http.StatusCreated, record)
}

func getAlbums(context *gin.Context) {
	keys := make([]Album, 0, len(albums))
	for _, album := range albums {
		keys = append(keys, album)
	}

	context.IndentedJSON(http.StatusOK, keys)
}

func getAlbumByID(context *gin.Context) {
	id := context.Param("id")

	record, exists := albums[id]

	if exists {
		context.IndentedJSON(http.StatusOK, record)
		return
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(context *gin.Context) {
	id := context.Param("id")
	var record Album

	_, exists := albums[id]

	if !exists {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	albums[id] = record
	context.IndentedJSON(http.StatusAccepted, record)
}

func deleteAlbumByID(context *gin.Context) {
	id := context.Param("id")

	_, exists := albums[id]

	if exists {
		delete(albums, id)
	}

	context.IndentedJSON(http.StatusNoContent, gin.H{})
}

// Sample data
var albums = map[string]Album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func Router() *gin.Engine {
	log.SetPrefix("api: ")
	log.SetFlags(0) // suppresses the time, source file, and line number

	router := gin.Default()

	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PATCH("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	return router
}
